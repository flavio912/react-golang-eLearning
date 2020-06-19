package users

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"
)

type UsersApp interface {
	Delegate(uuid gentypes.UUID) (gentypes.Delegate, error)
	GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy) ([]gentypes.Delegate, gentypes.PageInfo, error)
	CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, *string, error)

	Company(uuid gentypes.UUID) (gentypes.Company, error)
	GetCompaniesByUUID(uuids []gentypes.UUID) ([]gentypes.Company, error)
	CreateCompany(company gentypes.CreateCompanyInput) (gentypes.Company, error)
	UpdateCompany(company gentypes.UpdateCompanyInput) (gentypes.Company, error)
	GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error)
	CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) (bool, error)
	ApproveCompany(companyUUID gentypes.UUID) (gentypes.Company, error)

	CreateManager(managerDetails gentypes.CreateManagerInput) (gentypes.Manager, error)
	UpdateManager(input gentypes.UpdateManagerInput) (gentypes.Manager, error)
	DeleteManager(uuid gentypes.UUID) (bool, error)
	GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]gentypes.Manager, gentypes.PageInfo, error)
	GetManagersByUUID(uuids []gentypes.UUID) ([]gentypes.Manager, error)
	GetManagerIDsByCompany(
		companyUUID gentypes.UUID,
		page *gentypes.Page,
		filter *gentypes.ManagersFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.UUID, gentypes.PageInfo, error)

	CreateIndividual(input gentypes.CreateIndividualInput) (gentypes.User, error)

	ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	ManagerProfileUploadSuccess(token string) error

	GetCurrentUser() (gentypes.User, error)
	GetAddressesByIDs(ids []uint) ([]gentypes.Address, error)

	TakerActivity(courseTakerUUID gentypes.UUID, page *gentypes.Page) ([]gentypes.Activity, gentypes.PageInfo, error)
}

type usersAppImpl struct {
	grant           *middleware.Grant
	usersRepository user.UsersRepository
}

func NewUsersApp(grant *middleware.Grant) UsersApp {
	return &usersAppImpl{
		grant:           grant,
		usersRepository: user.NewUsersRepository(&grant.Logger),
	}
}
