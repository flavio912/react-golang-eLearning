package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
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

	var (
		bannerImagekey *string
		voiceoverKey   *string
		videoType      *gentypes.VideoType
		videoURL       *string
	)

	if lesson.BannerImageToken != nil {
		key, err := uploads.VerifyUploadSuccess(*lesson.BannerImageToken, "lessonImages")
		if err != nil {
			return models.Lesson{}, err
		}

		bannerImagekey = &key
	}
	if lesson.VoiceoverToken != nil {
		key, err := uploads.VerifyUploadSuccess(*lesson.VoiceoverToken, "voiceoverUploads")
		if err != nil {
			return models.Lesson{}, err
		}

		voiceoverKey = &key
	}
	if lesson.Video != nil {
		videoType = &lesson.Video.Type
		videoURL = &lesson.Video.URL
	}

	lessonModel := models.Lesson{
		Name:         lesson.Name,
		Tags:         tags,
		Description:  lesson.Description,
		Transcript:   lesson.Transcript,
		BannerKey:    bannerImagekey,
		VideoType:    videoType,
		VideoURL:     videoURL,
		VoiceoverKey: voiceoverKey,
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
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
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

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"name"}, "name ASC")
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

	lesson, err := c.GetLessonByUUID(input.UUID)
	if err != nil {
		return models.Lesson{}, err
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.Logger.LogMessage(sentry.LevelFatal, "UpdateLesson: Forced to recover")
		}
	}()

	if input.Name != nil {
		lesson.Name = *input.Name
	}
	if input.Description != nil {
		lesson.Description = *input.Description
	}
	if input.Tags != nil {
		tags, err := c.CheckTagsExist(*input.Tags)

		if err != nil {
			return models.Lesson{}, err
		}
		lesson.Tags = tags

		if err := tx.Delete(models.LessonTagsLink{}, "lesson_uuid = ?", lesson.UUID).Error; err != nil {
			c.Logger.Logf(sentry.LevelError, err, "Error updating tags linked with lesson %s", lesson.UUID)
			tx.Rollback()
			return models.Lesson{}, &errors.ErrDeleteFailed
		}

	}

	if err := tx.Model(&models.Lesson{}).Where("uuid = ?", lesson.UUID).Updates(&lesson).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Error updating lesson with UUID: %s", input.UUID)
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return models.Lesson{}, &errors.ErrWhileHandling
	}

	return lesson, nil
}

func (c *coursesRepoImpl) DeleteLesson(uuid gentypes.UUID) (bool, error) {
	query := database.GormDB.Begin()

	defer func() {
		if r := recover(); r != nil {
			query.Rollback()
			c.Logger.LogMessage(sentry.LevelFatal, "DeleteLesson: Forced to recover")
		}
	}()

	if err := query.Delete(models.LessonTagsLink{}, "lesson_uuid = ?", uuid).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Unable to remove tags linked with lesson: %s", uuid)
		query.Rollback()
		return false, &errors.ErrDeleteFailed
	}

	if err := query.Delete(models.Lesson{}, "uuid = ?", uuid).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Unable to delete lesson: %s", uuid)
		query.Rollback()
		return false, &errors.ErrDeleteFailed
	}

	if err := query.Commit().Error; err != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to commit transaction of deleting lesson %s", uuid)
		return false, &errors.ErrWhileHandling
	}

	return true, nil
}
