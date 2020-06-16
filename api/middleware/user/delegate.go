package user

import (
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/getsentry/sentry-go"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type UsersRepository interface {
	Delegate(uuid gentypes.UUID) (models.Delegate, error)
	GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy, companyUUID *gentypes.UUID) ([]models.Delegate, gentypes.PageInfo, error)
	CreateDelegate(
		delegateDetails gentypes.CreateDelegateInput,
		s3UploadKey *string,
		password *string,
		company models.Company,
		beforeCommit *func(delegate models.Delegate) bool,
	) (models.Delegate, error)

	Company(uuid gentypes.UUID) (models.Company, error)
	GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error)
	GetCompaniesByUUID(uuids []gentypes.UUID) ([]models.Company, error)
	CompanyExists(companyUUID gentypes.UUID) bool
	CreateCompany(company gentypes.CreateCompanyInput) (models.Company, error)
	UpdateCompany(input gentypes.UpdateCompanyInput) (models.Company, error)
	CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) error
	ApproveCompany(companyUUID gentypes.UUID) (models.Company, error)

	CreateManager(managerDetails gentypes.CreateManagerInput, companyUUID gentypes.UUID) (models.Manager, error)

	Individual(uuid gentypes.UUID) (models.Individual, error)
	CreateIndividual(input gentypes.CreateIndividualInput) (models.Individual, error)

	Manager(UUID gentypes.UUID) (models.Manager, error)
	GetManagersByUUID(uuids []gentypes.UUID) ([]models.Manager, error)
	GetManagerIDsByCompany(
		companyUUID gentypes.UUID,
		page *gentypes.Page,
		filter *gentypes.ManagersFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.UUID, gentypes.PageInfo, error)
	UpdateManager(input gentypes.UpdateManagerInput) (models.Manager, error)
	DeleteManager(uuid gentypes.UUID) (bool, error)
	GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]models.Manager, gentypes.PageInfo, error)
	UpdateManagerProfileKey(managerUUID gentypes.UUID, newKey *string) error

	GetAddressesByIDs(ids []uint) ([]models.Address, error)
}

type usersRepoImpl struct {
	Logger *logging.Logger
}

func NewUsersRepository(logger *logging.Logger) UsersRepository {
	return &usersRepoImpl{
		Logger: logger,
	}
}

func (u *usersRepoImpl) Delegate(uuid gentypes.UUID) (models.Delegate, error) {
	var delegate models.Delegate
	err := database.GormDB.Where("uuid = ?", uuid).Find(&delegate).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.Delegate{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, err, "Unable to find delegate")
		return models.Delegate{}, &errors.ErrWhileHandling
	}

	return delegate, nil
}

func filterDelegate(query *gorm.DB, filter *gentypes.DelegatesFilter) *gorm.DB {
	if filter != nil {
		query = middleware.FilterUser(query, &filter.UserFilter)

		if filter.TTC_ID != nil && *filter.TTC_ID != "" {
			query = query.Where("ttc_id ILIKE ?", "%%"+*filter.TTC_ID+"%%")
		}

		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
	}

	return query
}

func (u *usersRepoImpl) GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy, companyUUID *gentypes.UUID) ([]models.Delegate, gentypes.PageInfo, error) {

	var delegates []models.Delegate

	query := filterDelegate(database.GormDB, filter)

	// only get certain company's delegates
	if companyUUID != nil {
		query = query.Where("company_uuid = ?", *companyUUID)
	}

	// Count the total filtered dataset
	var count int32
	countErr := query.Model(&models.Delegate{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		u.Logger.Log(sentry.LevelError, countErr, "Unable to count delegates")
		return []models.Delegate{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title", "ttc_id"}, "created_at DESC")
	if orderErr != nil {
		return []models.Delegate{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&delegates)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Delegate{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to find delegates")
		return []models.Delegate{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return delegates, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(delegates)),
	}, nil
}

func (u *usersRepoImpl) generateTTCID(tx *gorm.DB, companyName string, delegateFName string, delegateLName string) (string, error) {
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
			u.Logger.Log(sentry.LevelError, err.Error, "TTC_ID find error")
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

	u.Logger.LogMessage(sentry.LevelError, "Iteration exceeded max")
	return "", &errors.ErrWhileHandling
}

func (u *usersRepoImpl) CreateDelegate(
	delegateDetails gentypes.CreateDelegateInput,
	s3UploadKey *string,
	password *string,
	company models.Company,
	beforeCommit *func(delegate models.Delegate) bool,
) (models.Delegate, error) {
	// Create a transaction to ensure that a new TTC_ID isn't created before we insert ours
	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ttcId, err := u.generateTTCID(tx, company.Name, delegateDetails.FirstName, delegateDetails.LastName)
	if err != nil {
		tx.Rollback()
		return models.Delegate{}, err
	}

	// Add link manually because gorm doesn't like blank associations
	var courseTaker = models.CourseTaker{}
	if err := tx.Create(&courseTaker).Error; err != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, err, "Unable to create courseTaker")
		return models.Delegate{}, &errors.ErrWhileHandling
	}

	delegate := models.Delegate{
		FirstName:     delegateDetails.FirstName,
		LastName:      delegateDetails.LastName,
		JobTitle:      delegateDetails.JobTitle,
		Telephone:     delegateDetails.Telephone,
		Password:      password,
		Email:         delegateDetails.Email,
		CompanyUUID:   company.UUID,
		TtcId:         ttcId,
		ProfileKey:    s3UploadKey,
		CourseTakerID: courseTaker.ID,
	}
	createErr := tx.Create(&delegate).Error
	if createErr != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, createErr, "Unable to create delegate")
		return models.Delegate{}, &errors.ErrWhileHandling
	}

	if beforeCommit != nil {
		success := (*beforeCommit)(delegate)
		if !success {
			tx.Rollback()
			return models.Delegate{}, &errors.ErrWhileHandling
		}
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "Error commiting create delegate transaction")
		return models.Delegate{}, &errors.ErrWhileHandling
	}

	return delegate, nil
}
