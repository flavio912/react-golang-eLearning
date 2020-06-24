package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) blogToGentype(blog models.Blog) gentypes.Blog {
	url := uploads.GetImgixURL(blog.HeaderImageKey)

	return gentypes.Blog{
		UUID:           blog.UUID,
		Title:          blog.Title,
		Body:           blog.Body,
		HeaderImageURL: url,
		Category: gentypes.Category{
			UUID: blog.CategoryUUID,
		},
		Author: gentypes.BlogAuthor{
			FirstName: blog.Author.FirstName,
			LastName:  blog.Author.LastName,
		},
	}
}

func (c *courseAppImpl) blogsToGentype(blogs []models.Blog) []gentypes.Blog {
	var gens []gentypes.Blog
	for _, blog := range blogs {
		gens = append(gens, c.blogToGentype(blog))
	}
	return gens
}

func (c *courseAppImpl) CreateBlog(input gentypes.CreateBlogInput) (gentypes.Blog, error) {
	if !c.grant.IsAdmin {
		return gentypes.Blog{}, &errors.ErrUnauthorized
	}

	if input.AuthorUUID == nil {
		input.AuthorUUID = &c.grant.Claims.UUID
	}
	blog, err := c.coursesRepository.CreateBlog(input)

	if err != nil {
		return gentypes.Blog{}, err
	}

	if input.HeaderImageToken != nil {
		err = c.UpdateBlogHeaderImage(blog.UUID, *input.HeaderImageToken)

		if err != nil {
			return gentypes.Blog{}, err
		}
	}

	if input.BodyImages != nil {
		err := c.BlogImagesUploadSuccess(blog.UUID, *input.BodyImages)

		if err != nil {
			return gentypes.Blog{}, err
		}
	}

	return c.blogToGentype(blog), err
}

func (c *courseAppImpl) BlogHeaderImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,
		imageMeta.ContentLength,
		[]string{"jpg", "png"},
		int32(15000000),
		"blog_header_image",
		"blogHeaderImage",
	)

	return url, successToken, err
}

func (c *courseAppImpl) UpdateBlogHeaderImage(blogUUID gentypes.UUID, token string) error {
	if !c.grant.IsAdmin {
		return &errors.ErrUnauthorized
	}

	s3key, err := uploads.VerifyUploadSuccess(token, "blogHeaderImage")
	if err != nil {
		return err
	}

	err = c.coursesRepository.UploadHeaderImage(blogUUID, s3key)
	if err != nil {
		return err
	}

	return nil
}

func (c *courseAppImpl) BlogBodyImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,
		imageMeta.ContentLength,
		[]string{"jpg", "png", "gif"}, // gif just for fun, can remove later
		int32(10000000),
		"blog_images",
		"blogBodyImage",
	)

	return url, successToken, err
}

// BlogImagesUploadSuccess verifies tokens of images and uploads them to related blog
func (c *courseAppImpl) BlogImagesUploadSuccess(blog gentypes.UUID, imgs []gentypes.BlogImageInput) error {
	if !c.grant.IsAdmin {
		return &errors.ErrUnauthorized
	}

	keyMap := make(map[string]string)
	for _, img := range imgs {
		s3key, err := uploads.VerifyUploadSuccess(img.Token, "blogBodyImage")
		if err != nil {
			return err
		}
		keyMap[img.JsonID] = s3key
	}

	err := c.coursesRepository.UploadBlogImages(blog, keyMap)
	if err != nil {
		for _, v := range keyMap {
			_ = uploads.DeleteImageFromKey(v)
		}
		return &errors.ErrWhileHandling
	}

	return nil
}

func (c *courseAppImpl) GetBlogBodyImages(blogUUID gentypes.UUID) ([]gentypes.BlogImage, error) {
	imgs, err := c.coursesRepository.GetBlogImages(blogUUID)

	if err != nil {
		return []gentypes.BlogImage{}, err
	}

	var gens []gentypes.BlogImage
	for _, img := range imgs {
		url := uploads.GetImgixURL(img.S3key)
		gens = append(gens, gentypes.BlogImage{
			JsonID: img.BodyID,
			Url:    url,
		})
	}

	return gens, nil
}

func (c *courseAppImpl) GetBlogsByUUID(uuids []string) ([]gentypes.Blog, error) {
	blogs, err := c.coursesRepository.GetBlogsByUUID(uuids)

	if err != nil {
		return []gentypes.Blog{}, err
	}

	return c.blogsToGentype(blogs), nil
}
