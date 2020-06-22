package middleware

import (
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"

	"github.com/getsentry/sentry-go"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type AdminRepository interface {
	Admins(uuids []gentypes.UUID) ([]models.Admin, error)
	AdminExists(uuid gentypes.UUID) bool
	AdminEmailExists(email string) bool
	PageAdmins(page *gentypes.Page, filter *gentypes.AdminFilter) ([]models.Admin, gentypes.PageInfo, error)
	CreateAdmin(input gentypes.CreateAdminInput) (models.Admin, error)
	UpdateAdmin(input gentypes.UpdateAdminInput) (models.Admin, error)
	DeleteAdmin(input gentypes.UUID) (bool, error)
}

type adminRepository struct {
	Logger *logging.Logger
}

func NewAdminRepository(logger *logging.Logger) AdminRepository {
	return &adminRepository{
		Logger: logger,
	}
}

// Take in a model, returns the gentype for it
func adminToGentype(modAdmin models.Admin) gentypes.Admin {
	return gentypes.Admin{
		UUID:      modAdmin.UUID,
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

func (a *adminRepository) Admins(uuids []gentypes.UUID) ([]models.Admin, error) {
	var admins []models.Admin
	q := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins)
	if q.Error != nil {
		a.Logger.Log(sentry.LevelError, q.Error, "Unable to get admins by UUID")
		return []models.Admin{}, &errors.ErrWhileHandling
	}

	if len(admins) == 0 {
		return []models.Admin{}, &errors.ErrNotFound
	}
	return admins, nil
}

// adminExists returns true if the given uuid is an admin
// NB: Uses a DB query, so use sparingly
func (a *adminRepository) AdminExists(uuid gentypes.UUID) bool {
	query := database.GormDB.Where("uuid = ?", uuid).First(&models.Admin{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}
		// If some other error occurs log it
		a.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin for UUID: %s", uuid)
		return false
	}

	return true
}

func (a *adminRepository) AdminEmailExists(email string) bool {
	query := database.GormDB.Where("email = ?", email).First(&models.Admin{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}
		// If some other error occurs log it
		a.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin for Email: %s", email)
		return false
	}
	return true
}

func (a *adminRepository) PageAdmins(page *gentypes.Page, filter *gentypes.AdminFilter) ([]models.Admin, gentypes.PageInfo, error) {
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
	countErr := query.Model(&models.Admin{}).Count(&count).Error
	if countErr != nil {
		a.Logger.Log(sentry.LevelError, countErr, "Unable to count records for admin")
		return []models.Admin{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query = query.Order("created_at DESC")

	query, limit, offset := GetPage(query, page)
	err := query.Find(&admins).Error
	if err != nil {
		a.Logger.Log(sentry.LevelError, err, "Unable to find admins")
		return []models.Admin{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return admins, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(admins)),
	}, nil
}

// CreateAdmin allows current admins to create new ones
func (a *adminRepository) CreateAdmin(input gentypes.CreateAdminInput) (models.Admin, error) {
	// Check if email already exists
	if a.AdminEmailExists(input.Email) {
		return models.Admin{}, &errors.ErrUserExists
	}

	admin := models.Admin{
		Email:     input.Email,
		Password:  &input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	query := database.GormDB.Create(&admin)
	if query.Error != nil {
		a.Logger.Log(sentry.LevelError, query.Error, "Unable to create admin")
		return models.Admin{}, &errors.ErrWhileHandling
	}

	return admin, nil
}

func (a *adminRepository) UpdateAdmin(input gentypes.UpdateAdminInput) (models.Admin, error) {
	var admin models.Admin
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&admin)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Admin{}, &errors.ErrAdminNotFound
		}

		a.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin to update with UUID: %s", input.UUID)
		return models.Admin{}, &errors.ErrWhileHandling
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
		return admin, nil
	}

	save := database.GormDB.Save(admin)
	if save.Error != nil {
		a.Logger.Logf(sentry.LevelError, save.Error, "Error updating Admin with UUID: %s", input.UUID)
		return models.Admin{}, &errors.ErrWhileHandling
	}

	return admin, nil
}

// DeleteAdmin allows admins to delete other admins
func (a *adminRepository) DeleteAdmin(uuid gentypes.UUID) (bool, error) {
	// Check admin exists
	if !a.AdminExists(uuid) {
		return false, &errors.ErrAdminNotFound
	}

	query := database.GormDB.Where("uuid = ?", uuid).Delete(models.Admin{})
	if query.Error != nil {
		a.Logger.Logf(sentry.LevelError, query.Error, "Error updating Admin with UUID: %s", uuid)
		return false, &errors.ErrDeleteFailed
	}

	return true, nil
}
