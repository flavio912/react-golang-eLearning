package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestUpdateCourseInfo(t *testing.T) {
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
		_, err := adminGrant.UpdateCourseInfo(1, inp)
		assert.Nil(t, err)

		info, err := adminGrant.GetCourseInfoFromID(1)
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

		prevInfo, err := adminGrant.GetCourseInfoFromID(1)
		assert.Nil(t, err)

		inp := middleware.CourseInfoInput{
			Color: helpers.StringPointer("#ffffff"),
		}
		_, err = adminGrant.UpdateCourseInfo(1, inp)
		assert.Nil(t, err)

		info, err := adminGrant.GetCourseInfoFromID(1)
		assert.Nil(t, err)
		assert.Equal(t, prevInfo.Name, info.Name)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		prepareTestDatabase()

		// Manager should fail
		inp := middleware.CourseInfoInput{
			Name: helpers.StringPointer("New Course name"),
		}
		_, err := managerGrant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// delegate should fail
		inp = middleware.CourseInfoInput{
			Name: helpers.StringPointer("New Course name"),
		}
		_, err = delegateGrant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

}

func TestComposeCourseinfo(t *testing.T) {
	t.Run("Gives correct model", func(t *testing.T) {
		prepareTestDatabase()

		var open = gentypes.Open
		inp := middleware.CourseInfoInput{
			Name:            helpers.StringPointer("Correct model course"),
			Price:           helpers.FloatPointer(32.3),
			Color:           helpers.StringPointer("#fff"),
			CategoryUUID:    &gentypes.UUID{},
			HowToComplete:   helpers.StringPointer("{}"),
			HoursToComplete: helpers.FloatPointer(12.3),
			WhatYouLearn: &[]string{
				"This cool thing",
				"This also cool thing",
			},
			Requirements: &[]string{
				"req 1",
				"req 2",
				"req 3",
			},
			AccessType:      &open,
			BackgroundCheck: helpers.BoolPointer(false),
			SpecificTerms:   helpers.StringPointer("Some specific stuff"),
		}

		info, err := adminGrant.ComposeCourseInfo(inp)
		assert.Nil(t, err)

		// Expected requirements
		req := []models.BulletPoint{
			models.BulletPoint{
				Text:    (*inp.Requirements)[0],
				OrderID: 0,
			},
			models.BulletPoint{
				Text:    (*inp.Requirements)[1],
				OrderID: 1,
			},
			models.BulletPoint{
				Text:    (*inp.Requirements)[2],
				OrderID: 2,
			},
		}

		whatLearn := []models.BulletPoint{
			models.BulletPoint{
				Text:    (*inp.WhatYouLearn)[0],
				OrderID: 0,
			},
			models.BulletPoint{
				Text:    (*inp.WhatYouLearn)[1],
				OrderID: 1,
			},
		}

		assert.Equal(t, req, info.Requirements)
		assert.Equal(t, whatLearn, info.WhatYouLearn)

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
