package gentypes

type Individual struct {
	UUID      UUID
	CreatedAt *string
	Email     string
	FirstName string
	LastName  string
	JobTitle  *string
	Telephone *string
	LastLogin string
}

type CreateIndividualInput struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	JobTitle  *string
	Telephone *string `valid:"optional,numeric"`
	Email     string  `valid:"email"`
	Password  string
}

type IndividualFilter struct {
	UserFilter
	Email *string `valid:"email"`
}
