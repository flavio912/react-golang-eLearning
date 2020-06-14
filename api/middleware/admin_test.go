package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestCreateAdmin(t *testing.T) {
	prepareTestDatabase()

	newAdmin := gentypes.CreateAdminInput{
		Email:     "admi1n@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}

	t.Run("Must be admin to add user", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&nonAdminGrant)
		_, err := adminFuncs.CreateAdmin(newAdmin)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Check Admin is created", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		createdAdmin, err := adminFuncs.CreateAdmin(newAdmin)
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
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		_, err := adminFuncs.CreateAdmin(newAdmin)
		assert.Nil(t, err)
		a, err := adminFuncs.CreateAdmin(newAdmin)
		assert.Equal(t, gentypes.Admin{}, a, "should return blank")
		assert.Equal(t, &errors.ErrUserExists, err)
	})
}

func TestUpdateAdmin(t *testing.T) {
	prepareTestDatabase()

	updateAdmin := gentypes.UpdateAdminInput{
		UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000001999"), // non-existant uuid
	}

	t.Run("Must be admin to update", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&nonAdminGrant)
		_, err := adminFuncs.UpdateAdmin(updateAdmin)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Admin must exist", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		_, err := adminFuncs.UpdateAdmin(updateAdmin)
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})

	t.Run("Check it updates admin record", func(t *testing.T) {
		testAdmin := gentypes.Admin{
			UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			FirstName: "NAME",
			LastName:  "LASTNAME",
			Email:     "email@email.com",
		}

		updateAdmin.UUID = testAdmin.UUID
		updateAdmin.FirstName = &testAdmin.FirstName
		updateAdmin.LastName = &testAdmin.LastName
		updateAdmin.Email = &testAdmin.Email

		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		updatedAdmin, err := adminFuncs.UpdateAdmin(updateAdmin)
		assert.Nil(t, err)
		assert.Equal(t, testAdmin, updatedAdmin)
	})
}

func TestDeleteAdmin(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin to delete", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&nonAdminGrant)
		_, err := adminFuncs.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Admin must exist", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		_, err := adminFuncs.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000999"))
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})

	t.Run("Check it deletes the admin", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		out, err := adminFuncs.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Nil(t, err)
		assert.True(t, out)

		// trying to delete again then causes not found
		_, err = adminFuncs.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})
}

func TestGetAdmins(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&nonAdminGrant)
		_, _, err := adminFuncs.GetAdmins(nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Should return all admins", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		admins, _, err := adminFuncs.GetAdmins(nil, nil)
		assert.Nil(t, err)
		assert.Len(t, admins, 4)
	})

	t.Run("Should filter", func(t *testing.T) {
		admin := gentypes.Admin{
			UUID:      gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000004"),
			Email:     "rodger@van.com",
			FirstName: "Rodger",
			LastName:  "Van",
		}

		filterTests := []struct {
			name   string
			filter middleware.AdminFilter
		}{
			{"Email", middleware.AdminFilter{Email: admin.Email}},
			{"FirstName", middleware.AdminFilter{Name: "Rodger"}},
			{"LastName", middleware.AdminFilter{Name: "Van"}},
			{"First and Last", middleware.AdminFilter{Name: "Rodger Van"}},
			{"Full", middleware.AdminFilter{Name: "Rodger Van", Email: admin.Email}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
				admins, _, err := adminFuncs.GetAdmins(nil, &test.filter)
				assert.Nil(t, err)
				require.Len(t, admins, 1)
				assert.Equal(t, admin, admins[0])
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			filter := middleware.AdminFilter{Email: ".com"}
			adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
			admins, _, err := adminFuncs.GetAdmins(nil, &filter)
			assert.Nil(t, err)
			require.Len(t, admins, 4)
		})
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		admins, pageinfo, err := adminFuncs.GetAdmins(&page, nil)
		assert.Nil(t, err)
		assert.Len(t, admins, 2)
		assert.Equal(t, gentypes.PageInfo{
			Total:  4,
			Offset: 0,
			Limit:  2,
			Given:  2,
		}, pageinfo)
	})
}

func TestGetAdminsByUUID(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&nonAdminGrant)
		_, err := adminFuncs.GetAdminsByUUID([]string{""})
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	uuids := []string{
		"00000000-0000-0000-0000-000000000001",
		"00000000-0000-0000-0000-000000000002",
	}

	t.Run("Should get admins", func(t *testing.T) {
		adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
		admins, err := adminFuncs.GetAdminsByUUID(uuids)
		assert.Nil(t, err)
		assert.Len(t, admins, 2)
	})

	t.Run("Should return errs", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
			admins, err := adminFuncs.GetAdminsByUUID([]string{"00000000-0000-0000-0000-000000000000"})
			assert.Equal(t, &errors.ErrNotFound, err)
			assert.Len(t, admins, 0)
		})
		t.Run("err while handling", func(t *testing.T) {
			// fake non validated uuid will cause db to bug out
			adminFuncs, _ := middleware.NewAdminRepository(&adminGrant)
			admins, err := adminFuncs.GetAdminsByUUID([]string{"asdfasdf-asdfasdf"})
			assert.Equal(t, &errors.ErrWhileHandling, err)
			assert.Len(t, admins, 0)
		})
	})
}
