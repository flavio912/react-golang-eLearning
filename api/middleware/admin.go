package middleware

import (
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// Take in a model, retuns the gentype for it
func adminModelToGentype(modAdmin models.Admin) gentypes.Admin {
	return gentypes.Admin{
		UUID:      modAdmin.UUID.String(),
		Email:     modAdmin.Email,
		FirstName: modAdmin.FirstName,
		LastName:  modAdmin.LastName,
	}
}

func adminModelsToGentypes(admins []models.Admin) []*gentypes.Admin {
	var returnAdmins []*gentypes.Admin
	for _, admin := range admins {
		newAdmin := adminModelToGentype(admin)
		returnAdmins = append(returnAdmins, &newAdmin)
	}
	return returnAdmins
}

type AdminFilter struct {
	Email string
	Name  string
}

// GetAdminsByUUID
func (g *Grant) GetAdminsByUUID(uuids []string) ([]*gentypes.Admin, error) {
	if !g.IsAdmin {
		return []*gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin
	err := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins).Error
	if err != nil {
		return []*gentypes.Admin{}, &errors.ErrNotFound
	}
	return adminModelsToGentypes(admins), nil
}

// GetAdmins
func (g *Grant) GetAdmins(page *gentypes.Page, filter *AdminFilter) ([]*gentypes.Admin, error) {
	if !g.IsAdmin {
		return []*gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin

	// TODO: LIKE querys should be replaced with elasticsearch
	query := database.GormDB
	if filter != nil {
		if filter.Email != "" {
			query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", filter.Email))
		}
		if filter.Name != "" {
			query = query.Where("first_name ILIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
			query = query.Where("last_name ILIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
		}
	}

	query = getPage(query, page)
	err := query.Find(&admins).Error
	if err != nil {
		return []*gentypes.Admin{}, err
	}

	return adminModelsToGentypes(admins), nil
}
