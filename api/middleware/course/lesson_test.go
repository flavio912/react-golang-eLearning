package course_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func TestCreateLesson(t *testing.T) {
	prepareTestDatabase()

	var newLessonInput = gentypes.CreateLessonInput{
		Name: "Test lesson",
		Text: "{}",
		Tags: nil,
	}

	t.Run("Check non-tagged lesson is created with no tags", func(t *testing.T) {
		lesson, err := courseRepo.CreateLesson(newLessonInput)

		assert.Nil(t, err)
		assert.Equal(t, newLessonInput.Name, lesson.Name)
		assert.Equal(t, newLessonInput.Text, lesson.Text)
	})
	tag, _ := courseRepo.CreateTag(gentypes.CreateTagInput{
		Name:  "Go",
		Color: "#fff",
	})
	newLessonInput.Tags = &[]gentypes.UUID{tag.UUID}

	t.Run("Check tagged lesson is created with tags", func(t *testing.T) {
		lesson, err := courseRepo.CreateLesson(newLessonInput)

		assert.Nil(t, err)
		assert.Equal(t, newLessonInput.Name, lesson.Name)
		assert.Equal(t, newLessonInput.Text, lesson.Text)

		foundTags, err := courseRepo.GetTagsByLessonUUID(lesson.UUID.String())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(foundTags))
		assert.Equal(t, tag.Name, foundTags[0].Name)
		assert.Equal(t, tag.Color, foundTags[0].Color)
	})
}

func TestGetLessonByUUID(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must show ErrNotFound if not found", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("10000000-0000-0000-0000-000000000001")
		lesson, err := courseRepo.GetLessonByUUID(uuid)

		assert.Equal(t, &errors.ErrLessonNotFound, err)
		assert.Equal(t, models.Lesson{}, lesson)
	})

	t.Run("Must get correct lesson", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		lesson, err := courseRepo.GetLessonByUUID(uuid)

		assert.Nil(t, err)
		assert.Equal(t, uuid, lesson.UUID)
	})
}

func TestGetLessonsByUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuids   []string
		wantErr interface{}
		wantLen int
	}{
		{
			"UUIDs must be valid",
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"yoloo",
			},
			&errors.ErrWhileHandling,
			0,
		},
		{
			"Must get only existed lessons",
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"00000000-0000-0000-0000-000000000033",
			},
			nil,
			1,
		},
		{
			"Must get all lessons",
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
			l, err := courseRepo.GetLessonsByUUID(test.uuids)

			assert.Equal(t, test.wantErr, err)
			assert.Len(t, l, test.wantLen)
		})
	}
}

func TestGetLessons(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return ALL lessons", func(t *testing.T) {
		lessons, _, err := courseRepo.GetLessons(nil, nil, nil)
		assert.Nil(t, err)
		// there are only 3 lessons in test_db
		assert.Len(t, lessons, 3)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		lessons, pageInfo, err := courseRepo.GetLessons(&page, nil, nil)
		assert.Nil(t, err)
		assert.Len(t, lessons, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 3, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := true
		order := gentypes.OrderBy{Field: "name", Ascending: &asc}

		lessons, _, err := courseRepo.GetLessons(nil, nil, &order)
		assert.Nil(t, err)
		assert.Len(t, lessons, 3)
		assert.Equal(t, "Dynamic Programming", lessons[0].Name)
	})

	t.Run("Should filter", func(t *testing.T) {
		lesson := gentypes.Lesson{
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Name: "Dynamic Programming",
			Tags: nil,
			Text: "{}",
		}
		uuidString := lesson.UUID.String()

		filterTests := []struct {
			name   string
			filter gentypes.LessonFilter
		}{
			{"uuid", gentypes.LessonFilter{UUID: &uuidString}},
			{"Name", gentypes.LessonFilter{Name: &lesson.Name}},
		}

		for _, test := range filterTests {
			t.Run(fmt.Sprintf("Should filter %s", test.name), func(t *testing.T) {
				lessons, _, err := courseRepo.GetLessons(nil, &test.filter, nil)
				assert.Nil(t, err)
				require.Len(t, lessons, 1)

				assert.Equal(t, lesson.UUID, lessons[0].UUID)
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
			lessons, _, err := courseRepo.GetLessons(nil, &filterTagTest, nil)

			assert.Nil(t, err)
			assert.Len(t, lessons, 2)
		})

	})

}

func TestUpdateLesson(t *testing.T) {
	prepareTestDatabase()

	t.Run("Lesson must exist", func(t *testing.T) {
		uuidZero := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")
		l, err := courseRepo.UpdateLesson(gentypes.UpdateLessonInput{UUID: uuidZero})
		assert.Equal(t, &errors.ErrLessonNotFound, err)
		assert.Equal(t, models.Lesson{}, l)
	})

	tags := []gentypes.UUID{gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003")}
	input := gentypes.UpdateLessonInput{
		UUID:  gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
		Title: helpers.StringPointer("Diagonalizing Matrices"),
		Text:  helpers.StringPointer(`{"ayy" : "yoo"}`),
		Tags:  &tags,
	}
	t.Run("Updates existing lesson", func(t *testing.T) {
		lesson, err := courseRepo.UpdateLesson(input)

		assert.Nil(t, err)
		assert.Equal(t, input.UUID, lesson.UUID)
		assert.Equal(t, *input.Title, lesson.Title)
		assert.Equal(t, *input.Text, lesson.Text)
		assert.Equal(t, tags[0].UUID, lesson.Tags[0].UUID.UUID)
		assert.Equal(t, "Fancy tag for cool people", lesson.Tags[0].Name)

		lesson, err = courseRepo.GetLessonByUUID(input.UUID)
		assert.Nil(t, err)

		_tags, tag_err := courseRepo.GetTagsByLessonUUID(input.UUID.String())
		lesson.Tags = _tags

		assert.Nil(t, tag_err)
		assert.Equal(t, input.UUID, lesson.UUID)
		assert.Equal(t, *input.Title, lesson.Title)
		assert.Equal(t, *input.Text, lesson.Text)
		assert.Equal(t, tags[0].UUID, lesson.Tags[0].UUID.UUID)
		assert.Equal(t, "Fancy tag for cool people", lesson.Tags[0].Name)
	})
}

func TestDeleteLesson(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must delete lesson", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002")

		b, err := courseRepo.DeleteLesson(gentypes.DeleteLessonInput{
			UUID: uuid,
		})

		assert.Nil(t, err)
		assert.True(t, b)

		lesson, get_err := courseRepo.GetLessonByUUID(uuid)

		assert.Equal(t, &errors.ErrLessonNotFound, get_err)
		assert.Equal(t, models.Lesson{}, lesson)

		tags, _ := courseRepo.GetTagsByLessonUUID(uuid.String())

		assert.Equal(t, []models.Tag{}, tags)
	})

}
