package users

import (
	"time"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/email"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
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
	GetManagersByUUID(uuids []gentypes.UUID) ([]gentypes.Manager, error)
	GetManagerIDsByCompany(
		companyUUID gentypes.UUID,
		page *gentypes.Page,
		filter *gentypes.ManagersFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.UUID, gentypes.PageInfo, error)
	CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) (bool, error)
	ApproveCompany(companyUUID gentypes.UUID) (gentypes.Company, error)

	CreateManager(managerDetails gentypes.CreateManagerInput) (gentypes.Manager, error)
	UpdateManager(input gentypes.UpdateManagerInput) (gentypes.Manager, error)
	DeleteManager(uuid gentypes.UUID) (bool, error)
	GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]gentypes.Manager, gentypes.PageInfo, error)

	CreateIndividual(input gentypes.CreateIndividualInput) (gentypes.User, error)

	ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	ManagerProfileUploadSuccess(token string) error

	GetCurrentUser() (gentypes.User, error)
	GetAddressesByIDs(ids []uint) ([]gentypes.Address, error)
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

func (u *usersAppImpl) delegateToGentype(delegate models.Delegate) gentypes.Delegate {
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

func (u *usersAppImpl) delegatesToGentype(delegates []models.Delegate) []gentypes.Delegate {
	var genDelegates []gentypes.Delegate
	for _, delegate := range delegates {
		genDelegates = append(genDelegates, u.delegateToGentype(delegate))
	}

	return genDelegates
}

func (u *usersAppImpl) CreateIndividual(input gentypes.CreateIndividualInput) (gentypes.User, error) {
	individual, err := u.usersRepository.CreateIndividual(input)
	user := u.IndividualToUser(individual)
	return user, err
}

func (u *usersAppImpl) GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy) ([]gentypes.Delegate, gentypes.PageInfo, error) {
	if u.grant.IsAdmin {
		delegates, pageInfo, err := u.usersRepository.GetDelegates(page, filter, orderBy, nil)
		return u.delegatesToGentype(delegates), pageInfo, err
	}

	if u.grant.IsManager {
		delegates, pageInfo, err := u.usersRepository.GetDelegates(page, filter, orderBy, &u.grant.Claims.Company)
		return u.delegatesToGentype(delegates), pageInfo, err
	}

	return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
}

func (u *usersAppImpl) CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, *string, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager {
		return gentypes.Delegate{}, nil, &errors.ErrUnauthorized
	}

	if err := delegateDetails.Validate(); err != nil {
		return gentypes.Delegate{}, nil, err
	}

	var (
		needsGeneratePass = delegateDetails.GeneratePassword != nil && *delegateDetails.GeneratePassword
		companyUUID       gentypes.UUID
		s3UploadKey       *string
		realPass          *string
		password          string
	)

	if u.grant.IsManager {
		companyUUID = u.grant.Claims.Company
	}
	if u.grant.IsAdmin {
		if delegateDetails.CompanyUUID == nil {
			return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
		}
		companyUUID = *delegateDetails.CompanyUUID
	}

	// Check if company exists
	if !u.usersRepository.CompanyExists(companyUUID) {
		return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
	}

	// Check if autogenerating password is required
	if needsGeneratePass {
		pass, err := auth.GenerateSecurePassword(10)
		if err != nil {
			u.grant.Logger.Log(sentry.LevelError, err, "Unable to generate secure password")
			return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
		}
		password = pass
		realPass = &pass
	}

	// Check if upload token is valid
	if delegateDetails.ProfileImageUploadToken != nil {
		tmpUploadKey, err := uploads.VerifyUploadSuccess(*delegateDetails.ProfileImageUploadToken, "profileImage")
		if err != nil {
			return gentypes.Delegate{}, nil, &errors.ErrUploadTokenInvalid
		}

		s3UploadKey = &tmpUploadKey
	}

	comp, err := u.usersRepository.Company(companyUUID)
	if err != nil {
		return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
	}

	sendEmails := func(delegate models.Delegate) bool {
		// Send transactional email
		// If not generated password, send an email to the user
		if !needsGeneratePass {
			token, err := auth.GenerateFinaliseDelegateToken(auth.FinaliseDelegateClaims{
				UUID: delegate.UUID,
			})
			if err != nil {
				u.grant.Logger.Log(sentry.LevelError, err, "Unable to generate finalise delegate token")
				return false
			}

			if delegate.Email == nil {
				u.grant.Logger.LogMessage(sentry.LevelError, "Delegate email is nil")
				return false
			}

			err = email.SendFinaliseAccountEmail(token, delegate.FirstName, *delegate.Email)
			if err != nil {
				u.grant.Logger.Log(sentry.LevelWarning, err, "Unable to send finalise account email")
				return false
			}
		}

		return true
	}
	delegate, err := u.usersRepository.CreateDelegate(
		delegateDetails,
		s3UploadKey,
		&password,
		comp,
		&sendEmails,
	)

	return u.delegateToGentype(delegate), realPass, err
}

