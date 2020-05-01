package middleware

import (
	"time"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

//companyToGentype converts a company model to gentype.
func (g *Grant) companyToGentype(company models.Company) gentypes.Company {
	if g.ManagesCompany(gentypes.UUID{UUID: company.UUID}) {
		createdAt := company.CreatedAt.Format(time.RFC3339)
		return gentypes.Company{
			CreatedAt: &createdAt,
			Approved:  &company.Approved,
			UUID:      gentypes.UUID{UUID: company.UUID},
			Name:      company.Name,
			AddressID: company.AddressID,
		}
	}

	if g.IsCompanyDelegate(gentypes.UUID{UUID: company.UUID}) {
		return gentypes.Company{
			UUID: gentypes.UUID{UUID: company.UUID},
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
	existsErr := database.GormDB.Where("uuid = ?", companyUUID).First(&company)
	if existsErr.Error != nil {
		if existsErr.RecordNotFound() {
			return false
		}
		glog.Errorf("Error while finding company: %s", existsErr.Error.Error())
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
	for _, uuid := range uuids {
		if g.ManagesCompany(uuid) {
			authorizedUUIDs = append(authorizedUUIDs, uuid)
		}
	}

	var companies []models.Company
	query := database.GormDB.Where("uuid IN (?)", authorizedUUIDs).Find(&companies)
	if query.Error != nil {
		return []gentypes.Company{}, getDBErrorType(query)
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

	query := database.GormDB.Select("uuid").Where("company_id = ?", companyUUID)
	query = filterManager(query, filter)

	var count int32
	err := query.Model(&models.Manager{}).Count(&count)
	if err.Error != nil {
		glog.Errorf("DB Error %s", err.Error.Error())
		glog.Errorf("Unable to count records for %s", companyUUID)
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "first_name", "last_name"}, "created_at DESC")
	if orderErr != nil {
		return []gentypes.UUID{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query.Find(&managers)
	if query.Error != nil {
		return []gentypes.UUID{}, gentypes.PageInfo{}, getDBErrorType(query)
	}

	for _, manager := range managers {
		managerUUIDs = append(managerUUIDs, gentypes.UUID{UUID: manager.UUID})
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
		glog.Errorf("DB Error %s", countErr.Error())
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}
	query, limit, offset := getPage(query, page)

	query.Find(&companies)

	var uuids = make([]gentypes.UUID, len(companies))
	for i, comp := range companies {
		uuids[i] = gentypes.UUID{UUID: comp.UUID}
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
		glog.Errorf("Unable to create company: %s", query.Error.Error())
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

		glog.Errorf("Unable to find company to update with UUID: %s - error: %s", input.UUID, query.Error.Error())
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
		glog.Errorf("Error updating company with UUID: %s - error: %s", input.UUID, save.Error.Error())
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.companyToGentype(company), nil

}

// CreateCompanyRequest creates a company and sets it to unapproved, for an admin to approve later
func CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateManagerInput) error {
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
				User: models.User{
					FirstName: manager.FirstName,
					LastName:  manager.LastName,
					JobTitle:  manager.JobTitle,
					Telephone: manager.Telephone,
					Email:     manager.Email,
					Password:  manager.Password,
					LastLogin: time.Now(),
				},
			}},
	}
	query := database.GormDB.Create(&compModel)
	if query.Error != nil {
		glog.Errorf("Unable to create company request: %s", query.Error.Error())
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
		glog.Errorf("Unable to approve company: %s", query.Error.Error())
		return gentypes.Company{}, &errors.ErrWhileHandling
	}

	return g.GetCompanyByUUID(companyUUID)
}
