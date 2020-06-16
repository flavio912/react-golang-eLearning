package user_test

import (
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestDelegate(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Can get delegate",
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			nil,
		},
		{
			"Should return ErrNotFound if not found",
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"), // does not exist
			&errors.ErrNotFound,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := usersRepo.Delegate(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.uuid, m.UUID)
			} else {
				// should return a blank delegate if it errors
				assert.Equal(t, models.Delegate{}, m)
			}
		})
	}
}

func TestGetDelegates(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return all delegates", func(t *testing.T) {
		delegates, _, err := usersRepo.GetDelegates(nil, nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, delegates, 5)
	})

	t.Run("Should return only delegates of company", func(t *testing.T) {
		var compUUID = gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		delegates, _, err := usersRepo.GetDelegates(nil, nil, nil, &compUUID)
		assert.Nil(t, err)
		assert.Greater(t, len(delegates), 0)
		for _, d := range delegates {
			assert.Equal(t, compUUID, d.CompanyUUID)
		}
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		delegates, pageInfo, err := usersRepo.GetDelegates(&page, nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, delegates, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 5, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		delegates, _, err := usersRepo.GetDelegates(nil, nil, &order, nil)
		assert.Nil(t, err)
		assert.Len(t, delegates, 5)
		assert.Equal(t, "David", delegates[0].FirstName)
	})

	t.Run("Should filter", func(t *testing.T) {
		delegate := gentypes.Delegate{
			UUID:        gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			FirstName:   "Delegate",
			LastName:    "Man",
			LastLogin:   "0001-01-01T00:00:00Z",
			Telephone:   helpers.StringPointer("7912935287"),
			JobTitle:    "Doer",
			Email:       helpers.StringPointer("del@delegates.com"),
			TTC_ID:      "delegate-test-1",
			CompanyUUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		}

		fullName := fmt.Sprintf("%s %s", delegate.FirstName, delegate.LastName)
		uuidString := delegate.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.DelegatesFilter
		}{
			{"Email", gentypes.DelegatesFilter{Email: delegate.Email}},
			{"FirstName", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &delegate.FirstName}}},
			{"LastName", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &delegate.LastName}}},
			{"First and Last", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &fullName}}},
			{"JobTitle", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{JobTitle: &delegate.JobTitle}}},
			{"uuid", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{UUID: &uuidString}}},
			{"ttc_id", gentypes.DelegatesFilter{TTC_ID: &delegate.TTC_ID}},
			{"Full", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &fullName}, Email: delegate.Email}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				delegates, _, err := usersRepo.GetDelegates(nil, &test.filter, nil, nil)
				assert.Nil(t, err)
				require.Len(t, delegates, 1)
				assert.Equal(t, delegate.FirstName, delegates[0].FirstName)
				assert.Equal(t, delegate.LastName, delegates[0].LastName)
				assert.Equal(t, delegate.UUID, delegates[0].UUID)
				assert.Equal(t, delegate.TTC_ID, delegates[0].TtcId)
				assert.Equal(t, delegate.CompanyUUID, delegates[0].CompanyUUID)
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.DelegatesFilter{Email: &email}
			delegates, _, err := usersRepo.GetDelegates(nil, &filter, nil, nil)
			assert.Nil(t, err)
			require.Len(t, delegates, 2)
		})
	})
}

func TestCreateDelegate(t *testing.T) {

	inp := gentypes.CreateDelegateInput{
		FirstName: "Smelly",
		LastName:  "Joe",
		Email:     helpers.StringPointer("testemail@devserver.london"),
	}

	t.Run("Should create delegate", func(t *testing.T) {
		prepareTestDatabase()

		comp, _ := usersRepo.Company(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		d, err := usersRepo.CreateDelegate(inp, helpers.StringPointer("/key/place"), helpers.StringPointer("MYPASSWORD"), comp, nil)
		assert.Nil(t, err)

		// generate TTC_ID uniquely
		assert.Equal(t, "testcompany-smellyjoe-1", d.TtcId)
		assert.Equal(t, inp.FirstName, d.FirstName)
		assert.Equal(t, inp.LastName, d.LastName)
		assert.Equal(t, inp.Email, d.Email)
	})

}
