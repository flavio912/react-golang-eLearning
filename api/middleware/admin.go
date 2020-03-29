package middleware

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// AdminErrors - All error responses from admin middleware
var (
	ErrUserNotFound = FullError{
		Type:     "ErrUserNotFound",
		Message:  "Could not find user",
		Title:    "User not found",
		HelpText: "Couldn't find a user with that email address",
	}
	ErrAuthFailed = FullError{
		Type:     "ErrAuthFailed",
		Message:  "Email or password incorrect",
		Title:    "Email or password incorrect",
		HelpText: "Please try again",
	}
)

// GetAccessToken - Get an access token from the users email and password
func GetAccessToken(email string, password string) (string, error) {

	admin := &models.Admin{}
	if err := database.GormDB.Where("email = ?", email).First(&admin); err != nil {
		return "", &ErrUserNotFound
	}

	token, err := admin.GenerateToken(password)
	if err != nil {
		return "", &ErrAuthFailed
	}

	return token, nil
}
