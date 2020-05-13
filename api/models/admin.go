package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// Admin - Model for admin users
type Admin struct {
	Base
	Email     string `gorm:"unique"`
	Password  string
	FirstName string
	LastName  string
}

// BeforeSave - Validate fields before save
func (admin *Admin) BeforeSave(scope *gorm.Scope) (err error) {
	// Make email lowercase
	scope.SetColumn("Email", strings.ToLower(admin.Email))

	// Hash the password
	if pw, err := auth.HashPassword(admin.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}

func (admin *Admin) getHash() string {
	return admin.Password
}

// FindUser - Find the user by their email address
func (*Admin) FindUser(email string) (IUser, error) {
	var admin Admin
	if err := database.GormDB.Where("email = ?", email).First(&admin).Error; err != nil {
		return &admin, err
	}
	return &admin, nil
}

// ValidatePassword - Check if a password and email for an admin is valid
func (*Admin) ValidatePassword(email string, password string) error {
	failedError := errors.New("Incorrect email or password")

	// Find the user
	var admin Admin
	if err := database.GormDB.Where("email = ?", email).First(&admin).Error; err != nil {
		return failedError
	}

	if err := auth.ValidatePassword(admin.Password, password); err == nil {
		// Success
		return nil
	}

	return failedError
}

/*GenerateToken - Create a JWT token for admins

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (admin *Admin) GenerateToken(password string) (string, error) {
	if err := admin.ValidatePassword(admin.Email, password); err != nil {
		return "", ErrPasswordInvalid
	}
	claims := auth.UserClaims{
		UUID: admin.UUID,
		Role: auth.AdminRole,
	}
	token, err := auth.GenerateToken(claims, helpers.Config.Jwt.AdminExpirationHours)
	return token, err
}
