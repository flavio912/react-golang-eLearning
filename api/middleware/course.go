package middleware

import (
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func nilStringToEmpty(item *string) string {
	if item == nil {
		return ""
	}
	return *item
}

func nilFloatToZero(item *float64) float64 {
	if item == nil {
		return 0
	}
	return *item
}

func (g *Grant) GetCourseInfoFromID(courseInfoID uint) (models.CourseInfo, error) {
	var info models.CourseInfo
	query := database.GormDB.Where("id = ?", courseInfoID).First(&info)
	if query.Error != nil {
		if query.RecordNotFound() {
			return info, &errors.ErrNotFound
		}
		glog.Warningf("Unable to get course info: %s", query.Error.Error())
		return info, &errors.ErrWhileHandling
	}
	return info, nil
}

func onlineCourseToGentype(course models.OnlineCourse) gentypes.OnlineCourse {
	return gentypes.OnlineCourse{
		gentypes.Course{
			UUID:         course.UUID.String(),
			CourseInfoID: course.CourseInfoID,
		},
	}
}

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
			Name:  nilStringToEmpty(courseInfo.Name),
			Price: nilFloatToZero(courseInfo.Price),
			Color: nilStringToEmpty(courseInfo.Color),
			//TAGS
			Excerpt:       nilStringToEmpty(courseInfo.Excerpt),
			Introduction:  nilStringToEmpty(courseInfo.Introduction),
			SpecificTerms: nilStringToEmpty(courseInfo.SpecificTerms),
		},
	}

	if courseInfo.CategoryUUID != nil {
		catUUID, err := uuid.Parse(*courseInfo.CategoryUUID)
		if err != nil {
			glog.Info("category uuid is invalid")
			return gentypes.OnlineCourse{}, &errors.ErrUUIDInvalid
		}
		newCourse.CourseInfo.CategoryUUID = &catUUID
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

	err := g.saveOnlineCourseStructure(newCourse.UUID, courseInfo.Structure)
	if err != nil {
		return gentypes.OnlineCourse{}, err
	}

	return onlineCourseToGentype(newCourse), nil
}

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
		onlineCourse, err := g.updateOnlineCourse(courseInfo)
		return onlineCourseToGentype(onlineCourse), err
	}

	return g.CreateOnlineCourse(courseInfo)

}

