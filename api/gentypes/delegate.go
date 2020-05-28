package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type Delegate struct {
	User
	TTC_ID          string
	CompanyUUID     UUID
	ProfileImageURL *string
}

type DelegatesFilter struct {
	UserFilter
	TTC_ID *string
}

func (d *DelegatesFilter) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}

type CreateDelegateInput struct {
	CreateUserInput
	CompanyUUID *UUID
}

func (m *CreateDelegateInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}
