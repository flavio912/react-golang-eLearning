package gentypes

import "github.com/asaskevich/govalidator"

type Category struct {
	UUID  UUID
	Name  string
	Color string
}

type CreateCategoryInput struct {
	Name  string
	Color string `valid:"hexcolor"`
}

type UpdateCategoryInput struct {
	UUID  UUID
	Name  *string
	Color *string `valid:"optional,hexcolor"`
}

func (c *UpdateCategoryInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
