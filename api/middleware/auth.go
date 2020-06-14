package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
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

func GetDelegateAccessToken(ttcId string, password string) (string, error) {
	d := &models.Delegate{}
	delegate, err := d.FindUser(ttcId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", &errors.ErrUserNotFound
		}

		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}

	token, err := delegate.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}
	return token, nil
}

func GetIndividualAccessToken(email string, password string) (string, error) {
	ind := &models.Individual{}
	individual, err := ind.FindUser(email)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", &errors.ErrUserNotFound
		}

		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}

	token, err := individual.GenerateToken(password)
	if err != nil {
		glog.Info(err.Error())
		return "", &errors.ErrAuthFailed
	}
	return token, nil
}

func (g *Grant) GenerateCSRFToken() (string, error) {
	token, err := auth.GenerateCSRFToken(auth.CSRFClaims{
		UUID: g.Claims.UUID,
	})

	if err != nil {
		g.Logger.Log(sentry.LevelInfo, err, "Unable to generate CSRF token for user")
		return "", &errors.ErrWhileHandling
	}

	return token, nil
}
