package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
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
			gentypes.Manager{CompanyUUID: managerGrant.Claims.Company},
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
			&errors.ErrUnauthorized,
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
			prepareTestDatabase()
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
			middleware.Grant{auth.UserClaims{}, true, false, true, false, false, logging.Logger{}},
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
			prepareTestDatabase()
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
		_, _, err := nonAdminGrant.GetManagers(nil, nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all managers", func(t *testing.T) {
		managers, _, err := adminGrant.GetManagers(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 5)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		managers, pageInfo, err := adminGrant.GetManagers(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 5, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		managers, _, err := adminGrant.GetManagers(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, managers, 5)
		assert.Equal(t, "Jimothy", managers[0].FirstName)
	})

	t.Run("Should filter", func(t *testing.T) {
		manager := gentypes.Manager{
			UUID:        gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			FirstName:   "Manager",
			LastName:    "Man",
			Telephone:   "7912938287",
			JobTitle:    "In Charge",
			Email:       "man@managers.com",
			CompanyUUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		}

		fullName := fmt.Sprintf("%s %s", manager.FirstName, manager.LastName)
		uuidString := manager.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.ManagersFilter
		}{
			{"Email", gentypes.ManagersFilter{Email: &manager.Email}},
			{"FirstName", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{Name: &manager.FirstName}}},
			{"LastName", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{Name: &manager.LastName}}},
			{"First and Last", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{Name: &fullName}}},
			{"JobTitle", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{JobTitle: &manager.JobTitle}}},
			{"uuid", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{UUID: &uuidString}}},
			{"Full", gentypes.ManagersFilter{UserFilter: gentypes.UserFilter{Name: &fullName}, Email: &manager.Email}},
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
			require.Len(t, managers, 5)
		})
	})
}

func TestUpdateManager(t *testing.T) {
	prepareTestDatabase()

	input := gentypes.UpdateManagerInput{
		UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		Email:     helpers.StringPointer("test@test.com"),
		FirstName: helpers.StringPointer("test"),
		LastName:  helpers.StringPointer("test2"),
		Telephone: helpers.StringPointer("test3"),
		JobTitle:  helpers.StringPointer("test4"),
	}

	t.Run("Updates existing manager", func(t *testing.T) {
		prepareTestDatabase()
		manager, err := adminGrant.UpdateManager(input)

		outputWant := gentypes.Manager{
			UUID:            input.UUID,
			FirstName:       *input.FirstName,
			LastName:        *input.LastName,
			Telephone:       *input.Telephone,
			JobTitle:        *input.JobTitle,
			CreatedAt:       manager.CreatedAt,
			LastLogin:       manager.LastLogin,
			Email:           *input.Email,
			ProfileImageURL: manager.ProfileImageURL,
			CompanyUUID:     manager.CompanyUUID,
		}

		assert.Nil(t, err)
		assert.Equal(t, outputWant, manager)

		manager, err = adminGrant.GetManagerByUUID(input.UUID)
		assert.Nil(t, err)
		assert.Equal(t, outputWant, manager)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		_, err := managerGrant.UpdateManager(gentypes.UpdateManagerInput{UUID: managerGrant.Claims.UUID})
		assert.Nil(t, err, "manager should be able to edit itself")

		_, err = managerGrant.UpdateManager(gentypes.UpdateManagerInput{
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
		})
		assert.Equal(t, &errors.ErrUnauthorized, err, "manager should not be able to edit other managers")

		_, err = delegateGrant.UpdateManager(gentypes.UpdateManagerInput{UUID: uuidZero})
		assert.Equal(t, &errors.ErrUnauthorized, err, "delegates must not be able to update manager")
	})

	t.Run("UUID must exist", func(t *testing.T) {
		_, err := adminGrant.UpdateManager(gentypes.UpdateManagerInput{UUID: uuidZero})
		assert.Equal(t, &errors.ErrManagerNotFound, err)
	})
}
