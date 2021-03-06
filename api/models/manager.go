package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// Manager - DB model for managers
type Manager struct {
	Base
	FirstName   string
	LastName    string
	JobTitle    string
	Telephone   string
	LastLogin   time.Time
	Password    string
	Email       string `gorm:"unique"`
	CompanyUUID gentypes.UUID
	ProfileKey  string
}

func (manager *Manager) getHash() string {
	return manager.Password
}

// FindUser - Find the user by their email address
func (*Manager) FindUser(email string) (*Manager, error) {
	var manager Manager
	if err := database.GormDB.Where("email = ?", email).First(&manager).Error; err != nil {
		return &manager, err
	}
	return &manager, nil
}

// ValidatePassword - Check if a password and email for a manager is valid
func (*Manager) ValidatePassword(email string, password string) error {
	failedError := &errors.ErrUnauthorized

	// Find the user
	m := &Manager{}
	manager, err := m.FindUser(email)
	if err != nil {
		return err
	}

	if err := auth.ValidatePassword(manager.getHash(), password); err == nil {
		// Success
		return nil
	}

	return failedError
}

/*GenerateToken - Create a JWT token for managers

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (manager *Manager) GenerateToken(password string) (string, error) {
	if err := manager.ValidatePassword(manager.Email, password); err != nil {
		return "", &errors.ErrUnauthorized
	}

	claims := auth.UserClaims{
		UUID:    manager.UUID,
		Role:    auth.ManagerRole,
		Company: manager.CompanyUUID,
	}
	token, err := auth.GenerateToken(claims, helpers.Config.Jwt.TokenExpirationHours)

	if err == nil {
		// Update last login time
		database.GormDB.Model(&manager).Update("last_login", time.Now())
	}
	return token, err
}

// BeforeCreate - Hash the given password
func (manager *Manager) BeforeCreate(scope *gorm.Scope) (err error) {
	if pw, err := auth.HashPassword(manager.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}
