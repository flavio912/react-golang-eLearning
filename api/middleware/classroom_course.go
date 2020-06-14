package middleware

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Classroom Course CRUD */

// SaveClassroomCourse is a wrapper around CreateClassroomCourse and UpdateClassroomCourse to
// update the course if a uuid is provided, otherwise it will create a new one
func (g *Grant) SaveClassroomCourse(course gentypes.SaveClassroomCourseInput) (gentypes.Course, error) {
	if course.ID != nil {
		return g.UpdateClassroomCourse(course)
	}
	return g.CreateClassroomCourse(course)
}

// CreateClassroomCourse makes a new classroom course
func (g *Grant) CreateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	// Validate
	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	classroomCourse := models.ClassroomCourse{}

	if courseInfo.MaxParticipants != nil {
		classroomCourse.MaxParticipants = *courseInfo.MaxParticipants
	}
	if courseInfo.StartDate != nil {
		classroomCourse.StartDate = *courseInfo.StartDate
	}
	if courseInfo.EndDate != nil {
		classroomCourse.EndDate = *courseInfo.EndDate
	}
	if courseInfo.Location != nil {
		classroomCourse.Location = *courseInfo.Location
	}

	var courseType = gentypes.ClassroomCourseType
	course, err := g.ComposeCourse(CourseInput{
		Name:            courseInfo.Name,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		CategoryUUID:    courseInfo.CategoryUUID,
		Excerpt:         courseInfo.Excerpt,
		Introduction:    courseInfo.Introduction,
		AccessType:      courseInfo.AccessType,
		BackgroundCheck: courseInfo.BackgroundCheck,
		SpecificTerms:   courseInfo.SpecificTerms,
		Tags:            courseInfo.Tags,
		CourseType:      &courseType,
	})
	if err != nil {
		return gentypes.Course{}, err
	}

	course.ClassroomCourse = classroomCourse

	query := database.GormDB.Create(&course)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to create classroom course")
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	return g.courseToGentype(course), nil
}

// UpdateClassroomCourse updates the given classroom course
func (g *Grant) UpdateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	// An id is required for this function
	if courseInfo.ID == nil {
		return gentypes.Course{}, &errors.ErrUUIDInvalid
	}

	// Update courseInfo
	course, err := g.UpdateCourse(*courseInfo.ID, CourseInput{
		Name:            courseInfo.Name,
		CategoryUUID:    courseInfo.CategoryUUID,
		Excerpt:         courseInfo.Excerpt,
		Introduction:    courseInfo.Introduction,
		BackgroundCheck: courseInfo.BackgroundCheck,
		AccessType:      courseInfo.AccessType,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		SpecificTerms:   courseInfo.SpecificTerms,
		Tags:            courseInfo.Tags,
	})
	if err != nil {
		return gentypes.Course{}, err
	}

	var updates models.ClassroomCourse
	if courseInfo.StartDate != nil {
		updates.StartDate = *courseInfo.StartDate
	}
	if courseInfo.EndDate != nil {
		updates.EndDate = *courseInfo.EndDate
	}
	if courseInfo.Location != nil {
		updates.Location = *courseInfo.Location
	}
	if courseInfo.MaxParticipants != nil {
		// If max participants is 0 it will not update
		updates.MaxParticipants = *courseInfo.MaxParticipants
	}

	courseModel := models.Course{ID: course.ID}
	courseModel.ClassroomCourse = updates

	q := database.GormDB.Model(models.Course{}).
		Where("id = ?", course.ID).
		Updates(&courseModel).
		Find(&courseModel)
	if q.Error != nil {
		g.Logger.Log(sentry.LevelError, q.Error, "Unable to update course")
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	return g.courseToGentype(courseModel), nil
}
