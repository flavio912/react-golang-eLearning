package models

import (
	"errors"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

// Delegate - DB model for delegates
type Delegate struct {
	User
	Ident string `gorm:"unique"` // User identifier e.g Fedex_tom_emmerson1
}

// ValidatePassword - Check if a password and ident for a delegate is valid
func (*Delegate) ValidatePassword(ident string, password string) error {
	failedError := errors.New("Incorrect email or password")

	// Find the user
	var delegate Delegate
	if err := database.GormDB.Where("ident = ?", ident).First(&delegate).Error; err != nil {
		return failedError
	}

	if err := auth.ValidatePassword(delegate.Password, password); err == nil {
		// Success
		return nil
	}

	return failedError
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (delegate *Delegate) GenerateToken(password string) (string, error) {
	claims := auth.UserClaims{
		UUID: gentypes.UUID{UUID: delegate.UUID},
		Role: auth.DelegateRole,
	}
	token, err := auth.GenerateToken(claims, 24)
	return token, err
}
