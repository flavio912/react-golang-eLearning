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

type DelegateFilter struct {
	UserFilter
	TTC_ID *string
}

func (d *DelegateFilter) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}

type CreateDelegateInput struct {
	CreateUserInput
	TTC_ID      string
	CompanyUUID *UUID
}

func (m *CreateDelegateInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}
