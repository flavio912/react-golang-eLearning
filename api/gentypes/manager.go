package gentypes

import (
	"github.com/asaskevich/govalidator"
)

// Manager - CompanyManager graphQL type
type Manager struct {
	User
}

type ManagersFilter struct {
	Email     *string `valid:"email"`
	Name      *string `valid:"-"`
	JobTitle  *string `valid:"-"`
	Telephone *string `valid:"numeric"`
	UUID      *string `valid:"uuidv4"`
}

func (m *ManagersFilter) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

type AddManagerInput struct {
	FirstName string
	LastName  string
	Email     string
	JobTitle  string
	Telephone string
	Password  string
}

// ManagerLoginInput - ManagerLogin graphQL input
type ManagerLoginInput struct {
	Email    string
	Password string
}
