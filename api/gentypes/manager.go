package gentypes

// Manager - CompanyManager graphQL type
type Manager struct {
	User
}

type AddManagerInput struct {
	FirstName string
	LastName  string
	Email     string
	JobTitle  string
	Telephone string
	Password  string
}

// ManagerLoginInput - ManagerLogin graphQL input
type ManagerLoginInput struct {
	Email    string
	Password string
}
