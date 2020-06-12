package middleware

import (
	"sort"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (g *Grant) courseToGentype(courseInfo models.Course) gentypes.Course {

	// Get bullet points
	var requirementModels []models.RequirementBullet
	if err := database.GormDB.Where("course_id = ?", courseInfo.ID).Find(&requirementModels).Error; err != nil {
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
	if err := database.GormDB.Where("course_id = ?", courseInfo.ID).Find(&learnModels).Error; err != nil {
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
	return gentypes.Course{
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
		CourseType:      courseInfo.CourseType,
	}
}

// TODO: Bulk load rather than making a million db calls
func (g *Grant) coursesToGentypes(courses []models.Course) []gentypes.Course {
	var genCourses []gentypes.Course
	for _, course := range courses {
		genCourses = append(genCourses, g.courseToGentype(course))
	}
	return genCourses
}

type CourseInput struct {
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
func (g *Grant) ComposeCourse(courseInfo CourseInput) (models.Course, error) {
	// TODO: validate course info input

	var tags []models.Tag
	if courseInfo.Tags != nil {
		_tags, err := g.CheckTagsExist(*courseInfo.Tags)
		if err != nil {
			return models.Course{}, err
		}
		tags = _tags
	}

	var requirements = composeRequirements(courseInfo.Requirements)
	var whatYouLearn = composeWhatYouLearn(courseInfo.WhatYouLearn)

	if courseInfo.CourseType == nil {
		g.Logger.LogMessage(sentry.LevelWarning, "ComposeCourseInfo requires a courseType")
		return models.Course{}, &errors.ErrWhileHandling
	}

	info := models.Course{
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

// UpdateCourse updates the course for a given courseID
func (g *Grant) UpdateCourse(courseID uint, infoChanges CourseInput) (gentypes.Course, error) {
	if !g.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	// Validate input
	_, err := govalidator.ValidateStruct(infoChanges)
	if err != nil {
		return gentypes.Course{}, err
	}

	var courseInfo models.Course
	courseInfo.ID = courseID
	if helpers.StringNotNilOrEmpty(infoChanges.ImageSuccessToken) {
		key, err := uploads.VerifyUploadSuccess(*infoChanges.ImageSuccessToken, "courseBannerImage")
		if err != nil {
			return gentypes.Course{}, err
		}
		courseInfo.ImageKey = &key
	}

	if infoChanges.Tags != nil {
		// Check each tag exists
		if tags, err := g.CheckTagsExist(*infoChanges.Tags); err == nil {
			courseInfo.Tags = tags
		} else {
			return gentypes.Course{}, err
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

		if err := tx.Delete(models.RequirementBullet{}, "course_id = ?", courseID).Error; err != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelError, err, "Unable to delete requirements for course")
			return gentypes.Course{}, &errors.ErrWhileHandling
		}

		courseInfo.Requirements = newRequirements
	}

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.WhatYouLearn != nil {
		var newWhatYouLearn = composeWhatYouLearn(infoChanges.WhatYouLearn)

		if err := tx.Delete(models.WhatYouLearnBullet{}, "course_id = ?", courseID).Error; err != nil {
			tx.Rollback()
			g.Logger.Log(sentry.LevelError, err, "Unable to delete whatYouLearn for course")
			return gentypes.Course{}, &errors.ErrWhileHandling
		}

		courseInfo.WhatYouLearn = newWhatYouLearn
	}

	query := tx.Model(&models.Course{}).Where("id = ?", courseID).Updates(&courseInfo)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to update course")
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return gentypes.Course{}, &errors.ErrWhileHandling
	}

	return g.courseToGentype(courseInfo), nil
}

func (g *Grant) Course(courseID uint) (models.Course, error) {
	var course models.Course
	query := database.GormDB.Where("id = ?", courseID).First(&course)
	if query.Error != nil {
		if query.RecordNotFound() {
			return course, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to get course")
		return course, &errors.ErrWhileHandling
	}
	return course, nil
}

// TODO: Optimise to use (IN) query
func (g *Grant) Courses(courseIDs []uint) ([]models.Course, error) {
	var courseModels []models.Course
	for _, id := range courseIDs {
		mod, err := g.Course(id)
		if err != nil {
			return []models.Course{}, err
		}
		courseModels = append(courseModels, mod)
	}
	return courseModels, nil
}

func (g *Grant) getOnlineCourseFromCourseID(courseID uint) (models.OnlineCourse, error) {
	var onlineCourse models.OnlineCourse
	query := database.GormDB.Where("course_id = ?", courseID).First(&onlineCourse)
	if query.Error != nil {
		if query.RecordNotFound() {
			return onlineCourse, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to get onlineCourse")
		return onlineCourse, &errors.ErrWhileHandling
	}
	return onlineCourse, nil
}

func (g *Grant) GetCourseFromID(courseID uint) (gentypes.Course, error) {
	courseModel, err := g.Course(courseID)
	return g.courseToGentype(courseModel), err
}

func (g *Grant) filterCourse(query *gorm.DB, filter *gentypes.CourseFilter) *gorm.DB {
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.AccessType != nil && *filter.AccessType != "" {
			query = query.Where("access_type = ?", *filter.AccessType)
		}
		if filter.Price != nil {
			query = query.Where("price = ?", *filter.Price)
		}
		if filter.AllowedToBuy != nil && *filter.AllowedToBuy {
			if !g.IsFullyApproved() {
				query = query.Where("access_type = ?", gentypes.Open)
			}
		}
	}

	return query
}

func (g *Grant) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error) {
	// Public function

	var courses []models.Course

	query := g.filterCourse(database.GormDB, filter)

	var count int32
	if err := query.Model(&models.Course{}).Count(&count).Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to count courses")
		return []gentypes.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := getOrdering(query, orderBy, []string{"name", "price"}, "created_at DESC")
	if orderErr != nil {
		g.Logger.Log(sentry.LevelError, orderErr, "Unable to order courses")
		return []gentypes.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := getPage(query, page)
	if err := query.Find(&courses).Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to find courses")
		return []gentypes.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return g.coursesToGentypes(courses), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(courses)),
	}, nil
}

func (g *Grant) isAuthorizedToBook(courses []models.Course) bool {
	if g.IsManager || g.IsIndividual {
		for _, course := range courses {
			if course.AccessType == gentypes.Restricted {
				if !g.IsFullyApproved() {
					return false
				}
			}
		}
		return true
	}
	return false
}
