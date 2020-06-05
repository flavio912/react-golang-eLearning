package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type Lesson struct {
	UUID  UUID
	Title string
	Tags  []Tag
	Text  string
}

type CreateLessonInput struct {
	Title string `valid:"required"`
	Tags  *[]UUID
	Text  string `valid:"json"`
}

func (c *CreateLessonInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
