package gentypes

import (
	"github.com/google/uuid"
)

// User - User graphQL interface
type User struct {
	UUID      uuid.UUID
	Email     string
	FirstName string
	LastName  string
	Telephone string
	JobTitle  string
	LastLogin string
	CompanyID uuid.UUID
}
