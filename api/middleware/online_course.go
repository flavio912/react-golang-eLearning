package middleware

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"

	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Online Course CRUD */

// CreateOnlineCourse creates a new online course
func (g *Grant) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	// Validate
	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	// Get courseInfo model
	var courseType = gentypes.OnlineCourseType
	infoModel, err := g.ComposeCourse(CourseInput{
		Name:            courseInfo.Name,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		Tags:            courseInfo.Tags,
		Excerpt:         courseInfo.Excerpt,
		Introduction:    courseInfo.Introduction,
		HowToComplete:   courseInfo.HowToComplete,
		HoursToComplete: courseInfo.HoursToComplete,
		WhatYouLearn:    courseInfo.WhatYouLearn,
		Requirements:    courseInfo.Requirements,
		SpecificTerms:   courseInfo.SpecificTerms,
		CourseType:      &courseType,
	})

	if err != nil {
		return gentypes.Course{}, err
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

	infoModel.OnlineCourse = models.OnlineCourse{}

	query := database.GormDB.Create(&infoModel)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to create course")
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	err = g.saveOnlineCourseStructure(infoModel.OnlineCourse.UUID, courseInfo.Structure)
	if err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to save course structure")
		return gentypes.Course{}, err
	}

	return g.courseToGentype(infoModel), nil
}

// UpdateOnlineCourse updates an existing online course
func (g *Grant) UpdateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	// Find Course
	if courseInfo.ID == nil {
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	// TODO: think about putting these two in a transaction
	course, err := g.UpdateCourse(*courseInfo.ID, CourseInput{
		Name:         courseInfo.Name,
		Price:        courseInfo.Price,
		Color:        courseInfo.Color,
		CategoryUUID: courseInfo.CategoryUUID,
		// TAGS
		Excerpt:           courseInfo.Excerpt,
		Introduction:      courseInfo.Introduction,
		HoursToComplete:   courseInfo.HoursToComplete,
		HowToComplete:     courseInfo.HowToComplete,
		WhatYouLearn:      courseInfo.WhatYouLearn,
		Requirements:      courseInfo.Requirements,
		AccessType:        courseInfo.AccessType,
		ImageSuccessToken: courseInfo.BannerImageSuccess,
		BackgroundCheck:   courseInfo.BackgroundCheck,
		SpecificTerms:     courseInfo.SpecificTerms,
	})

	if err != nil {
		return course, err
	}

	onlineCourse, err := g.getOnlineCourseFromCourseID(course.ID)
	if err != nil {
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	err = g.saveOnlineCourseStructure(onlineCourse.UUID, courseInfo.Structure)
	if err != nil {
		return gentypes.Course{}, err
	}

	return course, nil
}

// SaveOnlineCourse updates or creates a new onlineCourse dependant on the existance of a UUID key in the courseInfo input
func (g *Grant) SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	// If courseUUID given, update
	if courseInfo.ID != nil {
		// Update CourseInfo
		return g.UpdateOnlineCourse(courseInfo)
	}

	return g.CreateOnlineCourse(courseInfo)

}

func (g *Grant) saveOnlineCourseStructure(courseUUID gentypes.UUID, structure *[]gentypes.CourseItem) error {
	if !g.IsAdmin {
		g.Logger.LogMessage(sentry.LevelError, "Non admin tried to save online course structure (shouldn't be possible)")
		return &errors.ErrWhileHandling
	}

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
	query := tx.Where("online_course_id = ?", courseUUID).Delete(models.CourseStructure{})
	if query.Error != nil {
		tx.Rollback()
		g.Logger.Log(sentry.LevelError, query.Error, "Course delete before re add failed")
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
			_, err := g.UpdateModuleStructure(tx, courseItem, true)
			if err != nil {
				tx.Rollback()
				return err
			}
			structureItem.ModuleUUID = &courseItem.UUID
		case gentypes.LessonType:
			structureItem.LessonUUID = &courseItem.UUID
		case gentypes.TestType:
			structureItem.TestUUID = &courseItem.UUID
		}

		query := tx.Create(&structureItem)
		if query.Error != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelError, query.Error, "Failed to create the structure")
			return &errors.ErrWhileHandling
		}
	}

	if err := tx.Commit().Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Failed to commit new course structure")
		return &errors.ErrWhileHandling
	}
	return nil
}

// filterCoursesFromInfo takes a join of course_infos and online_courses or classroom_courses
// and filters by course info
func (g *Grant) filterCoursesFromInfo(query *gorm.DB, filter *gentypes.CourseFilter) *gorm.DB {
	// Non-admins can only see published courses
	if !g.IsAdmin {
		query = query.Where("course_infos.published = ?", true)
	}

	// Filter course info
	if filter != nil {
		if filter.Name != nil {
			query = query.Where("course_infos.name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.AccessType != nil {
			query = query.Where("course_infos.access_type = ?", *filter.AccessType)
		}
		if filter.BackgroundCheck != nil {
			query = query.Where("course_infos.background_check = ?", *filter.BackgroundCheck)
		}
		if filter.Price != nil {
			query = query.Where("course_infos.price = ?", *filter.Price)
		}
		if filter.AllowedToBuy != nil && !g.IsFullyApproved() {
			query = query.Not("course_infos.access_type = ?", gentypes.Restricted)
		}
	}
	return query
}

// func (g *Grant) GetOnlineCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.OnlineCourse, gentypes.PageInfo, error) {
// 	// TODO: allow delegates access to their assigned courses
// 	if !g.IsAdmin && !g.IsManager {
// 		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
// 	}

// 	var courses []models.OnlineCourse
// 	query := database.GormDB.Joins("JOIN course_infos ON course_infos.id = online_courses.course_info_id")

// 	if filter != nil {
// 		query = g.filterCoursesFromInfo(query, filter.CourseInfo)
// 	} else {
// 		query = g.filterCoursesFromInfo(query, nil)
// 	}

// 	query, err := getOrdering(query, orderBy, []string{"name", "access_type", "price"}, "created_at DESC")
// 	if err != nil {
// 		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, err
// 	}

// 	// Count total that can be retrieved by the current filter
// 	var total int32
// 	if err := query.Model(&models.OnlineCourse{}).Count(&total).Error; err != nil {
// 		g.Logger.Log(sentry.LevelError, err, "Unable to get online course count")
// 		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
// 	}

// 	query, limit, offset := getPage(query, page)

// 	query = query.Find(&courses)
// 	if query.Error != nil {
// 		g.Logger.Log(sentry.LevelError, err, "Unable to get courses")
// 		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
// 	}

// 	return onlineCoursesToGentypes(courses), gentypes.PageInfo{
// 		Total:  total,
// 		Given:  int32(len(courses)),
// 		Offset: offset,
// 		Limit:  limit,
// 	}, nil
// }
