package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"

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
