package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestCreateClassroomCourse(t *testing.T) {
	t.Run("Check classroom course created correctly", func(t *testing.T) {
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
			MaxParticipants: helpers.IntPointer(12),
			Location:        helpers.StringPointer("A cool new place"),
		}

		course, err := adminGrant.CreateClassroomCourse(inp)
		assert.Nil(t, err)

		assert.Equal(t, startTime, course.StartDate)
		assert.Equal(t, endTime, course.EndDate)
		assert.Equal(t, "A cool new place", course.Location)
		assert.Equal(t, 12, course.MaxParticipants)

		// Get course info
		info, err := adminGrant.GetCourseInfoFromID(course.CourseInfoID)
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
