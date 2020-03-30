package middleware

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
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

// GetAccessToken - Get an access token from the users email and password
func GetAccessToken(email string, password string) (string, error) {
	admin := &models.Admin{}

	if err := database.GormDB.Where("email = ?", email).First(&admin).Error; gorm.IsRecordNotFoundError(err) {
		glog.Infof("Login failed, user not found - Given: %s", email)
		return "", &ErrUserNotFound
	}

	token, err := admin.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &ErrAuthFailed
	}

	return token, nil
}

func adminModelToGentype(modAdmin models.Admin) gentypes.Admin {
	return gentypes.Admin{
		UUID:      modAdmin.UUID.String(),
		Email:     modAdmin.Email,
		FirstName: modAdmin.FirstName,
		LastName:  modAdmin.LastName,
	}
}

// GetAdmins - Get all of the admins
func GetAdmins(jwt string) ([]gentypes.Admin, error) {
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

		// Map admin model onto the gentype model
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

		return gentypes.Admin{
			UUID:      admin.UUID.String(),
			Email:     admin.Email,
			FirstName: admin.FirstName,
			LastName:  admin.LastName,
		}, nil

	}
	return gentypes.Admin{}, &ErrUnauthorized
}
