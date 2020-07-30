package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type Lesson struct {
	UUID           UUID
	Name           string
	Tags           []Tag
	Description    string
	BannerImageURL *string
	VoiceoverURL   *string
	Transcript     *string
	Video          *Video
}

type LessonFilter struct {
	UUID *string `valid:"uuidv4"`
	Name *string
	Tags *[]*UUID
}

func (l *LessonFilter) Validate() error {
	_, err := govalidator.ValidateStruct(l)
	return err
}

type CreateLessonInput struct {
	Name             string `valid:"required"`
	Tags             *[]UUID
	Description      string
	BannerImageToken *string
	VoiceoverToken   *string
	Transcript       *string
	Video            *Video
}

func (c *CreateLessonInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type UpdateLessonInput struct {
	UUID             UUID `valid:"required"`
	Name             *string
	Description      *string
	Tags             *[]UUID
	BannerImageToken *string
	VoiceoverToken   *string
	Transcript       *string
	Video            *Video
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
