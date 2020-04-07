package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
)

// Manager - DB model for managers
type Manager struct {
	User
	CompanyID  uuid.UUID
	ProfileKey string
}

func (manager *Manager) getHash() string {
	return manager.Password
}

// FindUser - Find the user by their email address
func (*Manager) FindUser(email string) (IUser, error) {
	var manager Manager
	if err := database.GormDB.Where("email = ?", email).First(&manager).Error; err != nil {
		return &manager, err
	}
	return &manager, nil
}

// ValidatePassword - Check if a password and email for a manager is valid
func (*Manager) ValidatePassword(email string, password string) error {
	failedError := errors.New("Incorrect email or password")

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
		return "", ErrPasswordInvalid
	}

	// Update last login time
	manager.LastLogin = time.Now()
	database.GormDB.Save(manager)

	claims := auth.UserClaims{
		UUID:    manager.UUID.String(),
		Role:    auth.ManagerRole,
		Company: manager.CompanyID.String(),
	}
	token, err := auth.GenerateToken(claims, 24)
	return token, err
}
