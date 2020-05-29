package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Module struct {
	UUID       gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Structure  []ModuleStructure
	Template   bool           // Is the module a template or custom module
	TemplateID *gentypes.UUID // The ID of the template used to create this FKEY
}

type ModuleStructure struct {
	Module     Module
	ModuleUUID gentypes.UUID
	Lesson     Lesson
	LessonUUID *gentypes.UUID
	Test       Test
	TestUUID   *gentypes.UUID
	Rank       string
}

type Lesson struct {
	UUID  gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title string
	Tags  []Tag  `gorm:"many2many:lesson_tags_link;"`
	Text  string `sql:"json"`
}

type LessonTagsLink struct {
	LessonUUID gentypes.UUID
	TagUUID    gentypes.UUID
}

func (LessonTagsLink) TableName() string {
	return "lesson_tags_link"
}

type Test struct {
	UUID gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}
