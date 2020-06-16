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
