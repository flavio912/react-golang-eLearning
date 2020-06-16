package application

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type AdminApp interface {
	Admins(uuids []gentypes.UUID) ([]gentypes.Admin, error)
	CreateAdmin(input gentypes.CreateAdminInput) (gentypes.Admin, error)
	UpdateAdmin(input gentypes.UpdateAdminInput) (gentypes.Admin, error)
	DeleteAdmin(uuid gentypes.UUID) (bool, error)
	PageAdmins(page *gentypes.Page, filter *gentypes.AdminFilter) ([]gentypes.Admin, gentypes.PageInfo, error)
}

type adminAppImpl struct {
	grant           *middleware.Grant
	adminRepository middleware.AdminRepository
}

func NewAdminApp(grant *middleware.Grant) AdminApp {
	return &adminAppImpl{
		grant:           grant,
		adminRepository: middleware.NewAdminRepository(&grant.Logger),
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

func (a *adminAppImpl) CreateAdmin(input gentypes.CreateAdminInput) (gentypes.Admin, error) {
	if !a.grant.IsAdmin {
		return gentypes.Admin{}, &errors.ErrUnauthorized
	}

	admin, err := a.adminRepository.CreateAdmin(input)
	return adminToGentype(admin), err
}

func (a *adminAppImpl) Admins(uuids []gentypes.UUID) ([]gentypes.Admin, error) {
	if !a.grant.IsAdmin {
		return []gentypes.Admin{}, &errors.ErrUnauthorized
	}

	admins, err := a.adminRepository.Admins(uuids)
	return adminsToGentypes(admins), err
}

func (a *adminAppImpl) UpdateAdmin(input gentypes.UpdateAdminInput) (gentypes.Admin, error) {
	if !a.grant.IsAdmin {
		return gentypes.Admin{}, &errors.ErrUnauthorized
	}

	adminModel, updateErr := a.adminRepository.UpdateAdmin(input)

	if updateErr != nil {
		return gentypes.Admin{}, updateErr
	}

	return adminToGentype(adminModel), nil
}

func (a *adminAppImpl) DeleteAdmin(uuid gentypes.UUID) (bool, error) {
	if !a.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	return a.adminRepository.DeleteAdmin(uuid)
}

func (a *adminAppImpl) PageAdmins(page *gentypes.Page, filter *gentypes.AdminFilter) ([]gentypes.Admin, gentypes.PageInfo, error) {
	if !a.grant.IsAdmin {
		return []gentypes.Admin{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	admins, pageInfo, err := a.adminRepository.PageAdmins(page, filter)
	return adminsToGentypes(admins), pageInfo, err
}
