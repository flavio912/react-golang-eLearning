package middleware

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

/* Course Info CRUD */

func courseInfoToGentype(courseInfo models.CourseInfo) gentypes.CourseInfo {
	// TODO: Make for admin only
	return gentypes.CourseInfo{
		ID:              courseInfo.ID,
		Name:            courseInfo.Name,
		AccessType:      courseInfo.AccessType,
		BackgroundCheck: courseInfo.BackgroundCheck,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		Introduction:    courseInfo.Introduction,
		Excerpt:         courseInfo.Excerpt,
		SpecificTerms:   courseInfo.SpecificTerms,
	}
}

type UpdateCourseInfoInput struct {
	Name              *string
	Price             *float64
	Color             *string `valid:"hexcolor"`
	CategoryUUID      *gentypes.UUID
	Tags              *[]gentypes.UUID
	Excerpt           *string `valid:"json"`
	Introduction      *string `valid:"json"`
	AccessType        *gentypes.AccessType
	ImageSuccessToken *string
	BackgroundCheck   *bool
	SpecificTerms     *string `valid:"json"`
}

// UpdateCourseInfo updates the courseInfo for a given courseInfoID
func (g *Grant) UpdateCourseInfo(courseInfoID uint, infoChanges UpdateCourseInfoInput) (gentypes.CourseInfo, error) {
	if !g.IsAdmin {
		return gentypes.CourseInfo{}, &errors.ErrUnauthorized
	}

	// Validate input
	_, err := govalidator.ValidateStruct(infoChanges)
	if err != nil {
		return gentypes.CourseInfo{}, err
	}

	var courseInfo models.CourseInfo
	courseInfo.ID = courseInfoID
	if helpers.StringNotNilOrEmpty(infoChanges.ImageSuccessToken) {
		key, err := uploads.VerifyUploadSuccess(*infoChanges.ImageSuccessToken, "courseBannerImage")
		if err != nil {
			return gentypes.CourseInfo{}, err
		}
		courseInfo.ImageKey = &key
	}
	// if infoChanges.Tags != nil {
	// 	courseInfo.Tags = infoChanges.Tags
	// }
	if infoChanges.Name != nil {
		courseInfo.Name = *infoChanges.Name
	}
	if infoChanges.Price != nil {
		courseInfo.Price = *infoChanges.Price
	}
	if infoChanges.Color != nil {
		courseInfo.Color = *infoChanges.Color
	}
	if infoChanges.CategoryUUID != nil {
		courseInfo.CategoryUUID = infoChanges.CategoryUUID
	}
	if infoChanges.Excerpt != nil {
		courseInfo.Excerpt = *infoChanges.Excerpt
	}
	if infoChanges.Introduction != nil {
		courseInfo.Introduction = *infoChanges.Introduction
	}
	if infoChanges.AccessType != nil {
		courseInfo.AccessType = *infoChanges.AccessType
	}
	if infoChanges.BackgroundCheck != nil {
		courseInfo.BackgroundCheck = *infoChanges.BackgroundCheck
	}
	if infoChanges.SpecificTerms != nil {
		courseInfo.SpecificTerms = *infoChanges.SpecificTerms
	}

	query := database.GormDB.Model(&models.CourseInfo{}).Where("id = ?", courseInfoID).Updates(&courseInfo)
	if query.Error != nil {
		glog.Errorf("Unable to update courseInfo: %s", query.Error.Error())
		//logging.Log(ctx, sentry.LevelError, "Unable to update courseInfo", query.Error)
		return gentypes.CourseInfo{}, &errors.ErrWhileHandling
	}

	return courseInfoToGentype(courseInfo), nil
}

// GetCourseInfoFromID -
func (g *Grant) GetCourseInfoFromID(courseInfoID uint) (gentypes.CourseInfo, error) {
	var info models.CourseInfo
	query := database.GormDB.Where("id = ?", courseInfoID).First(&info)
	if query.Error != nil {
		if query.RecordNotFound() {
			return courseInfoToGentype(info), &errors.ErrNotFound
		}
		glog.Warningf("Unable to get course info: %s", query.Error.Error())
		return courseInfoToGentype(info), &errors.ErrWhileHandling
	}
	return courseInfoToGentype(info), nil
}

/* Online Course CRUD */

func onlineCourseToGentype(course models.OnlineCourse) gentypes.OnlineCourse {
	return gentypes.OnlineCourse{
		Course: gentypes.Course{
			UUID:         course.UUID,
			CourseInfoID: course.CourseInfoID,
		},
	}
}

// ComposeCourseInfo creates a courseInfo model from given info
func ComposeCourseInfo(courseInfo UpdateCourseInfoInput) (models.CourseInfo, error) {
	info := models.CourseInfo{
		Name:  helpers.NilStringToEmpty(courseInfo.Name),
		Price: helpers.NilFloatToZero(courseInfo.Price),
		Color: helpers.NilStringToEmpty(courseInfo.Color),
		//TAGS
		Excerpt:       helpers.NilStringToEmpty(courseInfo.Excerpt),
		Introduction:  helpers.NilStringToEmpty(courseInfo.Introduction),
		SpecificTerms: helpers.NilStringToEmpty(courseInfo.SpecificTerms),
		CategoryUUID:  courseInfo.CategoryUUID,
	}

	if courseInfo.AccessType != nil {
		info.AccessType = *courseInfo.AccessType
	}
	if courseInfo.BackgroundCheck != nil {
		info.BackgroundCheck = *courseInfo.BackgroundCheck
	}

	return info, nil
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

	newCourse := models.OnlineCourse{
		CourseInfo: models.CourseInfo{
			Name:  helpers.NilStringToEmpty(courseInfo.Name),
			Price: helpers.NilFloatToZero(courseInfo.Price),
			Color: helpers.NilStringToEmpty(courseInfo.Color),
			//TAGS
			Excerpt:       helpers.NilStringToEmpty(courseInfo.Excerpt),
			Introduction:  helpers.NilStringToEmpty(courseInfo.Introduction),
			SpecificTerms: helpers.NilStringToEmpty(courseInfo.SpecificTerms),
		},
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
	_, err := g.UpdateCourseInfo(onlineCourse.CourseInfoID, UpdateCourseInfoInput{
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

	course.CourseInfo, err = ComposeCourseInfo(UpdateCourseInfoInput{
		Name:            courseInfo.Name,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		CategoryUUID:    courseInfo.CategoryUUID,
		Excerpt:         courseInfo.Excerpt,
		Introduction:    courseInfo.Introduction,
		AccessType:      courseInfo.AccessType,
		BackgroundCheck: courseInfo.BackgroundCheck,
		SpecificTerms:   courseInfo.SpecificTerms,
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
	_, err := g.UpdateCourseInfo(course.CourseInfoID, UpdateCourseInfoInput{
		Name:            courseInfo.Name,
		CategoryUUID:    courseInfo.CategoryUUID,
		Excerpt:         courseInfo.Excerpt,
		Introduction:    courseInfo.Introduction,
		BackgroundCheck: courseInfo.BackgroundCheck,
		AccessType:      courseInfo.AccessType,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		SpecificTerms:   courseInfo.SpecificTerms,
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

	fmt.Print("START DATE")
	fmt.Print(updates.StartDate)
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
