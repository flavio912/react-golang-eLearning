package users

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (u *usersAppImpl) DelegateToUser(delegate models.Delegate) gentypes.User {

	var uploadUrl *string
	if delegate.ProfileKey != nil {
		uploadUrl = helpers.StringPointer(uploads.GetImgixURL(*delegate.ProfileKey))
	}

	return gentypes.User{
		UUID:            delegate.UUID,
		Type:            gentypes.DelegateType,
		Email:           delegate.Email,
		FirstName:       delegate.FirstName,
		LastName:        delegate.LastName,
		Telephone:       delegate.Telephone,
		JobTitle:        &delegate.JobTitle,
		LastLogin:       delegate.LastLogin.Format(time.RFC3339),
		ProfileImageUrl: uploadUrl,
		CourseTakerUUID: &delegate.CourseTakerUUID,
	}
}

func (u *usersAppImpl) ManagerToUser(manager models.Manager) gentypes.User {
	genManager := u.managerToGentype(manager)
	return gentypes.User{
		Type:            gentypes.ManagerType,
		Email:           &genManager.Email,
		FirstName:       genManager.FirstName,
		LastName:        genManager.LastName,
		Telephone:       &genManager.Telephone,
		JobTitle:        &genManager.JobTitle,
		LastLogin:       genManager.LastLogin,
		ProfileImageUrl: genManager.ProfileImageURL,
	}
}

func (u *usersAppImpl) IndividualToUser(individual models.Individual) gentypes.User {
	return gentypes.User{
		Type:            gentypes.IndividualType,
		Email:           &individual.Email,
		FirstName:       individual.FirstName,
		LastName:        individual.LastName,
		Telephone:       individual.Telephone,
		JobTitle:        individual.JobTitle,
		LastLogin:       individual.LastLogin.String(),
		CourseTakerUUID: &individual.CourseTakerUUID,
	}
}

func (u *usersAppImpl) GetCurrentUser() (gentypes.User, error) {
	if u.grant.IsDelegate {
		delegate, err := u.usersRepository.Delegate(u.grant.Claims.UUID)
		return u.DelegateToUser(delegate), err
	}

	if u.grant.IsManager {
		manager, err := u.usersRepository.Manager(u.grant.Claims.UUID)
		return u.ManagerToUser(manager), err
	}

	if u.grant.IsIndividual {
		individual, err := u.usersRepository.Individual(u.grant.Claims.UUID)
		return u.IndividualToUser(individual), err
	}

	return gentypes.User{}, &errors.ErrUnauthorized
}

