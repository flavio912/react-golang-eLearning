package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// CreateLesson is an admin function for creating lessons directly
func (c *coursesRepoImpl) CreateLesson(lesson gentypes.CreateLessonInput) (models.Lesson, error) {
	// Validate input
	if err := lesson.Validate(); err != nil {
		return models.Lesson{}, err
	}

	// Get tags if they exist
	var tags []models.Tag
	if lesson.Tags != nil {
		_tags, err := c.CheckTagsExist(*lesson.Tags)
		if err != nil {
			return models.Lesson{}, err
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
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create lesson")
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	return lessonModel, nil
}

// GetLessonByUUID is an admin function that gets a lesson using it's UUID
func (c *coursesRepoImpl) GetLessonByUUID(UUID gentypes.UUID) (models.Lesson, error) {
	var lesson models.Lesson
	query := database.GormDB.Where("uuid = ?", UUID).First(&lesson)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Lesson{}, &errors.ErrLessonNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get lesson")
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	return lesson, nil
}

func (c *coursesRepoImpl) GetLessonsByUUID(uuids []string) ([]models.Lesson, error) {
	var lessons []models.Lesson

	query := database.GormDB.Where("uuid IN (?)", uuids).Find(&lessons)
	if query.Error != nil {
		if query.RecordNotFound() {
			return lessons, &errors.ErrLessonNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to find lessons")
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
			query = query.Where("title ILIKE ?", "%%"+*filter.Title+"%%")
		}
		if filter.Tags != nil && len(*filter.Tags) > 0 {
			query = query.Table("lessons").
				Joins("JOIN lesson_tags_link ON lesson_tags_link.lesson_uuid = lessons.uuid AND lesson_tags_link.tag_uuid IN (?)", *filter.Tags)
		}
	}

	return query
}

func (c *coursesRepoImpl) GetLessons(
	page *gentypes.Page,
	filter *gentypes.LessonFilter,
	orderBy *gentypes.OrderBy,
) ([]models.Lesson, gentypes.PageInfo, error) {

	var lessons []models.Lesson

	// Count the total filtered dataset
	var count int32
	query := filterLesson(database.GormDB, filter)
	countErr := query.Model(&models.Lesson{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		c.Logger.Log(sentry.LevelError, countErr, "Unable to count lessons")
		return []models.Lesson{}, gentypes.PageInfo{}, countErr
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"title"}, "title ASC")
	if orderErr != nil {
		return []models.Lesson{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&lessons)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Lesson{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to find lessons")
		return []models.Lesson{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return lessons, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(lessons)),
	}, nil
}

// UpdateLesson updates an existing lesson
func (c *coursesRepoImpl) UpdateLesson(input gentypes.UpdateLessonInput) (models.Lesson, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		return models.Lesson{}, err
	}

	var lesson models.Lesson
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&lesson)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Lesson{}, &errors.ErrLessonNotFound
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to find lesson to update with UUID: %s", input.UUID)
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	if input.Title != nil {
		lesson.Title = *input.Title
	}
	if input.Text != nil {
		lesson.Text = *input.Text
	}
	if input.Tags != nil {
		tags, err := c.CheckTagsExist(*input.Tags)

		if err != nil {
			return models.Lesson{}, err
		}
		lesson.Tags = tags

		remove := database.GormDB.Delete(models.LessonTagsLink{}, "lesson_uuid = ?", lesson.UUID)
		if remove.Error != nil {
			c.Logger.Logf(sentry.LevelError, remove.Error, "Error updating tags linked with lesson %s", lesson.UUID)
			return models.Lesson{}, &errors.ErrDeleteFailed
		}

	}

	save := database.GormDB.Model(&models.Lesson{}).Where("uuid = ?", lesson.UUID).Updates(&lesson)
	if save.Error != nil {
		c.Logger.Logf(sentry.LevelError, save.Error, "Error updating lesson with UUID: %s", input.UUID)
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	return lesson, nil
}

func (c *coursesRepoImpl) DeleteLesson(input gentypes.DeleteLessonInput) (bool, error) {
	query := database.GormDB.Begin().Delete(models.LessonTagsLink{}, "lesson_uuid = ?", input.UUID)
	if query.Error != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to remove tags linked with lesson: %s", input.UUID)
		return false, &errors.ErrDeleteFailed
	}

	query = query.Delete(models.Lesson{}, "uuid = ?", input.UUID)
	if query.Error != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to delete lesson: %s", input.UUID)
		return false, &errors.ErrDeleteFailed
	}

	if query.RowsAffected == 0 {
		c.Logger.Logf(sentry.LevelError, &errors.ErrLessonNotFound, "Unable to delete non-existant lesson: %s", input.UUID)
		return false, &errors.ErrLessonNotFound
	}

	if err := query.Commit().Error; err != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to commit transaction of deleting lesson %s", input.UUID)
		return false, &errors.ErrWhileHandling
	}

	return true, nil
}
