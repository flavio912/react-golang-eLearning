package middleware_test

import (
	"context"
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
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

func TestIsCompanyDelegate(t *testing.T) {
	assert.True(t, delegateGrant.IsCompanyDelegate(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")))
	assert.False(t, delegateGrant.IsCompanyDelegate(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002")))
	// should only happen for delegates
	assert.False(t, managerGrant.IsCompanyDelegate(realCompany))
	assert.False(t, adminGrant.IsCompanyDelegate(realCompany))
}

func TestManagesCompany(t *testing.T) {
	tests := []struct {
		name  string
		grant middleware.Grant
		uuid  gentypes.UUID
		want  bool
	}{
		{
			"Admin always true",
			adminGrant,
			uuidZero,
			true,
		},
		{
			"Delegate always false",
			delegateGrant,
			uuidZero,
			false,
		},
		{
			"Manager must be part of company",
			managerGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			true,
		},
		{
			"Should fail if not managers company",
			managerGrant,
			gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := test.grant.ManagesCompany(test.uuid)
			assert.Equal(t, test.want, ret)
		})
	}
}

func TestGetManagerIDsByCompany(t *testing.T) {
	prepareTestDatabase()

	company1 := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
	t.Run("Must be admin", func(t *testing.T) {
		ids, _, err := nonAdminGrant.GetManagerIDsByCompany(uuidZero, nil, nil, nil)
		assert.Len(t, ids, 0)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all managers", func(t *testing.T) {
		ids, _, err := adminGrant.GetManagerIDsByCompany(company1, nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 2)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(1)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		ids, pageInfo, err := adminGrant.GetManagerIDsByCompany(company1, &page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 1)
		assert.Equal(t, pageInfo, gentypes.PageInfo{Total: 2, Given: 1, Limit: limit})
	})

	t.Run("Should order", func(t *testing.T) {
		order := gentypes.OrderBy{Field: "first_name"}

		ids, _, err := adminGrant.GetManagerIDsByCompany(company1, nil, nil, &order)
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
				ids, _, err := adminGrant.GetManagerIDsByCompany(company1, nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, ids, 1)
				assert.Equal(t, manager.UUID, ids[0])
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			email := ".com"
			filter := gentypes.ManagersFilter{Email: &email}
			managers, _, err := adminGrant.GetManagerIDsByCompany(company1, nil, &filter, nil)
			assert.Nil(t, err)
			require.Len(t, managers, 2)
		})
	})
}

func TestGetCompaniesByUUID(t *testing.T) {
	prepareTestDatabase()

	t.Run("Admin can get companies", func(t *testing.T) {
		comp, err := adminGrant.GetCompaniesByUUID([]gentypes.UUID{realCompany})
		assert.Nil(t, err)
		assert.Len(t, comp, 1)
		assert.Equal(t, realCompany, comp[0].UUID)
	})

	// TODO: Test for managers + delegates access to companies
}

func TestCompany(t *testing.T) {
	prepareTestDatabase()

	t.Run("Admin can get company", func(t *testing.T) {
		company, err := adminGrant.Company(realCompany)
		assert.Nil(t, err)
		assert.Equal(t, realCompany, company.UUID)
	})

	t.Run("Get non-existant company", func(t *testing.T) {
		company, err := adminGrant.Company(fakeCompany)
		assert.Equal(t, &errors.ErrCompanyNotFound, err)
		assert.Equal(t, models.Company{}, company)
	})

	// TODO: Test if manager can get own company
}

func TestGetCompanyUUIDs(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		_, _, err := nonAdminGrant.GetCompanyUUIDs(nil, nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all companies", func(t *testing.T) {
		ids, _, err := adminGrant.GetCompanyUUIDs(nil, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 4)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		ids, pageInfo, err := adminGrant.GetCompanyUUIDs(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, ids, 2)
		assert.Equal(t, pageInfo, gentypes.PageInfo{Total: 4, Given: 2, Limit: limit})
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "name", Ascending: &asc}

		ids, _, err := adminGrant.GetCompanyUUIDs(nil, nil, &order)
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
				ids, _, err := adminGrant.GetCompanyUUIDs(nil, &test.filter, nil)
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
}

func TestCreateCompany(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		_, err := nonAdminGrant.CreateCompany(gentypes.CreateCompanyInput{})
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	// TODO
	// t.Run("Must validate input",  {
	// })

	approved := true

	t.Run("Check company and address are created", func(t *testing.T) {
		company, err := adminGrant.CreateCompany(newCompInput)
		assert.Nil(t, err)
		assert.Equal(t, gentypes.Company{
			UUID:      company.UUID,
			CreatedAt: company.CreatedAt,
			Name:      newCompInput.CompanyName,
			Approved:  &approved,
			AddressID: company.AddressID,
		}, company)

		// check address
		addresses, err := adminGrant.GetAddressesByIDs([]uint{company.AddressID})
		assert.Nil(t, err)
		require.Len(t, addresses, 1)
		assert.Equal(t, gentypes.Address{
			ID:           company.AddressID,
			AddressLine1: newCompInput.AddressLine1,
			AddressLine2: newCompInput.AddressLine2,
			County:       newCompInput.County,
			PostCode:     newCompInput.PostCode,
			Country:      newCompInput.Country,
		}, addresses[0])
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
		err := middleware.CreateCompanyRequest(context.Background(), newCompInput, managerInput)
		assert.Nil(t, err)

		// get the latest company
		ids, _, _ := adminGrant.GetCompanyUUIDs(nil, nil, nil)
		company, _ := adminGrant.Company(ids[len(ids)-1])
		assert.Equal(t, false, company.Approved)
		assert.Equal(t, newCompInput.CompanyName, company.Name)

		// check address
		addresses, err := adminGrant.GetAddressesByIDs([]uint{company.AddressID})
		assert.Nil(t, err)
		require.Len(t, addresses, 1, "Should have created an address")
		assert.Equal(t, gentypes.Address{
			ID:           company.AddressID,
			AddressLine1: newCompInput.AddressLine1,
			AddressLine2: newCompInput.AddressLine2,
			County:       newCompInput.County,
			PostCode:     newCompInput.PostCode,
			Country:      newCompInput.Country,
		}, addresses[0])

		// check it created a manager
		managers, _, err := adminGrant.GetManagerIDsByCompany(company.UUID, nil, nil, nil)
		require.Nil(t, err)
		require.Len(t, managers, 1, "Should return the new manager")
		manager, err := adminGrant.Manager(managers[0])
		assert.Nil(t, err)
		assert.Equal(t, gentypes.Manager{
			CreatedAt:       manager.CreatedAt,
			UUID:            manager.UUID,
			FirstName:       managerInput.FirstName,
			LastName:        managerInput.LastName,
			Telephone:       managerInput.Telephone,
			JobTitle:        managerInput.JobTitle,
			LastLogin:       manager.LastLogin,
			Email:           managerInput.Email,
			ProfileImageURL: manager.ProfileImageURL,
			CompanyUUID:     company.UUID,
		}, manager)
	})
}

func TestApproveCompany(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuid    string
		want    bool
		wantErr interface{}
	}{
		{
			"Must be admin",
			nonAdminGrant,
			"00000000-0000-0000-0000-000000000000",
			false,
			&errors.ErrUnauthorized,
		},
		{
			"company must exist",
			adminGrant,
			"00000000-0000-0000-0000-000000000000",
			false,
			&errors.ErrCompanyNotFound,
		},
		{
			"Should set approved to true",
			adminGrant,
			"00000000-0000-0000-0000-000000000004",
			true,
			nil,
		},
		{
			"Cannot be public",
			publicGrant,
			"00000000-0000-0000-0000-000000000004",
			false,
			&errors.ErrUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comp, err := test.grant.ApproveCompany(gentypes.MustParseToUUID(test.uuid))
			assert.Equal(t, test.wantErr, err)
			if err == nil {
				assert.Equal(t, gentypes.MustParseToUUID(test.uuid), comp.UUID)
				assert.Equal(t, &test.want, comp.Approved)
			}
		})
	}
}
