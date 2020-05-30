package gentypes

type UserFilter struct {
	UUID      *string `valid:"uuidv4"`
	Name      *string
	JobTitle  *string
	Telephone *string `valid:"numeric"`
}

type CreateUserInput struct {
	FirstName string `valid:"required,alpha"`
	LastName  string `valid:"required,alpha"`
	JobTitle  string `valid:"required"`
	Telephone string `valid:"numeric"`
	Password  string `valid:"required,stringlength(5|30)"`
}
