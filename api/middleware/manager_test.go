package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestCreateManager(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		wantErr interface{}
		want    gentypes.Manager
		input   gentypes.CreateManagerInput
	}{
		{
			"Users cannot create",
			delegateGrant,
			&errors.ErrUnauthorized,
			gentypes.Manager{},
			gentypes.CreateManagerInput{},
		},
		{
			"Admin must supply company uuid",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Manager{},
			gentypes.CreateManagerInput{},
		},
		{
			"Admin supplied company must exist",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Manager{},
			gentypes.CreateManagerInput{
				CompanyUUID: &uuidZero,
			},
		},
		{
			"Should use manager's company",
			managerGrant,
			nil,
			gentypes.Manager{CompanyID: managerGrant.Claims.Company},
			gentypes.CreateManagerInput{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.CreateManager(test.input)
			assert.Equal(t, test.wantErr, err)
			// generated fields
			test.want.UUID = m.UUID
			test.want.ProfileImageURL = m.ProfileImageURL
			test.want.CreatedAt = m.CreatedAt
			assert.Equal(t, test.want, m)
		})
	}
}

func TestGetManagerByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Delegates cannot get managers",
			delegateGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			&errors.ErrUnauthorized,
		},
		{
			"Managers cannot see other managers",
			managerGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"), // different to managerGrant.Claims.UUID
			&errors.ErrUnauthorized,
		},
		{
			"Managers can see their own info",
			managerGrant,
			managerGrant.Claims.UUID,
			nil,
		},
		{
			"Admins can see managers",
			adminGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			nil,
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
			m, err := test.grant.GetManagerByUUID(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.uuid, m.UUID)
			} else {
				// should return a blank manager if it errors
				assert.Equal(t, gentypes.Manager{}, m)
			}
		})
	}
}

func TestGetManagersByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuids   []string
		wantErr interface{}
		wantLen int
	}{
		{
			"Delegates cannot get managers",
			delegateGrant,
			[]string{"00000000-0000-0000-0000-000000000001"},
			&errors.ErrUnauthorized,
			0,
		},
		{
			"Managers cannot see other managers",
			managerGrant,
			[]string{"00000000-0000-0000-0000-000000000002"},
			nil,
			0,
		},
		{
			"Managers can see their own info",
			managerGrant,
			[]string{managerGrant.Claims.UUID.String(), "00000000-0000-0000-0000-000000000002"},
			nil,
			1,
		},
		{
			"Admins can see all managers",
			adminGrant,
			[]string{"00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-000000000002"},
			nil,
			2,
		},
		{
			"UUIDs must be valid",
			adminGrant,
			[]string{"00000000-0000-0000-0000-000000000001", "this is not a uuid"},
			&errors.ErrWhileHandling, // we don't check all of the uuids
			0,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.GetManagersByUUID(test.uuids)
			assert.Equal(t, test.wantErr, err)
			assert.Len(t, m, test.wantLen)
		})
	}
}

func TestGetManagerSelf(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name     string
		grant    middleware.Grant
		wantUUID string
		wantErr  interface{}
	}{
		{
			"Must be manager",
			middleware.Grant{auth.UserClaims{}, true, false, true},
			"",
			&errors.ErrUnauthorized,
		},
		{
			"Should return own manager",
			managerGrant,
			managerGrant.Claims.UUID.String(),
			nil,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.GetManagerSelf()
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.wantUUID, m.UUID.String())
			} else {
				// should return a blank manager if it errors
				assert.Equal(t, gentypes.Manager{}, m)
			}
		})
	}
}

func TestDeleteManager(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Delegates cannot delete",
			delegateGrant,
			uuidZero,
			&errors.ErrUnauthorized,
		},
		{
			"Manager cannot delete other managers",
			managerGrant,
			uuidZero, // does not match managerGrant.Claims.UUID
			&errors.ErrUnauthorized,
		},
		{
			"Manager can delete self",
			managerGrant,
			managerGrant.Claims.UUID,
			nil,
		},
		{
			"Admins can delete",
			adminGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			nil,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := test.grant.DeleteManager(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				// check deleted
				assert.Equal(t, true, ret)
				_, err := test.grant.GetManagerByUUID(test.uuid)
				assert.Equal(t, err, &errors.ErrNotFound)
			} else {
				// should return a blank manager if it errors
				assert.Equal(t, false, ret)
			}
		})
	}
}

func TestGetManagers(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
		_, _, err := nonAdminGrant.GetManagers(nil, nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all managers", func(t *testing.T) {
		managers, _, err := adminGrant.GetManagers(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 4)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		managers, pageInfo, err := adminGrant.GetManagers(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 2)
		assert.Equal(t, pageInfo, gentypes.PageInfo{Total: 4, Given: 2, Limit: limit})
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		managers, _, err := adminGrant.GetManagers(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, managers, 4)
		assert.Equal(t, "Jimothy", managers[0].FirstName)
	})

	t.Run("Should filter", func(t *testing.T) {
		manager := gentypes.Manager{
			User: gentypes.User{
				UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				Email:     "man@managers.com",
				FirstName: "Manager",
				LastName:  "Man",
				Telephone: "7912938287",
				JobTitle:  "In Charge",
			},
			CompanyID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		}

		fullName := fmt.Sprintf("%s %s", manager.FirstName, manager.LastName)
		uuidString := manager.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.ManagersFilter
		}{
			{"Email", gentypes.ManagersFilter{Email: &manager.Email}},
			{"FirstName", gentypes.ManagersFilter{Name: &manager.FirstName}},
			{"LastName", gentypes.ManagersFilter{Name: &manager.LastName}},
			{"First and Last", gentypes.ManagersFilter{Name: &fullName}},
			{"JobTitle", gentypes.ManagersFilter{JobTitle: &manager.JobTitle}},
			{"uuid", gentypes.ManagersFilter{UUID: &uuidString}},
			{"Full", gentypes.ManagersFilter{Name: &fullName, Email: &manager.Email}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				managers, _, err := adminGrant.GetManagers(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, managers, 1)
				manager.CreatedAt = managers[0].CreatedAt
				manager.ProfileImageURL = managers[0].ProfileImageURL
				assert.Equal(t, manager, managers[0])
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.ManagersFilter{Email: &email}
			managers, _, err := adminGrant.GetManagers(nil, &filter, nil)
			assert.Nil(t, err)
			require.Len(t, managers, 4)
		})
	})
}
