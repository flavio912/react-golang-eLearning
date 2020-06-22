package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Module struct {
	UUID      gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Structure []ModuleStructure
}

type ModuleStructure struct {
	Module     Module
	ModuleUUID gentypes.UUID `gorm:"primary_key"`
	Lesson     Lesson
	LessonUUID *gentypes.UUID
	Test       Test
	TestUUID   *gentypes.UUID
	Rank       string `gorm:"primary_key"`
}
