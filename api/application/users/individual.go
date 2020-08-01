package users

import (
	"time"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/email"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersAppImpl) individualToGentype(ind models.Individual) gentypes.Individual {
	created_at := ind.CreatedAt.Format(time.RFC3339)
	last_login := ind.LastLogin.Format(time.RFC3339)
	return gentypes.Individual{
		UUID:            ind.UUID,
		CreatedAt:       &created_at,
		Email:           ind.Email,
		FirstName:       ind.FirstName,
		LastName:        ind.LastName,
		JobTitle:        ind.JobTitle,
		Telephone:       ind.Telephone,
		LastLogin:       last_login,
		CourseTakerUUID: ind.CourseTakerUUID,
	}
}

func (u *usersAppImpl) individualsToGentype(inds []models.Individual) []gentypes.Individual {
	out := make([]gentypes.Individual, len(inds))
	for i, ind := range inds {
		out[i] = u.individualToGentype(ind)
	}
	return out
}

func (u *usersAppImpl) CreateIndividual(input gentypes.CreateIndividualInput) (gentypes.User, error) {
	individual, err := u.usersRepository.CreateIndividual(input)
	if err != nil {
		return gentypes.User{}, err
	}

	// Send transactional email
	err = email.SendAccountCompleteEmail(individual.FirstName, individual.Email)
	if err != nil {
		u.grant.Logger.Log(sentry.LevelWarning, err, "CreateIndividual: Unable to send complete email")
	}

	user := u.IndividualToUser(individual)
	return user, nil
}

func (u *usersAppImpl) UpdateIndividual(input gentypes.UpdateIndividualInput) (gentypes.User, error) {
	if !u.grant.IsAdmin && !u.grant.IsIndividual {
		return gentypes.User{}, &errors.ErrUnauthorized
	}

	if u.grant.IsIndividual && u.grant.Claims.UUID != input.UUID {
		return gentypes.User{}, &errors.ErrUnauthorized
	}

	ind, err := u.usersRepository.UpdateIndividual(input)
	user := u.IndividualToUser(ind)
	return user, err
}

func (u *usersAppImpl) Individual(uuid gentypes.UUID) (gentypes.Individual, error) {
	individual, err := u.usersRepository.Individual(uuid)
	return u.individualToGentype(individual), err
}

func (u *usersAppImpl) Individuals(page *gentypes.Page, filter *gentypes.IndividualFilter, orderBy *gentypes.OrderBy) ([]gentypes.Individual, gentypes.PageInfo, error) {
	individuals, pageInfo, err := u.usersRepository.Individuals(page, filter, orderBy)
	return u.individualsToGentype(individuals), pageInfo, err
}

func (u *usersAppImpl) DeleteIndividual(input gentypes.DeleteIndividualInput) (bool, error) {
	if !u.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	return u.usersRepository.DeleteIndividual(input.UUID)
}
