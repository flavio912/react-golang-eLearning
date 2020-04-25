package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
)

func TestCreateOnlineCourse(t *testing.T) {

	grant := adminGrant

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
		course, err := grant.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.CourseInfoID)
		assert.NotNil(t, course.UUID)

		info, err := grant.GetCourseInfoFromID(course.CourseInfoID)
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
		course, err := grant.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.CourseInfoID)
		assert.NotNil(t, course.UUID)

		info, err := grant.GetCourseInfoFromID(course.CourseInfoID)
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
		course, err := adminGrant.CreateOnlineCourse(gentypes.SaveOnlineCourseInput{
			CourseInput: gentypes.CourseInput{
				Name: helpers.StringPointer("course with fantastic tags"),
				Tags: &tags,
			},
		})

		assert.Nil(t, err)

		// Get tags
		tagsMap, err := adminGrant.GetTagsByCourseInfoIDs([]uint{course.CourseInfoID})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(tagsMap[course.CourseInfoID]))
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		prepareTestDatabase()

		// Manager should fail
		_, err := managerGrant.CreateOnlineCourse(inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// Delegate should fail
		_, err = delegateGrant.CreateOnlineCourse(inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

}

func TestGetOnlineCourses(t *testing.T) {

	// TODO: Infil -  Test other filters
	t.Run("Test name filter", func(t *testing.T) {
		prepareTestDatabase()

		name := "test"
		courses, pageInfo, err := adminGrant.GetOnlineCourses(nil, &gentypes.OnlineCourseFilter{
			CourseInfo: &gentypes.CourseInfoFilter{
				Name: &name,
			},
		}, nil)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(courses)) // Only one course has test in the name
		assert.Equal(t, gentypes.PageInfo{
			Given:  1,
			Total:  1,
			Limit:  middleware.MaxPageLimit,
			Offset: 0,
		}, pageInfo)
	})

	t.Run("Test Paging", func(t *testing.T) {
		prepareTestDatabase()
		courses, pageInfo, err := adminGrant.GetOnlineCourses(&gentypes.Page{
			Limit:  helpers.Int32Pointer(2),
			Offset: helpers.Int32Pointer(0),
		}, nil, nil)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(courses))
		assert.Equal(t, gentypes.PageInfo{
			Total:  3,
			Offset: 0,
			Limit:  2,
			Given:  2,
		}, pageInfo)
	})
}
