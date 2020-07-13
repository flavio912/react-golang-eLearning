package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// IsFullyApproved checks if a user is approved to view all restricted courses
func (c *courseAppImpl) IsFullyApproved() bool {
	if c.grant.IsAdmin {
		return true
	}
	if !c.grant.IsManager {
		return false
	}

	var company, err = c.usersRepository.Company(c.grant.Claims.Company)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to check if manager is approved")
		return false
	}

	return company.Approved
}

func (c *courseAppImpl) courseToGentype(courseInfo models.Course) gentypes.Course {

	// Get bullet points
	requirementModels, err := c.coursesRepository.RequirementBullets(courseInfo.ID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to get bullets")
		return gentypes.Course{}
	}

	var requirementBullets []string
	for _, bullet := range requirementModels {
		requirementBullets = append(requirementBullets, bullet.Text)
	}

	// Get WhatYouLearn bullet points
	learnModels, err := c.coursesRepository.LearnBullets(courseInfo.ID)
	if err != nil {
		return gentypes.Course{}
	}

	var learnBullets []string
	for _, bullet := range learnModels {
		learnBullets = append(learnBullets, bullet.Text)
	}

	var allowedToBuy = true
	if courseInfo.AccessType == gentypes.Restricted {
		allowedToBuy = application.IsFullyApproved(&c.usersRepository, c.grant)
	}

	return gentypes.Course{
		ID:              courseInfo.ID,
		Name:            courseInfo.Name,
		AccessType:      courseInfo.AccessType,
		BackgroundCheck: courseInfo.BackgroundCheck,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		Introduction:    courseInfo.Introduction,
		HowToComplete:   courseInfo.HowToComplete,
		HoursToComplete: courseInfo.HoursToComplete,
		WhatYouLearn:    learnBullets,
		Requirements:    requirementBullets,
		Excerpt:         courseInfo.Excerpt,
		SpecificTerms:   courseInfo.SpecificTerms,
		CategoryUUID:    courseInfo.CategoryUUID,
		AllowedToBuy:    allowedToBuy,
		CourseType:      courseInfo.CourseType,
	}
}

// TODO: Bulk load rather than making a million db calls
func (c *courseAppImpl) coursesToGentypes(courses []models.Course) []gentypes.Course {
	var genCourses []gentypes.Course
	for _, course := range courses {
		genCourses = append(genCourses, c.courseToGentype(course))
	}
	return genCourses
}

func (c *courseAppImpl) Course(courseID uint) (gentypes.Course, error) {
	course, err := c.coursesRepository.Course(courseID)
	return c.courseToGentype(course), err
}

func (c *courseAppImpl) Courses(courseIDs []uint) ([]gentypes.Course, error) {
	courses, err := c.coursesRepository.Courses(courseIDs)
	return c.coursesToGentypes(courses), err
}

func (c *courseAppImpl) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error) {
	courses, pageInfo, err := c.coursesRepository.GetCourses(page, filter, orderBy, application.IsFullyApproved(&c.usersRepository, c.grant))

	return c.coursesToGentypes(courses), pageInfo, err
}

func (c *courseAppImpl) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	course, err := c.coursesRepository.CreateOnlineCourse(courseInfo)
	return c.courseToGentype(course), err
}

func (c *courseAppImpl) SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	var course models.Course
	// If courseUUID given, update
	if courseInfo.ID != nil {
		// Update CourseInfo
		course, err = c.coursesRepository.UpdateOnlineCourse(courseInfo)
	} else {
		course, err = c.coursesRepository.CreateOnlineCourse(courseInfo)
	}

	return c.courseToGentype(course), err
}

// SaveClassroomCourse is a wrapper around CreateClassroomCourse and UpdateClassroomCourse to
// update the course if a uuid is provided, otherwise it will create a new one
func (c *courseAppImpl) SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	var course models.Course
	var err error
	if courseInfo.ID != nil {
		course, err = c.coursesRepository.UpdateClassroomCourse(courseInfo)
	} else {
		course, err = c.coursesRepository.CreateClassroomCourse(courseInfo)
	}

	return c.courseToGentype(course), err
}

func (c *courseAppImpl) CourseSyllabus(courseID uint) ([]gentypes.CourseItem, error) {
	return []gentypes.CourseItem{
		gentypes.CourseItem{
			Type: gentypes.TestType,
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
		},
	}, nil
}

func (c *courseAppImpl) SearchSyllabus(
	page *gentypes.Page,
	filter *gentypes.SyllabusFilter,
) ([]gentypes.CourseItem, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.CourseItem{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	results, pageInfo, err := c.coursesRepository.SearchSyllabus(page, filter)

	return results, pageInfo, err
}
