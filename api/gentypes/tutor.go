package gentypes

import "github.com/asaskevich/govalidator"

type Tutor struct {
	UUID         UUID
	Name         string
	CIN          string
	SignatureURL string
}

type CreateTutorInput struct {
	Name           string `valid:"required"`
	CIN            string `valid:"required"`
	SignatureToken *string
}

func (c *CreateTutorInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type UpdateTutorSignatureInput struct {
	FileSuccess UploadFileSuccess
	TutorUUID   UUID
}

type UpdateTutorInput struct {
	UUID           UUID `valid:"required"`
	Name           *string
	CIN            *string
	SignatureToken *string
}

func (u *UpdateTutorInput) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

type TutorFilter struct {
	Name *string
	CIN  *string
}
