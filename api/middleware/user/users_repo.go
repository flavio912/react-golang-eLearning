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
	UpdateDelegate(
		details gentypes.UpdateDelegateInput,
		s3UploadKey *string,
		password *string,
	) (models.Delegate, error)

	Company(uuid gentypes.UUID) (models.Company, error)
	GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error)
	GetCompaniesByUUID(uuids []gentypes.UUID) ([]models.Company, error)
	CompanyExists(companyUUID gentypes.UUID) bool
	CreateCompany(company gentypes.CreateCompanyInput, logoKey *string) (models.Company, error)
	UpdateCompany(input gentypes.UpdateCompanyInput) (models.Company, error)
	CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) error
	ApproveCompany(companyUUID gentypes.UUID) (models.Company, error)

	CreateManager(managerDetails gentypes.CreateManagerInput, companyUUID gentypes.UUID) (models.Manager, error)

	Individual(uuid gentypes.UUID) (models.Individual, error)
	Individuals(page *gentypes.Page, filter *gentypes.IndividualFilter, orderBy *gentypes.OrderBy) ([]models.Individual, gentypes.PageInfo, error)
	CreateIndividual(input gentypes.CreateIndividualInput) (models.Individual, error)
	UpdateIndividual(input gentypes.UpdateIndividualInput) (models.Individual, error)
	DeleteIndividual(uuid gentypes.UUID) (bool, error)

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

	UserFromCourseTaker(takerUUID gentypes.UUID) (*models.Delegate, *models.Individual)
	TakerActivity(courseTaker gentypes.UUID, page *gentypes.Page) ([]models.CourseTakerActivity, gentypes.PageInfo, error)
	TakerActivitys(courseTakers []gentypes.UUID, page *gentypes.Page) ([]models.CourseTakerActivity, gentypes.PageInfo, error)
	CreateTakerActivity(courseTaker gentypes.UUID, activityType gentypes.ActivityType, relatedCourseID *uint) (models.CourseTakerActivity, error)
	DeleteTakerActivity(activityUUID gentypes.UUID) error
	TakerActiveCourse(courseTaker gentypes.UUID, courseID uint) (models.ActiveCourse, error)
	TakerActiveCourses(courseTaker gentypes.UUID) ([]models.ActiveCourse, error)
	TakerHistoricalCourses(courseTaker gentypes.UUID) ([]models.HistoricalCourse, error)
	TakerHasActiveCourse(courseTaker gentypes.UUID, courseID uint) (bool, error)

	SaveTestMarks(mark models.TestMark) error
	TakerTestMarks(courseTaker gentypes.UUID, courseID uint) ([]models.TestMark, error)

	HistoricalCourse(uuid gentypes.UUID) (models.HistoricalCourse, error)
	CreateHistoricalCourse(course models.HistoricalCourse) (models.HistoricalCourse, error)
	UpdateHistoricalCourse(input UpdateHistoricalCourseInput) error

	// Returns true if all course takers are part of the company
	CompanyManagesCourseTakers(companyUUID gentypes.UUID, courseTakerUUIDs []gentypes.UUID) (bool, error)
}

type usersRepoImpl struct {
	Logger *logging.Logger
}

func NewUsersRepository(logger *logging.Logger) UsersRepository {
	return &usersRepoImpl{
		Logger: logger,
	}
}
