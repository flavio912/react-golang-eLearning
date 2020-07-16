package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CAANumber struct {
	UUID       gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt  time.Time
	Identifier string
	Used       bool
}

type CertificateType struct {
	UUID                    gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name                    string
	CreatedAt               time.Time
	CertificateBodyImageKey *string
	RegulationText          string
	RequiresCAANo           bool
	ShowTrainingSection     bool
}
