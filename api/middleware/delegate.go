package middleware

import (
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) delegateToGentype(delegate models.Delegate) gentypes.Delegate {
	createdAt := delegate.CreatedAt.Format(time.RFC3339)
	return gentypes.Delegate{
		User: gentypes.User{
			CreatedAt: &createdAt,
			UUID:      gentypes.UUID{UUID: delegate.UUID},
			Email:     delegate.Email,
			FirstName: delegate.FirstName,
			LastName:  delegate.LastName,
			JobTitle:  delegate.JobTitle,
			Telephone: delegate.Telephone,
		},
		CompanyUUID: gentypes.UUID{UUID: delegate.CompanyUUID},
		TTC_ID:      delegate.TtcId,
	}
}

func (g *Grant) delegateExists(email string, ttcId string) bool {
	query := database.GormDB.Where("email = ? or ttc_id = ?", email, ttcId).First(&models.Delegate{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}
		// If some other error occurs log it
		glog.Errorf("Unable to find delegate for Email: %s - error: %s", email, query.Error.Error())
		return false
	}
	return true

}

func (g *Grant) GetDelegateFromUUID(UUID gentypes.UUID) (gentypes.Delegate, error) {
	var delegate models.Delegate
	err := database.GormDB.Where("uuid = ?", UUID).First(&delegate).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return gentypes.Delegate{}, &errors.ErrNotFound
		}

		return gentypes.Delegate{}, err
	}

	if !g.IsAdmin && g.Claims.Company.UUID != delegate.CompanyUUID {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	return g.delegateToGentype(delegate), nil
}

func (g *Grant) CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, error) {
	if !g.IsAdmin && !g.IsManager {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	// check delegate does not exist
	if g.delegateExists(delegateDetails.Email, delegateDetails.TTC_ID) {
		return gentypes.Delegate{}, &errors.ErrUserExists
	}

	var companyUUID gentypes.UUID
	// If you're an admin you need to provide the company UUID
	if g.IsAdmin {
		if delegateDetails.CompanyUUID != nil {
			companyUUID = *delegateDetails.CompanyUUID
		} else {
			return gentypes.Delegate{}, &errors.ErrCompanyNotFound
		}
	} else {
		companyUUID = g.Claims.Company
	}

	// Check if company exists
	if !g.CompanyExists(companyUUID) {
		return gentypes.Delegate{}, &errors.ErrCompanyNotFound
	}

	delegate := models.Delegate{
		User: models.User{
			FirstName: delegateDetails.FirstName,
			LastName:  delegateDetails.LastName,
			Email:     delegateDetails.Email,
			JobTitle:  delegateDetails.JobTitle,
			Telephone: delegateDetails.Telephone,
			Password:  delegateDetails.Password,
		},
		CompanyUUID: companyUUID.UUID,
		TtcId:       delegateDetails.TTC_ID,
	}
	createErr := database.GormDB.Create(&delegate).Error
	if createErr != nil {
		return gentypes.Delegate{}, &errors.ErrWhileHandling
	}

	return g.delegateToGentype(delegate), nil
}
