package middleware

import (
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

// GetAdminsByUUID
func (g *grant) GetAdminsByUUID(uuids []string) ([]*gentypes.Admin, error) {
	if g.IsAdmin {
		var admins []models.Admin
		err := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins).Error
		if err != nil {
			return []*gentypes.Admin{}, &errors.ErrNotFound
		}
		return adminModelsToGentypes(admins), nil
	}
	return []*gentypes.Admin{}, nil
}

// GetAllAdmins gets all of the admins in the db, or returns an error
func (g *grant) GetAllAdmins() ([]*gentypes.Admin, error) {
	if !g.IsAdmin {
		return []*gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin
	err := database.GormDB.Find(&admins).Error
	if err != nil {
		return []*gentypes.Admin{}, err
	}
	var returnAdmins []*gentypes.Admin
	for _, admin := range admins {
		newAdmin := adminModelToGentype(admin)
		returnAdmins = append(returnAdmins, &newAdmin)
	}
	return returnAdmins, nil
}

// GetAdmin gets an admin object from a jwt
func (g *grant) GetAdmin(uuid *string) (gentypes.Admin, error) {
	if g.IsAdmin {
		// Get the admin user from the jwt
		var admin models.Admin
		err := database.GormDB.Where("uuid = ?", g.Claims.UUID).First(&admin).Error
		if err != nil {
			return gentypes.Admin{}, err
		}

		// An admin can get this info about any other admin
		return gentypes.Admin{
			UUID:      admin.UUID.String(),
			Email:     admin.Email,
			FirstName: admin.FirstName,
			LastName:  admin.LastName,
		}, nil

	}
	return gentypes.Admin{}, &errors.ErrUnauthorized
}
