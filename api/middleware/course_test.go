package middleware_test

import (
	"testing"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
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
		assert.Equal(t, name, info.Name)
		// assert.Equal(t, info.AccessType, accessType)
		// assert.Equal(t, info.BackgroundCheck, backgroundCheck)
		// assert.Equal(t, info.Price, price)
		// assert.Equal(t, info.Color, color)
		// assert.Equal(t, info.Introduction, introduction)
		// assert.Equal(t, info.Excerpt, excerpt)
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
	grant := &middleware.Grant{auth.UserClaims{}, true, false, false}
	t.Run("Updates existing course", func(t *testing.T) {
		prepareTestDatabase()
		open := gentypes.Open
		inp := middleware.UpdateCourseInfoInput{
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

		inp := middleware.UpdateCourseInfoInput{
			Color: stringPointer("#ffffff"),
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

		inp := middleware.UpdateCourseInfoInput{
			Name: stringPointer("New Course name"),
		}

		_, err := grant.UpdateCourseInfo(1, inp)
		assert.Equal(t, &errors.ErrUnauthorized, err)

		// Delegate should fail
		grant = &middleware.Grant{auth.UserClaims{}, false, false, true}

		inp = middleware.UpdateCourseInfoInput{
			Name: stringPointer("New Course name"),
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
func TestCreateClassroomCourse(t *testing.T) {
	t.Run("Check classroom course created correctly", func(t *testing.T) {
		grant := &middleware.Grant{auth.UserClaims{}, true, false, false}

		open := gentypes.Open
		startTime := gentypes.Time{time.Now()}
		endTime := gentypes.Time{time.Now().AddDate(0, 0, 4)}
		inp := gentypes.SaveClassroomCourseInput{
			CourseInput: gentypes.CourseInput{
				Name:            helpers.StringPointer("New Classroom course"),
				Excerpt:         helpers.StringPointer("{}"),
				Introduction:    helpers.StringPointer("{}"),
				BackgroundCheck: helpers.BoolPointer(true),
				AccessType:      &open,
				Price:           helpers.FloatPointer(23.33),
				Color:           helpers.StringPointer("#fff"),
				SpecificTerms:   helpers.StringPointer("{}"),
			},
			StartDate:       &startTime,
			EndDate:         &endTime,
			MaxParticipants: helpers.IntPointer(0),
			Location:        helpers.StringPointer("A cool new place"),
		}

		course, err := grant.CreateClassroomCourse(inp)
		assert.Nil(t, err)

		assert.Equal(t, course.StartDate, startTime)
		assert.Equal(t, course.EndDate, endTime)
		assert.Equal(t, course.Location, "A cool new place")
		// TODO: max participants

		// Get course info
		info, err := grant.GetCourseInfoFromID(course.CourseInfoID)
		assert.Nil(t, err)
		checkCourseInfoEqual(t, inp.CourseInput, info)
	})
}

func TestUpdateClassroomCourse(t *testing.T) {
	t.Run("Updates name, startDate, endDate and location", func(t *testing.T) {

		startTime := gentypes.Time{time.Now()}
		endTime := gentypes.Time{time.Now().AddDate(0, 1, 0)}
		uid, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000012")
		updates := gentypes.SaveClassroomCourseInput{
			CourseInput: gentypes.CourseInput{
				UUID: &uid,
				Name: helpers.StringPointer("New classroom name"),
			},
			StartDate:       &startTime,
			EndDate:         &endTime,
			Location:        helpers.StringPointer("A new place"),
			MaxParticipants: helpers.IntPointer(2),
		}
		course, err := adminGrant.UpdateClassroomCourse(updates)
		assert.Nil(t, err)
		assert.Equal(t, startTime.Unix(), course.StartDate.Unix())
		assert.Equal(t, endTime.Unix(), course.EndDate.Unix())
		assert.Equal(t, 2, course.MaxParticipants)
		assert.Equal(t, "A new place", course.Location)

		// Find course info
		info, err := adminGrant.GetCourseInfoFromID(course.CourseInfoID)
		assert.Nil(t, err)

		assert.Equal(t, "New classroom name", info.Name)
		assert.Equal(t, 12.01, info.Price) // Price shouldn't have changed
	})
}
