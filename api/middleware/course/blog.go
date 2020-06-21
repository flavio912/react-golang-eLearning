package course

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// CreateBlog creates a blog with author as an admin
func (c *coursesRepoImpl) CreateBlog(input gentypes.CreateBlogInput) (models.Blog, error) {
	if err := input.Validate(); err != nil {
		return models.Blog{}, err
	}

	if input.AuthorUUID == nil {
		return models.Blog{}, &errors.ErrWhileHandling
	}

	var admin models.Admin
	query := database.GormDB.Where("uuid = ?", *input.AuthorUUID).First(&admin)
	if query.Error != nil {
		if query.RecordNotFound() {
			c.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin %s", *input.AuthorUUID)
			return models.Blog{}, &errors.ErrAdminNotFound
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to create blog because of admin %s", *input.AuthorUUID)
		return models.Blog{}, &errors.ErrWhileHandling
	}

	blog := models.Blog{
		Title:        input.Title,
		Body:         input.Body,
		CategoryUUID: input.CategoryUUID,
		Author:       admin,
	}

	query = database.GormDB.Create(&blog)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create blog")
		return models.Blog{}, query.Error
	}

	return blog, nil
}

func (c *coursesRepoImpl) UploadHeaderImage(blogUUID gentypes.UUID, key string) error {
	query := database.GormDB.Model(&models.Blog{}).Where("uuid = ?", blogUUID).Update("header_image_key", key)
	if query.Error != nil {
		if query.RecordNotFound() {
			return errors.ErrBlogNotFound(blogUUID.String())
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to upload blog header image")
		return &errors.ErrWhileHandling
	}

	return nil
}

func (c *coursesRepoImpl) UploadBlogImages(blog gentypes.UUID, imgs map[string]string) error {
	query := database.GormDB.Begin()
	for k, v := range imgs {
		img := models.BlogImage{
			BlogUUID: blog,
			BodyID:   k,
			S3key:    v,
		}
		query = query.Create(&img)
	}

	if err := query.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to upload blog images")
		return err
	}

	return nil
}
