package course_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
)

func TestUpdateModule(t *testing.T) {
	lessonUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-00000000001")
	lesson2UUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000002")
	testUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")
	moduleUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")

	t.Run("Updates module structure correctly", func(t *testing.T) {
		prepareTestDatabase()
		modItem := []gentypes.ModuleItem{
			gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: lessonUUID,
			},
			gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: lesson2UUID,
			},
			gentypes.ModuleItem{
				Type: gentypes.ModuleTest,
				UUID: testUUID,
			},
		}

		updatedModule, err := courseRepo.UpdateModuleStructure(database.GormDB, moduleUUID, modItem)
		assert.Nil(t, err)
		assert.NotEqual(t, updatedModule.UUID, uuid.UUID{})

		_, err = courseRepo.UpdateModuleStructure(database.GormDB, moduleUUID, modItem)
		assert.Nil(t, err)

		// Get structure
		structure, err := courseRepo.GetModuleStructure(moduleUUID)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(structure))
		assert.Equal(t, lessonUUID, structure[0].UUID)
		assert.Equal(t, gentypes.ModuleLesson, structure[0].Type)
		assert.Equal(t, lesson2UUID, structure[1].UUID)
		assert.Equal(t, gentypes.ModuleLesson, structure[1].Type)
		assert.Equal(t, testUUID, structure[2].UUID)
		assert.Equal(t, gentypes.ModuleTest, structure[2].Type)
	})

	t.Run("Update template in place", func(t *testing.T) {
		prepareTestDatabase()
		modItems := []gentypes.ModuleItem{
			gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: lessonUUID,
			},
			gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: lesson2UUID,
			},
			gentypes.ModuleItem{
				Type: gentypes.ModuleTest,
				UUID: testUUID,
			},
		}
		module, err := courseRepo.UpdateModuleStructure(database.GormDB, moduleUUID, modItems)
		assert.Nil(t, err)
		assert.Equal(t, module.UUID.String(), moduleUUID.String())

		structure, err := courseRepo.GetModuleStructure(module.UUID)
		assert.Nil(t, err)
		assert.Equal(t, len(modItems), len(structure))
		for i, item := range modItems {
			assert.Equal(t, item.Type, structure[i].Type)
			assert.Equal(t, item.UUID, structure[i].UUID)
		}
	})
}
