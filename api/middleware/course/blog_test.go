package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateBlog(t *testing.T) {
	prepareTestDatabase()

	t.Run("Creates a blog", func(t *testing.T) {
		adminUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		input := gentypes.CreateBlogInput{
			Title:        "How to golang",
			Body:         "{}",
			CategoryUUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			AuthorUUID:   &adminUUID,
		}

		blog, err := courseRepo.CreateBlog(input)

		assert.Nil(t, err)
		assert.Equal(t, input.Title, blog.Title)
		assert.Equal(t, adminUUID, blog.Author.UUID)
		assert.Equal(t, input.CategoryUUID, blog.CategoryUUID)
	})
}
