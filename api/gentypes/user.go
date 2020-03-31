package gentypes

import (
	"time"
)

// User - User graphQL interface
type User struct {
	UUID      string
	Email     string
	FirstName string
	LastName  string
	Telephone string
	JobTitle  string
	LastLogin time.Time
}
