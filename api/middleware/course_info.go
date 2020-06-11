package middleware

import (
	"sort"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

/* Course Info CRUD */

func (g *Grant) courseInfoToGentype(courseInfo models.CourseInfo) gentypes.CourseInfo {

	// Get bullet points
	var requirementModels []models.RequirementBullet
	if err := database.GormDB.Where("course_info_id = ?", courseInfo.ID).Find(&requirementModels).Error; err != nil {
		// Unable to get courseInfo
	}

	sort.SliceStable(requirementModels, func(i, j int) bool {
		return requirementModels[i].OrderID < requirementModels[j].OrderID
	})

	var requirementBullets []string
	for _, bullet := range requirementModels {
		requirementBullets = append(requirementBullets, bullet.Text)
	}

	// Get WhatYouLearn bullet points
	var learnModels []models.WhatYouLearnBullet
	if err := database.GormDB.Where("course_info_id = ?", courseInfo.ID).Find(&learnModels).Error; err != nil {
		// Unable to get courseInfo
	}

	sort.SliceStable(learnModels, func(i, j int) bool {
		return learnModels[i].OrderID < learnModels[j].OrderID
	})

	var learnBullets []string
	for _, bullet := range learnModels {
		learnBullets = append(learnBullets, bullet.Text)
	}

	var allowedToBuy = true
	if courseInfo.AccessType == gentypes.Restricted {
		allowedToBuy = g.IsFullyApproved()
	}

	// TODO: Check if user has access to this course
	return gentypes.CourseInfo{
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
	HowToComplete     *string `valid:"json"`
	HoursToComplete   *float64
	WhatYouLearn      *[]string
	Requirements      *[]string
	AccessType        *gentypes.AccessType
	ImageSuccessToken *string
	BackgroundCheck   *bool
	SpecificTerms     *string `valid:"json"`
	CourseType        *gentypes.CourseType
}

// composeRequirements creates a slice of Bulletpoint models from a slice of strings (the bullet points)
func composeRequirements(requirements *[]string) []models.RequirementBullet {
	var requirementModels []models.RequirementBullet
	if requirements != nil {
		for index, reqText := range *requirements {
			requirementModels = append(requirementModels, models.RequirementBullet{Text: reqText, OrderID: index})
		}
	}

	return requirementModels
}

func composeWhatYouLearn(whatYouLearn *[]string) []models.WhatYouLearnBullet {
	var whatYouLearnModels []models.WhatYouLearnBullet
	if whatYouLearn != nil {
		for index, reqText := range *whatYouLearn {
			whatYouLearnModels = append(whatYouLearnModels, models.WhatYouLearnBullet{Text: reqText, OrderID: index})
		}
	}

	return whatYouLearnModels
}

// ComposeCourseInfo creates a courseInfo model from given info
func (g *Grant) ComposeCourseInfo(courseInfo CourseInfoInput) (models.CourseInfo, error) {
	// TODO: validate course info input

	var tags []models.Tag
	if courseInfo.Tags != nil {
		_tags, err := g.CheckTagsExist(*courseInfo.Tags)
		if err != nil {
			return models.CourseInfo{}, err
		}
		tags = _tags
	}

	var requirements = composeRequirements(courseInfo.Requirements)
	var whatYouLearn = composeWhatYouLearn(courseInfo.WhatYouLearn)

	if courseInfo.CourseType == nil {
		g.Logger.LogMessage(sentry.LevelWarning, "ComposeCourseInfo requires a courseType")
		return models.CourseInfo{}, &errors.ErrWhileHandling
	}

	info := models.CourseInfo{
		Name:            helpers.NilStringToEmpty(courseInfo.Name),
		Price:           helpers.NilFloatToZero(courseInfo.Price),
		Color:           helpers.NilStringToEmpty(courseInfo.Color),
		Tags:            tags,
		Excerpt:         helpers.NilStringToEmpty(courseInfo.Excerpt),
		Introduction:    helpers.NilStringToEmpty(courseInfo.Introduction),
		HowToComplete:   helpers.NilStringToEmpty(courseInfo.HowToComplete),
		HoursToComplete: helpers.NilFloatToZero(courseInfo.HoursToComplete),
		Requirements:    requirements,
		WhatYouLearn:    whatYouLearn,
		SpecificTerms:   helpers.NilStringToEmpty(courseInfo.SpecificTerms),
		CategoryUUID:    courseInfo.CategoryUUID,
		CourseType:      *courseInfo.CourseType,
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
		if tags, err := g.CheckTagsExist(*infoChanges.Tags); err == nil {
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
		courseInfo.CategoryUUID = infoChanges.CategoryUUID // TODO: Check if exists
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

	tx := database.GormDB.Begin()

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.Requirements != nil {
		var newRequirements = composeRequirements(infoChanges.Requirements)

		if err := tx.Delete(models.RequirementBullet{}, "course_info_id = ?", courseInfoID).Error; err != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelError, err, "Unable to delete requirements for course")
			return gentypes.CourseInfo{}, &errors.ErrWhileHandling
		}

		courseInfo.Requirements = newRequirements
	}

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.WhatYouLearn != nil {
		var newWhatYouLearn = composeWhatYouLearn(infoChanges.WhatYouLearn)

		if err := tx.Delete(models.WhatYouLearnBullet{}, "course_info_id = ?", courseInfoID).Error; err != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelError, err, "Unable to delete whatYouLearn for course")
			return gentypes.CourseInfo{}, &errors.ErrWhileHandling
		}

		courseInfo.WhatYouLearn = newWhatYouLearn
	}

	query := tx.Model(&models.CourseInfo{}).Where("id = ?", courseInfoID).Updates(&courseInfo)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to update courseInfo")
		return gentypes.CourseInfo{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return gentypes.CourseInfo{}, &errors.ErrWhileHandling
	}

	return g.courseInfoToGentype(courseInfo), nil
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
			return g.courseInfoToGentype(info), &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to get course info")
		return g.courseInfoToGentype(info), &errors.ErrWhileHandling
	}
	return g.courseInfoToGentype(info), nil
}
