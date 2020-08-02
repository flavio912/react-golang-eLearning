package user_test

import (
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var realCompany = gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001") // A company that exists
var fakeCompany = gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000999") // A company that doesn't exist

// func TestCompanyExists(t *testing.T) {
// 	prepareTestDatabase()

// 	t.Run("Company should exist", func(t *testing.T) {
// 		assert.True(t, adminGrant.CompanyExists(realCompany))
// 	})

// 	t.Run("Company should not exist", func(t *testing.T) {
// 		assert.False(t, adminGrant.CompanyExists(fakeCompany))
// 	})
// }

func TestGetManagerIDsByCompany(t *testing.T) {
	prepareTestDatabase()

	company1 := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")

	t.Run("Should return all managers", func(t *testing.T) {
		ids, _, err := usersRepo.GetManagerIDsByCompany(company1, nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 2)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(1)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		ids, pageInfo, err := usersRepo.GetManagerIDsByCompany(company1, &page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 1)
		assert.Equal(t, pageInfo, gentypes.PageInfo{Total: 2, Given: 1, Limit: limit})
	})

	t.Run("Should order", func(t *testing.T) {
		order := gentypes.OrderBy{Field: "first_name"}

		ids, _, err := usersRepo.GetManagerIDsByCompany(company1, nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, ids, 2)
		assert.Equal(t, "00000000-0000-0000-0000-000000000002", ids[0].String())
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
				ids, _, err := usersRepo.GetManagerIDsByCompany(company1, nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, ids, 1)
				assert.Equal(t, manager.UUID, ids[0])
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.ManagersFilter{Email: &email}
			managers, _, err := usersRepo.GetManagerIDsByCompany(company1, nil, &filter, nil)
			assert.Nil(t, err)
			require.Len(t, managers, 2)
		})
	})
}

func TestGetCompaniesByUUID(t *testing.T) {
	prepareTestDatabase()

	t.Run("Admin can get companies", func(t *testing.T) {
		comp, err := usersRepo.GetCompaniesByUUID([]gentypes.UUID{realCompany})
		assert.Nil(t, err)
		assert.Len(t, comp, 1)
		assert.Equal(t, realCompany, comp[0].UUID)
	})

	// TODO: Test for managers + delegates access to companies
}

func TestCompany(t *testing.T) {
	prepareTestDatabase()

	t.Run("Can get company", func(t *testing.T) {
		company, err := usersRepo.Company(realCompany)
		assert.Nil(t, err)
		assert.Equal(t, realCompany, company.UUID)
	})

	t.Run("Get non-existant company", func(t *testing.T) {
		company, err := usersRepo.Company(fakeCompany)
		assert.Equal(t, &errors.ErrCompanyNotFound, err)
		assert.Equal(t, models.Company{}, company)
	})

	// TODO: Test if manager can get own company
}

func TestGetCompanyUUIDs(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return all companies", func(t *testing.T) {
		ids, _, err := usersRepo.GetCompanyUUIDs(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 4)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		ids, pageInfo, err := usersRepo.GetCompanyUUIDs(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 2)
		assert.Equal(t, pageInfo, gentypes.PageInfo{Total: 4, Given: 2, Limit: limit})
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "name", Ascending: &asc}

		ids, _, err := usersRepo.GetCompanyUUIDs(nil, nil, &order)
		assert.Nil(t, err)
		require.Len(t, ids, 4)
		assert.Equal(t, "00000000-0000-0000-0000-000000000003", ids[0].String())
	})

	t.Run("Should filter", func(t *testing.T) {
		uuidString := "00000000-0000-0000-0000-000000000003"
		name := "ACME"
		approved := true

		filterTests := []struct {
			name    string
			filter  gentypes.CompanyFilter
			wantLen int
		}{
			{"UUID", gentypes.CompanyFilter{UUID: &uuidString}, 1},
			{"Name", gentypes.CompanyFilter{Name: &name}, 1},
			{"Approved", gentypes.CompanyFilter{Approved: &approved}, 3},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				ids, _, err := usersRepo.GetCompanyUUIDs(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, ids, test.wantLen)
			})
		}
	})
}

