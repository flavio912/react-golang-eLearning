package models

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
)

// Delegate - DB model for delegates
type Delegate struct {
	User
	Ident string `gorm:"unique"` // User identifier e.g tom_emmerson1
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (delegate *Delegate) GenerateToken(password string) (string, error) {
	claims := auth.UserClaims{
		UUID: delegate.UUID.String(),
		Role: "delegate",
	}
	token, err := auth.GenerateToken(claims, 24)
	return token, err
}
