package middleware

import (
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
