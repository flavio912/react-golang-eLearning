package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *courseAppImpl) blogToGentype(blog models.Blog) gentypes.Blog {
	return gentypes.Blog{
		UUID:           blog.UUID,
		Title:          blog.Title,
		Body:           blog.Body,
		HeaderImageURL: blog.HeaderImageURL,
		Category: gentypes.Category{
			UUID: blog.CategoryUUID,
		},
		Author: gentypes.BlogAuthor{
			FirstName: blog.Author.FirstName,
			LastName:  blog.Author.LastName,
		},
	}
}

func (c *courseAppImpl) CreateBlog(input gentypes.CreateBlogInput) (gentypes.Blog, error) {
	if !c.grant.IsAdmin {
		return gentypes.Blog{}, &errors.ErrUnauthorized
	}

	if input.AuthorUUID == nil {
		input.AuthorUUID = &c.grant.Claims.UUID
	}
	blog, err := c.coursesRepository.CreateBlog(input)
	return c.blogToGentype(blog), err
}
