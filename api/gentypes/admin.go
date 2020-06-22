package gentypes

import "github.com/asaskevich/govalidator"

type Admin struct {
	UUID      UUID
	Email     string
	FirstName string
	LastName  string
}

// Key gets the admin primary identifier
func (admin *Admin) Key() UUID {
	return admin.UUID
}

// AdminPage - a list of admins
type AdminPage struct {
	Edges    Admin
	PageInfo PageInfo
}

type AdminFilter struct {
	Email string
	Name  string
}

// AdminLoginInput -
type AdminLoginInput struct {
	Email    string
	Password string
}

type CreateAdminInput struct {
	FirstName string `valid:"alpha,required"`
	LastName  string `valid:"alpha,required"`
	Email     string `valid:"email,required"`
	Password  string `valid:"stringlength(8|30),required"`
}

func (m *CreateAdminInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

type UpdateAdminInput struct {
	UUID      UUID
	FirstName *string `valid:"alpha"`
	LastName  *string `valid:"alpha"`
	Email     *string `valid:"email"`
}

func (m *UpdateAdminInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

type DeleteAdminInput struct {
	UUID UUID
}