// duplicateModule copys a module and its stucture
// TODO: This really isn't nice, make more efficient
func duplicateModule(tx *gorm.DB, module models.Module, template bool, duplicateStructure bool) (models.Module, error) {
	// Duplicate module
	newModule := models.Module{
		Template:   template,
		TemplateID: module.TemplateID,
	}

	query := tx.Save(&newModule)
	if query.Error != nil {
		glog.Error("Unable to save duplicated module")
		return models.Module{}, &errors.ErrWhileHandling
	}

	if !duplicateStructure {
		return newModule, nil
	}

	for _, item := range module.Structure {
		structure := models.ModuleStructure{
			ModuleUUID: newModule.UUID,
			LessonUUID: item.LessonUUID,
			TestUUID:   item.TestUUID,
			Rank:       item.Rank,
		}
		query := tx.Save(&structure)
		if query.Error != nil {
			glog.Error("Unable to save module structure while duplicating")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return newModule, nil
}

func ReorderModule(tx *gorm.DB, moduleItem gentypes.CourseItem, duplicateTemplates bool) error {
	var moduleModel models.Module
	query := tx.Where("uuid = ?", moduleItem.UUID).First(&moduleModel)
	if query.Error != nil {
		if query.RecordNotFound() {
			glog.Infof("Could not find uuid: %s", moduleItem.UUID)
			return &errors.ErrNotFound
		}
		glog.Errorf("Unable to get module: %s", moduleItem.UUID)
		return &errors.ErrWhileHandling
	}

	// Module templates should be duplicated
	if duplicateTemplates && moduleModel.Template {
		moduleModel, err := duplicateModule(tx, moduleModel, false, false)
		if err != nil {
			return err
		}
		moduleItem.UUID = moduleModel.UUID.String()
	} else {
		query = tx.Where("module_uuid = ?", moduleItem.UUID).Delete(models.ModuleStructure{})
		if query.Error != nil {
			return &errors.ErrWhileHandling
		}
	}

	// ModuleUUID to real uuid
	modUUID, err := uuid.Parse(moduleItem.UUID)
	if err != nil {
		glog.Errorf("ModuleUUID is invalid: %s", err.Error())
	}

	for _, item := range moduleItem.Items {
		itemUUID, err := uuid.Parse(item.UUID)
		if err != nil {
			glog.Errorf("ItemUUID is invalid: %s", err.Error())
		}

		// TODO check if lessons + tests exist
		structureItem := models.ModuleStructure{
			ModuleUUID: modUUID,
		}
		if item.Type == gentypes.LessonType {
			structureItem.LessonUUID = &itemUUID
		}
		if item.Type == gentypes.TestType {
			structureItem.TestUUID = &itemUUID
		}

		if err := tx.Save(&structureItem).Error; err != nil {
			glog.Errorf("Unable to save structure item: %s", err)
			return &errors.ErrWhileHandling
		}
	}

	return nil
}

func (g *Grant) saveOnlineCourseStructure(courseUUID uuid.UUID, structure *[]gentypes.CourseItem) error {
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

		// String UUID to uuid.UUID
		itemUUID, err := uuid.Parse(courseItem.UUID)
		if err != nil {
			tx.Rollback()
			glog.Infof("Invalid Course Item UUID, %s", err.Error())
			return &errors.ErrUUIDInvalid
		}

		// TODO: Check if these items exist
		switch courseItem.Type {
		case gentypes.ModuleType:
			if err := ReorderModule(tx, courseItem, true); err != nil {
				tx.Rollback()
				return err
			}
			structureItem.ModuleUUID = &itemUUID
		case gentypes.LessonType:
			structureItem.LessonUUID = &itemUUID
		case gentypes.TestType:
			structureItem.TestUUID = &itemUUID
		}

		query := tx.Create(&structureItem)
		if query.Error != nil {
			tx.Rollback()
			return &errors.ErrWhileHandling
		}
	}

	return nil
}

type SaveCourseInfoInput struct {
	Name              *string
	Price             *float64
	Color             *string `valid:"hexcolor"`
	CategoryUUID      *uuid.UUID
	Tags              *[]string
	Excerpt           *string `valid:"json"`
	Introduction      *string `valid:"json"`
	AccessType        *gentypes.AccessType
	ImageSuccessToken *string
	BackgroundCheck   *bool
	SpecificTerms     *string
}

// SaveCourseInfo is not exported as it returns a model
func (g *Grant) SaveCourseInfo(courseInfoID uint, infoChanges SaveCourseInfoInput) error {
	if !g.IsAdmin {
		return &errors.ErrUnauthorized
	}

	var courseInfo models.CourseInfo
	courseInfo.ID = courseInfoID
	if infoChanges.ImageSuccessToken != nil {
		key, err := uploads.VerifyUploadSuccess(*infoChanges.ImageSuccessToken, "courseBannerImage")
		if err != nil {
			return err
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
		courseInfo.Color = *infoChanges.Name
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

	query := database.GormDB.Save(&courseInfo)
	if query.Error != nil {
		glog.Errorf("Unable to update courseInfo: %s", query.Error.Error())
		//logging.Log(ctx, sentry.LevelError, "Unable to update courseInfo", query.Error)
		return &errors.ErrWhileHandling
	}

	return nil
}

func (g *Grant) updateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.OnlineCourse, error) {
	if !g.IsAdmin {
		return models.OnlineCourse{}, &errors.ErrUnauthorized
	}

	_, err := uuid.Parse(*courseInfo.UUID)
	if err != nil {
		return models.OnlineCourse{}, &errors.ErrUUIDInvalid
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

	// Convert CategoryUUID to
	var catID uuid.UUID
	if courseInfo.CategoryUUID != nil {
		catID, err = uuid.Parse(*courseInfo.CategoryUUID)
	}

	// TODO: think about putting these two in a transaction
	err = g.SaveCourseInfo(onlineCourse.CourseInfoID, SaveCourseInfoInput{
		Name:         courseInfo.Name,
		Price:        courseInfo.Price,
		Color:        courseInfo.Color,
		CategoryUUID: &catID,
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
