package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestGetAddressesByIDs(t *testing.T) {
	prepareTestDatabase()

	t.Run("Check only admin can access", func(t *testing.T) {
		nonAdminGrant := middleware.Grant{auth.UserClaims{}, false, true, true}
		addresses, err := nonAdminGrant.GetAddressesByIDs([]uint{})
		assert.Len(t, addresses, 0)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	tests := []struct {
		name   string
		ids    []uint
		retLen int
	}{
		{"Empty IDs", []uint{}, 0},
		{"One ID", []uint{0}, 1},
		{"Two IDs", []uint{0, 1}, 2},
	}
	grant := middleware.Grant{auth.UserClaims{}, true, false, false}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addresses, err := grant.GetAddressesByIDs(test.ids)
			assert.Nil(t, err)
			assert.Len(t, addresses, test.retLen)
		})
	}
}
