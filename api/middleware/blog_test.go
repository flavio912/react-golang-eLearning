package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

var blogRepo = middleware.NewBlogRepository(&logging.Logger{})

func TestCreateBlog(t *testing.T) {
	prepareTestDatabase()

	t.Run("Validates input", func(t *testing.T) {
		invalidInput := gentypes.CreateBlogInput{
			Title: "How to golang",
			Body:  "not json",
		}

		blog, err := blogRepo.CreateBlog(invalidInput)

		assert.Equal(t, invalidInput.Validate(), err)
		assert.Equal(t, models.Blog{}, blog)
	})

	t.Run("Creates a blog", func(t *testing.T) {
		adminUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
		input := gentypes.CreateBlogInput{
			Title:        "How to golang",
			Body:         "{}",
			CategoryUUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			AuthorUUID:   &adminUUID,
		}

		blog, err := blogRepo.CreateBlog(input)

		assert.Nil(t, err)
		assert.Equal(t, input.Title, blog.Title)
		assert.Equal(t, adminUUID, blog.Author.UUID)
		assert.Equal(t, input.CategoryUUID, blog.Category.UUID)
	})
}

func TestGetBlogImages(t *testing.T) {
	prepareTestDatabase()
	keyMap := map[string]string{
		"img1": "key1",
		"img2": "key2",
	}
	blogUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
	_ = blogRepo.UploadBlogImages(blogUUID, keyMap)

	t.Run("Gets all images of blog", func(t *testing.T) {
		imgs, err := blogRepo.GetBlogImages(blogUUID)

		assert.Nil(t, err)
		assert.Len(t, imgs, 2)
	})
}

func TestDeleteBlogImages(t *testing.T) {
	prepareTestDatabase()

	keyMap := map[string]string{
		"img1": "key1",
		"img2": "key2",
		"img3": "key3",
		"img4": "key4",
	}
	blogUUID := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001")
	_ = blogRepo.UploadBlogImages(blogUUID, keyMap)

	t.Run("Deletes specific images of blog", func(t *testing.T) {
		to_delete := []string{
			"img1", "img3",
		}
		err := blogRepo.DeleteBlogImages(blogUUID, &to_delete)

		assert.Nil(t, err)

		imgs, err := blogRepo.GetBlogImages(blogUUID)

		assert.Nil(t, err)
		assert.Len(t, imgs, len(keyMap)-len(to_delete))
	})

	t.Run("Deletes all images of blog", func(t *testing.T) {
		err := blogRepo.DeleteBlogImages(blogUUID, nil)

		assert.Nil(t, err)

		blogs, err := blogRepo.GetBlogImages(blogUUID)

		assert.Nil(t, err)
		assert.Len(t, blogs, 0)
	})
}

func TestGetBlogsByUUIDs(t *testing.T) {
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
			"Must get only existed blogs",
			[]string{
				"00000000-0000-0000-0000-000000000001",
				"00000000-0000-0000-0000-000000000033",
			},
			nil,
			1,
		},
		{
			"Must get all blogs",
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
			b, err := blogRepo.GetBlogsByUUID(test.uuids)

			assert.Equal(t, test.wantErr, err)
			assert.Len(t, b, test.wantLen)
		})
	}
}

func TestGetBlogs(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return ALL blogs", func(t *testing.T) {
		blogs, _, err := blogRepo.GetBlogs(nil, nil)
		assert.Nil(t, err)
		assert.Len(t, blogs, 3)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		blogs, pageInfo, err := blogRepo.GetBlogs(&page, nil)
		assert.Nil(t, err)
		assert.Len(t, blogs, 2)
		assert.Equal(t, gentypes.PageInfo{Total: 3, Given: 2, Limit: limit}, pageInfo)
	})

	t.Run("Should order", func(t *testing.T) {
		asc := false
		order := gentypes.OrderBy{Field: "created_at", Ascending: &asc}

		blogs, _, err := blogRepo.GetBlogs(nil, &order)
		assert.Nil(t, err)
		assert.Len(t, blogs, 3)
		assert.Equal(t, "2020-03-08T13:53:37Z", blogs[0].CreatedAt.Format(time.RFC3339))
	})
}

func TestUpdateBlog(t *testing.T) {
	prepareTestDatabase()

	t.Run("Blog must exist", func(t *testing.T) {
		uuidZero := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")
		b, err := blogRepo.UpdateBlog(gentypes.UpdateBlogInput{UUID: uuidZero})
		assert.Equal(t, errors.ErrBlogNotFound(uuidZero.String()), err)
		assert.Equal(t, models.Blog{}, b)
	})

	t.Run("Updates existing blog", func(t *testing.T) {
		input := gentypes.UpdateBlogInput{
			UUID:         gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			Title:        helpers.StringPointer("Everything You Need To Know About CNNs"),
			Body:         helpers.StringPointer(`{"ayy" : "yoo"}`),
			CategoryUUID: helpers.UUIDPointer(gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002")),
		}

		blog, err := blogRepo.UpdateBlog(input)

		assert.Nil(t, err)
		assert.Equal(t, input.UUID, blog.UUID)
		assert.Equal(t, *input.Title, blog.Title)
		assert.Equal(t, *input.Body, blog.Body)
		assert.Equal(t, *input.CategoryUUID, blog.Category.UUID)
		assert.NotEqual(t, "", blog.UpdatedAt.Format(time.RFC3339))

		asc := true
		blogs, _, errs := blogRepo.GetBlogs(
			nil,
			&gentypes.OrderBy{
				Ascending: &asc,
				Field:     "updated_at",
			},
		)

		assert.Nil(t, errs)
		assert.Equal(t, blogs[0].UUID, blog.UUID)
		assert.Equal(t, blogs[0].UpdatedAt.Format(time.RFC3339), blog.UpdatedAt.Format(time.RFC3339))
	})
}
