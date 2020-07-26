package users

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/email"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

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
