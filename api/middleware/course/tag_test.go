package course_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateTag(t *testing.T) {

	t.Run("Create new, valid tag", func(t *testing.T) {
		prepareTestDatabase()

		tagInput := gentypes.CreateTagInput{
			Name:  "Im a totally new tag",
			Color: "#123",
		}

		tag, err := courseRepo.CreateTag(tagInput)

		assert.Nil(t, err)
		assert.Equal(t, tagInput.Name, tag.Name, "Tag should have correct name")
		assert.Equal(t, tagInput.Color, tag.Color, "Tag should have correct color")
	})

	t.Run("Try to create existing tag name", func(t *testing.T) {
		prepareTestDatabase()

		tagInput := gentypes.CreateTagInput{
			Name:  "existing tag",
			Color: "#234",
		}

		tag, err := courseRepo.CreateTag(tagInput)
		assert.NotNil(t, err)
		assert.Equal(t, &errors.ErrTagAlreadyExists, err, "Should return tag exists err")
		assert.Equal(t, models.Tag{}, tag)
	})
}

func TestManyCourseTags(t *testing.T) {
	prepareTestDatabase()
	courseToTags, err := courseRepo.ManyCourseTags([]uint{3, 2})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(courseToTags[2]))
	assert.Equal(t, 3, len(courseToTags[3]))

	// Check 2 has right tags
	var uuids []gentypes.UUID
	for _, i := range courseToTags[2] {
		uuids = append(uuids, i.UUID)
	}

	uuid, _ := gentypes.StringToUUID("00000000-0000-0000-0000-000000000002")
	assert.Contains(t, uuids, uuid)
	uuid, _ = gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")
	assert.Contains(t, uuids, uuid)

	// Check id 3 has right tags
	uuids = []gentypes.UUID{}
	for _, i := range courseToTags[3] {
		uuids = append(uuids, i.UUID)
	}

	uuid, _ = gentypes.StringToUUID("00000000-0000-0000-0000-000000000002")
	assert.Contains(t, uuids, uuid)
	uuid, _ = gentypes.StringToUUID("00000000-0000-0000-0000-000000000001")
	assert.Contains(t, uuids, uuid)
	uuid, _ = gentypes.StringToUUID("00000000-0000-0000-0000-000000000003")
	assert.Contains(t, uuids, uuid)

}

func TestGetTagsByLessonUUID(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name    string
		uuid    string
		wantErr interface{}
		wantLen int
	}{
		{
			"UUID must be valid",
			"ayyyyyyy",
			&errors.ErrWhileHandling,
			0,
		},
		{
			"Get all tags",
			"00000000-0000-0000-0000-000000000002",
			nil,
			3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tags, err := courseRepo.GetTagsByLessonUUID(test.uuid)

			assert.Equal(t, test.wantErr, err)
			assert.Len(t, tags, test.wantLen)
		})
	}
}
