package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// Admin - Model for admin users
type Admin struct {
	Base
	Email    string
	Password string
	Name     string
}

// AdminClaims - JWT claims
type AdminClaims struct {
	jwt.StandardClaims
	UUID uuid.UUID
	Role string
}

// BeforeSave - Hash the given password
func (admin *Admin) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := helpers.HashPassword(admin.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (admin *Admin) GenerateToken(password string) (string, error) {

	claims := &AdminClaims{
		UUID: admin.UUID,
		Role: "admin",
	}
	token, err := auth.GenerateToken(claims, helpers.Config.Jwt.AdminSecret)
	return token, err
}
