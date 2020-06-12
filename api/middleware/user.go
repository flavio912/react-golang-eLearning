package middleware

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) GetCurrentUserByUUID(uuid gentypes.UUID) (gentypes.User, error) {
	if g.IsDelegate {
		delegate, err := g.GetDelegateByUUID(uuid)

		return gentypes.User{
			Type:            gentypes.DelegateType,
			Email:           delegate.Email,
			FirstName:       delegate.FirstName,
			LastName:        delegate.LastName,
			Telephone:       delegate.Telephone,
			JobTitle:        &delegate.JobTitle,
			LastLogin:       delegate.LastLogin,
			ProfileImageUrl: delegate.ProfileImageURL,
		}, err
	}

	if g.IsManager {
		manager, err := g.GetManagerByUUID(uuid)

		return gentypes.User{
			Type:            gentypes.ManagerType,
			Email:           &manager.Email,
			FirstName:       manager.FirstName,
			LastName:        manager.LastName,
			Telephone:       &manager.Telephone,
			JobTitle:        &manager.JobTitle,
			LastLogin:       manager.LastLogin,
			ProfileImageUrl: manager.ProfileImageURL,
		}, err
	}

	if g.IsIndividual {
		// individual, err := g.GetIndividualByUUID(uuid)
	}

	return gentypes.User{}, &errors.ErrUnauthorized
}

func (g *Grant) UsersByUUID(uuids []gentypes.UUID) ([]gentypes.User, error) {
	// Check delegates
	var delegates models.Delegate
	database.GormDB.Where("uuid IN (?)", uuids).Find(&delegates)

	// Check individuals

	// Check managers

}
