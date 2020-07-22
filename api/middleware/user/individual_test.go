package user_test

import (
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestIndividuals(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return all individuals", func(t *testing.T) {
		inds, _, err := usersRepo.Individuals(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, inds, 2)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(1)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		inds, pageInfo, err := usersRepo.Individuals(&page, nil, nil)

		assert.Nil(t, err)
		assert.Len(t, inds, 1)
		assert.Equal(t, gentypes.PageInfo{Total: 2, Given: 1, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		inds, _, err := usersRepo.Individuals(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, inds, 2)
		assert.Equal(t, "Funny", inds[0].FirstName)
	})

	t.Run("Should filter", func(t *testing.T) {
		individual := gentypes.Individual{
			UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			FirstName: "Nice",
			LastName:  "Sharron",
			Email:     "individual@individual.com",
			Telephone: helpers.StringPointer("07912935287"),
			JobTitle:  helpers.StringPointer("Cool Person"),
		}

		fullName := fmt.Sprintf("%s %s", individual.FirstName, individual.LastName)
		uuidString := individual.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.IndividualFilter
		}{
			{"Email", gentypes.IndividualFilter{Email: &individual.Email}},
			{"FirstName", gentypes.IndividualFilter{UserFilter: gentypes.UserFilter{Name: &individual.FirstName}}},
			{"LastName", gentypes.IndividualFilter{UserFilter: gentypes.UserFilter{Name: &individual.LastName}}},
			{"First and Last", gentypes.IndividualFilter{UserFilter: gentypes.UserFilter{Name: &fullName}}},
			{"uuid", gentypes.IndividualFilter{UserFilter: gentypes.UserFilter{UUID: &uuidString}}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				inds, _, err := usersRepo.Individuals(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, inds, 1)
				assert.Equal(t, individual.FirstName, inds[0].FirstName)
				assert.Equal(t, individual.LastName, inds[0].LastName)
				assert.Equal(t, individual.Email, inds[0].Email)
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.IndividualFilter{Email: &email}
			inds, _, err := usersRepo.Individuals(nil, &filter, nil)
			assert.Nil(t, err)
			require.Len(t, inds, 2)
		})
	})
}

func TestUpdateIndividual(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must validate input", func(t *testing.T) {
		input := gentypes.UpdateIndividualInput{
			Email: helpers.StringPointer("not an email"),
		}

		ind, err := usersRepo.UpdateIndividual(input)
		ok, val_err := govalidator.ValidateStruct(input)

		assert.False(t, ok)
		assert.Equal(t, val_err, err)
		assert.Equal(t, models.Individual{}, ind)
	})

	t.Run("Cannot update non-existant individual", func(t *testing.T) {
		ind, err := usersRepo.UpdateIndividual(gentypes.UpdateIndividualInput{
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"),
		})

		assert.Equal(t, &errors.ErrNotFound, err)
		assert.Equal(t, models.Individual{}, ind)
	})

	t.Run("Updates an individual", func(t *testing.T) {
		input := gentypes.UpdateIndividualInput{
			UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			FirstName: helpers.StringPointer("Elon"),
			LastName:  helpers.StringPointer("Musk"),
			JobTitle:  helpers.StringPointer("CEO of SpaceX"),
			Telephone: helpers.StringPointer("07912935269"),
			Email:     helpers.StringPointer("elon.musk@spacex.com"),
			Password:  helpers.StringPointer("iamironman"),
		}

		ind, err := usersRepo.UpdateIndividual(input)

		assert.Nil(t, err)
		assert.Equal(t, *input.FirstName, ind.FirstName)
		assert.Equal(t, *input.LastName, ind.LastName)
		assert.Equal(t, *input.JobTitle, *ind.JobTitle)
		assert.Equal(t, *input.Telephone, *ind.Telephone)
		assert.Equal(t, *input.Email, ind.Email)
	})
}
