package gentypes

// Admin - The admin graphQL type
type Admin struct {
	UUID      string
	Email     string
	FirstName string
	LastName  string
}

// AdminPage - a list of admins
type AdminPage struct {
	Edges    Admin
	PageInfo PageInfo
}

// AdminLoginInput -
type AdminLoginInput struct {
	Email    string
	Password string
}
