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

func TestUpdateModule(t *testing.T) {
	lessonUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-00000000001")
	lesson2UUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000002")
	testUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")
	templateModuleUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}

	t.Run("Duplicates + updates template correctly", func(t *testing.T) {
		prepareTestDatabase()
		// Get module to check for changes after
		mod, err := grant.GetModuleByUUID(templateModuleUUID)
		assert.Nil(t, err)

		modItem := gentypes.CourseItem{
			Type: gentypes.ModuleType,
			UUID: templateModuleUUID,
			Items: []gentypes.ModuleItem{
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: lessonUUID,
				},
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: lesson2UUID,
				},
			},
		}

		updatedModule, err := grant.UpdateModuleStructure(database.GormDB, modItem, true)
		assert.Nil(t, err)
		assert.False(t, updatedModule.Template)
		assert.NotEqual(t, updatedModule.UUID, uuid.UUID{})
		assert.NotNil(t, updatedModule.TemplateID)
		assert.Equal(t, templateModuleUUID.String(), (*updatedModule.TemplateID).String())

		_, err = grant.UpdateModuleStructure(database.GormDB, modItem, true)
		assert.Nil(t, err)

		// Get structure
		structure, err := grant.GetModuleStructure(updatedModule.UUID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(structure.Items))
		assert.Equal(t, lessonUUID, structure.Items[0].UUID)
		assert.Equal(t, gentypes.LessonType, structure.Items[0].Type)
		assert.Equal(t, lesson2UUID, structure.Items[1].UUID)
		assert.Equal(t, gentypes.LessonType, structure.Items[1].Type)

		templateMod, err := grant.GetModuleByUUID(templateModuleUUID)
		assert.Nil(t, err)
		assert.Equal(t, mod, templateMod) // The template model shouldn't have changed
	})

	t.Run("Update template in place", func(t *testing.T) {
		prepareTestDatabase()
		modItem := gentypes.CourseItem{
			Type: gentypes.ModuleType,
			UUID: templateModuleUUID,
			Items: []gentypes.ModuleItem{
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: lessonUUID,
				},
				gentypes.ModuleItem{
					Type: gentypes.LessonType,
					UUID: lesson2UUID,
				},
				gentypes.ModuleItem{
					Type: gentypes.TestType,
					UUID: testUUID,
				},
			},
		}
		module, err := grant.UpdateModuleStructure(database.GormDB, modItem, false)
		assert.Nil(t, err)
		assert.Equal(t, module.UUID.String(), templateModuleUUID.String())

		structure, err := grant.GetModuleStructure(module.UUID)
		assert.Nil(t, err)
		assert.Equal(t, len(modItem.Items), len(structure.Items))
		for i, item := range modItem.Items {
			assert.Equal(t, item.Type, structure.Items[i].Type)
			assert.Equal(t, item.UUID, structure.Items[i].UUID)
		}
	})
}