var newCompInput = gentypes.CreateCompanyInput{
	CompanyName:  "Big Cat House",
	AddressLine1: "64 Zoo Lane",
	AddressLine2: "",
	County:       "York",
	PostCode:     "YO108JD",
	Country:      "UK",
	ContactEmail: "email@email.com",
}

func TestCreateCompany(t *testing.T) {
	prepareTestDatabase()

	t.Run("Check company and address are created", func(t *testing.T) {
		company, err := usersRepo.CreateCompany(newCompInput)
		assert.Nil(t, err)
		assert.Equal(t, newCompInput.CompanyName, company.Name)
		assert.Equal(t, true, company.Approved)

		// check address
		addresses, err := usersRepo.GetAddressesByIDs([]uint{company.AddressID})
		assert.Nil(t, err)
		require.Len(t, addresses, 1)
		assert.Equal(t, newCompInput.AddressLine1, addresses[0].AddressLine1)
		assert.Equal(t, newCompInput.AddressLine2, addresses[0].AddressLine2)
		assert.Equal(t, company.AddressID, addresses[0].ID)
		assert.Equal(t, newCompInput.Country, addresses[0].Country)
		assert.Equal(t, newCompInput.PostCode, addresses[0].PostCode)
	})
}

func TestCreateCompanyRequest(t *testing.T) {
	prepareTestDatabase()

	managerInput := gentypes.CreateManagerInput{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "uinqad@tes.asd",
		JobTitle:  "Big Man",
		Telephone: "79865563351",
		Password:  "bigpass",
	}

	t.Run("Should create company, address, and manager", func(t *testing.T) {
		err := usersRepo.CreateCompanyRequest(newCompInput, managerInput)
		assert.Nil(t, err)

		// get the latest company
		ids, _, _ := usersRepo.GetCompanyUUIDs(nil, nil, nil)
		company, _ := usersRepo.Company(ids[len(ids)-1])
		assert.Equal(t, false, company.Approved)
		assert.Equal(t, newCompInput.CompanyName, company.Name)

		// check address
		addresses, err := usersRepo.GetAddressesByIDs([]uint{company.AddressID})
		assert.Nil(t, err)
		require.Len(t, addresses, 1, "Should have created an address")
		assert.Equal(t, newCompInput.AddressLine1, addresses[0].AddressLine1)
		assert.Equal(t, newCompInput.AddressLine2, addresses[0].AddressLine2)
		assert.Equal(t, newCompInput.PostCode, addresses[0].PostCode)
		assert.Equal(t, newCompInput.Country, addresses[0].Country)
		assert.Equal(t, newCompInput.County, addresses[0].County)

		// check it created a manager
		managers, _, err := usersRepo.GetManagerIDsByCompany(company.UUID, nil, nil, nil)
		require.Nil(t, err)
		require.Len(t, managers, 1, "Should return the new manager")
		manager, err := usersRepo.Manager(managers[0])
		assert.Nil(t, err)
		assert.Equal(t, manager.Email, managerInput.Email)
		assert.Equal(t, manager.JobTitle, managerInput.JobTitle)
		assert.Equal(t, manager.Telephone, managerInput.Telephone)
	})
}

func TestApproveCompany(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuid    string
		want    bool
		wantErr interface{}
	}{
		{
			"company must exist",
			"00000000-0000-0000-0000-000000000000",
			false,
			&errors.ErrCompanyNotFound,
		},
		{
			"Should set approved to true",
			"00000000-0000-0000-0000-000000000004",
			true,
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp, err := usersRepo.ApproveCompany(gentypes.MustParseToUUID(test.uuid))
			assert.Equal(t, test.wantErr, err)
			if err == nil {
				assert.Equal(t, gentypes.MustParseToUUID(test.uuid), comp.UUID)
				assert.Equal(t, test.want, comp.Approved)
			}
		})
	}
}
