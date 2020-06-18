package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type BlogAuthor struct {
	FirstName     string
	LastName      string
	ProfilePicURL *string
}

type Blog struct {
	UUID           UUID
	Title          string
	Body           string
	Category       Category
	HeaderImageURL string
	Author         BlogAuthor
}

type CreateBlogInput struct {
	Title          string
	Body           string `valid:"json"`
	CategoryUUID   UUID
	HeaderImageURL *string
	AuthorUUID     *UUID
}

func (c *CreateBlogInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
