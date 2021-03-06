package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type Module struct {
	UUID         gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt    time.Time
	Name         string
	Tags         []Tag `gorm:"many2many:module_tags_link"`
	Description  string
	Transcript   string
	VoiceoverKey *string
	VideoType    *gentypes.VideoType
	VideoURL     *string
	BannerKey    *string
	Structure    []ModuleStructure
}

type ModuleTag struct {
	ModuleUUID gentypes.UUID
	TagUUID    gentypes.UUID
}

func (ModuleTag) TableName() string {
	return "module_tags_link"
}

type ModuleStructure struct {
	Module     Module
	ModuleUUID gentypes.UUID `gorm:"primary_key;type:uuid;"`
	Lesson     Lesson
	LessonUUID *gentypes.UUID
	Test       Test
	TestUUID   *gentypes.UUID
	Rank       string `gorm:"primary_key"`
}
