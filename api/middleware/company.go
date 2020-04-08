package middleware

import (
	"time"

	"github.com/golang/glog"
	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

//companyToGentypes converts a company model to gentype.
func (g *Grant) companyToGentypes(company models.Company) gentypes.Company {
	if g.ManagesCompany(company.UUID.String()) {
		createdAt := company.CreatedAt.Format(time.RFC3339)
		return gentypes.Company{
			CreatedAt: &createdAt,
			UUID:      company.UUID,
			Name:      company.Name,
		}
	}

	if g.IsCompanyDelegate(company.UUID.String()) {
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
		genCompanies = append(genCompanies, g.companyToGentypes(comp))
	}
	return genCompanies
}

//IsCompanyDelegate returns true if the grant user is a delegate of the given company uuid
func (g *Grant) IsCompanyDelegate(companyUUID string) bool {
	if g.IsDelegate {
		if companyUUID != "" && g.Claims.Company == companyUUID {
			return true
		}
	}
	return false
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

	return g.companiesToGentype(companies), nil
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

	return g.companyToGentypes(company), nil
}

// GetManagerIDsByCompany returns the uuids for the managers of a company
func (g *Grant) GetManagerIDsByCompany(companyUUID string, page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]uuid.UUID, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []uuid.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var (
		managerUUIDs []uuid.UUID
		managers     []models.Manager
	)

	query := database.GormDB.Select("uuid").Where("company_id = ?", companyUUID)

	var count int32
	err := query.Model(&models.Manager{}).Count(&count)
	if err.Error != nil {
		glog.Errorf("DB Error %s", err.Error.Error())
		glog.Errorf("Unable to count records for %s", companyUUID)
		return []uuid.UUID{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "first_name", "last_name"})
	if orderErr != nil {
		return []uuid.UUID{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query.Find(&managers)
	if query.Error != nil {
		return []uuid.UUID{}, gentypes.PageInfo{}, getDBErrorType(query)
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

func (g *Grant) GetCompanyUUIDs(page *gentypes.Page, filter *gentypes.CompanyFilter, orderBy *gentypes.OrderBy) ([]string, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []string{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var companies []models.Company

	query := database.GormDB.Select("uuid").Model(&models.Company{})

	var count int32
	query = query.Model(&models.Manager{}).Count(&count)
	if query.Error != nil {
		glog.Errorf("DB Error %s", query.Error.Error())
		return []string{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, err := getOrdering(query, orderBy, []string{"created_at", "name"})
	if err != nil {
		return []string{}, gentypes.PageInfo{}, err
	}

	query, limit, offset := getPage(query, page)

	query.Find(&companies)

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