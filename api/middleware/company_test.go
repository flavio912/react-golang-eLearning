package middleware_test

import (
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var realCompany = "00000000-0000-0000-0000-000000000001" // A company that exists
var fakeCompany = "00000000-0000-0000-0000-000000000999" // A company that doesn't exist

func TestCompanyExists(t *testing.T) {
	prepareTestDatabase()

	grant := &middleware.Grant{auth.UserClaims{}, true, true, true}
	t.Run("Company should exist", func(t *testing.T) {
		id, _ := uuid.Parse(realCompany)
		assert.True(t, grant.CompanyExists(id))
	})

	t.Run("Company should not exist", func(t *testing.T) {
		id, _ := uuid.Parse(fakeCompany)
		assert.False(t, grant.CompanyExists(id))
	})
}

func TestIsCompanyDelegate(t *testing.T) {
	assert.True(t, delegateGrant.IsCompanyDelegate("00000000-0000-0000-0000-000000000001"))
	assert.False(t, delegateGrant.IsCompanyDelegate("00000000-0000-0000-0000-000000000002"))
	// should only happen for delegates
	assert.False(t, managerGrant.IsCompanyDelegate("00000000-0000-0000-0000-000000000001"))
	assert.False(t, adminGrant.IsCompanyDelegate(""))
}

func TestManagesCompany(t *testing.T) {
	tests := []struct {
		name  string
		grant middleware.Grant
		uuid  string
		want  bool
	}{
		{
			"Admin always true",
			adminGrant,
			"",
			true,
		},
		{
			"Delegate always false",
			delegateGrant,
			"",
			false,
		},
		{
			"Manager must be part of company",
			managerGrant,
			"00000000-0000-0000-0000-000000000001",
			true,
		},
		{
			"Should fail if not managers company",
			managerGrant,
			"00000000-0000-0000-0000-000000000002",
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

	company1 := "00000000-0000-0000-0000-000000000001"
	t.Run("Must be admin", func(t *testing.T) {
		grant := middleware.Grant{auth.UserClaims{}, false, true, true}
		ids, _, err := grant.GetManagerIDsByCompany("", nil, nil, nil)
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
			User: gentypes.User{
				UUID:      uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Email:     "man@managers.com",
				FirstName: "Manager",
				LastName:  "Man",
				Telephone: "7912938287",
				JobTitle:  "In Charge",
			},
			CompanyID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Admin can get companies", func(t *testing.T) {
		comp, err := grant.GetCompaniesByUUID([]string{realCompany})
		assert.Nil(t, err)
		assert.Len(t, comp, 1)
		assert.Equal(t, realCompany, comp[0].UUID.String())
	})

	// TODO: Test for managers + delegates access to companies
}

func TestGetCompanyByUUID(t *testing.T) {
	prepareTestDatabase()

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Admin can get company", func(t *testing.T) {
		company, err := grant.GetCompanyByUUID(realCompany)
		assert.Nil(t, err)
		assert.Equal(t, realCompany, company.UUID.String())
	})

	t.Run("Get non-existant company", func(t *testing.T) {
		company, err := grant.GetCompanyByUUID(fakeCompany)
		assert.Equal(t, &errors.ErrCompanyNotFound, err)
		assert.Equal(t, gentypes.Company{}, company)
	})

	// TODO: Test if manager can get own company
}

func TestGetCompanyUUIDs(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
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
		assert.Equal(t, "00000000-0000-0000-0000-000000000003", ids[0])
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
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
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

	managerInput := gentypes.AddManagerInput{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "uinqad@tes.asd",
		JobTitle:  "Big Man",
		Telephone: "79865563351",
		Password:  "bigpass",
	}

	t.Run("Should create company, address, and manager", func(t *testing.T) {
		err := middleware.CreateCompanyRequest(newCompInput, managerInput)
		assert.Nil(t, err)

		// get the latest company
		ids, _, _ := adminGrant.GetCompanyUUIDs(nil, nil, nil)
		company, _ := adminGrant.GetCompanyByUUID(ids[len(ids)-1])
		approved := false
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
		managers, _, err := adminGrant.GetManagerIDsByCompany(company.UUID.String(), nil, nil, nil)
		require.Nil(t, err)
		require.Len(t, managers, 1, "Should return the new manager")
		manager, err := adminGrant.GetManagerByUUID(managers[0].String())
		assert.Nil(t, err)
		assert.Equal(t, gentypes.Manager{
			User: gentypes.User{
				CreatedAt: manager.CreatedAt,
				UUID:      manager.UUID,
				Email:     managerInput.Email,
				FirstName: managerInput.FirstName,
				LastName:  managerInput.LastName,
				Telephone: managerInput.Telephone,
				JobTitle:  managerInput.JobTitle,
				LastLogin: manager.LastLogin,
			},
			ProfileImageURL: manager.ProfileImageURL,
			CompanyID:       company.UUID,
		}, manager)
	})
}
