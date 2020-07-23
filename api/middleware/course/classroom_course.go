package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Classroom Course CRUD */

// CreateClassroomCourse makes a new classroom course
func (c *coursesRepoImpl) CreateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error) {
	// Validate
	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return models.Course{}, err
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
	var expMonths *uint
	if courseInfo.ExpiresInMonths != nil {
		l := uint(*courseInfo.ExpiresInMonths)
		expMonths = &l
	}
	course, err := c.ComposeCourse(CourseInput{
		Name:                 courseInfo.Name,
		Price:                courseInfo.Price,
		Color:                courseInfo.Color,
		Tags:                 courseInfo.Tags,
		Excerpt:              courseInfo.Excerpt,
		Introduction:         courseInfo.Introduction,
		HowToComplete:        courseInfo.HowToComplete,
		HoursToComplete:      courseInfo.HoursToComplete,
		WhatYouLearn:         courseInfo.WhatYouLearn,
		Requirements:         courseInfo.Requirements,
		SpecificTerms:        courseInfo.SpecificTerms,
		CourseType:           &courseType,
		ExpirationToEndMonth: courseInfo.ExpirationToEndMonth,
		ExpiresInMonths:      expMonths,
	})
	if err != nil {
		return models.Course{}, err
	}

	course.ClassroomCourse = &classroomCourse

	query := database.GormDB.Create(&course)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create classroom course")
		return models.Course{}, &errors.ErrWhileHandling
	}

	return course, nil
}

// UpdateClassroomCourse updates the given classroom course
func (c *coursesRepoImpl) UpdateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error) {
	// An id is required for this function
	if courseInfo.ID == nil {
		return models.Course{}, &errors.ErrUUIDInvalid
	}

	// Update courseInfo
	var expMonths *uint
	if courseInfo.ExpiresInMonths != nil {
		l := uint(*courseInfo.ExpiresInMonths)
		expMonths = &l
	}
	course, err := c.UpdateCourse(uint(*courseInfo.ID), CourseInput{
		Name:                 courseInfo.Name,
		CategoryUUID:         courseInfo.CategoryUUID,
		Excerpt:              courseInfo.Excerpt,
		Introduction:         courseInfo.Introduction,
		BackgroundCheck:      courseInfo.BackgroundCheck,
		AccessType:           courseInfo.AccessType,
		Price:                courseInfo.Price,
		Color:                courseInfo.Color,
		SpecificTerms:        courseInfo.SpecificTerms,
		Tags:                 courseInfo.Tags,
		ExpirationToEndMonth: courseInfo.ExpirationToEndMonth,
		ExpiresInMonths:      expMonths,
	})
	if err != nil {
		return models.Course{}, err
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
	courseModel.ClassroomCourse = &updates

	q := database.GormDB.Model(models.Course{}).
		Where("id = ?", course.ID).
		Updates(&courseModel).
		Find(&courseModel)
	if q.Error != nil {
		c.Logger.Log(sentry.LevelError, q.Error, "Unable to update course")
		return models.Course{}, &errors.ErrWhileHandling
	}

	return courseModel, nil
}
