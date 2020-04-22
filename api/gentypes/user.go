package gentypes

// User - User graphQL interface
type User struct {
	CreatedAt *string
	UUID      UUID
	Email     string
	FirstName string
	LastName  string
	Telephone string
	JobTitle  string
	LastLogin string
}
