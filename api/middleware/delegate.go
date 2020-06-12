package middleware

import (
	"fmt"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/email"

	"github.com/getsentry/sentry-go"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) delegateToGentype(delegate models.Delegate) gentypes.Delegate {
	var profileURL *string
	if delegate.ProfileKey != nil {
		url := uploads.GetImgixURL(*delegate.ProfileKey)
		profileURL = &url
	}

	createdAt := delegate.CreatedAt.Format(time.RFC3339)
	lastLogin := delegate.LastLogin.Format(time.RFC3339)
	return gentypes.Delegate{
		CreatedAt:       &createdAt,
		UUID:            delegate.UUID,
		FirstName:       delegate.FirstName,
		LastName:        delegate.LastName,
		JobTitle:        delegate.JobTitle,
		Telephone:       delegate.Telephone,
		Email:           delegate.Email,
		CompanyUUID:     delegate.CompanyUUID,
		TTC_ID:          delegate.TtcId,
		ProfileImageURL: profileURL,
		LastLogin:       lastLogin,
	}
}

func (g *Grant) delegatesToGentype(delegates []models.Delegate) []gentypes.Delegate {
	var genDelegates []gentypes.Delegate
	for _, delegate := range delegates {
		genDelegates = append(genDelegates, g.delegateToGentype(delegate))
	}

	return genDelegates
}

func (g *Grant) DelegateExists(email string, ttcId string) bool {
	// Only managers and admins can check a delegate exists
	if !g.IsManager || !g.IsAdmin {
		return false
	}

	query := database.GormDB.Where("email = ? or ttc_id = ?", email, ttcId)

	// Managers can only check inside their own company
	if g.IsManager {
		query = query.Where("company_uuid = ?", g.Claims.Company)
	}

	query = query.First(&models.Delegate{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find delegate for Email: %s", email)
		return false
	}

	return true
}

func (g *Grant) Delegate(uuid gentypes.UUID) (models.Delegate, error) {
	var delegate models.Delegate
	err := database.GormDB.Where("uuid = ?", uuid).Find(&delegate).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.Delegate{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, err, "Unable to find delegate")
		return models.Delegate{}, &errors.ErrWhileHandling
	}

	if !g.IsAdmin &&
		!(g.IsManager && g.Claims.Company == delegate.CompanyUUID) &&
		!(g.IsDelegate && g.Claims.UUID == delegate.UUID) {
		return models.Delegate{}, &errors.ErrUnauthorized
	}

	return delegate, nil
}

func (g *Grant) GetDelegateByUUID(UUID gentypes.UUID) (gentypes.Delegate, error) {
	delegate, err := g.Delegate(UUID)

	if err != nil {
		return gentypes.Delegate{}, err
	}

	return g.delegateToGentype(delegate), nil
}

func filterDelegate(query *gorm.DB, filter *gentypes.DelegatesFilter) *gorm.DB {
	if filter != nil {
		query = filterUser(query, &filter.UserFilter)

		if filter.TTC_ID != nil && *filter.TTC_ID != "" {
			query = query.Where("ttc_id ILIKE ?", "%%"+*filter.TTC_ID+"%%")
		}

		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
	}

	return query
}

