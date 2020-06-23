package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Blog struct {
	Base
	Title          string
	Body           string `sql:"json"`
	Category       Category
	CategoryUUID   gentypes.UUID
	HeaderImageKey string
	Author         Admin `gorm:"foreignkey:BlogUUID"`
}

type BlogImage struct {
	BlogUUID gentypes.UUID
	BodyID   string
	S3key    string
}
