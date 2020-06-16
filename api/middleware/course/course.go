package course

import (
	"sort"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

type CoursesRepository interface {
	Course(courseID uint) (models.Course, error)
	Courses(courseIDs []uint) ([]models.Course, error)
	UpdateCourse(courseID uint, infoChanges CourseInput) (models.Course, error)
	ComposeCourse(courseInfo CourseInput) (models.Course, error)
	GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy, fullyApproved bool) ([]models.Course, gentypes.PageInfo, error)

	CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error)
	UpdateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error)

	CreateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error)
	UpdateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error)

	RequirementBullets(courseID uint) ([]models.RequirementBullet, error)
	LearnBullets(courseID uint) ([]models.WhatYouLearnBullet, error)

	GetLessons(page *gentypes.Page, filter *gentypes.LessonFilter, orderBy *gentypes.OrderBy) ([]models.Lesson, gentypes.PageInfo, error)
	CreateLesson(lesson gentypes.CreateLessonInput) (models.Lesson, error)
	GetLessonByUUID(UUID gentypes.UUID) (models.Lesson, error)
	GetLessonsByUUID(uuids []string) ([]models.Lesson, error)
	UpdateLesson(input gentypes.UpdateLessonInput) (models.Lesson, error)
	DeleteLesson(input gentypes.DeleteLessonInput) (bool, error)

	CheckTagsExist(tags []gentypes.UUID) ([]models.Tag, error)
	CreateTag(input gentypes.CreateTagInput) (models.Tag, error)
	GetTagsByCourseInfoIDs(ids []uint) (map[uint][]models.Tag, error)
	GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]models.Tag, error)
	GetTagsByLessonUUID(uuid string) ([]models.Tag, error)

	GetModuleByUUID(moduleUUID gentypes.UUID) (models.Module, error)
	GetModuleStructure(moduleUUID gentypes.UUID) (gentypes.CourseItem, error)
	UpdateModuleStructure(tx *gorm.DB, moduleItem gentypes.CourseItem, duplicateTemplates bool) (models.Module, error)
}

type coursesRepoImpl struct {
	Logger *logging.Logger
}

func NewCoursesRepository(logger *logging.Logger) CoursesRepository {
	return &coursesRepoImpl{
		Logger: logger,
	}
}

func (c *coursesRepoImpl) Course(courseID uint) (models.Course, error) {
	var course models.Course
	query := database.GormDB.Where("id = ?", courseID).First(&course)
	if query.Error != nil {
		if query.RecordNotFound() {
			return course, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get course")
		return course, &errors.ErrWhileHandling
	}
	return course, nil
}

// TODO: Optimise to use (IN) query
func (c *coursesRepoImpl) Courses(courseIDs []uint) ([]models.Course, error) {
	var courseModels []models.Course
	for _, id := range courseIDs {
		mod, err := c.Course(id)
		if err != nil {
			return []models.Course{}, err
		}
		courseModels = append(courseModels, mod)
	}
	return courseModels, nil
}

type CourseInput struct {
	Name              *string
	Price             *float64
	Color             *string `valid:"hexcolor"`
	CategoryUUID      *gentypes.UUID
	Tags              *[]gentypes.UUID
	Excerpt           *string
	Introduction      *string
	HowToComplete     *string
	HoursToComplete   *float64
	WhatYouLearn      *[]string
	Requirements      *[]string
	AccessType        *gentypes.AccessType
	ImageSuccessToken *string
	BackgroundCheck   *bool
	SpecificTerms     *string
	CourseType        *gentypes.CourseType
}

// UpdateCourse updates the course for a given courseID
func (c *coursesRepoImpl) UpdateCourse(courseID uint, infoChanges CourseInput) (models.Course, error) {
	// Validate input
	_, err := govalidator.ValidateStruct(infoChanges)
	if err != nil {
		return models.Course{}, err
	}

	var courseInfo models.Course
	courseInfo.ID = courseID
	if helpers.StringNotNilOrEmpty(infoChanges.ImageSuccessToken) {
		key, err := uploads.VerifyUploadSuccess(*infoChanges.ImageSuccessToken, "courseBannerImage")
		if err != nil {
			return models.Course{}, err
		}
		courseInfo.ImageKey = &key
	}

	if infoChanges.Tags != nil {
		// Check each tag exists
		if tags, err := c.CheckTagsExist(*infoChanges.Tags); err == nil {
			courseInfo.Tags = tags
		} else {
			return models.Course{}, err
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
			c.Logger.Log(sentry.LevelError, err, "Unable to delete requirements for course")
			return models.Course{}, &errors.ErrWhileHandling
		}

		courseInfo.Requirements = newRequirements
	}

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.WhatYouLearn != nil {
		var newWhatYouLearn = composeWhatYouLearn(infoChanges.WhatYouLearn)

		if err := tx.Delete(models.WhatYouLearnBullet{}, "course_id = ?", courseID).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to delete whatYouLearn for course")
			return models.Course{}, &errors.ErrWhileHandling
		}

		courseInfo.WhatYouLearn = newWhatYouLearn
	}

	query := tx.Model(&models.Course{}).Where("id = ?", courseID).Updates(&courseInfo)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to update course")
		return models.Course{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return models.Course{}, &errors.ErrWhileHandling
	}

	return courseInfo, nil
}

