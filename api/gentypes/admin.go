package gentypes

type Admin struct {
	UUID      string
	Email     string
	FirstName string
	LastName  string
}

// Key gets the admin primary identifier
func (admin *Admin) Key() string {
	return admin.UUID
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
