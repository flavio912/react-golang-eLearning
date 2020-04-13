package middleware_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestAddManager(t *testing.T) {
	prepareTestDatabase()
	fakeUUID := "this is not a uuid"

	tests := []struct {
		name    string
		grant   middleware.Grant
		wantErr interface{}
		want    gentypes.Manager
		input   gentypes.AddManagerInput
	}{
		{
			"Users cannot create",
			delegateGrant,
			&errors.ErrUnauthorized,
			gentypes.Manager{},
			gentypes.AddManagerInput{},
		},
		{
			"Admin must supply company uuid",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Manager{},
			gentypes.AddManagerInput{},
		},
		{
			"Admin must supply valid company uuid",
			adminGrant,
			&errors.ErrUUIDInvalid,
			gentypes.Manager{},
			gentypes.AddManagerInput{
				CompanyUUID: &fakeUUID,
			},
		},
		{
			"Admin supplied company must exist",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Manager{},
			gentypes.AddManagerInput{
				CompanyUUID: &uuidZero,
			},
		},
		{
			"Should use manager's company",
			managerGrant,
			nil,
			gentypes.Manager{CompanyID: uuid.MustParse(managerGrant.Claims.Company)},
			gentypes.AddManagerInput{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.AddManager(test.input)
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
		uuid    string
		wantErr interface{}
	}{
		{
			"Delegates cannot get managers",
			delegateGrant,
			"00000000-0000-0000-0000-000000000001",
			&errors.ErrUnauthorized,
		},
		{
			"Managers cannot see other managers",
			managerGrant,
			"00000000-0000-0000-0000-000000000002", // different to managerGrant.Claims.UUID
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
			"00000000-0000-0000-0000-000000000001",
			nil,
		},
		{
			"UUID must be valid",
			adminGrant,
			"this is not a uuid",
			&errors.ErrUUIDInvalid,
		},
		{
			"Should return ErrNotFound if not found",
			adminGrant,
			"00000000-0000-0000-0000-000000000000", // does not exist
			&errors.ErrNotFound,
		},
	}

	// these only check the uuid returned is correct
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.GetManagerByUUID(test.uuid)
			assert.Equal(t, test.wantErr, err)
			if test.wantErr == nil {
				assert.Equal(t, test.uuid, m.UUID.String())
			} else {
				// should return a blank manager if it errors
				assert.Equal(t, gentypes.Manager{}, m)
			}
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
			managerGrant.Claims.UUID,
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
		uuid    string
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
			"00000000-0000-0000-0000-000000000002",
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
