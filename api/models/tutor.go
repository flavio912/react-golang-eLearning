package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Tutor struct {
	UUID         gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name         string
	CIN          uint
	SignatureKey string
}
