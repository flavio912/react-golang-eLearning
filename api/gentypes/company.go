package gentypes

import (
	"github.com/asaskevich/govalidator"
)

type Company struct {
	CreatedAt    *string `valid:"rfc3339"`
	IsContract   bool
	ContactEmail string
	ContactPhone *string
	Approved     *bool
	UUID         UUID
	Name         string
	AddressID    uint
	LogoURL      *string
}

type CompanyFilter struct {
	UUID     *string
	Name     *string
	Approved *bool
}

type OrderBy struct {
	Ascending *bool //defaults to false, thus decending
	Field     string
}

type CreateCompanyInput struct {
	CompanyName  string  `valid:"required"`
	AddressLine1 string  `valid:"required"`
	AddressLine2 string  `valid:"required"`
	County       string  `valid:"required"`
	PostCode     string  `valid:"required,stringlength(6|7)"` // 6 or 7 depending on whether space in middle
	Country      string  `valid:"required"`
	ContactEmail string  `valid:"required,email"`
	ContactPhone *string `valid:"numeric"`
	LogoToken    *string
}

func (c *CreateCompanyInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type UpdateCompanyInput struct {
	UUID         UUID
	CompanyName  *string
	Approved     *bool
	IsContract   *bool
	AddressLine1 *string
	AddressLine2 *string
	County       *string `valid:"optional"`
	PostCode     *string `valid:"optional,alphanum,stringlength(6|7)"` // 6 or 7 depending on whether space in middle
	Country      *string `valid:"optional"`
	ContactEmail *string `valid:"optional,email"`
}

func (c *UpdateCompanyInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

type CreateCompanyRequestManager struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	JobTitle  string `valid:"required"`
	Telephone string `valid:"required,numeric"`
	Email     string `valid:"required,email"`
}

func (c *CreateCompanyRequestManager) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
