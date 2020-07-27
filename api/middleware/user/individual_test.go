package user_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
)

func TestIndividual(t *testing.T) {
	prepareTestDatabase()

	individualUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012")
	t.Run("Gets correct user", func(t *testing.T) {
		ind, err := usersRepo.Individual(individualUUID)
		assert.Nil(t, err)
		assert.Equal(t, individualUUID, ind.UUID)
		assert.Equal(t, "individual@individual.com", ind.Email)
	})

}

func TestCreateIndividual(t *testing.T) {
	prepareTestDatabase()

	t.Run("Creates user with taker", func(t *testing.T) {
		ind, err := usersRepo.CreateIndividual(gentypes.CreateIndividualInput{
			FirstName: "Tony",
			LastName:  "Orange",
			Email:     "jim@jim.com",
			Password:  "iamapassword",
		})

		assert.Nil(t, err)
		assert.NotEqual(t, gentypes.UUID{}, ind.UUID)
		assert.Equal(t, "jim@jim.com", ind.Email)
		assert.Equal(t, "Tony", ind.FirstName)
		assert.Equal(t, "Orange", ind.LastName)
		assert.NotEqual(t, gentypes.UUID{}, ind.CourseTakerUUID)
	})

}