func (g *Grant) GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy) ([]gentypes.Delegate, gentypes.PageInfo, error) {
	if !g.IsAdmin && !g.IsManager {
		return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var delegates []models.Delegate

	query := filterDelegate(database.GormDB, filter)
	// only get manager's company's delegates
	if g.IsManager {
		query = query.Where("company_uuid = ?", g.Claims.Company.UUID)
	}

	// Count the total filtered dataset
	var count int32
	countErr := query.Model(&models.Delegate{}).Limit(MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		g.Logger.Log(sentry.LevelError, countErr, "Unable to count delegates")
		return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title", "ttc_id"}, "created_at DESC")
	if orderErr != nil {
		return []gentypes.Delegate{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query = query.Find(&delegates)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to find delegates")
		return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return g.delegatesToGentype(delegates), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(delegates)),
	}, nil
}

func generateTTCID(tx *gorm.DB, g *Grant, companyName string, delegateFName string, delegateLName string) (string, error) {

	var (
		baseID = fmt.Sprintf("%s-%s%s", slug.Make(companyName), slug.Make(delegateFName), slug.Make(delegateLName))
		newID  = baseID
		iter   = 1 // Starts at 1 because company-fnamelname-0 looks funny
	)

	// TODO: Could use a LIKE query to speed this up
	for iter < 20 {
		// Check if ttcID created already exists
		var delegate models.Delegate
		err := tx.Where("ttc_id = ?", newID).First(&delegate)
		if !err.RecordNotFound() && err.Error != nil {
			g.Logger.Log(sentry.LevelError, err.Error, "TTC_ID find error")
			return "", &errors.ErrWhileHandling
		}

		// If record doesn't exist return TTCID
		if err.RecordNotFound() {
			return newID, nil
		}

		// If not found keep trying new ones
		newID = fmt.Sprintf("%s-%d", baseID, iter)
		iter = iter + 1
	}

	g.Logger.LogMessage(sentry.LevelError, "Iteration exceeded max")
	return "", &errors.ErrWhileHandling
}

func (g *Grant) GenerateFinaliseDelegateToken(delegateUUID gentypes.UUID) (string, error) {
	if !g.IsAdmin && !g.IsManager {
		return "", &errors.ErrUnauthorized
	}

	delegate, err := g.Delegate(delegateUUID)
	if err != nil {
		return "", err
	}

	// Only allow managers that are in the same company as the delegate
	if g.IsManager && delegate.CompanyUUID != g.Claims.Company {
		g.Logger.LogMessage(sentry.LevelWarning, "SEC: Attempt to get finalise delegate token for another company's delegate")
		return "", &errors.ErrUnauthorized
	}

	// Check delegate doesn't already have a password
	if delegate.Password != nil {
		return "", &errors.ErrDelegateFinalised
	}

	token, err := auth.GenerateFinaliseDelegateToken(auth.FinaliseDelegateClaims{
		delegate.UUID,
	})

	if err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to generate finalise delegate token")
		return "", &errors.ErrWhileHandling
	}

	return token, nil
}

func (g *Grant) CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, *string, error) {
	if !g.IsAdmin && !g.IsManager {
		return gentypes.Delegate{}, nil, &errors.ErrUnauthorized
	}

	if err := delegateDetails.Validate(); err != nil {
		return gentypes.Delegate{}, nil, err
	}

	var companyUUID gentypes.UUID
	// If you're an admin you need to provide the company UUID
	if g.IsAdmin {
		if delegateDetails.CompanyUUID != nil {
			companyUUID = *delegateDetails.CompanyUUID
		} else {
			return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
		}
	} else {
		companyUUID = g.Claims.Company
	}

	// Check if company exists
	if !g.companyExists(companyUUID) {
		return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
	}

	// Get company
	comp, err := g.GetCompanyByUUID(companyUUID)
	if err != nil {
		return gentypes.Delegate{}, nil, err
	}

	var (
		s3UploadKey       *string
		password          *string
		realPass          *string
		needsGeneratePass = delegateDetails.GeneratePassword != nil && *delegateDetails.GeneratePassword
	)

	// Check if autogenerating password is required
	if needsGeneratePass {
		pass, err := auth.GenerateSecurePassword(10)
		if err != nil {
			g.Logger.Log(sentry.LevelError, err, "Unable to generate secure password")
			return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
		}
		password = &pass
		newPass := pass
		realPass = &newPass // So that this pointer is not altered later on (by BeforeCreate)
	}

	// Check if upload token is valid
	if delegateDetails.ProfileImageUploadToken != nil {
		tmpUploadKey, err := uploads.VerifyUploadSuccess(*delegateDetails.ProfileImageUploadToken, "profileImage")
		if err != nil {
			return gentypes.Delegate{}, nil, &errors.ErrUploadTokenInvalid
		}

		s3UploadKey = &tmpUploadKey
	}

	// Create a transaction to ensure that a new TTC_ID isn't created before we insert ours
	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ttcId, err := generateTTCID(tx, g, comp.Name, delegateDetails.FirstName, delegateDetails.LastName)
	if err != nil {
		tx.Rollback()
		return gentypes.Delegate{}, nil, err
	}

	delegate := models.Delegate{
		FirstName:   delegateDetails.FirstName,
		LastName:    delegateDetails.LastName,
		JobTitle:    delegateDetails.JobTitle,
		Telephone:   delegateDetails.Telephone,
		Password:    password,
		Email:       delegateDetails.Email,
		CompanyUUID: companyUUID,
		TtcId:       ttcId,
		ProfileKey:  s3UploadKey,
	}
	createErr := tx.Create(&delegate).Error
	if createErr != nil {
		tx.Rollback()
		g.Logger.Log(sentry.LevelError, createErr, "Unable to create delegate")
		return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Error commiting create delegate transaction")
		return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
	}

	// Send transactional email
	// If not generated password, send an email to the user
	if !needsGeneratePass {
		token, err := g.GenerateFinaliseDelegateToken(delegate.UUID)
		if err != nil {
			tx.Rollback()
			return gentypes.Delegate{}, nil, err
		}

		if delegate.Email == nil {
			tx.Rollback()
			g.Logger.LogMessage(sentry.LevelError, "Delegate email is nil")
			return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
		}

		err = email.SendFinaliseAccountEmail(token, delegate.FirstName, *delegate.Email)
		if err != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelWarning, err, "Unable to send finalise account email")
			return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
		}

	}

	return g.delegateToGentype(delegate), realPass, nil
}
