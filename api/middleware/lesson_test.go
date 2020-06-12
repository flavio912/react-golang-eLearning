package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
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

		assert.Equal(t, &errors.ErrLessonNotFound, err)
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
		uuids   []string
		wantErr interface{}
		wantLen int
	}{
		{
			"Must be admin",
			nonAdminGrant,
			[]string{
				"00000000-0000-0000-0000-000000000001",
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
			"UUIDs must be valid",
			adminGrant,
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"yoloo",
			},
			&errors.ErrWhileHandling,
			0,
		},
		{
			"Must get only existed lessons",
			adminGrant,
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"00000000-0000-0000-0000-000000000033",
			},
			nil,
			1,
		},
		{
			"Must get all lessons",
			adminGrant,
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"00000000-0000-0000-0000-000000000002",
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

func TestGetLessons(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		lessons, info, err := nonAdminGrant.GetLessons(nil, nil, nil)
		assert.Equal(t, &errors.ErrUnauthorized, err)
		assert.Equal(t, []gentypes.Lesson{}, lessons)
		assert.Equal(t, gentypes.PageInfo{}, info)
	})

	t.Run("Should return ALL lessons", func(t *testing.T) {
		lessons, _, err := adminGrant.GetLessons(nil, nil, nil)
		assert.Nil(t, err)
		// there are only 3 lessons in test_db
		assert.Len(t, lessons, 3)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		lessons, pageInfo, err := adminGrant.GetLessons(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, lessons, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 3, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "title", Ascending: &asc}

		lessons, _, err := adminGrant.GetLessons(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, lessons, 3)
		assert.Equal(t, "Dynamic Programming", lessons[0].Title)
	})

	t.Run("Should filter", func(t *testing.T) {
		lesson := gentypes.Lesson{
			UUID:  gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Title: "Dynamic Programming",
			Tags:  nil,
			Text:  "{}",
		}
		uuidString := lesson.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.LessonFilter
		}{
			{"uuid", gentypes.LessonFilter{UUID: &uuidString}},
			{"Title", gentypes.LessonFilter{Title: &lesson.Title}},
		}

		for _, test := range filterTests {
			t.Run(fmt.Sprintf("Should filter %s", test.name), func(t *testing.T) {
				lessons, _, err := adminGrant.GetLessons(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, lessons, 1)

				assert.Equal(t, lesson, lessons[0])
			})
		}
		t.Run("Should filter by tags and return multiple", func(t *testing.T) {
			tag := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002")
			tags := []*gentypes.UUID{
				&tag,
			}
			filterTagTest := gentypes.LessonFilter{
				Tags: &tags,
			}
			lessons, _, err := adminGrant.GetLessons(nil, &filterTagTest, nil)

			assert.Nil(t, err)
			assert.Len(t, lessons, 2)
		})

	})

}

func TestUpdateLesson(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must be admin", func(t *testing.T) {
		l, err := nonAdminGrant.UpdateLesson(gentypes.UpdateLessonInput{})
		assert.Equal(t, &errors.ErrUnauthorized, err)
		assert.Equal(t, gentypes.Lesson{}, l)
	})

	t.Run("Input must be valid", func(t *testing.T) {
		invalidInput := gentypes.UpdateLessonInput{
			Text: helpers.StringPointer("{"),
		}

		l, err := adminGrant.UpdateLesson(invalidInput)
		assert.Equal(t, invalidInput.Validate(), err)
		assert.Equal(t, gentypes.Lesson{}, l)
	})

	t.Run("Lesson must exist", func(t *testing.T) {
		l, err := adminGrant.UpdateLesson(gentypes.UpdateLessonInput{UUID: uuidZero})
		assert.Equal(t, &errors.ErrLessonNotFound, err)
		assert.Equal(t, gentypes.Lesson{}, l)
	})

	tags := []gentypes.UUID{gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003")}
	input := gentypes.UpdateLessonInput{
		UUID:  gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
		Title: helpers.StringPointer("Diagonalizing Matrices"),
		Text:  helpers.StringPointer(`{"ayy" : "yoo"}`),
		Tags:  &tags,
	}
	t.Run("Updates existing lesson", func(t *testing.T) {
		lesson, err := adminGrant.UpdateLesson(input)

		expected := gentypes.Lesson{
			UUID:  input.UUID,
			Title: *input.Title,
			Text:  *input.Text,
			Tags: []gentypes.Tag{
				{
					UUID:  tags[0],
					Name:  "Fancy tag for cool people",
					Color: "#123",
				},
			},
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, lesson)

		lesson, err = adminGrant.GetLessonByUUID(input.UUID)
		assert.Nil(t, err)

		_tags, tag_err := adminGrant.GetTagsByLessonUUID(input.UUID.String())
		lesson.Tags = _tags

		assert.Nil(t, tag_err)
		assert.Equal(t, expected, lesson)
	})
}
