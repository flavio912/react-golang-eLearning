package users

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func CreateIndividual(g *middleware.Grant, input gentypes.CreateIndividualInput) (gentypes.User, error) {
	individual, err := g.CreateIndividual(input)
	user := g.IndividualToUser(individual)
	return user, err
}
