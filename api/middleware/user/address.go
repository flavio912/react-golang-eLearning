package user

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersRepoImpl) GetAddressesByIDs(ids []uint) ([]models.Address, error) {
	// if !g.IsAdmin {
	// 	//TODO: Check if the users own company is in here, if so allow
	// 	//TODO: Check if individuals address is in here, is so allow
	// 	return []gentypes.Address{}, &errors.ErrUnauthorized
	// }

	var addresses []models.Address
	query := database.GormDB.Where("id IN (?)", ids).Find(&addresses)
	if query.Error != nil {
		u.Logger.Log(sentry.LevelError, query.Error, "Unable to load addresses")
		return []models.Address{}, &errors.ErrWhileHandling
	}

	return addresses, nil
}
