package gentypes

import (
	"github.com/google/uuid"
)

type Company struct {
	UUID uuid.UUID
	Name string
}
