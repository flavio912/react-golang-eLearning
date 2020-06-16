package gentypes

type CreateIndividualInput struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	JobTitle  *string
	Telephone *string `valid:"optional,numeric"`
	Email     string  `valid:"email"`
	Password  string
}
