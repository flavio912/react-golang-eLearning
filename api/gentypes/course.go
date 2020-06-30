package gentypes

type AccessType string

const (
	Restricted AccessType = "restricted"
	Open       AccessType = "open"
)

type CourseElement string

const (
	ModuleType CourseElement = "module"
	TestType   CourseElement = "test"
	LessonType CourseElement = "lesson"
)

type CourseType string

const (
	ClassroomCourseType CourseType = "classroom"
	OnlineCourseType    CourseType = "online"
)

type CourseInput struct {
	ID                 *uint
	Name               *string
	CategoryUUID       *UUID
	Excerpt            *string
	Introduction       *string
	HowToComplete      *string
	HoursToComplete    *float64
	WhatYouLearn       *[]string
	Requirements       *[]string
	BackgroundCheck    *bool
	AccessType         *AccessType
	Price              *float64
	Color              *string `valid:"hexcolor"`
	Tags               *[]UUID
	SpecificTerms      *string
	BannerImageSuccess *string
}

type CourseItem struct {
	Type CourseElement
	UUID UUID
}

type Course struct {
	ID              uint
	Name            string
	AccessType      AccessType
	CourseType      CourseType
	AllowedToBuy    bool // Helper field, true if current user is allowed to buy this course
	BackgroundCheck bool
	Price           float64
	Tags            []Tag
	Color           string `valid:"hexcolor"`
	Introduction    string
	HowToComplete   string
	HoursToComplete float64
	WhatYouLearn    []string
	Requirements    []string
	Excerpt         string
	SpecificTerms   string
	CategoryUUID    *UUID
}

type ActiveCourse struct {
	CourseID       uint
	CurrentAttempt uint
	MinutesTracked float64
}

type CourseFilter struct {
	Name            *string
	AccessType      *AccessType
	BackgroundCheck *bool
	Price           *float64
	AllowedToBuy    *bool
}

type PurchaseCoursesInput struct {
	Courses                []int32
	Users                  []UUID
	ExtraInvoiceEmail      *string `valid:"email"`
	AcceptedTerms          bool
	BackgroundCheckConfirm *bool
}

type PurchaseCoursesResponse struct {
	TransactionComplete bool
	StripeClientSecret  *string
}
