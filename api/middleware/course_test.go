package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestCreateOnlineCourse(t *testing.T) {
	prepareTestDatabase()

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}

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
			Name: &name,
			Excerpt: &excerpt.
			Introduction: &introduction,
			BackgroundCheck: &backgroundCheck,
			
		},
	}

	t.Run("Create course with full info", func(t *testing.T) {
		course, err := grant.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.CourseInfoID)
		assert.NotNil(t, course.UUID)

		info, err := grant.GetCourseInfoFromID(course.CourseInfoID)
		assert.Nil(t, err)
		assert.Equal(t, info.Name, name)
	})

}
