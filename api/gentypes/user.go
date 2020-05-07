package gentypes

// User - User graphQL interface
type User struct {
	CreatedAt *string
	UUID      UUID
	Email     string
	FirstName string
	LastName  string
	Telephone string
	JobTitle  string
	LastLogin string
}

type UserFilter struct {
	UUID      *string `valid:"uuidv4"`
	Email     *string
	Name      *string
	JobTitle  *string
	Telephone *string `valid:"numeric"`
}

type CreateUserInput struct {
	FirstName string `valid:"required,alpha"`
	LastName  string `valid:"required,alpha"`
	Email     string `valid:"required,email"`
	JobTitle  string `valid:"required"`
	Telephone string `valid:"required,numeric"`
	Password  string `valid:"required,stringlength(5|30)"`
}
