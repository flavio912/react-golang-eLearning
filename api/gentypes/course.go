package gentypes

import "github.com/asaskevich/govalidator"

type AccessType string

const (
	Restricted AccessType = "restricted"
	Open       AccessType = "open"
)

type CourseStatus string

const (
	CourseComplete   CourseStatus = "complete"
	CourseFailed     CourseStatus = "failed"
	CourseIncomplete CourseStatus = "incomplete"
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
	ID                   *int32
	Name                 *string
	CategoryUUID         *UUID
	Excerpt              *string
	Introduction         *string
	HowToComplete        *string
	HoursToComplete      *float64
	WhatYouLearn         *[]string
	Requirements         *[]string
	BackgroundCheck      *bool
	AccessType           *AccessType
	Price                *float64
	Color                *string `valid:"hexcolor"`
	Tags                 *[]UUID
	SpecificTerms        *string
	BannerImageSuccess   *string
	CertificateType      *UUID
	ExpirationToEndMonth *bool
	ExpiresInMonths      *int32
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
	BannerImageURL  *string
}

type ActiveCourse struct {
	CourseID       uint
	MinutesTracked float64
	Status         CourseStatus
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

type DeleteCourseInput struct {
	ID int32 `valid:"required"`
}

func (d *DeleteCourseInput) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}
