package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

var newLesonInput = gentypes.CreateLessonInput{
	Title: "Test lesson",
	Text:  "{}",
	Tags:  nil,
}

func TestCreateLesson(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		_, err := nonAdminGrant.CreateLesson(gentypes.CreateLessonInput{})
		assert.Equal(t, &errors.ErrUnauthorized, err)
	})

	//TODO
	// t.Run("Must validate input", {

	// })

	//TODO
	// t.Run("Tags must exist", {

	// })

	t.Run("Check lesson is created", func(t *testing.T) {
		lesson, err := adminGrant.CreateLesson(newLesonInput)

		assert.Nil(t, err)
		assert.Equal(t, gentypes.Lesson{
			UUID:  lesson.UUID,
			Title: newLesonInput.Title,
			Text:  newLesonInput.Text,
			Tags:  lesson.Tags,
		}, lesson)

	})
}
