package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func TestCreateAdminUser(t *testing.T) {
	newAdmin := &models.Admin{
		Email:     "test@test.com",
		Password:  "test",
		FirstName: "FName",
		LastName:  "LName",
	}
	err := database.GormDB.Create(newAdmin)
	if err.Error != nil {
		t.Fatal(err.GetErrors())
	}

	admin := &models.Admin{}
	q := database.GormDB.Where("email = ?", "test@test.com").First(admin)
	if q.Error != nil {
		if q.RecordNotFound() {
			t.Error("GORM didn't create a user")
		}

		t.Errorf("GORM Errored:\n%#v", q.GetErrors())
	}
}

func TestAddAdmin(t *testing.T) {
	// fake grant
	grant := &middleware.Grant{auth.UserClaims{}, true, true, true}

	newAdmin := gentypes.AddAdminInput{
		Email:     "admi1n@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}

	t.Run("Must be admin to add user", func(t *testing.T) {
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
		_, err := nonAdminGrant.AddAdmin(newAdmin)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Check Admin is created", func(t *testing.T) {
		createdAdmin, err := grant.AddAdmin(newAdmin)
		assert.Nil(t, err)
		assert.Equal(t, gentypes.Admin{
			UUID:      createdAdmin.UUID,
			Email:     newAdmin.Email,
			FirstName: newAdmin.FirstName,
			LastName:  newAdmin.LastName,
		}, createdAdmin)
	})

	newAdmin.Email = "email2@admin.com"
	t.Run("Check email must be unique", func(t *testing.T) {
		_, err := grant.AddAdmin(newAdmin)
		assert.Nil(t, err)
		a, err := grant.AddAdmin(newAdmin)
		assert.Equal(t, gentypes.Admin{}, a, "should return blank")
		assert.Equal(t, &errors.ErrUserExists, err)
	})
}

func TestUpdateAdmin(t *testing.T) {
	grant := &middleware.Grant{auth.UserClaims{}, true, true, true}

	updateAdmin := gentypes.UpdateAdminInput{
		UUID: "00000000-0000-0000-0000-000000000000",
	}

	t.Run("Must be admin to update", func(t *testing.T) {
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
		_, err := nonAdminGrant.UpdateAdmin(updateAdmin)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Admin must exist", func(t *testing.T) {
		_, err := grant.UpdateAdmin(updateAdmin)
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})

	// TODO refactor out test fixtures
	newAdmin := gentypes.AddAdminInput{
		Email:     "admin@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}
	admin, err := grant.AddAdmin(newAdmin)
	require.Nil(t, err, "Failed to add admin user")
	updateAdmin.UUID = admin.UUID

	t.Run("Check it updates admin record", func(t *testing.T) {
		testAdmin := gentypes.Admin{
			UUID:      updateAdmin.UUID,
			FirstName: "NAME",
			LastName:  "LASTNAME",
			Email:     "email@email.com",
		}

		updateAdmin.FirstName = &testAdmin.FirstName
		updateAdmin.LastName = &testAdmin.LastName
		updateAdmin.Email = &testAdmin.Email

		updatedAdmin, err := grant.UpdateAdmin(updateAdmin)
		assert.Nil(t, err)
		assert.Equal(t, testAdmin, updatedAdmin)
	})
}

func TestDeleteAdmin(t *testing.T) {
	grant := &middleware.Grant{auth.UserClaims{}, true, true, true}

	t.Run("Must be admin to delete", func(t *testing.T) {
		nonAdminGrant := &middleware.Grant{auth.UserClaims{}, false, true, true}
		_, err := nonAdminGrant.DeleteAdmin("00000000-0000-0000-0000-000000000000")
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Admin must exist", func(t *testing.T) {
		_, err := grant.DeleteAdmin("00000000-0000-0000-0000-000000000000")
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})

	// TODO refactor out test fixtures
	newAdmin := gentypes.AddAdminInput{
		Email:     "admin@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}
	admin, err := grant.AddAdmin(newAdmin)
	require.Nil(t, err, "Failed to add admin user")

	t.Run("Check it deletes the admin", func(t *testing.T) {
		out, err := grant.DeleteAdmin(admin.UUID)
		assert.Nil(t, err)
		assert.True(t, out)
		// trying to delete again then causes not found
		_, err = grant.DeleteAdmin(admin.UUID)
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})
}