// companyToGentype converts a company model to gentype.
func (u *usersAppImpl) companyToGentype(company models.Company) gentypes.Company {
	if u.grant.ManagesCompany(company.UUID) {
		createdAt := company.CreatedAt.Format(time.RFC3339)
		return gentypes.Company{
			CreatedAt: &createdAt,
			Approved:  &company.Approved,
			UUID:      company.UUID,
			Name:      company.Name,
			AddressID: company.AddressID,
		}
	}

	if u.grant.IsCompanyDelegate(company.UUID) {
		return gentypes.Company{
			UUID: company.UUID,
			Name: company.Name,
		}
	}

	return gentypes.Company{}
}

func (u *usersAppImpl) companiesToGentype(companies []models.Company) []gentypes.Company {
	var genCompanies []gentypes.Company
	for _, comp := range companies {
		genCompanies = append(genCompanies, u.companyToGentype(comp))
	}
	return genCompanies
}

func (u *usersAppImpl) Company(uuid gentypes.UUID) (gentypes.Company, error) {
	if !u.grant.IsAdmin && !(u.grant.IsManager && u.grant.Claims.Company == uuid) {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	company, err := u.usersRepository.Company(uuid)
	return u.companyToGentype(company), err
}

func (u *usersAppImpl) GetCompaniesByUUID(uuids []gentypes.UUID) ([]gentypes.Company, error) {
	if u.grant.IsAdmin {
		companies, err := u.usersRepository.GetCompaniesByUUID(uuids)
		return u.companiesToGentype(companies), err
	}

	if u.grant.IsManager {
		if len(uuids) == 1 && uuids[0] == u.grant.Claims.Company {
			companies, err := u.usersRepository.GetCompaniesByUUID(uuids)
			return u.companiesToGentype(companies), err
		}
	}

	return []gentypes.Company{}, &errors.ErrUnauthorized
}

func (u *usersAppImpl) GetManagersByUUID(uuids []gentypes.UUID) ([]gentypes.Manager, error) {
	// Manager can get own uuid, admin can get any
	if !(len(uuids) == 1 && u.grant.Claims.UUID == uuids[0]) && !u.grant.IsAdmin {
		return []gentypes.Manager{}, &errors.ErrUnauthorized
	}

	managers, err := u.usersRepository.GetManagersByUUID(uuids)
	return u.managersToGentype(managers), err
}

func (u *usersAppImpl) GetManagerIDsByCompany(
	companyUUID gentypes.UUID,
	page *gentypes.Page,
	filter *gentypes.ManagersFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.UUID, gentypes.PageInfo, error) {
	if !u.grant.IsAdmin {
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	return u.usersRepository.GetManagerIDsByCompany(companyUUID, page, filter, orderBy)
}

func (u *usersAppImpl) Delegate(uuid gentypes.UUID) (gentypes.Delegate, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager && !u.grant.IsDelegate {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	delegate, err := u.usersRepository.Delegate(uuid)

	if !u.grant.IsAdmin &&
		!(u.grant.IsManager && u.grant.Claims.Company == delegate.CompanyUUID) &&
		!(u.grant.IsDelegate && u.grant.Claims.UUID == delegate.UUID) {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	return u.delegateToGentype(delegate), err
}

func (u *usersAppImpl) CreateManager(managerDetails gentypes.CreateManagerInput) (gentypes.Manager, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	var compUUID gentypes.UUID
	// If you're an admin you need to provide the company UUID
	if u.grant.IsAdmin {
		if managerDetails.CompanyUUID != nil {
			compUUID = *managerDetails.CompanyUUID
		} else {
			return gentypes.Manager{}, &errors.ErrCompanyNotFound
		}
	}

	// If you're a manager the company UUID will be selected from the one in your JWT claims
	if u.grant.IsManager {
		compUUID = u.grant.Claims.Company
	}

	manager, err := u.usersRepository.CreateManager(managerDetails, compUUID)
	return u.managerToGentype(manager), err
}

func (u *usersAppImpl) UpdateManager(input gentypes.UpdateManagerInput) (gentypes.Manager, error) {
	if !(u.grant.IsManager && u.grant.Claims.UUID == input.UUID) && !u.grant.IsAdmin {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	manager, err := u.usersRepository.UpdateManager(input)
	return u.managerToGentype(manager), err
}

func (u *usersAppImpl) DeleteManager(uuid gentypes.UUID) (bool, error) {
	// managers can delete themselves
	// admins can delete any manager
	if !(u.grant.IsManager && u.grant.Claims.UUID == uuid) && !u.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	// TODO: delete profile image from S3

	return u.usersRepository.DeleteManager(uuid)
}

func (u *usersAppImpl) CreateCompany(company gentypes.CreateCompanyInput) (gentypes.Company, error) {
	if !u.grant.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	comp, err := u.usersRepository.CreateCompany(company)
	return u.companyToGentype(comp), err
}

func (u *usersAppImpl) UpdateCompany(company gentypes.UpdateCompanyInput) (gentypes.Company, error) {
	if !u.grant.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	comp, err := u.usersRepository.UpdateCompany(company)
	return u.companyToGentype(comp), err
}

func (u *usersAppImpl) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error) {
	// TODO add back auth
	return u.usersRepository.GetCompanyUUIDs(page, filter, orderBy)
}

func (u *usersAppImpl) GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]gentypes.Manager, gentypes.PageInfo, error) {
	if !u.grant.IsAdmin {
		return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	managers, pageInfo, err := u.usersRepository.GetManagers(page, filter, orderBy)
	return u.managersToGentype(managers), pageInfo, err
}

func (u *usersAppImpl) CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) (bool, error) {
	// Public

	if err := company.Validate(); err != nil {
		return false, err
	}
	if err := manager.Validate(); err != nil {
		return false, err
	}

	err := u.usersRepository.CreateCompanyRequest(company, manager)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *usersAppImpl) ApproveCompany(companyUUID gentypes.UUID) (gentypes.Company, error) {
	if !u.grant.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	company, err := u.usersRepository.ApproveCompany(companyUUID)
	return u.companyToGentype(company), err
}

func addressToGentype(address models.Address) gentypes.Address {
	return gentypes.Address{
		ID:           address.ID,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		County:       address.County,
		PostCode:     address.PostCode,
		Country:      address.Country,
	}
}

func addressesToGentypes(addresses []models.Address) []gentypes.Address {
	var genAddresses []gentypes.Address
	for _, address := range addresses {
		genAddresses = append(genAddresses, addressToGentype(address))
	}
	return genAddresses
}

func (u *usersAppImpl) GetAddressesByIDs(ids []uint) ([]gentypes.Address, error) {
	if !u.grant.IsAdmin {
		return []gentypes.Address{}, &errors.ErrUnauthorized
	}

	addresses, err := u.usersRepository.GetAddressesByIDs(ids)
	return addressesToGentypes(addresses), err
}
