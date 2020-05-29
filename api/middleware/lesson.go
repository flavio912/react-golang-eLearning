package middleware

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// lessonToGenType converts a Lesson model to gentype.
func (g *Grant) lessonToGenType(lesson models.Lesson) gentypes.Lesson {
	tags := tagsToGentypes(lesson.Tags)
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
	tags, err := g.CheckTagsExist(*lesson.Tags)
	// if err != nil {
	// 	g.Logger.Log(sentry.LevelError, err, "Unable to create lesson due to lack of tags existence")
	// 	return gentypes.Lesson{}, &errors.ErrTagsNotFound
	// }

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

	return g.lessonToGenType(lessonModel), nil
}