func (c *coursesRepoImpl) RequirementBullets(courseID uint) ([]models.RequirementBullet, error) {
	var requirementModels []models.RequirementBullet
	if err := database.GormDB.Where("course_id = ?", courseID).Find(&requirementModels).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []models.RequirementBullet{}, nil
		}

		c.Logger.Log(sentry.LevelError, err, "Unable to get requirements")
		return []models.RequirementBullet{}, &errors.ErrWhileHandling
	}

	sort.SliceStable(requirementModels, func(i, j int) bool {
		return requirementModels[i].OrderID < requirementModels[j].OrderID
	})

	return requirementModels, nil
}

func (c *coursesRepoImpl) LearnBullets(courseID uint) ([]models.WhatYouLearnBullet, error) {
	var learnModels []models.WhatYouLearnBullet
	if err := database.GormDB.Where("course_id = ?", courseID).Find(&learnModels).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []models.WhatYouLearnBullet{}, nil
		}

		c.Logger.Log(sentry.LevelError, err, "Unable to get learn models")
		return []models.WhatYouLearnBullet{}, &errors.ErrWhileHandling
	}

	sort.SliceStable(learnModels, func(i, j int) bool {
		return learnModels[i].OrderID < learnModels[j].OrderID
	})

	return learnModels, nil
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
func (c *coursesRepoImpl) ComposeCourse(courseInfo CourseInput) (models.Course, error) {
	// TODO: validate course info input

	var tags []models.Tag
	if courseInfo.Tags != nil {
		_tags, err := c.CheckTagsExist(*courseInfo.Tags)
		if err != nil {
			return models.Course{}, err
		}
		tags = _tags
	}

	var requirements = composeRequirements(courseInfo.Requirements)
	var whatYouLearn = composeWhatYouLearn(courseInfo.WhatYouLearn)

	if courseInfo.CourseType == nil {
		c.Logger.LogMessage(sentry.LevelWarning, "ComposeCourseInfo requires a courseType")
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

func (c *coursesRepoImpl) getOnlineCourseFromCourseID(courseID uint) (models.OnlineCourse, error) {
	var onlineCourse models.OnlineCourse
	query := database.GormDB.Where("course_id = ?", courseID).First(&onlineCourse)
	if query.Error != nil {
		if query.RecordNotFound() {
			return onlineCourse, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get onlineCourse")
		return onlineCourse, &errors.ErrWhileHandling
	}
	return onlineCourse, nil
}

func filterCourse(query *gorm.DB, filter *gentypes.CourseFilter, fullyApproved bool) *gorm.DB {
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
			if !fullyApproved {
				query = query.Where("access_type = ?", gentypes.Open)
			}
		}
	}

	return query
}

func (c *coursesRepoImpl) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy, fullyApproved bool) ([]models.Course, gentypes.PageInfo, error) {
	// Public function
	var courses []models.Course

	query := filterCourse(database.GormDB, filter, fullyApproved)

	var count int32
	if err := query.Model(&models.Course{}).Count(&count).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to count courses")
		return []models.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"name", "price"}, "created_at DESC")
	if orderErr != nil {
		c.Logger.Log(sentry.LevelError, orderErr, "Unable to order courses")
		return []models.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := middleware.GetPage(query, page)
	if err := query.Find(&courses).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to find courses")
		return []models.Course{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return courses, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(courses)),
	}, nil
}
