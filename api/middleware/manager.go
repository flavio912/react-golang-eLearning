package middleware

import (
	"github.com/golang/glog"
	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func managerToGentype(manager models.Manager) gentypes.Manager {
	return gentypes.Manager{
		User: gentypes.User{
			UUID:      manager.UUID,
			Email:     manager.Email,
			FirstName: manager.FirstName,
			LastName:  manager.LastName,
			JobTitle:  manager.JobTitle,
			Telephone: manager.Telephone,
			CompanyID: manager.CompanyID,
		},
	}
}

func managersToGentype(managers []models.Manager) []gentypes.Manager {
	var genManagers []gentypes.Manager
	for _, manager := range managers {
		genManagers = append(genManagers, managerToGentype(manager))
	}
	return genManagers
}

func (g *Grant) GetManagersByUUID(uuids []string) ([]gentypes.Manager, error) {
	var managers []gentypes.Manager
	if g.IsAdmin {
		db := database.GormDB.Where("uuid IN (?)", uuids).Find(&managers)
		if db.Error != nil {
			if db.RecordNotFound() {
				return managers, &errors.ErrNotFound
			}
			glog.Errorf("DB Error: %s", db.Error.Error())
			return managers, &errors.ErrWhileHandling
		}

		return managers, nil
	}

	return managers, &errors.ErrUnauthorized
}

func (g *Grant) GetManagerByUUID(uuid string) (gentypes.Manager, error) {
	// Admins can get any manager data
	// Managers can only get their own uuid
	if g.IsAdmin || (g.IsManager && uuid == g.Claims.UUID) {
		var manager models.Manager
		err := database.GormDB.Where("uuid = ?", uuid).First(&manager).Error
		if err != nil {
			return gentypes.Manager{}, err
		}

		return managerToGentype(manager), nil
	}
	return gentypes.Manager{}, &errors.ErrUnauthorized
}

func (g *Grant) GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter) ([]gentypes.Manager, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var managers []models.Manager

	query := database.GormDB
	if filter != nil {
		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("first_name ILIKE ?", "%%"+*filter.Name+"%%").Or("last_name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.UUID != nil && *filter.UUID != "" {
			query = query.Where("uuid = ?", *filter.UUID)
		}
		if filter.JobTitle != nil && *filter.JobTitle != "" {
			query = query.Where("job_title ILIKE ?", "%%"+*filter.JobTitle+"%%")
		}
	}

	// Count the total filtered dataset
	var count int32
	countErr := query.Model(&models.Manager{}).Limit(MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		glog.Errorf("Count query failed: %s", countErr.Error())
		return []gentypes.Manager{}, gentypes.PageInfo{}, countErr
	}

	query = query.Order("created_at DESC")
	query, limit, offset := getPage(query, page)
	err := query.Find(&managers).Error
	if err != nil {
		return []gentypes.Manager{}, gentypes.PageInfo{}, err
	}

	return managersToGentype(managers), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(managers)),
	}, nil
}

func (g *Grant) GetManagerSelf() (gentypes.Manager, error) {
	if !g.IsManager {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	manager, err := g.GetManagerByUUID(g.Claims.UUID)
	if err != nil {
		return gentypes.Manager{}, err
	}

	return manager, nil
}

func (g *Grant) AddManager(managerDetails gentypes.AddManagerInput) (gentypes.Manager, error) {
	if !g.IsAdmin {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	_uuid, err := uuid.Parse(managerDetails.CompanyUUID)
	if err != nil {
		return gentypes.Manager{}, &errors.ErrUUIDInvalid
	}

	// Check if company exists
	var company models.Company
	existsErr := database.GormDB.Where("uuid = ?", _uuid).First(&company)
	if existsErr.Error != nil {
		if existsErr.RecordNotFound() {
			return gentypes.Manager{}, &errors.ErrCompanyNotFound
		}
		return gentypes.Manager{}, &errors.ErrWhileHandling
	}

	// TODO: Validate input better and return useful details
	manager := models.Manager{
		User: models.User{
			CompanyID: _uuid,
			FirstName: managerDetails.FirstName,
			LastName:  managerDetails.LastName,
			Email:     managerDetails.Email,
			JobTitle:  managerDetails.JobTitle,
			Telephone: managerDetails.Telephone,
			Password:  managerDetails.Password,
		},
	}
	createErr := database.GormDB.Create(&manager).Error
	if createErr != nil {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	return managerToGentype(manager), nil
}
