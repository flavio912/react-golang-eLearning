package middleware

import (
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// GetAdminAccessToken gets an access token from the users email and password
func GetAdminAccessToken(email string, password string) (string, error) {
	a := &models.Admin{}
	admin, err := a.FindUser(email)
	token, err := admin.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}

	return token, nil
}

// GetManagerAccessToken gets an access token from the users email and password
func GetManagerAccessToken(email string, password string) (string, error) {
	m := &models.Manager{}
	manager, err := m.FindUser(email)
	token, err := manager.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}
	return token, nil
}
