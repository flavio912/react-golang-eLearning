package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func stringPointer(str string) *string {
	_string := str
	return &_string
}

func floatPointer(flo float64) *float64 {
	_float := flo
	return &_float
}

func boolPointer(boolean bool) *bool {
	_boolean := boolean
	return &_boolean
}

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
		course, err := grant.CreateOnlineCourse(inp)
		assert.Nil(t, err)
		assert.NotNil(t, course.CourseInfoID)
		assert.NotNil(t, course.UUID)

		info, err := grant.GetCourseInfoFromID(course.CourseInfoID)
		assert.Nil(t, err)
		assert.Equal(t, info.Name, name)
		assert.Equal(t, info.AccessType, accessType)
		assert.Equal(t, info.BackgroundCheck, backgroundCheck)
		assert.Equal(t, info.Price, price)
		assert.Equal(t, info.Color, color)
		assert.Equal(t, info.Introduction, introduction)
		assert.Equal(t, info.Excerpt, excerpt)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		// Manager should fail
		grant := &middleware.Grant{auth.UserClaims{}, false, true, false}
		_, err := grant.CreateOnlineCourse(inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// Delegate should fail
		grant = &middleware.Grant{auth.UserClaims{}, false, false, true}
		_, err = grant.CreateOnlineCourse(inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

}

func TestUpdateCourseInfo(t *testing.T) {
	prepareTestDatabase()

	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Updates existing course", func(t *testing.T) {

		open := gentypes.Open
		inp := middleware.UpdateCourseInfoInput{
			Name:            stringPointer("UpdatedCourse"),
			Price:           floatPointer(43.4),
			Color:           stringPointer("#ffffff"),
			Excerpt:         stringPointer("{}"),
			Introduction:    stringPointer("{}"),
			AccessType:      &open,
			BackgroundCheck: boolPointer(false),
			SpecificTerms:   stringPointer("{}"),
		}
		err := grant.UpdateCourseInfo(1, inp)
		assert.Nil(t, err)

		info, err := grant.GetCourseInfoFromID(1)
		assert.Nil(t, err)
		assert.Equal(t, *inp.Name, info.Name)
		assert.Equal(t, *inp.AccessType, info.AccessType)
		assert.Equal(t, *inp.BackgroundCheck, info.BackgroundCheck)
		assert.Equal(t, *inp.Price, info.Price)
		assert.Equal(t, *inp.Color, info.Color)
		assert.Equal(t, *inp.Introduction, info.Introduction)
		assert.Equal(t, *inp.Excerpt, info.Excerpt)
		assert.Equal(t, *inp.SpecificTerms, info.SpecificTerms)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		// Manager should fail
		grant := &middleware.Grant{auth.UserClaims{}, false, true, false}

		inp := middleware.UpdateCourseInfoInput{
			Name: stringPointer("New Course name"),
		}

		err := grant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// Delegate should fail
		grant = &middleware.Grant{auth.UserClaims{}, false, false, true}

		inp = middleware.UpdateCourseInfoInput{
			Name: stringPointer("New Course name"),
		}

		err = grant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

}


