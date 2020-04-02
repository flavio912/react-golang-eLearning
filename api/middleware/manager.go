package middleware

import (
	"fmt"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func managerToGentype(manager models.Manager) gentypes.Manager {
	return gentypes.Manager{
		User: gentypes.User{
			UUID:      manager.UUID.String(),
			Email:     manager.Email,
			FirstName: manager.FirstName,
			LastName:  manager.LastName,
			JobTitle:  manager.JobTitle,
			Telephone: manager.Telephone,
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

func (g *Grant) GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter) ([]gentypes.Manager, error) {
	if !g.IsAdmin {
		return []gentypes.Manager{}, &errors.ErrUnauthorized
	}

	var managers []models.Manager

	// TODO: Like calls should be replaced with elasticsearch querys
	query := database.GormDB
	if filter != nil {
		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Email))
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

	query = getPage(query, page)
	err := query.Find(&managers).Error
	if err != nil {
		return []gentypes.Manager{}, err
	}

	return managersToGentype(managers), nil
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

	// TODO: Validate input better and return useful details
	manager := models.Manager{
		User: models.User{
			FirstName: managerDetails.FirstName,
			LastName:  managerDetails.LastName,
			Email:     managerDetails.Email,
			JobTitle:  managerDetails.JobTitle,
			Telephone: managerDetails.Telephone,
			Password:  managerDetails.Password,
		},
	}
	err := database.GormDB.Create(&manager).Error
	if err != nil {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	return managerToGentype(manager), nil
}
