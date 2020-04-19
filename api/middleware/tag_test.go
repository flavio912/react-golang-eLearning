package middleware_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

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

		tag, err := adminGrant.CreateTag(tagInput)

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

		tag, err := adminGrant.CreateTag(tagInput)
		assert.NotNil(t, err)
		assert.Equal(t, &errors.ErrTagAlreadyExists, err, "Should return tag exists err")
		assert.Equal(t, gentypes.Tag{}, tag)
	})

	t.Run("Access Control Tests", func(t *testing.T) {
		prepareTestDatabase()

		tagInput := gentypes.CreateTagInput{
			Name:  "Fancy new tag",
			Color: "#234",
		}

		tag, err := delegateGrant.CreateTag(tagInput)
		assert.Equal(t, &errors.ErrUnauthorized, err, "Delegate cannot create tags")
		assert.Equal(t, gentypes.Tag{}, tag)

		tag, err = managerGrant.CreateTag(tagInput)
		assert.Equal(t, &errors.ErrUnauthorized, err, "Manager cannot create tags")
		assert.Equal(t, gentypes.Tag{}, tag)

	})
}

func TestGetTagsByCourseInfoIDs(t *testing.T) {
	prepareTestDatabase()
	courseToTags, err := adminGrant.GetTagsByCourseInfoIDs([]uint{3, 2})
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
