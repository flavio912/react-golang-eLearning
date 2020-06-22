package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Individual struct {
	Base
	FirstName       string
	LastName        string
	JobTitle        *string
	Telephone       *string
	LastLogin       time.Time
	Password        string
	Email           string `gorm:"unique"`
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID
}

func (individual *Individual) getHash() string {
	return individual.Password
}

// FindUser - Find the user by their email address
func (*Individual) FindUser(email string) (*Individual, error) {
	var individual Individual
	if err := database.GormDB.Where("email = ?", email).First(&individual).Error; err != nil {
		return &individual, err
	}
	return &individual, nil
}

// ValidatePassword - Check if a password and email is valid
func (*Individual) ValidatePassword(email string, password string) error {

	// Find the user
	m := &Individual{}
	individual, err := m.FindUser(email)
	if err != nil {
		return err
	}

	if err := auth.ValidatePassword(individual.getHash(), password); err == nil {
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
		return "", &errors.ErrUnauthorized
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

// BeforeCreate - Hash the given password
func (individual *Individual) BeforeCreate(scope *gorm.Scope) (err error) {
	if pw, err := auth.HashPassword(individual.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}
