package middleware

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) delegateToGentype(delegate models.Delegate) gentypes.Delegate {
	createdAt := delegate.CreatedAt.Format(time.RFC3339)
	return gentypes.Delegate{
		User: gentypes.User{
			CreatedAt: &createdAt,
			UUID:      delegate.UUID,
			Email:     delegate.Email,
			FirstName: delegate.FirstName,
			LastName:  delegate.LastName,
			JobTitle:  delegate.JobTitle,
			Telephone: delegate.Telephone,
		},
		CompanyUUID: delegate.CompanyUUID,
		TTC_ID:      delegate.TtcId,
	}
}

func (g *Grant) delegatesToGentype(delegates []models.Delegate) []gentypes.Delegate {
	var genDelegates []gentypes.Delegate
	for _, delegate := range delegates {
		genDelegates = append(genDelegates, g.delegateToGentype(delegate))
	}

	return genDelegates
}

func (g *Grant) delegateExists(email string, ttcId string) bool {
	query := database.GormDB.Where("email = ? or ttc_id = ?", email, ttcId).First(&models.Delegate{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find delegate for Email: %s", email)
		return false
	}

	return true
}

func (g *Grant) GetDelegateByUUID(UUID gentypes.UUID) (gentypes.Delegate, error) {
	var delegate models.Delegate
	err := database.GormDB.Where("uuid = ?", UUID).First(&delegate).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return gentypes.Delegate{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, err, "Unable to find delegate")
		return gentypes.Delegate{}, &errors.ErrWhileHandling
	}

	if !g.IsAdmin &&
		!(g.IsManager && g.Claims.Company == delegate.CompanyUUID) &&
		!(g.IsDelegate && g.Claims.UUID == delegate.UUID) {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	return g.delegateToGentype(delegate), nil
}

func filterDelegate(query *gorm.DB, filter *gentypes.DelegatesFilter) *gorm.DB {
	if filter != nil {
		query = filterUser(query, &filter.UserFilter)
		if filter.TTC_ID != nil && *filter.TTC_ID != "" {
			query = query.Where("ttc_id ILIKE ?", "%%"+*filter.TTC_ID+"%%")
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
		iter   = 0
	)

	for iter < 20 {
		// Check if ttcID created already exists
		var delegate models.Delegate
		err := tx.Where("ttc_id = ?", newID).First(&delegate)
		if err.Error != nil {
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

func (g *Grant) CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, error) {
	if !g.IsAdmin && !g.IsManager {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	// TODO: Generate and populate TTC_ID

	var companyUUID gentypes.UUID
	// If you're an admin you need to provide the company UUID
	if g.IsAdmin {
		if delegateDetails.CompanyUUID != nil {
			companyUUID = *delegateDetails.CompanyUUID
		} else {
			return gentypes.Delegate{}, &errors.ErrCompanyNotFound
		}
	} else {
		companyUUID = g.Claims.Company
	}

	// Check if company exists
	if !g.CompanyExists(companyUUID) {
		return gentypes.Delegate{}, &errors.ErrCompanyNotFound
	}

	// Get company
	comp, err := g.GetCompanyByUUID(companyUUID)
	if err != nil {
		return gentypes.Delegate{}, err
	}

	// Create a transaction to ensure that a new TTC_ID isn't created before
	// we insert ours

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ttcId, err := generateTTCID(tx, g, comp.Name, delegateDetails.FirstName, delegateDetails.LastName)
	if err != nil {
		return gentypes.Delegate{}, err
	}

	delegate := models.Delegate{
		User: models.User{
			FirstName: delegateDetails.FirstName,
			LastName:  delegateDetails.LastName,
			JobTitle:  delegateDetails.JobTitle,
			Telephone: delegateDetails.Telephone,
			Password:  delegateDetails.Password,
		},
		Email:       delegateDetails.Email,
		CompanyUUID: companyUUID,
		TtcId:       ttcId,
	}
	createErr := database.GormDB.Create(&delegate).Error
	if createErr != nil {
		g.Logger.Log(sentry.LevelError, createErr, "Unable to create delegate")
		return gentypes.Delegate{}, &errors.ErrWhileHandling
	}

	return g.delegateToGentype(delegate), nil
}
