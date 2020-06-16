package user

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersRepoImpl) Company(uuid gentypes.UUID) (models.Company, error) {
	// if !g.ManagesCompany(uuid) {
	// 	return models.Company{}, &errors.ErrUnauthorized
	// }

	var company models.Company
	query := database.GormDB.Where("uuid = ?", uuid).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Company{}, &errors.ErrCompanyNotFound
		}

		u.Logger.Logf(sentry.LevelError, query.Error, "Error finding company by uuid: %s", uuid)
		return models.Company{}, &errors.ErrWhileHandling
	}

	return company, nil
}

// companyExists checks if a companyUUID exists in the DB
func (u *usersRepoImpl) CompanyExists(companyUUID gentypes.UUID) bool {
	var company models.Company
	query := database.GormDB.Where("uuid = ?", companyUUID).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		u.Logger.Logf(sentry.LevelError, query.Error, "Error while finding company: %s", companyUUID)
		return false
	}
	return true
}

func (u *usersRepoImpl) GetCompaniesByUUID(uuids []gentypes.UUID) ([]models.Company, error) {
	// Check that all requested uuid's are allowed to be returned to the user
	// var authorizedUUIDs []gentypes.UUID

	// if g.IsManager {
	// 	for _, uuid := range uuids {
	// 		if g.ManagesCompany(uuid) {
	// 			authorizedUUIDs = append(authorizedUUIDs, uuid)
	// 		}
	// 		if len(uuids) > 0 && len(authorizedUUIDs) == 0 {
	// 			return []models.Company{}, &errors.ErrUnauthorized
	// 		}
	// 	}
	// } else if g.IsAdmin {
	// 	authorizedUUIDs = uuids
	// } else {
	// 	return []models.Company{}, &errors.ErrUnauthorized
	// }

	var companies []models.Company
	query := database.GormDB.Where("uuid IN (?)", uuids).Find(&companies)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Company{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Error finding companies by uuid")
		return []models.Company{}, &errors.ErrWhileHandling
	}

	return companies, nil
}

// GetManagerIDsByCompany returns the uuids for the managers of a company
func (u *usersRepoImpl) GetManagerIDsByCompany(
	companyUUID gentypes.UUID,
	page *gentypes.Page,
	filter *gentypes.ManagersFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.UUID, gentypes.PageInfo, error) {
	// if !g.IsAdmin {
	// 	return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	// }

	var (
		managerUUIDs []gentypes.UUID
		managers     []models.Manager
	)

	query := database.GormDB.Select("uuid").Where("company_uuid = ?", companyUUID)
	query = filterManager(query, filter)

	var count int32
	countQuery := query.Model(&models.Manager{}).Count(&count)
	if countQuery.Error != nil {
		u.Logger.Logf(sentry.LevelError, countQuery.Error, "Error trying to count managers for company uuid: %s", companyUUID)
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"created_at", "first_name", "last_name"}, "created_at DESC")
	if orderErr != nil {
		return []gentypes.UUID{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query.Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrCompanyNotFound
		}

		u.Logger.Logf(sentry.LevelError, countQuery.Error, "Error getting managers for company uuid: %s", companyUUID)
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	for _, manager := range managers {
		managerUUIDs = append(managerUUIDs, manager.UUID)
	}

	return managerUUIDs, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(managerUUIDs)),
	}, nil
}

func (u *usersRepoImpl) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error) {
	// if !g.IsAdmin {
	// 	return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	// }

	var companies []models.Company

	query := database.GormDB.Select("uuid").Model(&models.Company{})

	query, err := middleware.GetOrdering(query, orderBy, []string{"created_at", "name"}, "created_at DESC")
	if err != nil {
		return []gentypes.UUID{}, gentypes.PageInfo{}, err
	}

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.UUID != nil && *filter.UUID != "" {
			query = query.Where("uuid = ?", *filter.UUID)
		}
		if filter.Approved != nil {
			query = query.Where("approved = ?", *filter.Approved)
		}
	}

	var count int32
	countErr := query.Count(&count).Error
	if countErr != nil {
		u.Logger.Log(sentry.LevelError, countErr, "Error getting company uuids")
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}
	query, limit, offset := middleware.GetPage(query, page)

	query.Find(&companies)

	var uuids = make([]gentypes.UUID, len(companies))
	for i, comp := range companies {
		uuids[i] = comp.UUID
	}
	return uuids, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(companies)),
	}, nil

}

