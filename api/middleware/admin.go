package middleware

import (
	"fmt"

	"github.com/golang/glog"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// Take in a model, retuns the gentype for it
func adminToGentype(modAdmin models.Admin) gentypes.Admin {
	return gentypes.Admin{
		UUID:      modAdmin.UUID.String(),
		Email:     modAdmin.Email,
		FirstName: modAdmin.FirstName,
		LastName:  modAdmin.LastName,
	}
}

func adminsToGentypes(admins []models.Admin) []gentypes.Admin {
	var returnAdmins []gentypes.Admin
	for _, admin := range admins {
		newAdmin := adminToGentype(admin)
		returnAdmins = append(returnAdmins, newAdmin)
	}
	return returnAdmins
}

type AdminFilter struct {
	Email string
	Name  string
}

// GetAdminsByUUID
func (g *Grant) GetAdminsByUUID(uuids []string) ([]gentypes.Admin, error) {
	if !g.IsAdmin {
		return []gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin
	err := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins).Error
	if err != nil {
		return []gentypes.Admin{}, &errors.ErrNotFound
	}
	return adminsToGentypes(admins), nil
}

// adminExists returns true if the given uuid is an admin
// NB: Uses a DB query, so use sparingly
func (g *Grant) adminExists(uuid string) bool {
	query := database.GormDB.Where("uuid = ?", uuid).First(&models.Admin{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}
		// If some other error occurs log it
		glog.Errorf("Unable to find admin for UUID: %s - error: %s", uuid, query.Error.Error())
		return false
	}

	return true
}

func (g *Grant) adminEmailExists(email string) bool {
	query := database.GormDB.Where("email = ?", email).First(&models.Admin{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}
		// If some other error occurs log it
		glog.Errorf("Unable to find admin for Email: %s - error: %s", email, query.Error.Error())
		return false
	}
	return true
}

// GetAdmins
func (g *Grant) GetAdmins(page *gentypes.Page, filter *AdminFilter) ([]gentypes.Admin, error) {
	if !g.IsAdmin {
		return []gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin

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

	query, _, _ = getPage(query, page)
	err := query.Find(&admins).Error
	if err != nil {
		return []gentypes.Admin{}, err
	}

	return adminsToGentypes(admins), nil
}

// AddAdmin allows current admins to create new ones
func (g *Grant) AddAdmin(input gentypes.AddAdminInput) (gentypes.Admin, error) {
	if !g.IsAdmin {
		return gentypes.Admin{}, &errors.ErrUnauthorized
	}

	// Check if email already exists
	if g.adminEmailExists(input.Email) {
		return gentypes.Admin{}, &errors.ErrUserExists
	}

	admin := models.Admin{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	query := database.GormDB.Create(&admin)

	// Check for errors
	if query.Error != nil {
		glog.Errorf("Unable to create admin: %s", query.Error.Error())
		return gentypes.Admin{}, &errors.ErrWhileHandling
	}

	return adminToGentype(admin), nil
}

// DeleteAdmin allows admins to delete other admins
func (g *Grant) DeleteAdmin(uuid string) (bool, error) {
	if !g.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	// Check admin exists
	if !g.adminExists(uuid) {
		return false, &errors.ErrAdminNotFound
	}

	query := database.GormDB.Where("uuid = ?", uuid).Delete(models.Admin{})
	if query.Error != nil {
		glog.Errorf("Unable to delete admin: %s", query.Error.Error())
		return false, &errors.ErrDeleteFailed
	}

	return true, nil
}
