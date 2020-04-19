package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestUpdateCourseInfo(t *testing.T) {
	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Updates existing course", func(t *testing.T) {
		prepareTestDatabase()
		open := gentypes.Open
		inp := middleware.CourseInfoInput{
			Name:            helpers.StringPointer("UpdatedCourse"),
			Price:           helpers.FloatPointer(43.4),
			Color:           helpers.StringPointer("#ffffff"),
			Excerpt:         helpers.StringPointer("{}"),
			Introduction:    helpers.StringPointer("{}"),
			AccessType:      &open,
			BackgroundCheck: helpers.BoolPointer(false),
			SpecificTerms:   helpers.StringPointer("{}"),
		}
		_, err := grant.UpdateCourseInfo(1, inp)
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

	t.Run("Doesn't update nil fields", func(t *testing.T) {
		prepareTestDatabase()

		prevInfo, err := grant.GetCourseInfoFromID(1)
		assert.Nil(t, err)

		inp := middleware.CourseInfoInput{
			Color: helpers.StringPointer("#ffffff"),
		}
		_, err = grant.UpdateCourseInfo(1, inp)
		assert.Nil(t, err)

		info, err := grant.GetCourseInfoFromID(1)
		assert.Nil(t, err)
		assert.Equal(t, prevInfo.Name, info.Name)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		prepareTestDatabase()

		// Manager should fail
		grant := &middleware.Grant{auth.UserClaims{}, false, true, false}

		inp := middleware.CourseInfoInput{
			Name: helpers.StringPointer("New Course name"),
		}

		_, err := grant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// Delegate should fail
		grant = &middleware.Grant{auth.UserClaims{}, false, false, true}

		inp = middleware.CourseInfoInput{
			Name: helpers.StringPointer("New Course name"),
		}

		_, err = grant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

}

func checkCourseInfoEqual(t *testing.T, inpInfo gentypes.CourseInput, outInfo gentypes.CourseInfo) {
	if inpInfo.Name != nil {
		assert.Equal(t, *inpInfo.Name, outInfo.Name)
	}
	if inpInfo.Excerpt != nil {
		assert.Equal(t, *inpInfo.Excerpt, outInfo.Excerpt)
	}
	if inpInfo.Introduction != nil {
		assert.Equal(t, *inpInfo.Introduction, outInfo.Introduction)
	}
	if inpInfo.BackgroundCheck != nil {
		assert.Equal(t, *inpInfo.BackgroundCheck, outInfo.BackgroundCheck)
	}
	if inpInfo.AccessType != nil {
		assert.Equal(t, *inpInfo.AccessType, outInfo.AccessType)
	}
	if inpInfo.Price != nil {
		assert.Equal(t, *inpInfo.Price, outInfo.Price)
	}
	if inpInfo.Color != nil {
		assert.Equal(t, *inpInfo.Color, outInfo.Color)
	}
	if inpInfo.SpecificTerms != nil {
		assert.Equal(t, *inpInfo.SpecificTerms, outInfo.SpecificTerms)
	}
}
