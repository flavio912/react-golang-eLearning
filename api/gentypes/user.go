package gentypes

type UserType string

const (
	ManagerType    UserType = "manager"
	DelegateType   UserType = "delegate"
	IndividualType UserType = "individual"
)

type ActivityType string

const (
	ActivityNewCourse ActivityType = "newCourse"
	ActivityActivated ActivityType = "activated"
	ActivityCompleted ActivityType = "completedCourse"
	ActivityFailed    ActivityType = "failedCourse"
)

type Activity struct {
	UUID            UUID
	CreatedAt       string
	ActivityType    ActivityType
	CourseTakerUUID UUID
	CourseID        *uint
}

// User - User graphQL interface
type User struct {
	UUID            UUID
	CreatedAt       *string
	Type            UserType
	Email           *string
	FirstName       string
	LastName        string
	Telephone       *string
	JobTitle        *string
	LastLogin       string
	ProfileImageUrl *string
	CourseTakerUUID *UUID
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
