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

type CourseInput struct {
	UUID               *UUID
	Name               *string
	CategoryUUID       *UUID
	Excerpt            *string `valid:"json"`
	Introduction       *string `valid:"json"`
	BackgroundCheck    *bool
	AccessType         *AccessType
	Price              *float64
	Color              *string `valid:"hexcolor"`
	Tags               *[]UUID
	SpecificTerms      *string `valid:"json"`
	BannerImageSuccess *string
}

type SaveClassroomCourseInput struct {
	CourseInput
	TutorUUID       *UUID
	MaxParticipants *int
	StartDate       *Time
	EndDate         *Time
	Location        *string
}

type SaveOnlineCourseInput struct {
	CourseInput
	Structure *[]CourseItem
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

type CourseInfo struct {
	ID              uint
	Name            string
	AccessType      AccessType
	BackgroundCheck bool
	Price           float64
	Tags            []Tag
	Color           string `valid:"hexcolor"`
	Introduction    string `valid:"json"`
	Excerpt         string `valid:"json"`
	SpecificTerms   string `valid:"json"`
}

type Course struct {
	UUID         UUID
	CourseInfoID uint
}

type OnlineCourse struct {
	Course
}

type CourseInfoFilter struct {
	Name            *string
	AccessType      *AccessType
	BackgroundCheck *bool
	Price           *float64
}

type OnlineCourseFilter struct {
	CourseInfo CourseInfoFilter
}

type ClassroomCourse struct {
	Course
	StartDate       Time
	EndDate         Time
	Location        string
	MaxParticipants int
}

type Tag struct {
	UUID  UUID
	Name  string
	Color string
}

type CreateTagInput struct {
	Name  string
	Color string `valid:"hexcolor"`
}
