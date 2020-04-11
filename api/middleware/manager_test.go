package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestGrant_AddManager(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		wantErr interface{}
		want    gentypes.Manager
		input   gentypes.AddManagerInput
	}{
		{
			"Users cannot create",
			userGrant,
			&errors.ErrUnauthorized,
			gentypes.Manager{},
			gentypes.AddManagerInput{},
		},
		{
			"Admin must supply uuid",
			adminGrant,
			&errors.ErrCompanyNotFound,
			gentypes.Manager{},
			gentypes.AddManagerInput{},
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := test.grant.AddManager(test.input)
			assert.Equal(t, test.want, m)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
