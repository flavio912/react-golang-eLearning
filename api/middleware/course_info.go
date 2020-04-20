package middleware

import (
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
	// TODO: Check if user has access to this course
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

type CourseInfoInput struct {
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

// ComposeCourseInfo creates a courseInfo model from given info
func ComposeCourseInfo(courseInfo CourseInfoInput) (models.CourseInfo, error) {

	var tags []models.Tag
	if courseInfo.Tags != nil {
		_tags, err := CheckTagsExist(*courseInfo.Tags)
		if err != nil {
			return models.CourseInfo{}, err
		}
		tags = _tags
	}

	info := models.CourseInfo{
		Name:          helpers.NilStringToEmpty(courseInfo.Name),
		Price:         helpers.NilFloatToZero(courseInfo.Price),
		Color:         helpers.NilStringToEmpty(courseInfo.Color),
		Tags:          tags,
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

// UpdateCourseInfo updates the courseInfo for a given courseInfoID
func (g *Grant) UpdateCourseInfo(courseInfoID uint, infoChanges CourseInfoInput) (gentypes.CourseInfo, error) {
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

	if infoChanges.Tags != nil {
		// Check each tag exists
		if tags, err := CheckTagsExist(*infoChanges.Tags); err == nil {
			courseInfo.Tags = tags
		} else {
			return gentypes.CourseInfo{}, err
		}
	}
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
	// TODO: This can be relaxed but should check whether they have access to that course
	if !g.IsAdmin {
		return gentypes.CourseInfo{}, &errors.ErrUnauthorized
	}

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
