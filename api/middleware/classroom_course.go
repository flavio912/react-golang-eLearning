package middleware

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Classroom Course CRUD */

func classroomCourseToGentype(classroomCourse models.ClassroomCourse) gentypes.ClassroomCourse {
	return gentypes.ClassroomCourse{
		Course: gentypes.Course{
			UUID:         classroomCourse.UUID,
			CourseInfoID: classroomCourse.CourseInfoID,
		},
		StartDate:       classroomCourse.StartDate,
		EndDate:         classroomCourse.EndDate,
		Location:        classroomCourse.Location,
		MaxParticipants: classroomCourse.MaxParticipants,
	}
}

func classroomCoursesToGentypes(courses []models.ClassroomCourse) []gentypes.ClassroomCourse {
	var _courses = make([]gentypes.ClassroomCourse, len(courses))
	for i, course := range courses {
		_courses[i] = classroomCourseToGentype(course)
	}
	return _courses
}

// SaveClassroomCourse is a wrapper around CreateClassroomCourse and UpdateClassroomCourse to
// update the course if a uuid is provided, otherwise it will create a new one
func (g *Grant) SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.ClassroomCourse, error) {
	if courseInfo.UUID != nil {
		return g.UpdateClassroomCourse(courseInfo)
	}
	return g.CreateClassroomCourse(courseInfo)
}

// CreateClassroomCourse makes a new classroom course
func (g *Grant) CreateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.ClassroomCourse, error) {
	if !g.IsAdmin {
		return gentypes.ClassroomCourse{}, &errors.ErrUnauthorized
	}

	// Validate
	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.ClassroomCourse{}, err
	}

	course := models.ClassroomCourse{}

	if courseInfo.MaxParticipants != nil {
		course.MaxParticipants = *courseInfo.MaxParticipants
	}
	if courseInfo.StartDate != nil {
		course.StartDate = *courseInfo.StartDate
	}
	if courseInfo.EndDate != nil {
		course.EndDate = *courseInfo.EndDate
	}
	if courseInfo.Location != nil {
		course.Location = *courseInfo.Location
	}

	course.CourseInfo, err = ComposeCourseInfo(CourseInfoInput{
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
	})
	if err != nil {
		return gentypes.ClassroomCourse{}, err
	}

	query := database.GormDB.Create(&course)
	if query.Error != nil {
		glog.Errorf("Unable to create classroom course: %s", query.Error.Error())
		return gentypes.ClassroomCourse{}, &errors.ErrWhileHandling
	}

	return classroomCourseToGentype(course), nil
}

// UpdateClassroomCourse updates the given classroom course
func (g *Grant) UpdateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.ClassroomCourse, error) {
	if !g.IsAdmin {
		return gentypes.ClassroomCourse{}, &errors.ErrUnauthorized
	}

	// A uuid is required for this function
	if courseInfo.UUID == nil {
		return gentypes.ClassroomCourse{}, &errors.ErrUUIDInvalid
	}

	// Find the course
	var course models.ClassroomCourse
	query := database.GormDB.Where("uuid = ?", *courseInfo.UUID).Find(&course)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.ClassroomCourse{}, &errors.ErrNotFound
		}
		glog.Errorf("Unable to get course while updating: %s", query.Error.Error())
		return gentypes.ClassroomCourse{}, &errors.ErrWhileHandling
	}

	// Update courseInfo
	_, err := g.UpdateCourseInfo(course.CourseInfoID, CourseInfoInput{
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
		return gentypes.ClassroomCourse{}, err
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

	q := database.GormDB.Model(models.ClassroomCourse{}).
		Where("uuid = ?", *courseInfo.UUID).
		Updates(&updates).
		Find(&course)
	if q.Error != nil {
		glog.Errorf("Unable to update course: %s", q.Error.Error())
		return gentypes.ClassroomCourse{}, &errors.ErrWhileHandling
	}

	return classroomCourseToGentype(course), nil
}

func (g *Grant) GetClassroomCourses(
	page *gentypes.Page,
	filter *gentypes.ClassroomCourseFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.ClassroomCourse, gentypes.PageInfo, error) {
	// TODO: allow delegates access to their assigned courses
	if !g.IsAdmin && !g.IsManager {
		return []gentypes.ClassroomCourse{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var courses []models.ClassroomCourse
	query := database.GormDB.Joins("JOIN course_infos ON course_infos.id = classroom_courses.course_info_id")
	query = g.filterCoursesFromInfo(query, filter.CourseInfo)

	query, err := getOrdering(query, orderBy, []string{"name", "access_type", "price"})
	if err != nil {
		return []gentypes.ClassroomCourse{}, gentypes.PageInfo{}, err
	}

	// Count total that can be retrieved by the current filter
	var total int32
	if err := query.Model(&models.ClassroomCourse{}).Count(&total).Error; err != nil {
		glog.Errorf("Unable to get classroom course count: %s", err)
		return []gentypes.ClassroomCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := getPage(query, page)

	query = query.Find(&courses)
	if query.Error != nil {
		glog.Errorf("Unable to get courses: %s", query.Error.Error())
		return []gentypes.ClassroomCourse{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return classroomCoursesToGentypes(courses), gentypes.PageInfo{
		Total:  total,
		Given:  int32(len(courses)),
		Offset: offset,
		Limit:  limit,
	}, nil
}
