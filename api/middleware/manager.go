package middleware

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func managerModelToGentype(manager models.Manager) gentypes.Manager {
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

func (g *Grant) GetManagersByUUID(uuids []string) ([]*gentypes.Manager, error) {
	var managers []*gentypes.Manager
	if g.IsAdmin {
		err := database.GormDB.Where("uuid IN (?)", uuids).Find(&managers).Error
		return managers, err
	}
	return managers, &errors.ErrUnauthorized
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

	return managerModelToGentype(manager), nil
}
