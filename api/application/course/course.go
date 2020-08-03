package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
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

	var bannerUrl *string
	if courseInfo.ImageKey != nil {
		url := uploads.GetImgixURL(*courseInfo.ImageKey)
		bannerUrl = &url
	}

	return gentypes.Course{
		ID:                   courseInfo.ID,
		Name:                 courseInfo.Name,
		AccessType:           courseInfo.AccessType,
		BackgroundCheck:      courseInfo.BackgroundCheck,
		Price:                courseInfo.Price,
		Color:                courseInfo.Color,
		Introduction:         courseInfo.Introduction,
		HowToComplete:        courseInfo.HowToComplete,
		HoursToComplete:      courseInfo.HoursToComplete,
		WhatYouLearn:         learnBullets,
		Requirements:         requirementBullets,
		Excerpt:              courseInfo.Excerpt,
		SpecificTerms:        courseInfo.SpecificTerms,
		CategoryUUID:         courseInfo.CategoryUUID,
		AllowedToBuy:         allowedToBuy,
		CourseType:           courseInfo.CourseType,
		BannerImageURL:       bannerUrl,
		ExpiresInMonths:      courseInfo.ExpiresInMonths,
		ExpirationToEndMonth: courseInfo.ExpirationToEndMonth,
		Published:            courseInfo.Published,
		CertificateTypeUUID:  courseInfo.CertificateTypeUUID,
	}
}

func (c *courseAppImpl) SetCoursePublished(courseID uint, published bool) error {
	if !c.grant.IsAdmin {
		return &errors.ErrUnauthorized
	}

	_, err := c.coursesRepository.UpdateCourse(courseID, course.CourseInput{
		Published: &published,
	})
	return err
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

	// Only admins can view unpublished courses
	if course.Published != true && !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	return c.courseToGentype(course), err
}

func (c *courseAppImpl) Courses(courseIDs []uint) ([]gentypes.Course, error) {
	showUnpublished := false
	if c.grant.IsAdmin {
		showUnpublished = true
	}

	courses, err := c.coursesRepository.Courses(courseIDs, showUnpublished)
	return c.coursesToGentypes(courses), err
}

func (c *courseAppImpl) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error) {
	showUnpublished := false
	if c.grant.IsAdmin {
		showUnpublished = true
	}

	courses, pageInfo, err := c.coursesRepository.GetCourses(page, filter, orderBy, application.IsFullyApproved(&c.usersRepository, c.grant), showUnpublished)

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
	if !application.GrantCanViewSyllabus(&c.usersRepository, c.grant, courseID) {
		return []gentypes.CourseItem{}, &errors.ErrUnauthorized
	}

	// Check that its an online course
	onlineCourse, err := c.coursesRepository.OnlineCourse(courseID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelDebug, err, "CourseSyllabus: OnlineCourse: unable to get online part of course")
		return []gentypes.CourseItem{}, err
	}

	structures, structErr := c.coursesRepository.OnlineCourseStructure(onlineCourse.UUID)

	if structErr != nil {
		c.grant.Logger.Log(sentry.LevelWarning, err, "CourseSyllabus: Unable to get course structure")
		return []gentypes.CourseItem{}, structErr
	}

	courseItems := make([]gentypes.CourseItem, len(structures))
	for i, structure := range structures {
		switch {
		case structure.ModuleUUID != nil:
			courseItems[i] = gentypes.CourseItem{
				Type: gentypes.ModuleType,
				UUID: *structure.ModuleUUID,
			}
		case structure.LessonUUID != nil:
			courseItems[i] = gentypes.CourseItem{
				Type: gentypes.LessonType,
				UUID: *structure.LessonUUID,
			}
		case structure.TestUUID != nil:
			courseItems[i] = gentypes.CourseItem{
				Type: gentypes.TestType,
				UUID: *structure.TestUUID,
			}
		default:
			c.grant.Logger.LogMessage(sentry.LevelFatal, "Structure element not recognised")
			return []gentypes.CourseItem{}, &errors.ErrWhileHandling
		}
	}

	return courseItems, nil
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

func (c *courseAppImpl) DeleteCourse(input gentypes.DeleteCourseInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return false, err
	}

	return c.coursesRepository.DeleteCourse(uint(input.ID))
}

func (c *courseAppImpl) CourseBannerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,             // The actual file type
		imageMeta.ContentLength,        // The actual file content length
		[]string{"jpg", "png", "jpeg"}, // Allowed file types
		int32(20000000),                // Max file size = 20MB
		"courseBanners",                // Save files in the "answers" s3 directory
		"courseBannerImage",            // Unique identifier for this type of upload request
	)

	return url, successToken, err
}
