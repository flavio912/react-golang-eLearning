package middleware

import (
	"github.com/golang/glog"
	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

//companyToGentypes converts a company model to gentype.
func companyToGentypes(company models.Company) gentypes.Company {

	return gentypes.Company{
		UUID: company.UUID,
		Name: company.Name,
	}
}

func companiesToGentype(companies []models.Company) []gentypes.Company {
	var genCompanies []gentypes.Company
	for _, comp := range companies {
		genCompanies = append(genCompanies, companyToGentypes(comp))
	}
	return genCompanies
}

// ManagesCompany is an access-control helper to work out if the current grant
// is authorized to manage the given company uuid.
func (g *Grant) ManagesCompany(uuid string) bool {
	if g.IsAdmin {
		return true
	}
	if g.IsManager {
		if uuid != "" && g.Claims.Company == uuid {
			return true
		}
	}
	return false
}

func (g *Grant) GetCompaniesByUUID(uuids []string) ([]gentypes.Company, error) {
	// Check that all requested uuid's are allowed to be returned to the user
	var authorizedUUIDs []string
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

	return companiesToGentype(companies), nil
}

func (g *Grant) GetCompanyByUUID(uuid string) (gentypes.Company, error) {
	if !g.ManagesCompany(uuid) {
		return gentypes.Company{}, &errors.ErrUnauthorized
	}

	var company models.Company
	query := database.GormDB.Where("uuid = ?", uuid).First(&company)
	if query.Error != nil {
		return gentypes.Company{}, getDBErrorType(query)
	}

	return companyToGentypes(company), nil
}

// GetManagerIDsByCompany returns the uuids for the managers of a company
func (g *Grant) GetManagerIDsByCompany(companyUUID string, page *gentypes.Page, filter *gentypes.ManagersFilter) ([]uuid.UUID, gentypes.PageInfo, error) {
	if !g.ManagesCompany(companyUUID) {
		return []uuid.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var (
		managerUUIDs []uuid.UUID
		managers     []models.Manager
	)

	query := database.GormDB.Select("uuid").Where("company_id = ?", companyUUID)
	query, limit, offset := getPage(query, page)
	query.Find(&managers)
	if query.Error != nil {
		return []uuid.UUID{}, gentypes.PageInfo{}, getDBErrorType(query)
	}

	for _, manager := range managers {
		managerUUIDs = append(managerUUIDs, manager.UUID)
	}

	var count int32
	err := query.Model(&models.Manager{}).Count(&count)
	if err.Error != nil {
		glog.Errorf("DB Error %s", err.Error.Error())
		glog.Errorf("Unable to count records for %s", companyUUID)
		return []uuid.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return managerUUIDs, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(managerUUIDs)),
	}, nil
}

func (g *Grant) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter) ([]string, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []string{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var companies []models.Company

	query := database.GormDB.Select("uuid").Model(&models.Company{})
	query, limit, offset := getPage(query, page)

	query.Find(&companies)

	var count int32
	query = query.Model(&models.Manager{}).Count(&count)
	if query.Error != nil {
		glog.Errorf("DB Error %s", query.Error.Error())
		return []string{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	var uuids = make([]string, len(companies))
	for i, comp := range companies {
		uuids[i] = comp.UUID.String()
	}
	return uuids, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(companies)),
	}, nil

}
