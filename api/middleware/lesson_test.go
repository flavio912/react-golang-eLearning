package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
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

func TestGetLessonByUUID(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		lesson, err := nonAdminGrant.GetLessonByUUID(uuid)

		assert.Equal(t, &errors.ErrUnauthorized, err)
		assert.Equal(t, gentypes.Lesson{}, lesson)
	})

	t.Run("Must show ErrNotFound if not found", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("10000000-0000-0000-0000-000000000001")
		lesson, err := adminGrant.GetLessonByUUID(uuid)

		assert.Equal(t, &errors.ErrNotFound, err)
		assert.Equal(t, gentypes.Lesson{}, lesson)
	})

	t.Run("Must get correct lesson", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		lesson, err := adminGrant.GetLessonByUUID(uuid)

		assert.Nil(t, err)
		assert.Equal(t, uuid, lesson.UUID)
	})
}

func TestGetLessonsByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		grant   middleware.Grant
		uuids   []gentypes.UUID
		wantErr interface{}
		wantLen int
	}{
		{
			"Must be admin",
			nonAdminGrant,
			[]gentypes.UUID{
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			},
			&errors.ErrUnauthorized,
			0,
		},
		// {
		// 	"Must show ErrNotFound if not found *all*",
		// 	adminGrant,
		// 	[]gentypes.UUID{
		// 		gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000022"),
		// 		gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000033"),
		// 	},
		// 	&errors.ErrNotFound,
		// 	0,
		// },
		{
			"Must get only existed lessons",
			adminGrant,
			[]gentypes.UUID{
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000033"),
			},
			nil,
			1,
		},
		{
			"Must get all managers",
			adminGrant,
			[]gentypes.UUID{
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			},
			nil,
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l, err := test.grant.GetLessonsByUUID(test.uuids)

			assert.Equal(t, test.wantErr, err)
			assert.Len(t, l, test.wantLen)
		})
	}
}
