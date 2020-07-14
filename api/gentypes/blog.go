package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type BlogAuthor struct {
	FirstName     string
	LastName      string
	ProfilePicURL *string
}

type BlogImage struct {
	JsonID string
	Url    string
}

type Blog struct {
	CreatedAt      string
	UpdatedAt      *string
	UUID           UUID
	Title          string
	Body           string
	Category       Category
	HeaderImageURL string
	Author         BlogAuthor
	BlogBodyImages *[]BlogImage
}

type BlogImageInput struct {
	JsonID string
	Token  string
}

type CreateBlogInput struct {
	Title            string
	Body             string `valid:"json"`
	CategoryUUID     UUID
	HeaderImageToken *string
	AuthorUUID       *UUID
	BodyImages       *[]BlogImageInput
}

func (c *CreateBlogInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type UpdateBlogImageInput struct {
	JsonID string
	Token  *string
}

type UpdateBlogInput struct {
	UUID             UUID `valid:"required"`
	Title            *string
	Body             *string `valid:"json"`
	CategoryUUID     *UUID
	HeaderImageToken *string
	BodyImages       *[]UpdateBlogImageInput
}

func (u *UpdateBlogInput) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

type UpdateBlogHeaderImageInput struct {
	FileSucess UploadFileSuccess
	BlogUUID   UUID
}
