package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Blog struct {
	Base
	Title          string
	Body           string `sql:"json"`
	Category       Category
	CategoryUUID   gentypes.UUID
	HeaderImageURL string
	Author         Admin `gorm:"foreignkey:BlogUUID"`
}
