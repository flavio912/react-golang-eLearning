package user_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func TestCreateManager(t *testing.T) {

	compUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002")
	input := gentypes.CreateManagerInput{
		CompanyUUID: &compUUID,
		FirstName:   "Timmy",
		LastName:    "Orange",
		Email:       "emailface@email.com",
	}

	t.Run("Creates manager with required fields", func(t *testing.T) {
		prepareTestDatabase()
		manager, err := usersRepo.CreateManager(input, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Nil(t, err)
		assert.Equal(t, input.Email, manager.Email)
		assert.Equal(t, input.FirstName, manager.FirstName)
		assert.Equal(t, input.LastName, manager.LastName)
		assert.Equal(t, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"), manager.CompanyUUID) // Should ignore manager input uuid
	})

	t.Run("Doesn't allow invalid company", func(t *testing.T) {
		prepareTestDatabase()
		manager, err := usersRepo.CreateManager(input, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000333"))
		assert.Equal(t, &errors.ErrCompanyNotFound, err)
		assert.Equal(t, models.Manager{}, manager)
	})

	t.Run("Doesn't allow duplicate users", func(t *testing.T) {
		prepareTestDatabase()
		usersRepo.CreateManager(input, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		manager, err := usersRepo.CreateManager(input, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Equal(t, &errors.ErrUserExists, err)
		assert.Equal(t, models.Manager{}, manager)
	})

}

func TestManager(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Should return ErrNotFound if not found",
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"), // does not exist
			&errors.ErrNotFound,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := usersRepo.Manager(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.uuid, m.UUID)
			} else {
				// should return a blank manager if it errors
				assert.Equal(t, models.Manager{}, m)
			}
		})
	}
}

func TestGetManagersByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuids   []gentypes.UUID
		wantErr interface{}
		wantLen int
	}{
		{
			"Can get valid users",
			[]gentypes.UUID{gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")},
			nil, // we don't check all of the uuids
			1,
		},
		{
			"Returns none",
			[]gentypes.UUID{gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")},
			nil, // we don't check all of the uuids
			0,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prepareTestDatabase()
			m, err := usersRepo.GetManagersByUUID(test.uuids)
			assert.Equal(t, test.wantErr, err)
			assert.Len(t, m, test.wantLen)
		})
	}
}

func TestDeleteManager(t *testing.T) {

	tests := []struct {
		name    string
		uuid    gentypes.UUID
		wantErr interface{}
	}{
		{
			"Can delete",
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			nil,
		},
		{
			"Fails on non existant uuid",
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"),
			&errors.ErrUserNotFound,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prepareTestDatabase()
			ret, err := usersRepo.DeleteManager(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				// check deleted
				assert.Equal(t, true, ret)
				_, err := usersRepo.Manager(test.uuid)
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

	t.Run("Should return all managers", func(t *testing.T) {
		managers, _, err := usersRepo.GetManagers(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 5)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		managers, pageInfo, err := usersRepo.GetManagers(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, managers, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 5, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "first_name", Ascending: &asc}

		managers, _, err := usersRepo.GetManagers(nil, nil, &order)
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
				managers, _, err := usersRepo.GetManagers(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, managers, 1)
				assert.Equal(t, manager.UUID, managers[0].UUID)
				assert.Equal(t, manager.FirstName, managers[0].FirstName)
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.ManagersFilter{Email: &email}
			managers, _, err := usersRepo.GetManagers(nil, &filter, nil)
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
		manager, err := usersRepo.UpdateManager(input)

		assert.Nil(t, err)

		manager, err = usersRepo.Manager(input.UUID)
		assert.Nil(t, err)
		assert.Equal(t, *input.Email, manager.Email)
		assert.Equal(t, *input.FirstName, manager.FirstName)
		assert.Equal(t, *input.LastName, manager.LastName)
		assert.Equal(t, input.UUID, manager.UUID)
		assert.Equal(t, *input.Telephone, manager.Telephone)
		assert.Equal(t, *input.JobTitle, manager.JobTitle)
	})

	t.Run("UUID must exist", func(t *testing.T) {
		_, err := usersRepo.UpdateManager(gentypes.UpdateManagerInput{UUID: gentypes.UUID{}})
		assert.Equal(t, &errors.ErrManagerNotFound, err)
	})
}
