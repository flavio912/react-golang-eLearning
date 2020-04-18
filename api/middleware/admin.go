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
		UUID:      gentypes.UUID{UUID: modAdmin.UUID},
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
	q := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins)
	if q.Error != nil {
		return []gentypes.Admin{}, &errors.ErrWhileHandling
	}
	if len(admins) == 0 {
		return []gentypes.Admin{}, &errors.ErrNotFound
	}
	return adminsToGentypes(admins), nil
}

// adminExists returns true if the given uuid is an admin
// NB: Uses a DB query, so use sparingly
func (g *Grant) adminExists(uuid gentypes.UUID) bool {
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
func (g *Grant) GetAdmins(page *gentypes.Page, filter *AdminFilter) ([]gentypes.Admin, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.Admin{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var admins []models.Admin

	query := database.GormDB
	if filter != nil {
		if filter.Email != "" {
			query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", filter.Email))
		}
		if filter.Name != "" {
			query = query.Where("first_name || ' ' || last_name ILIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
		}
	}

	var count int32
	countErr := query.Model(&models.Manager{}).Count(&count).Error
	if countErr != nil {
		glog.Errorf("Unable to count records for admin. error: %s", countErr.Error())
		return []gentypes.Admin{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := getPage(query, page)
	err := query.Find(&admins).Error
	if err != nil {
		return []gentypes.Admin{}, gentypes.PageInfo{}, err
	}

	return adminsToGentypes(admins), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(admins)),
	}, nil
}

// CreateAdmin allows current admins to create new ones
func (g *Grant) CreateAdmin(input gentypes.CreateAdminInput) (gentypes.Admin, error) {
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

func (g *Grant) UpdateAdmin(input gentypes.UpdateAdminInput) (gentypes.Admin, error) {
	if !g.IsAdmin {
		return gentypes.Admin{}, &errors.ErrUnauthorized
	}

	var admin models.Admin
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&admin)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Admin{}, &errors.ErrAdminNotFound
		}

		glog.Errorf("Unable to find admin to update with UUID: %s - error: %s", input.UUID, query.Error.Error())
		return gentypes.Admin{}, &errors.ErrWhileHandling
	}

	changed := false
	if input.Email != nil && *input.Email != admin.Email {
		changed = true
		admin.Email = *input.Email
	}

	if input.FirstName != nil && *input.FirstName != admin.FirstName {
		changed = true
		admin.FirstName = *input.FirstName
	}

	if input.LastName != nil && *input.LastName != admin.LastName {
		changed = true
		admin.LastName = *input.LastName
	}

	if !changed {
		return adminToGentype(admin), nil
	}

	save := database.GormDB.Save(admin)
	if save.Error != nil {
		glog.Errorf("Error updating Admin with UUID: %s - error: %s", input.UUID, save.Error.Error())
		return gentypes.Admin{}, &errors.ErrWhileHandling
	}

	return adminToGentype(admin), nil
}

// DeleteAdmin allows admins to delete other admins
func (g *Grant) DeleteAdmin(uuid gentypes.UUID) (bool, error) {
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
