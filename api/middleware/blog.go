package middleware

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type BlogRepository interface {
	CreateBlog(input gentypes.CreateBlogInput) (models.Blog, error)
	UpdateBlog(input gentypes.UpdateBlogInput) (models.Blog, error)
	UploadHeaderImage(blogUUID gentypes.UUID, key string) error
	UploadBlogImages(blog gentypes.UUID, imgs map[string]string) error
	DeleteBlogImages(blogUUID gentypes.UUID) error
	GetBlogImages(blogUUID gentypes.UUID) ([]models.BlogImage, error)
	GetBlogsByUUID(uuids []string) ([]models.Blog, error)
	GetBlogs(page *gentypes.Page, orderBy *gentypes.OrderBy) ([]models.Blog, gentypes.PageInfo, error)
}

type blogRepoImpl struct {
	Logger *logging.Logger
}

func NewBlogRepository(logger *logging.Logger) BlogRepository {
	return &blogRepoImpl{
		Logger: logger,
	}
}

// CreateBlog creates a blog with author as an admin
func (c *blogRepoImpl) CreateBlog(input gentypes.CreateBlogInput) (models.Blog, error) {
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

	var category models.Category
	query = database.GormDB.Where("uuid = ?", input.CategoryUUID).First(&category)
	if query.Error != nil {
		if query.RecordNotFound() {
			c.Logger.Logf(sentry.LevelError, query.Error, "Unable to find category %s", input.CategoryUUID)
			return models.Blog{}, &errors.ErrNotFound
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to create blog because of category %s", input.CategoryUUID)
		return models.Blog{}, &errors.ErrWhileHandling
	}

	blog := models.Blog{
		Title:    input.Title,
		Body:     input.Body,
		Category: category,
		Author:   admin,
	}

	query = database.GormDB.Create(&blog)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create blog")
		return models.Blog{}, query.Error
	}

	return blog, nil
}

func (c *blogRepoImpl) UploadHeaderImage(blogUUID gentypes.UUID, key string) error {
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

func (c *blogRepoImpl) UploadBlogImages(blog gentypes.UUID, imgs map[string]string) error {
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

// DeleteBlogImages deletes images inside blog body
func (c *blogRepoImpl) DeleteBlogImages(blogUUID gentypes.UUID) error {
	query := database.GormDB.Delete(models.BlogImage{}, "blog_uuid = ?", blogUUID)
	if query.Error != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to delete blog's images: %s", blogUUID)
		return &errors.ErrDeleteFailed
	}

	return nil
}

func (c *blogRepoImpl) GetBlogImages(blogUUID gentypes.UUID) ([]models.BlogImage, error) {
	var imgs []models.BlogImage
	query := database.GormDB.Where("blog_uuid = ?", blogUUID).Find(&imgs)

	if query.Error != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to get blog images of uuid %s", blogUUID)
		return []models.BlogImage{}, query.Error
	}

	return imgs, nil
}

func (c *blogRepoImpl) GetBlogsByUUID(uuids []string) ([]models.Blog, error) {
	var blogs []models.Blog
	query := database.GormDB.Where("uuid IN (?)", uuids).Find(&blogs)

	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get blogs")
		return []models.Blog{}, &errors.ErrWhileHandling
	}

	return blogs, nil
}

func (c *blogRepoImpl) GetBlogs(
	page *gentypes.Page,
	orderBy *gentypes.OrderBy,
) ([]models.Blog, gentypes.PageInfo, error) {
	var blogs []models.Blog

	var count int32
	query := database.GormDB
	countErr := query.Model(&models.Blog{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		c.Logger.Log(sentry.LevelError, countErr, "Unable to count blogs")
		return blogs, gentypes.PageInfo{}, countErr
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"created_at", "updated_at"}, "created_at DESC")
	if orderErr != nil {
		return blogs, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&blogs)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to find blogs")
		return []models.Blog{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return blogs, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(blogs)),
	}, nil
}

func (c *blogRepoImpl) UpdateBlog(input gentypes.UpdateBlogInput) (models.Blog, error) {
	// Validate input
	if err := input.Validate(); err != nil {
		return models.Blog{}, err
	}

	var blog models.Blog
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&blog)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Blog{}, errors.ErrBlogNotFound(input.UUID.String())
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to find blog to update with UUID: %s", input.UUID)
		return models.Blog{}, &errors.ErrWhileHandling
	}

	if input.Title != nil {
		blog.Title = *input.Title
	}
	if input.Body != nil {
		blog.Body = *input.Body
	}
	if input.CategoryUUID != nil {
		var category models.Category
		query = database.GormDB.Where("uuid = ?", *input.CategoryUUID).First(&category)
		if query.Error != nil {
			if query.RecordNotFound() {
				return models.Blog{}, &errors.ErrNotFound
			}

			c.Logger.Logf(sentry.LevelError, query.Error, "Unable to update blog %s because of category: %s", input.UUID, *input.CategoryUUID)
			return models.Blog{}, &errors.ErrWhileHandling
		}

		blog.Category = category
	}

	save := database.GormDB.Save(&blog)
	if save.Error != nil {
		c.Logger.Logf(sentry.LevelError, save.Error, "Error updating blog with UUID: %s", input.UUID)
		return models.Blog{}, &errors.ErrWhileHandling
	}

	return blog, nil
}
