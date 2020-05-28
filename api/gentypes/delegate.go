package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type Delegate struct {
	User
	TTC_ID          string
	Email           string
	CompanyUUID     UUID
	ProfileImageURL *string
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
	CreateUserInput
	Email       string `valid:"email"`
	CompanyUUID *UUID
}

func (m *CreateDelegateInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}
