package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Lesson struct {
	UUID         gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name         string
	Tags         []Tag `gorm:"many2many:lesson_tags_link;"`
	Description  string
	Transcript   *string
	VoiceoverKey *string
	VideoType    *gentypes.VideoType
	VideoURL     *string
	BannerKey    *string
}

type LessonTagsLink struct {
	LessonUUID gentypes.UUID
	TagUUID    gentypes.UUID
}

func (LessonTagsLink) TableName() string {
	return "lesson_tags_link"
}
