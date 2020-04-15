package middleware_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var templateModuleUUID = "00000000-0000-0000-0000-000000000001"

func TestUpdateModule(t *testing.T) {
	prepareTestDatabase()

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Duplicates + updates template correctly", func(t *testing.T) {
		// Get module to check for changes after
		mod, err := grant.GetModuleByUUID(templateModuleUUID)
		assert.Nil(t, err)

		modItem := gentypes.CourseItem{
			Type: gentypes.ModuleType,
			UUID: templateModuleUUID,
			Items: []gentypes.ModuleItem{
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: "00000000-0000-0000-0000-000000000001",
				},
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: "00000000-0000-0000-0000-000000000002",
				},
			},
		}

		updatedModule, err := grant.UpdateModuleStructure(database.GormDB, modItem, true)
		assert.Nil(t, err)
		assert.False(t, updatedModule.Template)
		assert.NotEqual(t, updatedModule.UUID, uuid.UUID{})
		assert.NotNil(t, updatedModule.TemplateID)
		assert.Equal(t, templateModuleUUID, (*updatedModule.TemplateID).String())

		_, err = grant.UpdateModuleStructure(database.GormDB, modItem, true)
		assert.Nil(t, err)

		// Get structure
		structure, err := grant.GetModuleStructure(updatedModule.UUID.String())
		assert.Nil(t, err)
		assert.Equal(t, 2, len(structure.Items))
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", structure.Items[0].UUID)
		assert.Equal(t, gentypes.LessonType, structure.Items[0].Type)
		assert.Equal(t, "00000000-0000-0000-0000-000000000002", structure.Items[1].UUID)
		assert.Equal(t, gentypes.LessonType, structure.Items[1].Type)

		templateMod, err := grant.GetModuleByUUID(templateModuleUUID)
		assert.Nil(t, err)
		assert.Equal(t, mod, templateMod) // The template model shouldn't have changed
	})

	t.Run("Updates template in place", func(t *testing.T) {

	})
}
