package course_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func TestCreateLesson(t *testing.T) {
	prepareTestDatabase()

	var newLessonInput = gentypes.CreateLessonInput{
		Title: "Test lesson",
		Text:  "{}",
		Tags:  nil,
	}

	t.Run("Check non-tagged lesson is created with no tags", func(t *testing.T) {
		lesson, err := courseRepo.CreateLesson(newLessonInput)

		assert.Nil(t, err)
		assert.Equal(t, newLessonInput.Title, lesson.Title)
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
		assert.Equal(t, newLessonInput.Title, lesson.Title)
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

		assert.Equal(t, &errors.ErrNotFound, err)
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
		order := gentypes.OrderBy{Field: "title", Ascending: &asc}

		lessons, _, err := courseRepo.GetLessons(nil, nil, &order)
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