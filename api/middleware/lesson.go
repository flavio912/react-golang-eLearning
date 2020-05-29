package middleware

import (
	"github.com/getsentry/sentry-go"
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
