package user

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
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

	TakerActivity(courseTaker gentypes.UUID) ([]models.CourseTakerActivity, error)
	TakerActivitys(courseTakers []gentypes.UUID) ([]models.CourseTakerActivity, error)
	CreateTakerActivity(courseTaker gentypes.UUID, activityType gentypes.ActivityType, relatedCourseID *uint) (models.CourseTakerActivity, error)
	DeleteTakerActivity(activityUUID gentypes.UUID) error
}

type usersRepoImpl struct {
	Logger *logging.Logger
}

func NewUsersRepository(logger *logging.Logger) UsersRepository {
	return &usersRepoImpl{
		Logger: logger,
	}
}
