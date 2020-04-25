package middleware

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Online Course CRUD */

func onlineCourseToGentype(course models.OnlineCourse) gentypes.OnlineCourse {
	return gentypes.OnlineCourse{
		Course: gentypes.Course{
			UUID:         course.UUID,
			CourseInfoID: course.CourseInfoID,
		},
	}
}

func onlineCoursesToGentypes(courses []models.OnlineCourse) []gentypes.OnlineCourse {
	var _courses = make([]gentypes.OnlineCourse, len(courses))
	for i, course := range courses {
		_courses[i] = onlineCourseToGentype(course)
	}
	return _courses
}

// CreateOnlineCourse creates a new online course
func (g *Grant) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.OnlineCourse, error) {
	if !g.IsAdmin {
		return gentypes.OnlineCourse{}, &errors.ErrUnauthorized
	}

	// Validate
	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.OnlineCourse{}, err
	}

	// Get courseInfo model
	infoModel, err := ComposeCourseInfo(CourseInfoInput{
		Name:          courseInfo.Name,
		Price:         courseInfo.Price,
		Color:         courseInfo.Color,
		Tags:          courseInfo.Tags,
		Excerpt:       courseInfo.Excerpt,
		Introduction:  courseInfo.Introduction,
		SpecificTerms: courseInfo.SpecificTerms,
	})
	if err != nil {
		return gentypes.OnlineCourse{}, err
	}

	newCourse := models.OnlineCourse{
		CourseInfo: infoModel,
	}

	if courseInfo.CategoryUUID != nil {
		newCourse.CourseInfo.CategoryUUID = courseInfo.CategoryUUID
	}
	if courseInfo.AccessType != nil {
		newCourse.CourseInfo.AccessType = *courseInfo.AccessType
	}
	if courseInfo.BackgroundCheck != nil {
		newCourse.CourseInfo.BackgroundCheck = *courseInfo.BackgroundCheck
	}

	query := database.GormDB.Create(&newCourse)
	if query.Error != nil {
		glog.Errorf("Unable to create course: %s", query.Error.Error())
		return gentypes.OnlineCourse{}, &errors.ErrWhileHandling
	}

	err = g.saveOnlineCourseStructure(newCourse.UUID, courseInfo.Structure)
	if err != nil {
		return gentypes.OnlineCourse{}, err
	}

	return onlineCourseToGentype(newCourse), nil
}

// UpdateOnlineCourse updates an existing online course
func (g *Grant) UpdateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.OnlineCourse, error) {
	if !g.IsAdmin {
		return models.OnlineCourse{}, &errors.ErrUnauthorized
	}

	// Find Course
	var onlineCourse models.OnlineCourse
	query := database.GormDB.Where("uuid = ?", courseInfo.UUID).First(&onlineCourse)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.OnlineCourse{}, &errors.ErrNotFound
		}
		return models.OnlineCourse{}, &errors.ErrWhileHandling
	}

	// TODO: think about putting these two in a transaction
	_, err := g.UpdateCourseInfo(onlineCourse.CourseInfoID, CourseInfoInput{
		Name:         courseInfo.Name,
		Price:        courseInfo.Price,
		Color:        courseInfo.Color,
		CategoryUUID: courseInfo.CategoryUUID,
		// TAGS
		Excerpt:           courseInfo.Excerpt,
		Introduction:      courseInfo.Introduction,
		AccessType:        courseInfo.AccessType,
		ImageSuccessToken: courseInfo.BannerImageSuccess,
		BackgroundCheck:   courseInfo.BackgroundCheck,
		SpecificTerms:     courseInfo.SpecificTerms,
	})

	if err != nil {
		return onlineCourse, err
	}

	err = g.saveOnlineCourseStructure(onlineCourse.UUID, courseInfo.Structure)
	if err != nil {
		return onlineCourse, err
	}

	return onlineCourse, nil
}

// SaveOnlineCourse updates or create a new onlineCourse dependant on the existance of a UUID key in the courseInfo input
func (g *Grant) SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.OnlineCourse, error) {
	if !g.IsAdmin {
		return gentypes.OnlineCourse{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.OnlineCourse{}, err
	}

	// If courseUUID given, update
	if courseInfo.UUID != nil {
		// Update CourseInfo
		onlineCourse, err := g.UpdateOnlineCourse(courseInfo)
		return onlineCourseToGentype(onlineCourse), err
	}

	return g.CreateOnlineCourse(courseInfo)

}

func (g *Grant) saveOnlineCourseStructure(courseUUID gentypes.UUID, structure *[]gentypes.CourseItem) error {
	if !g.IsAdmin {
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
			return &errors.ErrWhileHandling
		}
	}

	return nil
}

// HasFullRestrictedAccess returns true if the user has access to all restricted courses
func (g *Grant) HasFullRestrictedAccess() bool {
	if g.IsAdmin {
		return true
	}

	// If a managers company is authorized a manager can view all restricted courses
	if g.IsManager {
		company := models.Company{}
		query := database.GormDB.Where("uuid = ?", g.Claims.Company).First(&company)
		if query.Error != nil {
			glog.Errorf("Unable to get manager's company: %s", query.Error.Error())
			return false
		}

		if company.Approved == true {
			return true
		}
	}

	// Delegates cannot access restricted courses unless specifically assigned them
	if g.IsDelegate {
		return false
	}

	return false
}

func (g *Grant) GetOnlineCourses(page *gentypes.Page, filter *gentypes.OnlineCourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.OnlineCourse, gentypes.PageInfo, error) {
	// TODO: allow delegates access to their assigned courses
	if !g.IsAdmin && !g.IsManager {
		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var courses []models.OnlineCourse
	query := database.GormDB.Joins("JOIN course_infos ON course_infos.id = online_courses.course_info_id")

	// Non-admins can only see published courses
	if !g.IsAdmin {
		query.Where("course_infos.published = ?", true)
	}

	// Filter course info
	if filter != nil && filter.CourseInfo != nil {
		if filter.CourseInfo.Name != nil {
			query = query.Where("course_infos.name ILIKE ?", "%%"+*filter.CourseInfo.Name+"%%")
		}
		if filter.CourseInfo.AccessType != nil {
			query = query.Where("course_infos.access_type = ?", *filter.CourseInfo.AccessType)
		}
		if filter.CourseInfo.BackgroundCheck != nil {
			query = query.Where("course_infos.background_check = ?", *filter.CourseInfo.BackgroundCheck)
		}
		if filter.CourseInfo.Price != nil {
			query = query.Where("course_infos.price = ?", *filter.CourseInfo.Price)
		}
		fmt.Print(courses)
	}

	// TODO: If you're a delegate you should only be allowed to see courses you're assigined too
	// Filter out restricted courses
	if !g.HasFullRestrictedAccess() {
		query = query.Not("course_infos.access_type = ?", "restricted")
	}

	query, err := getOrdering(query, orderBy, []string{"name", "access_type", "price"})
	if err != nil {
		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, err
	}

	// Count total that can be retrieved by the current filter
	var total int32
	if err := query.Model(&models.OnlineCourse{}).Count(&total).Error; err != nil {
		glog.Errorf("Unable to get online course count: %s", err)
		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := getPage(query, page)

	query = query.Find(&courses)
	if query.Error != nil {
		glog.Errorf("Unable to get courses: %s", query.Error.Error())
		return []gentypes.OnlineCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return onlineCoursesToGentypes(courses), gentypes.PageInfo{
		Total:  total,
		Given:  int32(len(courses)),
		Offset: offset,
		Limit:  limit,
	}, nil
}
