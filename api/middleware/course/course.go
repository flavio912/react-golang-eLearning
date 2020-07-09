package course

import (
	"sort"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
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
	ManyOnlineCourseStructures(onlineCourseUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.CourseStructure, error)
	OnlineCourseStructure(onlineCourseUUID gentypes.UUID) ([]models.CourseStructure, error)
	OnlineCourse(courseID uint) (models.OnlineCourse, error)

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

	CheckTagsExist(tags []gentypes.UUID) ([]models.Tag, error)
	CreateTag(input gentypes.CreateTagInput) (models.Tag, error)
	ManyCourseTags(ids []uint) (map[uint][]models.Tag, error)
	GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]models.Tag, error)
	GetTagsByLessonUUID(uuid string) ([]models.Tag, error)

	CreateModule(input CreateModuleInput) (models.Module, error)
	UpdateModule(input UpdateModuleInput) (models.Module, error)
	GetModuleByUUID(moduleUUID gentypes.UUID) (models.Module, error)
	GetModuleStructure(moduleUUID gentypes.UUID) ([]gentypes.ModuleItem, error)
	UpdateModuleStructure(tx *gorm.DB, moduleUUID gentypes.UUID, moduleStructure []gentypes.ModuleItem) (models.Module, error)
	IsModuleInCourses(courseIDs []uint, moduleUUID gentypes.UUID) (bool, error)

	Test(testUUID gentypes.UUID) (models.Test, error)
	Tests(
		page *gentypes.Page,
		filter *gentypes.TestFilter,
		orderBy *gentypes.OrderBy,
	) ([]models.Test, gentypes.PageInfo, error)
	ManyTests(testUUIDs []gentypes.UUID) (map[gentypes.UUID]models.Test, error)
	CreateTest(input CreateTestInput) (models.Test, error)
	UpdateTest(input UpdateTestInput) (models.Test, error)
	TestQuestions(testUUID gentypes.UUID) ([]models.Question, error)
	ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.BasicAnswer, error)

	CourseTests(onlineCourseUUID gentypes.UUID) ([]models.Test, error)

	Question(uuid gentypes.UUID) (models.Question, error)
	Questions(page *gentypes.Page, filter *gentypes.QuestionFilter, orderBy *gentypes.OrderBy) ([]models.Question, gentypes.PageInfo, error)
	CreateQuestion(input CreateQuestionArgs) (models.Question, error)
	UpdateQuestion(input UpdateQuestionArgs) (models.Question, error)
	DeleteQuestion(input gentypes.UUID) (bool, error)

	CreateTestMarks(mark models.TestMark) error
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

	updates := make(map[string]interface{})

	if helpers.StringNotNilOrEmpty(infoChanges.ImageSuccessToken) {
		key, err := uploads.VerifyUploadSuccess(*infoChanges.ImageSuccessToken, "courseBannerImage")
		if err != nil {
			return models.Course{}, err
		}
		updates["image_key"] = key
	}

	if infoChanges.Name != nil {
		glog.Errorf("GOT name update: %s", *infoChanges.Name)
		updates["name"] = *infoChanges.Name
	}
	if infoChanges.Price != nil {
		updates["price"] = *infoChanges.Price
	}
	if infoChanges.Color != nil {
		updates["color"] = *infoChanges.Color
	}
	if infoChanges.CategoryUUID != nil {
		updates["category_uuid"] = infoChanges.CategoryUUID // TODO: Check if exists
	}
	if infoChanges.Excerpt != nil {
		updates["excerpt"] = *infoChanges.Excerpt
	}
	if infoChanges.Introduction != nil {
		updates["introduction"] = *infoChanges.Introduction
	}
	if infoChanges.HowToComplete != nil {
		updates["how_to_complete"] = *infoChanges.HowToComplete
	}
	if infoChanges.HoursToComplete != nil {
		updates["hours_to_complete"] = *infoChanges.HoursToComplete
	}
	if infoChanges.AccessType != nil {
		updates["access_type"] = *infoChanges.AccessType
	}
	if infoChanges.BackgroundCheck != nil {
		updates["background_check"] = *infoChanges.BackgroundCheck
	}
	if infoChanges.SpecificTerms != nil {
		updates["specific_terms"] = *infoChanges.SpecificTerms
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.Logger.LogMessage(sentry.LevelFatal, "UpdateCourse: Forced to recover")
		}
	}()

	if infoChanges.Tags != nil {
		// Check each tag exists
		if tags, err := c.CheckTagsExist(*infoChanges.Tags); err == nil {
			repErr := tx.Model(models.Course{ID: courseID}).Association("Tags").Replace(tags).Error
			if repErr != nil {
				c.Logger.Log(sentry.LevelError, repErr, "Could not replace tags")
				tx.Rollback()
				return models.Course{}, &errors.ErrWhileHandling
			}
		} else {
			return models.Course{}, err
		}
	}

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.Requirements != nil {
		var newRequirements = composeRequirements(infoChanges.Requirements)

		if err := tx.Delete(models.RequirementBullet{}, "course_id = ?", courseID).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to delete requirements for course")
			return models.Course{}, &errors.ErrWhileHandling
		}

		repErr := tx.Model(&models.Course{ID: courseID}).Association("Requirements").Replace(newRequirements).Error
		if repErr != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, repErr, "Unable to replace requirements")
			return models.Course{}, &errors.ErrWhileHandling
		}
	}

	// If requirements changed, remove all old ones and repopulate
	if infoChanges.WhatYouLearn != nil {
		var newWhatYouLearn = composeWhatYouLearn(infoChanges.WhatYouLearn)

		if err := tx.Delete(models.WhatYouLearnBullet{}, "course_id = ?", courseID).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to delete whatYouLearn for course")
			return models.Course{}, &errors.ErrWhileHandling
		}

		repErr := tx.Model(&models.Course{ID: courseID}).Association("WhatYouLearn").Replace(newWhatYouLearn).Error
		if repErr != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, repErr, "Unable to replace whatYouLearn")
			return models.Course{}, &errors.ErrWhileHandling
		}
	}

	query := tx.Model(&models.Course{}).Where("id = ?", courseID).Updates(updates)
	if query.Error != nil {
		tx.Rollback()
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to update course")
		return models.Course{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return models.Course{}, &errors.ErrWhileHandling
	}

	course, err := c.Course(courseID)
	if err != nil {
		return models.Course{}, &errors.ErrNotFound
	}
	return course, nil
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

// ManyOnlineCourseStructures maps many given onlineCourseUUID to a slice of their respective course structures
func (c *coursesRepoImpl) ManyOnlineCourseStructures(onlineCourseUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.CourseStructure, error) {
	var structureItems []models.CourseStructure
	query := database.GormDB.Where("online_course_uuid IN (?)", onlineCourseUUIDs).Order("online_course_uuid, rank ASC").Find(&structureItems)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get online course structures")
		return map[gentypes.UUID][]models.CourseStructure{}, &errors.ErrWhileHandling
	}

	var syllabuses = make(map[gentypes.UUID][]models.CourseStructure)
	for _, item := range structureItems {
		id := item.OnlineCourseUUID
		syllabuses[id] = append(syllabuses[id], item)
	}

	return syllabuses, nil
}

// OnlineCourseStructure gets ordered structure items for a course
func (c *coursesRepoImpl) OnlineCourseStructure(onlineCourseUUID gentypes.UUID) ([]models.CourseStructure, error) {
	structures, err := c.ManyOnlineCourseStructures([]gentypes.UUID{onlineCourseUUID})
	if err != nil {
		return []models.CourseStructure{}, err
	}

	if _, ok := structures[onlineCourseUUID]; ok {
		return structures[onlineCourseUUID], nil
	}

	return []models.CourseStructure{}, nil
}
