package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
)

// Delegate - DB model for delegates
type Delegate struct {
	Base
	FirstName   string
	LastName    string
	JobTitle    string
	Telephone   *string
	LastLogin   time.Time
	Password    *string
	Email       *string
	TtcId       string `gorm:"unique"` // User identifier e.g Fedex_tom_emmerson1
	CompanyUUID gentypes.UUID
	ProfileKey  *string // S3 Upload key for the profile image
}

/*GenerateToken - Create a JWT token for delegates

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (delegate *Delegate) GenerateToken(password string) (string, error) {
	if err := delegate.ValidatePassword(delegate.TtcId, password); err != nil {
		return "", &errors.ErrUnauthorized
	}

	claims := auth.UserClaims{
		UUID:    delegate.UUID,
		Role:    auth.DelegateRole,
		Company: delegate.CompanyUUID,
	}

	token, err := auth.GenerateToken(claims, 24)

	if err == nil {
		database.GormDB.Model(&delegate).Update("last_login", time.Now())
	}
	return token, err
}

func (delegate *Delegate) getHash() *string {
	return delegate.Password
}

func (*Delegate) ValidatePassword(ttcId string, password string) error {
	failedError := &errors.ErrUnauthorized

	// Find the user
	d := &Delegate{}
	delegate, err := d.FindUser(ttcId)
	if err != nil {
		return err
	}

	if delegate.getHash() == nil {
		return failedError
	}

	if err := auth.ValidatePassword(*delegate.getHash(), password); err == nil {
		// Success
		return nil
	}

	return failedError
}

// FindUser - Find the user by their ttc id
func (*Delegate) FindUser(ttcId string) (*Delegate, error) {
	var delegate Delegate
	if err := database.GormDB.Where("ttc_id = ?", ttcId).First(&delegate).Error; err != nil {
		return &delegate, err
	}
	return &delegate, nil
}

// BeforeCreate - Hash the given password
func (delegate *Delegate) BeforeCreate(scope *gorm.Scope) (err error) {
	if delegate.Password != nil {
		if pw, err := auth.HashPassword(*delegate.Password); err == nil {
			scope.SetColumn("Password", pw)
		}
	}
	return
}
