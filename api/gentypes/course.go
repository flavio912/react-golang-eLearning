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
	UUID               *string `valid:"uuidv4"`
	Name               *string
	CategoryUUID       *string `valid:"uuidv4"`
	Excerpt            *string `valid:"json"`
	Introduction       *string `valid:"json"`
	BackgroundCheck    *bool
	AccessType         *AccessType
	Price              *float64
	Color              *string `valid:"hexcolor"`
	Tags               *[]*string
	SpecificTerms      *string `valid:"json"`
	BannerImageSuccess *string
}

type SaveClassroomCourseInput struct {
	CourseInput
	TutorUUID       *string `valid:"uuidv4"`
	MaxParticipants *int
}

type SaveOnlineCourseInput struct {
	CourseInput
	Structure *[]CourseItem
}

type CourseItem struct {
	Type  StructureElement
	UUID  string
	Items []ModuleItem
}

type ModuleItem struct {
	Type StructureElement
	UUID string
}

type Course struct {
	UUID         string
	CourseInfoID uint
}

type OnlineCourse struct {
	Course
}

type ClassroomCourse struct {
	Course
}
