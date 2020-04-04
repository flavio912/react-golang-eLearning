package gentypes

import "github.com/asaskevich/govalidator"

type Admin struct {
	UUID      string
	Email     string
	FirstName string
	LastName  string
}

// Key gets the admin primary identifier
func (admin *Admin) Key() string {
	return admin.UUID
}

// AdminPage - a list of admins
type AdminPage struct {
	Edges    Admin
	PageInfo PageInfo
}

// AdminLoginInput -
type AdminLoginInput struct {
	Email    string
	Password string
}

type AddAdminInput struct {
	FirstName string `valid:"alpha,required"`
	LastName  string `valid:"alpha,required"`
	Email     string `valid:"email,required"`
	Password  string `valid:"stringlength(8|30),required"`
}

func (m *AddManagerInput) Validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}
