package course_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
)

var courseRepo = course.NewCoursesRepository(&logging.Logger{})

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

		course, err := courseRepo.CreateClassroomCourse(inp)
		assert.Nil(t, err)
		assert.Equal(t, *inp.CourseInput.Name, course.Name)
		// TODO: UPDATE TEST
		// assert.Equal(t, startTime, course.StartDate)
		// assert.Equal(t, endTime, course.EndDate)
		// assert.Equal(t, "A cool new place", course.Location)
		// assert.Equal(t, 12, course.MaxParticipants)

	})
}

func TestUpdateClassroomCourse(t *testing.T) {
	t.Run("Updates name", func(t *testing.T) {

		startTime := gentypes.Time{time.Now()}
		endTime := gentypes.Time{time.Now().AddDate(0, 1, 0)}
		var id = uint(1)
		updates := gentypes.SaveClassroomCourseInput{
			CourseInput: gentypes.CourseInput{
				ID:   &id,
				Name: helpers.StringPointer("New classroom name"),
			},
			StartDate:       &startTime,
			EndDate:         &endTime,
			Location:        helpers.StringPointer("A new place"),
			MaxParticipants: helpers.IntPointer(2),
		}
		course, err := courseRepo.UpdateClassroomCourse(updates)
		assert.Nil(t, err)
		assert.Equal(t, *updates.Name, course.Name)
		// TODO Update test
	})
}
