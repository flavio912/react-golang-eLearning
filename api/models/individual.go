package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Individual struct {
	User
	Email string `gorm:"unique"`
}

func (individual *Individual) getHash() string {
	return individual.Password
}

// FindUser - Find the user by their email address
func (*Individual) FindUser(email string) (IUser, error) {
	var individual Individual
	if err := database.GormDB.Where("email = ?", email).First(&individual).Error; err != nil {
		return &individual, err
	}
	return &individual, nil
}

// ValidatePassword - Check if a password and email for a manager is valid
func (*Individual) ValidatePassword(email string, password string) error {

	// Find the user
	m := &Individual{}
	manager, err := m.FindUser(email)
	if err != nil {
		return err
	}
	if err := auth.ValidatePassword(manager.getHash(), password); err == nil {
		// Success
		return nil
	}

	return &errors.ErrUnauthorized
}

/*GenerateToken - Create a JWT token for individuals

This function purposely takes in and verifies the password
(possibly even a second time), so that the token can in no
circumstances be given without the password - @temmerson
*/
func (individual *Individual) GenerateToken(password string) (string, error) {
	if err := individual.ValidatePassword(individual.Email, password); err != nil {
		return "", ErrPasswordInvalid
	}

	// Update last login time
	individual.LastLogin = time.Now()
	database.GormDB.Save(individual)

	claims := auth.UserClaims{
		UUID: individual.UUID,
		Role: auth.IndividualRole,
	}
	token, err := auth.GenerateToken(claims, 24)
	return token, err
}
