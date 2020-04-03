package middleware

import (
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

type CompanyToManagers map[string][]gentypes.Manager

func (g *Grant) GetManagersByCompany(uuids []string) (CompanyToManagers, error) {
	var (
		authorizedUUIDs []string
		managers        []models.Manager
		compToMan       = CompanyToManagers{}
	)

	for _, uuid := range uuids {
		if g.ManagesCompany(uuid) {
			authorizedUUIDs = append(authorizedUUIDs, uuid)
		}
	}

	query := database.GormDB.Where("company_id IN (?)", authorizedUUIDs).Find(&managers)
	if query.Error != nil {
		return compToMan, getDBErrorType(query)
	}

	// Sort managers into correct pairings
	for _, manager := range managers {
		compToMan[manager.CompanyID.String()] = append(compToMan[manager.CompanyID.String()], managerToGentype(manager))
	}
	return compToMan, nil
}
