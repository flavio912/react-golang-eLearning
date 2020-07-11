package gentypes

import "github.com/asaskevich/govalidator"

type Tutor struct {
	UUID         UUID
	Name         string
	CIN          int32
	SignatureURL string
}

type CreateTutorInput struct {
	Name           string `valid:"required"`
	CIN            int32
	SignatureToken *string
}

func (c *CreateTutorInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
