package course

import (
	"fmt"
	"strconv"

	"github.com/getsentry/sentry-go"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Online Course CRUD */

// ManyOnlineCourseStructures maps many given onlineCourseUUID to a slice of their respective course structures
func (c *coursesRepoImpl) ManyOnlineCourseStructures(onlineCourseUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.CourseStructure, error) {
	var structureItems []models.CourseStructure
	query := database.GormDB.Where("online_course_uuid IN (?)", onlineCourseUUIDs).Order("online_course_uuid, rank ASC").Find(&structureItems)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get online course structures")
		return map[gentypes.UUID][]models.CourseStructure{}, &errors.ErrWhileHandling
	}

	var syllabuses = make(map[gentypes.UUID][]models.CourseStructure)
	for _, item := range structureItems {
		id := item.OnlineCourseUUID
		syllabuses[id] = append(syllabuses[id], item)
	}

	return syllabuses, nil
}

// OnlineCourseStructure gets ordered structure items for a course
func (c *coursesRepoImpl) OnlineCourseStructure(onlineCourseUUID gentypes.UUID) ([]models.CourseStructure, error) {
	structures, err := c.ManyOnlineCourseStructures([]gentypes.UUID{onlineCourseUUID})
	if err != nil {
		return []models.CourseStructure{}, err
	}

	if _, ok := structures[onlineCourseUUID]; ok {
		return structures[onlineCourseUUID], nil
	}

	return []models.CourseStructure{}, nil
}

// OnlineCourse gets an onlineCourse from the courseID
func (c *coursesRepoImpl) OnlineCourse(courseID uint) (models.OnlineCourse, error) {
	var onlineCourse models.OnlineCourse
	query := database.GormDB.Where("course_id = ?", courseID).Find(&onlineCourse)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.OnlineCourse{}, &errors.ErrNotFound
		}
		c.Logger.Log(sentry.LevelError, query.Error, "Could not find online course")
		return models.OnlineCourse{}, &errors.ErrWhileHandling
	}

	return onlineCourse, nil
}

// CreateOnlineCourse creates a new online course
func (c *coursesRepoImpl) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error) {

	// Get courseInfo model
	var courseType = gentypes.OnlineCourseType
	var expMonths *uint
	if courseInfo.ExpiresInMonths != nil {
		l := uint(*courseInfo.ExpiresInMonths)
		expMonths = &l
	}

	infoModel, err := c.ComposeCourse(CourseInput{
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
		CertificateType:      courseInfo.CertificateType,
	})

	if err != nil {
		return models.Course{}, err
	}

	if courseInfo.CategoryUUID != nil {
		infoModel.CategoryUUID = courseInfo.CategoryUUID
	}
	if courseInfo.AccessType != nil {
		infoModel.AccessType = *courseInfo.AccessType
	}
	if courseInfo.BackgroundCheck != nil {
		infoModel.BackgroundCheck = *courseInfo.BackgroundCheck
	}

	infoModel.OnlineCourse = &models.OnlineCourse{}

	query := database.GormDB.Create(&infoModel)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create course")
		return models.Course{}, &errors.ErrWhileHandling
	}

	err = c.saveOnlineCourseStructure(infoModel.OnlineCourse.UUID, courseInfo.Structure)
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to save course structure")
		return models.Course{}, err
	}

	return infoModel, nil
}

// UpdateOnlineCourse updates an existing online course
func (c *coursesRepoImpl) UpdateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error) {
	// Find Course
	if courseInfo.ID == nil {
		c.Logger.LogMessage(sentry.LevelWarning, "No ID given to update onlinecourse")
		return models.Course{}, &errors.ErrWhileHandling
	}

	// TODO: think about putting these two in a transaction
	var expMonths *uint
	if courseInfo.ExpiresInMonths != nil {
		l := uint(*courseInfo.ExpiresInMonths)
		expMonths = &l
	}
	course, err := c.UpdateCourse(uint(*courseInfo.ID), CourseInput{
		Name:         courseInfo.Name,
		Price:        courseInfo.Price,
		Color:        courseInfo.Color,
		CategoryUUID: courseInfo.CategoryUUID,
		// TAGS
		Excerpt:              courseInfo.Excerpt,
		Introduction:         courseInfo.Introduction,
		HoursToComplete:      courseInfo.HoursToComplete,
		HowToComplete:        courseInfo.HowToComplete,
		WhatYouLearn:         courseInfo.WhatYouLearn,
		Requirements:         courseInfo.Requirements,
		AccessType:           courseInfo.AccessType,
		ImageSuccessToken:    courseInfo.BannerImageSuccess,
		BackgroundCheck:      courseInfo.BackgroundCheck,
		SpecificTerms:        courseInfo.SpecificTerms,
		ExpirationToEndMonth: courseInfo.ExpirationToEndMonth,
		ExpiresInMonths:      expMonths,
		CertificateType:      courseInfo.CertificateType,
	})

	if err != nil {
		return course, err
	}

	onlineCourse, err := c.getOnlineCourseFromCourseID(course.ID)
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, fmt.Sprintf("UpdateOnlineCourse: cannot find course %d", course.ID))
		return models.Course{}, &errors.ErrWhileHandling
	}

	err = c.saveOnlineCourseStructure(onlineCourse.UUID, courseInfo.Structure)
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "UpdateOnlineCourse: cannot save structure")
		return models.Course{}, err
	}

	return course, nil
}

func (c *coursesRepoImpl) saveOnlineCourseStructure(courseUUID gentypes.UUID, structure *[]gentypes.CourseItem) error {
	if structure == nil {
		glog.Info("No structure to update")
		return nil
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Naive Implementation of ordering, just delete and re-add everything
	// Should use JIRA method
	query := tx.Where("online_course_uuid = ?", courseUUID).Delete(models.CourseStructure{})
	if query.Error != nil {
		tx.Rollback()
		c.Logger.Log(sentry.LevelError, query.Error, "Course delete before re add failed")
		return &errors.ErrWhileHandling
	}

	for i, courseItem := range *structure {
		structureItem := models.CourseStructure{
			OnlineCourseUUID: courseUUID,
			Rank:             strconv.Itoa(i),
		}

		// TODO: Check if these items exist
		switch courseItem.Type {
		case gentypes.ModuleType:
			structureItem.ModuleUUID = &courseItem.UUID
		case gentypes.LessonType:
			structureItem.LessonUUID = &courseItem.UUID
		case gentypes.TestType:
			structureItem.TestUUID = &courseItem.UUID
		}

		query := tx.Create(&structureItem)
		if query.Error != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, query.Error, "Failed to create the structure")
			return &errors.ErrWhileHandling
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Failed to commit new course structure")
		return &errors.ErrWhileHandling
	}
	return nil
}
