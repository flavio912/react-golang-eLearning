package users

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

// companyToGentype converts a company model to gentype.
func (u *usersAppImpl) companyToGentype(company models.Company) gentypes.Company {
	var url *string
	if company.LogoKey != nil {
		temp := uploads.GetImgixURL(*company.LogoKey)
		url = &temp
	}

	if u.grant.ManagesCompany(company.UUID) {
		createdAt := company.CreatedAt.Format(time.RFC3339)
		return gentypes.Company{
			CreatedAt:    &createdAt,
			Approved:     &company.Approved,
			UUID:         company.UUID,
			Name:         company.Name,
			AddressID:    company.AddressID,
			IsContract:   company.IsContract,
			ContactEmail: company.ContactEmail,
			ContactPhone: company.ContactPhone,
			LogoURL:      url,
		}
	}

	if u.grant.IsCompanyDelegate(company.UUID) {
		return gentypes.Company{
			UUID:    company.UUID,
			Name:    company.Name,
			LogoURL: url,
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

func (u *usersAppImpl) CreateCompany(company gentypes.CreateCompanyInput) (gentypes.Company, error) {
	if !u.grant.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	var logoKey *string
	if company.LogoToken != nil {
		key, err := uploads.VerifyUploadSuccess(*logoKey, "profileImage")
		if err != nil {
			return gentypes.Company{}, err
		}
		logoKey = &key
	}

	comp, err := u.usersRepository.CreateCompany(company, logoKey)
	return u.companyToGentype(comp), err
}

func (u *usersAppImpl) UpdateCompany(input gentypes.UpdateCompanyInput) (gentypes.Company, error) {
	if !u.grant.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	var logoKey *string

	company, err := u.usersRepository.Company(input.UUID)
	if err != nil {
		return gentypes.Company{}, err
	}

	if input.LogoToken != nil {
		if company.LogoKey != nil {
			err := uploads.DeleteImageFromKey(*company.LogoKey)
			if err != nil {
				return gentypes.Company{}, err
			}
		}
		key, err := uploads.VerifyUploadSuccess(*input.LogoToken, "profileImage")
		if err != nil {
			return gentypes.Company{}, err
		}

		logoKey = &key
	}

	company, err = u.usersRepository.UpdateCompany(input, logoKey)
	return u.companyToGentype(company), err
}

func (u *usersAppImpl) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error) {
	if !u.grant.IsAdmin {
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

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
