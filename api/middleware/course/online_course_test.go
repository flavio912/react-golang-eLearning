package course_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
)

func TestCreateOnlineCourse(t *testing.T) {

	name := "Pies"
	excerpt := "{}"
	introduction := "{}"
	backgroundCheck := false
	accessType := gentypes.Open
	price := 34.3
	color := "#ffffff"

	inp := gentypes.SaveOnlineCourseInput{
		CourseInput: gentypes.CourseInput{
			Name: &name,
		},
	}

	t.Run("Create course with name", func(t *testing.T) {
		prepareTestDatabase()
		course, err := courseRepo.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.ID)

		info, err := courseRepo.Course(course.ID)
		assert.Nil(t, err)
		assert.Equal(t, info.Name, name)
	})

	inp = gentypes.SaveOnlineCourseInput{
		CourseInput: gentypes.CourseInput{
			Name:            &name,
			Excerpt:         &excerpt,
			Introduction:    &introduction,
			BackgroundCheck: &backgroundCheck,
			AccessType:      &accessType,
			Price:           &price,
			Color:           &color,
		},
	}

	t.Run("Create course with full info", func(t *testing.T) {
		prepareTestDatabase()
		course, err := courseRepo.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.ID)

		info, err := courseRepo.Course(course.ID)
		assert.Nil(t, err)
		assert.Equal(t, name, info.Name)
		assert.Equal(t, info.AccessType, accessType)
		assert.Equal(t, info.BackgroundCheck, backgroundCheck)
		assert.Equal(t, info.Price, price)
		assert.Equal(t, info.Color, color)
		assert.Equal(t, info.Introduction, introduction)
		assert.Equal(t, info.Excerpt, excerpt)
	})

	t.Run("Create course with tags", func(t *testing.T) {
		prepareTestDatabase()

		tagUUID, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")
		tagUUID2, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000002")
		tags := []gentypes.UUID{
			tagUUID,
			tagUUID2,
		}
		course, err := courseRepo.CreateOnlineCourse(gentypes.SaveOnlineCourseInput{
			CourseInput: gentypes.CourseInput{
				Name: helpers.StringPointer("course with fantastic tags"),
				Tags: &tags,
			},
		})

		assert.Nil(t, err)

		// Get tags
		tagsMap, err := courseRepo.ManyCourseTags([]uint{course.ID})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(tagsMap[course.ID]))
	})

}
