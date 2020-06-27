package course_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
)

func TestUpdateModuleStructure(t *testing.T) {
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

func TestCreateModule(t *testing.T) {
	inputs := []struct {
		Name    string
		Input   course.CreateModuleInput
		WantErr error
	}{
		{
			Name: "Create without tags",
			Input: course.CreateModuleInput{
				Name:         "Cheesecake",
				Description:  "pies",
				Transcript:   "I like cakes",
				VoiceoverKey: helpers.StringPointer("/places/orange.mp3"),
				BannerKey:    helpers.StringPointer("/images/banner.png"),
				Video: &course.VideoInput{
					VideoType: gentypes.WistiaVideo,
					VideoURL:  "http://video.com/video",
				},
			},
			WantErr: nil,
		},
		{
			Name: "Create without video",
			Input: course.CreateModuleInput{
				Name:         "Cheesecake",
				Description:  "pies",
				Transcript:   "I like cakes",
				VoiceoverKey: helpers.StringPointer("/places/orange.mp3"),
				BannerKey:    helpers.StringPointer("/images/banner.png"),
			},
			WantErr: nil,
		},
		{
			Name: "Create with syllabus",
			Input: course.CreateModuleInput{
				Name:         "Cheesecake",
				Description:  "pies",
				Transcript:   "I like cakes",
				VoiceoverKey: helpers.StringPointer("/places/orange.mp3"),
				BannerKey:    helpers.StringPointer("/images/banner.png"),
				Syllabus: &[]gentypes.ModuleItem{
					gentypes.ModuleItem{
						Type: gentypes.ModuleLesson,
						UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003"),
					},
					gentypes.ModuleItem{
						Type: gentypes.ModuleTest,
						UUID: gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9"),
					},
				},
			},
			WantErr: nil,
		},
		{
			Name: "Create with syllabus items that don't exist",
			Input: course.CreateModuleInput{
				Name:         "Cheesecake",
				Description:  "pies",
				Transcript:   "I like cakes",
				VoiceoverKey: helpers.StringPointer("/places/orange.mp3"),
				BannerKey:    helpers.StringPointer("/images/banner.png"),
				Syllabus: &[]gentypes.ModuleItem{
					gentypes.ModuleItem{
						Type: gentypes.ModuleLesson,
						UUID: gentypes.MustParseToUUID("44262d13-fd9d-4235-8691-76b58a8375ad"), // Doesn't exist
					},
					gentypes.ModuleItem{
						Type: gentypes.ModuleTest,
						UUID: gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9"),
					},
				},
			},
			WantErr: nil,
		},
	}

	for _, test := range inputs {
		t.Run(test.Name, func(t *testing.T) {
			prepareTestDatabase()

			assert := assert.New(t)
			module, err := courseRepo.CreateModule(test.Input)
			assert.Equal(test.WantErr, err)

			assert.Equal(test.Input.Name, module.Name)
			assert.Equal(test.Input.Description, module.Description)
			assert.Equal(test.Input.Transcript, module.Transcript)
			assert.Equal(test.Input.VoiceoverKey, module.VoiceoverKey)
			assert.Equal(test.Input.BannerKey, module.BannerKey)

			if test.Input.Video != nil {
				assert.Equal((*test.Input.Video).VideoType, *module.VideoType)
				assert.Equal((*test.Input.Video).VideoURL, *module.VideoURL)
			}

			// Check structure
			items, _ := courseRepo.GetModuleStructure(module.UUID)

			var itemUUIDs = make([]gentypes.UUID, len(items))
			for i, item := range items {
				itemUUIDs[i] = item.UUID
			}

			if test.Input.Syllabus != nil {
				for _, item := range *test.Input.Syllabus {
					assert.Contains(itemUUIDs, item.UUID)
				}
			} else {
				assert.Len(items, 0)
			}
		})
	}
}
