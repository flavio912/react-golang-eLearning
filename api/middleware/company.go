package middleware

import (
	"context"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

//companyToGentype converts a company model to gentype.
func (g *Grant) companyToGentype(company models.Company) gentypes.Company {
	if g.ManagesCompany(company.UUID) {
		createdAt := company.CreatedAt.Format(time.RFC3339)
		return gentypes.Company{
			CreatedAt: &createdAt,
			Approved:  &company.Approved,
			UUID:      company.UUID,
			Name:      company.Name,
			AddressID: company.AddressID,
		}
	}

	if g.IsCompanyDelegate(company.UUID) {
		return gentypes.Company{
			UUID: company.UUID,
			Name: company.Name,
		}
	}

	return gentypes.Company{}
}

func (g *Grant) companiesToGentype(companies []models.Company) []gentypes.Company {
	var genCompanies []gentypes.Company
	for _, comp := range companies {
		genCompanies = append(genCompanies, g.companyToGentype(comp))
	}
	return genCompanies
}

// CompanyExists checks is a companyUUID exists in the DB
func (g *Grant) CompanyExists(companyUUID gentypes.UUID) bool {
	var company models.Company
	query := database.GormDB.Where("uuid = ?", companyUUID).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Error while finding company: %s", companyUUID)
		return false
	}
	return true
}

//IsCompanyDelegate returns true if the grant user is a delegate of the given company uuid
func (g *Grant) IsCompanyDelegate(companyUUID gentypes.UUID) bool {
	return g.IsDelegate && g.Claims.Company == companyUUID
}

// ManagesCompany is an access-control helper to work out if the current grant
// is authorized to manage the given company uuid.
func (g *Grant) ManagesCompany(uuid gentypes.UUID) bool {
	return g.IsAdmin || g.IsManager && g.Claims.Company == uuid
}

func (g *Grant) GetCompaniesByUUID(uuids []gentypes.UUID) ([]gentypes.Company, error) {
	// Check that all requested uuid's are allowed to be returned to the user
	var authorizedUUIDs []gentypes.UUID

	if g.IsManager {
		for _, uuid := range uuids {
			if g.ManagesCompany(uuid) {
				authorizedUUIDs = append(authorizedUUIDs, uuid)
			}
			if len(uuids) > 0 && len(authorizedUUIDs) == 0 {
				return []gentypes.Company{}, &errors.ErrUnauthorized
			}
		}
	} else if g.IsAdmin {
		authorizedUUIDs = uuids
	} else {
		return []gentypes.Company{}, &errors.ErrUnauthorized
	}

	var companies []models.Company
	query := database.GormDB.Where("uuid IN (?)", authorizedUUIDs).Find(&companies)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.Company{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Error finding companies by uuid")
		return []gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.companiesToGentype(companies), nil
}

func (g *Grant) GetCompanyByUUID(uuid gentypes.UUID) (gentypes.Company, error) {
	if !g.ManagesCompany(uuid) {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	var company models.Company
	query := database.GormDB.Where("uuid = ?", uuid).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Company{}, &errors.ErrCompanyNotFound
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Error finding company by uuid: %s", uuid)
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.companyToGentype(company), nil
}

// GetManagerIDsByCompany returns the uuids for the managers of a company
func (g *Grant) GetManagerIDsByCompany(
	companyUUID gentypes.UUID,
	page *gentypes.Page,
	filter *gentypes.ManagersFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.UUID, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var (
		managerUUIDs []gentypes.UUID
		managers     []models.Manager
	)

	query := database.GormDB.Select("uuid").Where("company_uuid = ?", companyUUID)
	query = filterManager(query, filter)

	var count int32
	countQuery := query.Model(&models.Manager{}).Count(&count)
	if countQuery.Error != nil {
		g.Logger.Logf(sentry.LevelError, countQuery.Error, "Error trying to count managers for company uuid: %s", companyUUID)
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "first_name", "last_name"}, "created_at DESC")
	if orderErr != nil {
		return []gentypes.UUID{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query.Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrCompanyNotFound
		}

		g.Logger.Logf(sentry.LevelError, countQuery.Error, "Error getting managers for company uuid: %s", companyUUID)
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

func (g *Grant) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]gentypes.UUID, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var companies []models.Company

	query := database.GormDB.Select("uuid").Model(&models.Company{})

	query, err := getOrdering(query, orderBy, []string{"created_at", "name"}, "created_at DESC")
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
		g.Logger.Log(sentry.LevelError, countErr, "Error getting company uuids")
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}
	query, limit, offset := getPage(query, page)

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
func (g *Grant) CreateCompany(company gentypes.CreateCompanyInput) (gentypes.Company, error) {
	if !g.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	// Validate input
	if err := company.Validate(); err != nil {
		return gentypes.Company{}, err
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
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to create company")
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.companyToGentype(compModel), nil
}

func (g *Grant) UpdateCompany(input gentypes.UpdateCompanyInput) (gentypes.Company, error) {
	if !g.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	var company models.Company
	query := database.GormDB.Preload("Address").Where("uuid = ?", input.UUID).First(&company)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Company{}, &errors.ErrCompanyNotFound
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find company to update with UUID: %s", input.UUID)
		return gentypes.Company{}, &errors.ErrWhileHandling
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
		g.Logger.Logf(sentry.LevelError, save.Error, "Unable to find company to update with UUID: %s", input.UUID)
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.companyToGentype(company), nil

}

// CreateCompanyRequest creates a company and sets it to unapproved, for an admin to approve later
func CreateCompanyRequest(ctx context.Context, company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) error {
	// Validate input
	if err := company.Validate(); err != nil {
		return err
	}
	if err := manager.Validate(); err != nil {
		return err
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
		Approved: false,
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
		logger := logging.GetLoggerFromCtx(ctx)
		logger.Log(sentry.LevelError, query.Error, "Unable to create company request")
		return &errors.ErrWhileHandling
	}

	return nil
}

// ApproveCompany sets a company's status to approved so they can access the manager
// dashboard etc
func (g *Grant) ApproveCompany(companyUUID gentypes.UUID) (gentypes.Company, error) {
	if !g.IsAdmin {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	if !g.CompanyExists(companyUUID) {
		return gentypes.Company{}, &errors.ErrCompanyNotFound
	}

	query := database.GormDB.Model(&models.Company{}).Where("uuid = ?", companyUUID).Update("approved", true)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to approve company")
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.GetCompanyByUUID(companyUUID)
}
