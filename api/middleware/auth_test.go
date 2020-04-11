package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestGetAdminAccessToken(t *testing.T) {
	t.Run("Correct info", func(t *testing.T) {
		token, err := middleware.GetAdminAccessToken("rodger@van.com", "iamasuperadmin")
		assert.Nil(t, err)

		// should return a valid token
		claim, err := auth.ValidateToken(token)
		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000004", claim.UUID)
		assert.Equal(t, auth.AdminRole, claim.Role)
	})

	t.Run("Bad password", func(t *testing.T) {
		token, err := middleware.GetAdminAccessToken("rodger@van.com", "notmypassword")
		assert.Equal(t, "", token)
		assert.Equal(t, &errors.ErrAuthFailed, err)
	})

	t.Run("Bad email", func(t *testing.T) {
		token, err := middleware.GetAdminAccessToken("idonot@exist.com", "notmypassword")
		assert.Equal(t, "", token)
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})
}

func TestGetManagerAccessToken(t *testing.T) {
	t.Run("Correct info", func(t *testing.T) {
		token, err := middleware.GetManagerAccessToken("man@managers.com", "iamamanager")
		assert.Nil(t, err)

		// should return a valid token
		claim, err := auth.ValidateToken(token)
		assert.Nil(t, err)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", claim.UUID)
		assert.Equal(t, auth.ManagerRole, claim.Role)
	})

	t.Run("Bad password", func(t *testing.T) {
		token, err := middleware.GetManagerAccessToken("man@managers.com", "notmypassword")
		assert.Equal(t, "", token)
		assert.Equal(t, &errors.ErrAuthFailed, err)
	})

	t.Run("Bad email", func(t *testing.T) {
		token, err := middleware.GetManagerAccessToken("idonot@exist.com", "notmypassword")
		assert.Equal(t, "", token)
		assert.Equal(t, &errors.ErrUserNotFound, err)
	})
}
