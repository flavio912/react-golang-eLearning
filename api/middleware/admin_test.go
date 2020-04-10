package middleware_test

import (
	"testing"

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
	// t.Logf("created admin: %v", admin)
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

	t.Run("Check Admin is created", func(t *testing.T) {
		createdAdmin, err := grant.AddAdmin(newAdmin)
		require.Nil(t, err)
		require.Equal(t, gentypes.Admin{
			UUID:      createdAdmin.UUID,
			Email:     newAdmin.Email,
			FirstName: newAdmin.FirstName,
			LastName:  newAdmin.LastName,
		}, createdAdmin)
	})

	newAdmin.Email = "email2@admin.com"
	t.Run("Check email must be unique", func(t *testing.T) {
		_, err := grant.AddAdmin(newAdmin)
		require.Nil(t, err)
		a, err := grant.AddAdmin(newAdmin)
		require.Equal(t, gentypes.Admin{}, a, "should return blank")
		require.Equal(t, err, &errors.ErrUserExists)
	})
}
