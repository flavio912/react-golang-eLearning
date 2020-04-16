package gentypes

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// Manager - CompanyManager graphQL type
type Manager struct {
	User
	ProfileImageURL *string
	CompanyID       uuid.UUID
}

type ManagersFilter struct {
	Email     *string `valid:"-"`
	Name      *string `valid:"-"`
	JobTitle  *string `valid:"-"`
	Telephone *string `valid:"numeric"`
	UUID      *string `valid:"uuidv4"`
}

func (m *ManagersFilter) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

type CreateManagerInput struct {
	CompanyUUID *string `valid:"uuidv4"`
	FirstName   string  `valid:"required,alpha"`
	LastName    string  `valid:"required,alpha"`
	Email       string  `valid:"required,email"`
	JobTitle    string  `valid:"required"`
	Telephone   string  `valid:"required,numeric"`
	Password    string  `valid:"required,stringlength(5|30)"`
}

func (m *CreateManagerInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

type DeleteManagerInput struct {
	UUID string `valid:"uuidv4"`
}

func (m *DeleteManagerInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

// ManagerLoginInput - ManagerLogin graphQL input
type ManagerLoginInput struct {
	Email    string
	Password string
}
