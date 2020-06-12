package middleware

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) DelegateToUser(delegate models.Delegate) gentypes.User {
	genDelegate := g.delegateToGentype(delegate)
	return gentypes.DelegateToUser(genDelegate)
}

func (g *Grant) ManagerToUser(delegate models.Delegate) gentypes.User {
	genDelegate := g.delegateToGentype(delegate)
	return gentypes.User{
		Type:            gentypes.DelegateType,
		Email:           genDelegate.Email,
		FirstName:       genDelegate.FirstName,
		LastName:        genDelegate.LastName,
		Telephone:       genDelegate.Telephone,
		JobTitle:        &genDelegate.JobTitle,
		LastLogin:       genDelegate.LastLogin,
		ProfileImageUrl: genDelegate.ProfileImageURL,
	}
}

func (g *Grant) GetCurrentUser() (gentypes.User, error) {
	if g.IsDelegate {
		delegate, err := g.GetDelegateByUUID(g.Claims.UUID)
		return gentypes.DelegateToUser(delegate), err
	}

	if g.IsManager {
		manager, err := g.GetManagerByUUID(g.Claims.UUID)
		return gentypes.ManagerToUser(manager), err
	}

	if g.IsIndividual {
		// individual, err := g.GetIndividualByUUID(uuid)
	}

	return gentypes.User{}, &errors.ErrUnauthorized
}
