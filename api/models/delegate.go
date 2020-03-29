package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// Delegate - DB model for delegates
type Delegate struct {
	User
	Ident string `gorm:"unique"` // User identifier e.g tom_emmerson1
}

// DelegateClaims - The JWT claims for a Delegate
type DelegateClaims struct {
	jwt.StandardClaims
	UUID uuid.UUID
	Role string
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (delegate *Delegate) GenerateToken(password string) (string, error) {
	claims := &DelegateClaims{
		UUID: delegate.UUID,
		Role: "delegate",
	}
	token, err := auth.GenerateToken(claims, helpers.Config.Jwt.DelegateSecret)
	return token, err
}
