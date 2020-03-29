package middleware

import (
	"errors"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"golang.org/x/crypto/bcrypt"
)

// ValidateAdminPassword - Check if a password and email for an admin is valid
func ValidateAdminPassword(email string, password string) error {
	failedError := errors.New("Incorrect email or password")

	// Find the user
	var admin models.Admin
	if err := database.GormDB.Where("email = ?", email).First(&admin); err != nil {
		return failedError
	}

	// Compare given password to the hashed value
	byteHash := []byte(admin.Password)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return failedError
	}

	return nil
}
