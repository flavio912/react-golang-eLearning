package models

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

// Delegate - DB model for delegates
type Delegate struct {
	User
	TtcId       string `gorm:"unique"` // User identifier e.g Fedex_tom_emmerson1
	CompanyUUID uuid.UUID
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (delegate *Delegate) GenerateToken(password string) (string, error) {
	claims := auth.UserClaims{
		UUID:    gentypes.UUID{UUID: delegate.UUID},
		Role:    auth.DelegateRole,
		Company: gentypes.UUID{UUID: delegate.CompanyUUID},
	}
	token, err := auth.GenerateToken(claims, 24)
	return token, err
}

func (delegate *Delegate) getHash() string {
	return delegate.Password
}

func (*Delegate) ValidatePassword(ttcId string, password string) error {
	failedError := errors.New("Incorrect TTC id or password")

	// Find the user
	d := &Delegate{}
	delegate, err := d.FindUser(ttcId)
	if err != nil {
		return err
	}
	if err := auth.ValidatePassword(delegate.getHash(), password); err == nil {
		// Success
		return nil
	}

	return failedError
}

// FindUser - Find the user by their ttc id
func (*Delegate) FindUser(ttcId string) (IUser, error) {
	var delegate Delegate
	if err := database.GormDB.Where("ttc_id = ?", ttcId).First(&delegate).Error; err != nil {
		return &delegate, err
	}
	return &delegate, nil
}
