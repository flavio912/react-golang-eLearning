package models

import (
	"time"

	"github.com/google/uuid"
)

// Base contains fields present in all records: ID
type Base struct {
	UUID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
