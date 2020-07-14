package gentypes

import (
	"github.com/asaskevich/govalidator"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Delegate struct {
	CreatedAt       *string
	UUID            UUID
	FirstName       string
	LastName        string
	Telephone       *string
	JobTitle        string
	LastLogin       string
	TTC_ID          string
	Email           *string
	CompanyUUID     UUID
	ProfileImageURL *string
	CourseTakerUUID UUID
}

type DelegatesFilter struct {
	UserFilter
	Email  *string
	TTC_ID *string
}

func (d *DelegatesFilter) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}

type CreateDelegateInput struct {
	FirstName               string  `valid:"required,alpha"`
	LastName                string  `valid:"required,alpha"`
	JobTitle                string  `valid:"required"`
	Telephone               *string `valid:"numeric"`
	Email                   *string `valid:"email"`
	ProfileImageUploadToken *string
	CompanyUUID             *UUID
	GeneratePassword        *bool
}

func (m *CreateDelegateInput) Validate() error {

	// If no email, user must specify 'generatePassword' = true
	if m.Email == nil {
		if m.GeneratePassword != nil {
			if !*m.GeneratePassword {
				return &errors.ErrNoEmailProvided
			}
		} else {
			return &errors.ErrNoEmailProvided
		}
	}

	_, err := govalidator.ValidateStruct(m)
	return err
}

type UpdateDelegateInput struct {
	UUID                    UUID `valid:"required"`
	CompanyUUID             *UUID
	FirstName               *string `valid:"alpha"`
	LastName                *string `valid:"alpha"`
	JobTitle                *string
	Email                   *string `valid:"email"`
	Telephone               *string `valid:"numeric"`
	ProfileImageUploadToken *string
	NewPassword             *string
}

func (u *UpdateDelegateInput) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

type DelegateLoginInput struct {
	TTC_ID   string
	Password string
	NoResp   *bool
}
