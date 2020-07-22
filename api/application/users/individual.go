package users

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersAppImpl) individualToGentype(ind models.Individual) gentypes.Individual {
	created_at := ind.CreatedAt.Format(time.RFC3339)
	last_login := ind.LastLogin.Format(time.RFC3339)
	return gentypes.Individual{
		UUID:      ind.UUID,
		CreatedAt: &created_at,
		Email:     ind.Email,
		FirstName: ind.FirstName,
		LastName:  ind.LastName,
		JobTitle:  ind.JobTitle,
		Telephone: ind.Telephone,
		LastLogin: last_login,
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
	user := u.IndividualToUser(individual)
	return user, err
}

func (u *usersAppImpl) UpdateIndividual(input gentypes.UpdateIndividualInput) (gentypes.User, error) {
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
