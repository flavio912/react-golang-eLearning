package middleware

import (
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// AdminErrors - All error responses from admin middleware
var (
	ErrUserNotFound = FullError{
		Type:     "ErrUserNotFound",
		Message:  "Could not find user",
		Title:    "User not found",
		HelpText: "Couldn't find a user with that email address",
	}
	ErrAuthFailed = FullError{
		Type:     "ErrAuthFailed",
		Message:  "Email or password incorrect",
		Title:    "Email or password incorrect",
		HelpText: "Please try again",
	}
	ErrUnauthorized = SimpleError{
		Type:    "ErrUnauthorized",
		Message: "User does not have permissions for request",
	}
	ErrTokenInvalid = SimpleError{
		Type:    "ErrTokenInvalid",
		Message: "Given token is broken or expired",
	}
)

func adminModelToGentype(modAdmin models.Admin) gentypes.Admin {
	return gentypes.Admin{
		UUID:      modAdmin.UUID.String(),
		Email:     modAdmin.Email,
		FirstName: modAdmin.FirstName,
		LastName:  modAdmin.LastName,
	}
}

func modelsToGentypes(admins []models.Admin) []gentypes.Admin {
	var returnAdmins []gentypes.Admin
	for _, admin := range admins {
		newAdmin := adminModelToGentype(admin)
		returnAdmins = append(returnAdmins, newAdmin)
	}
	return returnAdmins
}

// TODO create middleware function for validation of the tokens

func GetAdminsByUUID(jwt string, uuids []string) ([]gentypes.Admin, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Info(err.Error())
		return []gentypes.Admin{}, &ErrTokenInvalid
	}
	if claims.Role == auth.AdminRole {
		var admins []models.Admin
		err := database.GormDB.Where("uuid IN (?)", uuids).Find(&admins).Error
		if err != nil {
			return []gentypes.Admin{}, nil
		}
		return modelsToGentypes(admins), nil
	}
	return []gentypes.Admin{}, nil
}

// GetAllAdmins - Get all of the admins
func GetAllAdmins(jwt string) ([]gentypes.Admin, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Info(err.Error())
		return []gentypes.Admin{}, &ErrTokenInvalid
	}
	if claims.Role == auth.AdminRole {
		var admins []models.Admin
		err := database.GormDB.Find(&admins).Error
		if err != nil {
			return []gentypes.Admin{}, err
		}

		var returnAdmins []gentypes.Admin
		for _, admin := range admins {
			newAdmin := adminModelToGentype(admin)
			returnAdmins = append(returnAdmins, newAdmin)
		}
		return returnAdmins, nil
	}
	return []gentypes.Admin{}, &ErrUnauthorized
}

// GetAdmin - Get an admin object from a jwt
func GetAdmin(jwt string, uuid *string) (gentypes.Admin, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Info(err.Error())
		return gentypes.Admin{}, &ErrTokenInvalid
	}
	if claims.Role == auth.AdminRole {
		// Get the admin user from the jwt
		var admin models.Admin
		err := database.GormDB.Where("uuid = ?", claims.UUID).First(&admin).Error
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
	return gentypes.Admin{}, &ErrUnauthorized
}
