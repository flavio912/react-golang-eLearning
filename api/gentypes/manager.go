package gentypes

// Manager - CompanyManager graphQL type
type Manager struct {
	User
}

// ManagerLoginInput - ManagerLogin graphQL input
type ManagerLoginInput struct {
	Email    string
	Password string
}
