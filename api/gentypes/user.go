package gentypes

type UserType string

const (
	ManagerType    UserType = "manager"
	DelegateType   UserType = "delegate"
	IndividualType UserType = "individual"
)

// User - User graphQL interface
type User struct {
	CreatedAt       *string
	UUID            *UUID
	Type            UserType
	Email           *string
	FirstName       string
	LastName        string
	Telephone       *string
	JobTitle        *string
	LastLogin       string
	ProfileImageUrl *string
}

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
