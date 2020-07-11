package application

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

type BlogApp interface {
	CreateBlog(input gentypes.CreateBlogInput) (gentypes.Blog, error)
	UpdateBlog(input gentypes.UpdateBlogInput) (gentypes.Blog, error)
	BlogHeaderImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	UpdateBlogHeaderImage(blogUUID gentypes.UUID, token string) (string, error)
	BlogBodyImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	GetBlogBodyImages(blogUUID gentypes.UUID) ([]gentypes.BlogImage, error)
	GetBlogsByUUID(uuids []string) ([]gentypes.Blog, error)
	GetBlogs(page *gentypes.Page, orderBy *gentypes.OrderBy) ([]gentypes.Blog, gentypes.PageInfo, error)
}

type blogAppImpl struct {
	grant          *middleware.Grant
	blogRepository middleware.BlogRepository
}

func NewBlogApp(grant *middleware.Grant) BlogApp {
	return &blogAppImpl{
		grant:          grant,
		blogRepository: middleware.NewBlogRepository(&grant.Logger),
	}
}

func (b *blogAppImpl) blogToGentype(blog models.Blog) gentypes.Blog {
	url := uploads.GetImgixURL(blog.HeaderImageKey)

	createdAt := blog.CreatedAt.Format(time.RFC3339)
	updatedAt := blog.UpdatedAt.Format(time.RFC3339)
	return gentypes.Blog{
		CreatedAt:      createdAt,
		UpdatedAt:      &updatedAt,
		UUID:           blog.UUID,
		Title:          blog.Title,
		Body:           blog.Body,
		HeaderImageURL: url,
		Category: gentypes.Category{
			UUID:  blog.Category.UUID,
			Name:  blog.Category.Name,
			Color: blog.Category.Color,
		},
		Author: gentypes.BlogAuthor{
			FirstName: blog.Author.FirstName,
			LastName:  blog.Author.LastName,
		},
	}
}

func (b *blogAppImpl) blogsToGentype(blogs []models.Blog) []gentypes.Blog {
	var gens []gentypes.Blog
	for _, blog := range blogs {
		gens = append(gens, b.blogToGentype(blog))
	}
	return gens
}

func (b *blogAppImpl) CreateBlog(input gentypes.CreateBlogInput) (gentypes.Blog, error) {
	if !b.grant.IsAdmin {
		return gentypes.Blog{}, &errors.ErrUnauthorized
	}

	if input.AuthorUUID == nil {
		input.AuthorUUID = &b.grant.Claims.UUID
	}
	blog, err := b.blogRepository.CreateBlog(input)

	if err != nil {
		return gentypes.Blog{}, err
	}

	if input.HeaderImageToken != nil {
		key, err := b.UpdateBlogHeaderImage(blog.UUID, *input.HeaderImageToken)

		if err != nil {
			return gentypes.Blog{}, err
		}

		blog.HeaderImageKey = key
	}

	if input.BodyImages != nil {
		err := b.BlogImagesUploadSuccess(blog.UUID, *input.BodyImages)

		if err != nil {
			return gentypes.Blog{}, err
		}
	}

	return b.blogToGentype(blog), err
}

func (b *blogAppImpl) BlogHeaderImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !b.grant.IsAdmin {
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

func (b *blogAppImpl) UpdateBlogHeaderImage(blogUUID gentypes.UUID, token string) (string, error) {
	if !b.grant.IsAdmin {
		return "", &errors.ErrUnauthorized
	}

	s3key, err := uploads.VerifyUploadSuccess(token, "blogHeaderImage")
	if err != nil {
		return "", err
	}

	err = b.blogRepository.UploadHeaderImage(blogUUID, s3key)
	if err != nil {
		return "", err
	}

	return s3key, nil
}

func (b *blogAppImpl) BlogBodyImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !b.grant.IsAdmin {
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
func (b *blogAppImpl) BlogImagesUploadSuccess(blog gentypes.UUID, imgs []gentypes.BlogImageInput) error {
	if !b.grant.IsAdmin {
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

	err := b.blogRepository.UploadBlogImages(blog, keyMap)
	if err != nil {
		for _, v := range keyMap {
			_ = uploads.DeleteImageFromKey(v)
		}
		return &errors.ErrWhileHandling
	}

	return nil
}

func (b *blogAppImpl) GetBlogBodyImages(blogUUID gentypes.UUID) ([]gentypes.BlogImage, error) {
	imgs, err := b.blogRepository.GetBlogImages(blogUUID)

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

func (b *blogAppImpl) GetBlogsByUUID(uuids []string) ([]gentypes.Blog, error) {
	blogs, err := b.blogRepository.GetBlogsByUUID(uuids)

	if err != nil {
		return []gentypes.Blog{}, err
	}

	return b.blogsToGentype(blogs), nil
}

func (b *blogAppImpl) GetBlogs(
	page *gentypes.Page,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Blog, gentypes.PageInfo, error) {
	blogs, pageInfo, err := b.blogRepository.GetBlogs(page, orderBy)

	return b.blogsToGentype(blogs), pageInfo, err
}

func (b *blogAppImpl) UpdateBlog(input gentypes.UpdateBlogInput) (gentypes.Blog, error) {
	if !b.grant.IsAdmin {
		return gentypes.Blog{}, &errors.ErrUnauthorized
	}

	blog, err := b.blogRepository.UpdateBlog(input)

	if input.HeaderImageToken != nil {
		key, err := b.UpdateBlogHeaderImage(blog.UUID, *input.HeaderImageToken)
		if err != nil {
			return gentypes.Blog{}, err
		}

		blog.HeaderImageKey = key
	}
	if input.BodyImages != nil {
		olds, err := b.blogRepository.GetBlogImages(blog.UUID)
		if err != nil {
			return gentypes.Blog{}, err
		}

		// If an existing img is not provided, it will be deleted
		var to_delete []string
		for _, oldie := range olds {
			to_keep := false
			for _, img := range *input.BodyImages {
				to_keep = to_keep || oldie.BodyID == img.JsonID
			}
			if !to_keep {
				to_delete = append(to_delete, oldie.BodyID)
			}
		}
		err = b.blogRepository.DeleteBlogImages(blog.UUID, &to_delete)
		if err != nil {
			return gentypes.Blog{}, err
		}

		// It should ignore images that are existant and upload new ones
		var imgs []gentypes.BlogImageInput
		for _, img := range *input.BodyImages {
			if img.Token != nil {
				imgs = append(imgs, gentypes.BlogImageInput{
					JsonID: img.JsonID,
					Token:  *img.Token,
				})
			}
		}
		err = b.BlogImagesUploadSuccess(blog.UUID, imgs)
		if err != nil {
			return gentypes.Blog{}, err
		}
	}

	return b.blogToGentype(blog), err
}
