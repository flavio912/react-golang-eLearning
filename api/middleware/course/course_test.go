package course_test

import (
	"fmt"
	"os"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/testhelpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

var fixtures *testfixtures.Loader

func TestMain(m *testing.M) {
	var err error
	fixtures, err = testhelpers.SetupTestDatabase(true, "middleware_test")
	if err != nil {
		panic("Failed to init test db")
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Unable to load fixtures for test: %s", err.Error())
		panic("cannot load test fixtures")
	}
}

func TestUpdateCourse(t *testing.T) {
	t.Run("Updates existing course", func(t *testing.T) {
		prepareTestDatabase()
		open := gentypes.Open
		inp := course.CourseInput{
			Name:         helpers.StringPointer("UpdatedCourse"),
			Price:        helpers.FloatPointer(43.4),
			Color:        helpers.StringPointer("#fffffa"),
			Excerpt:      helpers.StringPointer("this is a cool excerpt"),
			Introduction: helpers.StringPointer("i am an introduction oifdsf"),
			AccessType:   &open,
			WhatYouLearn: &[]string{
				"This cool thing",
				"This also cool thing",
			},
			Requirements: &[]string{
				"req 1",
				"req 2",
				"req 3",
			},
			BackgroundCheck: helpers.BoolPointer(false),
			SpecificTerms:   helpers.StringPointer("Some specific terms"),
		}
		course, err := courseRepo.UpdateCourse(1, inp)
		assert.Nil(t, err)
		assert.NotEqual(t, models.Course{}, course)

		info, err := courseRepo.Course(1)
		assert.Nil(t, err)
		assert.Equal(t, *inp.Name, info.Name)
		assert.Equal(t, *inp.AccessType, info.AccessType)
		assert.Equal(t, *inp.BackgroundCheck, info.BackgroundCheck)
		assert.Equal(t, *inp.Price, info.Price)
		assert.Equal(t, *inp.Color, info.Color)
		assert.Equal(t, *inp.Introduction, info.Introduction)
		assert.Equal(t, *inp.Excerpt, info.Excerpt)

		var learnStrings []string
		bullets, err := courseRepo.LearnBullets(1)
		assert.Nil(t, err)
		for _, learnBullet := range bullets {
			learnStrings = append(learnStrings, learnBullet.Text)
		}
		assert.Equal(t, *inp.WhatYouLearn, learnStrings)

		var requireBullets []string
		reqBullets, err := courseRepo.RequirementBullets(1)
		assert.Nil(t, err)
		for _, bullet := range reqBullets {
			requireBullets = append(requireBullets, bullet.Text)
		}
		assert.Equal(t, *inp.Requirements, requireBullets)
		assert.Equal(t, *inp.SpecificTerms, info.SpecificTerms)
	})

	t.Run("Doesn't update nil fields", func(t *testing.T) {
		prepareTestDatabase()

		prevInfo, err := courseRepo.Course(1)
		assert.Nil(t, err)

		inp := course.CourseInput{
			Color: helpers.StringPointer("#ffffff"),
		}
		_, err = courseRepo.UpdateCourse(1, inp)
		assert.Nil(t, err)

		info, err := courseRepo.Course(1)
		assert.Nil(t, err)
		assert.Equal(t, prevInfo.Name, info.Name)
	})

}

func TestComposeCourse(t *testing.T) {
	t.Run("Gives correct model", func(t *testing.T) {
		prepareTestDatabase()

		var open = gentypes.Open
		var courseType = gentypes.ClassroomCourseType
		inp := course.CourseInput{
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
			CourseType:      &courseType,
		}

		info, err := courseRepo.ComposeCourse(inp)
		assert.Nil(t, err)

		// Expected requirements
		req := []models.RequirementBullet{
			models.RequirementBullet{
				Text:    (*inp.Requirements)[0],
				OrderID: 0,
			},
			models.RequirementBullet{
				Text:    (*inp.Requirements)[1],
				OrderID: 1,
			},
			models.RequirementBullet{
				Text:    (*inp.Requirements)[2],
				OrderID: 2,
			},
		}

		whatLearn := []models.WhatYouLearnBullet{
			models.WhatYouLearnBullet{
				Text:    (*inp.WhatYouLearn)[0],
				OrderID: 0,
			},
			models.WhatYouLearnBullet{
				Text:    (*inp.WhatYouLearn)[1],
				OrderID: 1,
			},
		}

		assert.Equal(t, req, info.Requirements)
		assert.Equal(t, whatLearn, info.WhatYouLearn)
		assert.Equal(t, courseType, info.CourseType)
	})
}

func checkCourseInfoEqual(t *testing.T, inpInfo gentypes.CourseInput, outInfo gentypes.Course) {
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

func TestManyOnlineCourseStructures(t *testing.T) {
	t.Run("Gets single", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003")
		mapping, err := courseRepo.ManyOnlineCourseStructures([]gentypes.UUID{uuid})
		assert.Nil(t, err)

		assert.Equal(t, "0", mapping[uuid][0].Rank)
		assert.Equal(t, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"), *mapping[uuid][0].LessonUUID)

		assert.Equal(t, "1", mapping[uuid][1].Rank)
		assert.Equal(t, gentypes.MustParseToUUID("2a7e551a-0291-422d-8508-c0ee8ff4c67e"), *mapping[uuid][1].TestUUID)
	})

	// TODO: Add test for multiple
}
