package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var adminRepository = middleware.NewAdminRepository(&logging.Logger{})

func TestCreateAdmin(t *testing.T) {
	prepareTestDatabase()
	newAdmin := gentypes.CreateAdminInput{
		Email:     "admi1n@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}

	t.Run("Check Admin is created", func(t *testing.T) {
		createdAdmin, err := adminRepository.CreateAdmin(newAdmin)
		assert.Nil(t, err)
		assert.Equal(t, newAdmin.Email, createdAdmin.Email)
		assert.Equal(t, newAdmin.FirstName, createdAdmin.FirstName)
		assert.Equal(t, newAdmin.LastName, createdAdmin.LastName)
	})

	newAdmin.Email = "email2@admin.com"
	t.Run("Check email must be unique", func(t *testing.T) {
		_, err := adminRepository.CreateAdmin(newAdmin)
		assert.Nil(t, err)
		a, err := adminRepository.CreateAdmin(newAdmin)
		assert.Equal(t, models.Admin{}, a, "should return blank")
		assert.Equal(t, &errors.ErrUserExists, err)
	})
}

func TestUpdateAdmin(t *testing.T) {
	prepareTestDatabase()

	updateAdmin := gentypes.UpdateAdminInput{
		UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000001999"), // non-existant uuid
	}

	t.Run("Admin must exist", func(t *testing.T) {
		_, err := adminRepository.UpdateAdmin(updateAdmin)
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
		updatedAdmin, err := adminRepository.UpdateAdmin(updateAdmin)
		assert.Nil(t, err)
		assert.Equal(t, testAdmin.Email, updatedAdmin.Email)
		assert.Equal(t, testAdmin.UUID, updatedAdmin.UUID)
		assert.Equal(t, testAdmin.FirstName, updatedAdmin.FirstName)
		assert.Equal(t, testAdmin.LastName, updatedAdmin.LastName)
	})
}

func TestDeleteAdmin(t *testing.T) {
	prepareTestDatabase()

	t.Run("Admin must exist", func(t *testing.T) {
		_, err := adminRepository.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000999"))
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})

	t.Run("Check it deletes the admin", func(t *testing.T) {
		out, err := adminRepository.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Nil(t, err)
		assert.True(t, out)

		// trying to delete again then causes not found
		_, err = adminRepository.DeleteAdmin(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"))
		assert.Equal(t, &errors.ErrAdminNotFound, err)
	})
}

func TestPageAdmins(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return all admins", func(t *testing.T) {
		admins, _, err := adminRepository.PageAdmins(nil, nil)
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
			filter gentypes.AdminFilter
		}{
			{"Email", gentypes.AdminFilter{Email: admin.Email}},
			{"FirstName", gentypes.AdminFilter{Name: "Rodger"}},
			{"LastName", gentypes.AdminFilter{Name: "Van"}},
			{"First and Last", gentypes.AdminFilter{Name: "Rodger Van"}},
			{"Full", gentypes.AdminFilter{Name: "Rodger Van", Email: admin.Email}},
		}

		for _, test := range filterTests {
			t.Run(test.name, func(t *testing.T) {
				admins, _, err := adminRepository.PageAdmins(nil, &test.filter)
				assert.Nil(t, err)
				require.Len(t, admins, 1)
				assert.Equal(t, admin.UUID, admins[0].UUID)
				assert.Equal(t, admin.Email, admins[0].Email)
			})
		}

		t.Run("return mutiple", func(t *testing.T) {
			filter := gentypes.AdminFilter{Email: ".com"}
			admins, _, err := adminRepository.PageAdmins(nil, &filter)
			assert.Nil(t, err)
			require.Len(t, admins, 4)
		})
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		admins, pageinfo, err := adminRepository.PageAdmins(&page, nil)
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

func TestAdmins(t *testing.T) {
	prepareTestDatabase()

	uuids := []gentypes.UUID{
		gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
	}

	t.Run("Should get admins", func(t *testing.T) {
		admins, err := adminRepository.Admins(uuids)
		assert.Nil(t, err)
		assert.Len(t, admins, 2)
	})

	t.Run("Should return errs", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			admins, err := adminRepository.Admins([]gentypes.UUID{gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")})
			assert.Equal(t, &errors.ErrNotFound, err)
			assert.Len(t, admins, 0)
		})
	})
}
