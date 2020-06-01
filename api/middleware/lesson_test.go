package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateLesson(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		_, err := nonAdminGrant.CreateLesson(gentypes.CreateLessonInput{})
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	t.Run("Must validate input", func(t *testing.T) {
		invalidInput := gentypes.CreateLessonInput{
			Title: "",
			Text:  "",
			Tags:  nil,
		}

		_, err := adminGrant.CreateLesson(invalidInput)
		val_err := invalidInput.Validate()

		assert.Equal(t, err, val_err)
	})

	var newLessonInput = gentypes.CreateLessonInput{
		Title: "Test lesson",
		Text:  "{}",
		Tags:  nil,
	}

	t.Run("Check non-tagged lesson is created with no tags", func(t *testing.T) {
		lesson, err := adminGrant.CreateLesson(newLessonInput)

		assert.Nil(t, err)
		assert.Equal(t, gentypes.Lesson{
			UUID:  lesson.UUID,
			Title: newLessonInput.Title,
			Text:  newLessonInput.Text,
			Tags:  nil,
		}, lesson)

	})
	tag, _ := adminGrant.CreateTag(gentypes.CreateTagInput{
		Name:  "Go",
		Color: "#fff",
	})
	newLessonInput.Tags = &[]gentypes.UUID{tag.UUID}

	t.Run("Check tagged lesson is created with tags", func(t *testing.T) {
		lesson, err := adminGrant.CreateLesson(newLessonInput)

		assert.Nil(t, err)
		assert.Equal(t, newLessonInput.Title, lesson.Title)
		assert.Equal(t, newLessonInput.Text, lesson.Text)

		assert.NotNil(t, lesson.Tags)
		assert.Equal(t, tag, lesson.Tags[0])
	})
}
