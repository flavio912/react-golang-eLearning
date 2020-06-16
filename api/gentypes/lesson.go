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

type LessonFilter struct {
	UUID  *string `valid:"uuidv4"`
	Title *string
	Tags  *[]*UUID
}

func (l *LessonFilter) Validate() error {
	_, err := govalidator.ValidateStruct(l)
	return err
}

type CreateLessonInput struct {
	Title string `valid:"required"`
	Tags  *[]UUID
	Text  string
}

func (c *CreateLessonInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type UpdateLessonInput struct {
	UUID  UUID `valid:"required"`
	Title *string
	Text  *string
	Tags  *[]UUID
}

func (u *UpdateLessonInput) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

type DeleteLessonInput struct {
	UUID UUID `valid:"required"`
}

func (d *DeleteLessonInput) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}
