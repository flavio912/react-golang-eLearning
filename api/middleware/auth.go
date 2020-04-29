package middleware

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// GetAdminAccessToken gets an access token from the users email and password
func GetAdminAccessToken(email string, password string) (string, error) {
	a := &models.Admin{}
	admin, err := a.FindUser(email)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", &errors.ErrAdminNotFound
		}

		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}

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
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", &errors.ErrUserNotFound
		}

		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}

	token, err := manager.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}
	return token, nil
}

// HasFullRestrictedAccess returns true if the user has access to all restricted courses
func (g *Grant) HasFullRestrictedAccess() bool {
	if g.IsAdmin {
		return true
	}

	// If a managers company is authorized a manager can view all restricted courses
	if g.IsManager {
		company := models.Company{}
		query := database.GormDB.Where("uuid = ?", g.Claims.Company).First(&company)
		if query.Error != nil {
			glog.Errorf("Unable to get manager's company: %s", query.Error.Error())
			return false
		}

		if company.Approved == true {
			return true
		}
	}

	// Delegates cannot access restricted courses unless specifically assigned them
	if g.IsDelegate {
		return false
	}

	return false
}
