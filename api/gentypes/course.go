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

type CourseInfo struct {
	ID              uint
	Name            string
	AccessType      AccessType
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

type Course struct {
	UUID         UUID
	CourseInfoID uint
}

type CourseInfoFilter struct {
	Name            *string
	AccessType      *AccessType
	BackgroundCheck *bool
	Price           *float64
}