// CreateCompany is an admin function for creating companys directly
func (u *usersRepoImpl) CreateCompany(company gentypes.CreateCompanyInput) (models.Company, error) {
	// if !g.IsAdmin {
	// 	return gentypes.Company{}, &errors.ErrUnauthorized
	// }

	// Validate input
	if err := company.Validate(); err != nil {
		return models.Company{}, err
	}

	compModel := models.Company{
		Name: company.CompanyName,
		Address: models.Address{
			AddressLine1: company.AddressLine1,
			AddressLine2: company.AddressLine2,
			County:       company.County,
			PostCode:     company.PostCode,
			Country:      company.Country,
		},
		Approved: true,
	}

	query := database.GormDB.Create(&compModel)
	if query.Error != nil {
		u.Logger.Log(sentry.LevelError, query.Error, "Unable to create company")
		return models.Company{}, &errors.ErrWhileHandling
	}

	return compModel, nil
}

func (u *usersRepoImpl) UpdateCompany(input gentypes.UpdateCompanyInput) (models.Company, error) {
	// if !g.IsAdmin {
	// 	return gentypes.Company{}, &errors.ErrUnauthorized
	// }

	var company models.Company
	query := database.GormDB.Preload("Address").Where("uuid = ?", input.UUID).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Company{}, &errors.ErrCompanyNotFound
		}

		u.Logger.Logf(sentry.LevelError, query.Error, "Unable to find company to update with UUID: %s", input.UUID)
		return models.Company{}, &errors.ErrWhileHandling
	}

	if input.CompanyName != nil {
		company.Name = *input.CompanyName
	}
	if input.Approved != nil {
		company.Approved = *input.Approved
	}
	if input.AddressLine1 != nil {
		company.Address.AddressLine1 = *input.AddressLine1
	}
	if input.AddressLine2 != nil {
		company.Address.AddressLine2 = *input.AddressLine2
	}
	if input.PostCode != nil {
		company.Address.PostCode = *input.PostCode
	}
	if input.County != nil {
		company.Address.County = *input.County
	}
	if input.Country != nil {
		company.Address.Country = *input.Country
	}

	save := database.GormDB.Save(&company)
	if save.Error != nil {
		u.Logger.Logf(sentry.LevelError, save.Error, "Unable to find company to update with UUID: %s", input.UUID)
		return models.Company{}, &errors.ErrWhileHandling
	}

	return company, nil

}

// CreateCompanyRequest creates a company and sets it to unapproved, for an admin to approve later
func (u *usersRepoImpl) CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) error {

	compModel := models.Company{
		Name: company.CompanyName,
		Address: models.Address{
			AddressLine1: company.AddressLine1,
			AddressLine2: company.AddressLine2,
			County:       company.County,
			PostCode:     company.PostCode,
			Country:      company.Country,
		},
		Approved:   false,
		IsContract: false,
		Managers: []models.Manager{
			models.Manager{
				FirstName: manager.FirstName,
				LastName:  manager.LastName,
				JobTitle:  manager.JobTitle,
				Telephone: manager.Telephone,
				Password:  manager.Password,
				LastLogin: time.Now(),
				Email:     manager.Email,
			}},
	}
	query := database.GormDB.Create(&compModel)
	if query.Error != nil {
		u.Logger.Log(sentry.LevelError, query.Error, "Unable to create company request")
		return &errors.ErrWhileHandling
	}

	return nil
}

// ApproveCompany sets a company's status to approved so they can access the manager
// dashboard etc
func (u *usersRepoImpl) ApproveCompany(companyUUID gentypes.UUID) (models.Company, error) {
	if !u.CompanyExists(companyUUID) {
		return models.Company{}, &errors.ErrCompanyNotFound
	}

	query := database.GormDB.Model(&models.Company{}).Where("uuid = ?", companyUUID).Update("approved", true)
	if query.Error != nil {
		u.Logger.Log(sentry.LevelError, query.Error, "Unable to approve company")
		return models.Company{}, &errors.ErrWhileHandling
	}

	comp, err := u.Company(companyUUID)
	if err != nil {
		return models.Company{}, &errors.ErrWhileHandling
	}

	return comp, nil
}
