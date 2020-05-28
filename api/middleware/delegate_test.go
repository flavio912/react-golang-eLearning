package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestGetDelegateByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Delegates cannot get delegates",
			delegateGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			&errors.ErrUnauthorized,
		},
		{
			"Delegates can see their own info",
			delegateGrant,
			delegateGrant.Claims.UUID,
			nil,
		},
		{
			"Admins can see delegates",
			adminGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			nil,
		},
		{
			"Managers of the same company can see its delegates",
			managerGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			nil,
		},
		{
			"Managers of different company cannot see delegate",
			managerGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003"),
			&errors.ErrUnauthorized,
		},

		{
			"Should return ErrNotFound if not found",
			adminGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"), // does not exist
			&errors.ErrNotFound,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.GetDelegateByUUID(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.uuid, m.UUID)
			} else {
				// should return a blank delegate if it errors
				assert.Equal(t, gentypes.Delegate{}, m)
			}
		})
	}
}

func TestGetDelegates(t *testing.T) {
	prepareTestDatabase()

	t.Run("Delegates can't access", func(t *testing.T) {
		_, _, err := delegateGrant.GetDelegates(nil, nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all delegates for admins", func(t *testing.T) {
		delegates, _, err := adminGrant.GetDelegates(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, delegates, 5)
	})

	t.Run("Should return only delegates of manager's company", func(t *testing.T) {
		delegates, _, err := managerGrant.GetDelegates(nil, nil, nil)
		assert.Nil(t, err)
		assert.Greater(t, len(delegates), 0)
		for _, d := range delegates {
			assert.Equal(t, managerGrant.Claims.Company, d.CompanyUUID)
		}
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		delegates, pageInfo, err := adminGrant.GetDelegates(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, delegates, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 5, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		delegates, _, err := adminGrant.GetDelegates(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, delegates, 5)
		assert.Equal(t, "David", delegates[0].FirstName)
	})

	t.Run("Should filter", func(t *testing.T) {
		delegate := gentypes.Delegate{
			User: gentypes.User{
				UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				FirstName: "Delegate",
				LastName:  "Man",
				Telephone: "7912935287",
				JobTitle:  "Doer",
			},
			Email:       "del@delegates.com",
			TTC_ID:      "delegate-test-1",
			CompanyUUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		}

		fullName := fmt.Sprintf("%s %s", delegate.FirstName, delegate.LastName)
		uuidString := delegate.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.DelegatesFilter
		}{
			{"Email", gentypes.DelegatesFilter{Email: &delegate.Email}},
			{"FirstName", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &delegate.FirstName}}},
			{"LastName", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &delegate.LastName}}},
			{"First and Last", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &fullName}}},
			{"JobTitle", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{JobTitle: &delegate.JobTitle}}},
			{"uuid", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{UUID: &uuidString}}},
			{"ttc_id", gentypes.DelegatesFilter{TTC_ID: &delegate.TTC_ID}},
			{"Full", gentypes.DelegatesFilter{UserFilter: gentypes.UserFilter{Name: &fullName}, Email: &delegate.Email}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				delegates, _, err := adminGrant.GetDelegates(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, delegates, 1)
				delegate.CreatedAt = delegates[0].CreatedAt
				delegate.ProfileImageURL = delegates[0].ProfileImageURL
				assert.Equal(t, delegate, delegates[0])
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.DelegatesFilter{Email: &email}
			delegates, _, err := adminGrant.GetDelegates(nil, &filter, nil)
			assert.Nil(t, err)
			require.Len(t, delegates, 2)
		})
	})
}

func TestCreateDelegate(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		wantErr interface{}
		want    gentypes.Delegate
		input   gentypes.CreateDelegateInput
	}{
		{
			"Users cannot create",
			delegateGrant,
			&errors.ErrUnauthorized,
			gentypes.Delegate{},
			gentypes.CreateDelegateInput{},
		},
		{
			"Admin must supply company uuid",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Delegate{},
			gentypes.CreateDelegateInput{},
		},
		{
			"Admin supplied company must exist",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Delegate{},
			gentypes.CreateDelegateInput{
				CompanyUUID: &uuidZero,
			},
		},
		{
			"Should use manager's company",
			managerGrant,
			nil,
			gentypes.Delegate{
				TTC_ID:      "testcompany-angrytim",
				CompanyUUID: managerGrant.Claims.Company,
				User: gentypes.User{
					FirstName: "Angry",
					LastName:  "Tim",
				},
			},
			gentypes.CreateDelegateInput{
				CreateUserInput: gentypes.CreateUserInput{
					FirstName: "Angry",
					LastName:  "Tim",
				},
			},
		},
		{
			"TTCID should be generated uniquely",
			managerGrant,
			nil,
			gentypes.Delegate{
				TTC_ID:      "testcompany-smellyjoe-1",
				CompanyUUID: managerGrant.Claims.Company,
				User: gentypes.User{
					FirstName: "Smelly",
					LastName:  "Joe",
				},
			},
			gentypes.CreateDelegateInput{
				CreateUserInput: gentypes.CreateUserInput{
					FirstName: "Smelly",
					LastName:  "Joe",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d, err := test.grant.CreateDelegate(test.input)
			assert.Equal(t, test.wantErr, err)
			// generated fields
			test.want.UUID = d.UUID
			test.want.ProfileImageURL = d.ProfileImageURL
			test.want.CreatedAt = d.CreatedAt
			assert.Equal(t, test.want, d)
		})
	}
}
