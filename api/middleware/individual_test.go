package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

func TestIndividual(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name           string
		uuid           gentypes.UUID
		grant          middleware.Grant
		wantErr        error
		wantIndividual models.Individual
	}{
		{
			name:           "Delegates can't access",
			uuid:           gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			grant:          delegateGrant,
			wantErr:        &errors.ErrUnauthorized,
			wantIndividual: models.Individual{},
		},
		{
			name:           "Managers can't access",
			uuid:           gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			grant:          managerGrant,
			wantErr:        &errors.ErrUnauthorized,
			wantIndividual: models.Individual{},
		},
		{
			name:           "Public can't access",
			uuid:           gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			grant:          publicGrant,
			wantErr:        &errors.ErrUnauthorized,
			wantIndividual: models.Individual{},
		},
		{
			name:           "Individuals can't get other info",
			uuid:           gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000013"),
			grant:          individualGrant,
			wantErr:        &errors.ErrUnauthorized,
			wantIndividual: models.Individual{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ind, err := test.grant.Individual(test.uuid)
			assert.Equal(t, test.wantIndividual, ind)
			assert.Equal(t, test.wantErr, err)
		})
	}

	t.Run("Gets correct user", func(t *testing.T) {
		ind, err := individualGrant.Individual(individualGrant.Claims.UUID)
		assert.Nil(t, err)
		assert.Equal(t, individualGrant.Claims.UUID, ind.UUID)
		assert.Equal(t, "individual@individual.com", ind.Email)
	})

}
