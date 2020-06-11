package gentypes

type AccessType string

const (
	Restricted AccessType = "restricted"
	Open       AccessType = "open"
)

type StructureElement string

const (
	ModuleType StructureElement = "module"
	TestType   StructureElement = "test"
	LessonType StructureElement = "lesson"
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
	Excerpt            *string `valid:"json"`
	Introduction       *string `valid:"json"`
	HowToComplete      *string `valid:"json"`
	HoursToComplete    *float64
	WhatYouLearn       *[]string
	Requirements       *[]string
	BackgroundCheck    *bool
	AccessType         *AccessType
	Price              *float64
	Color              *string `valid:"hexcolor"`
	Tags               *[]UUID
	SpecificTerms      *string `valid:"json"`
	BannerImageSuccess *string
}

type CourseItem struct {
	Type  StructureElement
	UUID  UUID
	Items []ModuleItem
}

type ModuleItem struct {
	Type StructureElement
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
	Introduction    string `valid:"json"`
	HowToComplete   string `valid:"json"`
	HoursToComplete float64
	WhatYouLearn    []string
	Requirements    []string
	Excerpt         string `valid:"json"`
	SpecificTerms   string `valid:"json"`
	CategoryUUID    *UUID
}

type CourseFilter struct {
	Name            *string
	AccessType      *AccessType
	BackgroundCheck *bool
	Price           *float64
	AllowedToBuy    *bool
}

type PurchaseCoursesInput struct {
	Courses                []UUID
	Users                  []UUID
	ExtraInvoiceEmail      *string `valid:"email"`
	AcceptedTerms          bool
	BackgroundCheckConfirm *bool
}

type PurchaseCoursesResponse struct {
	TransactionComplete bool
	StripeClientSecret  *string
}
