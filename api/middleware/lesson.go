package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// lessonToGentype converts a Lesson model to gentype.
func (g *Grant) lessonToGentype(lesson models.Lesson) gentypes.Lesson {
	var tags []gentypes.Tag
	if lesson.Tags != nil {
		tags = tagsToGentypes(lesson.Tags)
	}
	return gentypes.Lesson{
		UUID:  lesson.UUID,
		Title: lesson.Title,
		Tags:  tags,
		Text:  lesson.Text,
	}
}

func (g *Grant) lessonsToGentype(lessons []models.Lesson) []gentypes.Lesson {
	var genLessons []gentypes.Lesson
	for _, lesson := range lessons {
		genLessons = append(genLessons, g.lessonToGentype(lesson))
	}
	return genLessons
}

// CreateLesson is an admin function for creating lessons directly
func (g *Grant) CreateLesson(lesson gentypes.CreateLessonInput) (gentypes.Lesson, error) {
	if !g.IsAdmin {
		return gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	// Validate input
	if err := lesson.Validate(); err != nil {
		return gentypes.Lesson{}, err
	}

	// Get tags if they exist
	var tags []models.Tag
	if lesson.Tags != nil {
		_tags, err := g.CheckTagsExist(*lesson.Tags)
		if err != nil {
			return gentypes.Lesson{}, err
		}
		tags = _tags
	}

	lessonModel := models.Lesson{
		Title: lesson.Title,
		Tags:  tags,
		Text:  lesson.Text,
	}

	query := database.GormDB.Create(&lessonModel)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to create lesson")
		return gentypes.Lesson{}, &errors.ErrWhileHandling
	}

	return g.lessonToGentype(lessonModel), nil
}

// GetLessonByUUID is an admin function that gets a lesson using it's UUID
func (g *Grant) GetLessonByUUID(UUID gentypes.UUID) (gentypes.Lesson, error) {

	if !g.IsAdmin {
		return gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	var lesson models.Lesson
	query := database.GormDB.Where("uuid = ?", UUID).First(&lesson)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Lesson{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to get lesson")
		return gentypes.Lesson{}, &errors.ErrWhileHandling
	}

	return g.lessonToGentype(lesson), nil
}

func (g *Grant) GetLessonsByUUID(uuids []string) ([]gentypes.Lesson, error) {
	var lessons []gentypes.Lesson
	if !g.IsAdmin {
		return lessons, &errors.ErrUnauthorized
	}

	query := database.GormDB.Where("uuid IN (?)", uuids).Find(&lessons)
	if query.Error != nil {
		if query.RecordNotFound() {
			return lessons, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to find lessons")
		return lessons, &errors.ErrWhileHandling
	}

	return lessons, nil
}

func filterLesson(query *gorm.DB, filter *gentypes.LessonFilter) *gorm.DB {
	if filter != nil {
		if filter.UUID != nil && *filter.UUID != "" {
			query = query.Where("uuid = ?", *filter.UUID)
		}
		if filter.Title != nil && *filter.Title != "" {
			query = query.Where("title = ?", *filter.Title)
		}
	}

	return query
}

func (g *Grant) GetLessons(
	page *gentypes.Page,
	filter *gentypes.LessonFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Lesson, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.Lesson{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var lessons []models.Lesson

	// Count the total filtered dataset
	var count int32
	query := filterLesson(database.GormDB, filter)
	countErr := query.Model(&models.Lesson{}).Limit(MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		g.Logger.Log(sentry.LevelError, countErr, "Unable to count lessons")
		return []gentypes.Lesson{}, gentypes.PageInfo{}, countErr
	}

	query, orderErr := getOrdering(query, orderBy, []string{"title"}, "title ASC")
	if orderErr != nil {
		return []gentypes.Lesson{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query = query.Find(&lessons)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.Lesson{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to find lessons")
		return []gentypes.Lesson{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return g.lessonsToGentype(lessons), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(lessons)),
	}, nil
}
