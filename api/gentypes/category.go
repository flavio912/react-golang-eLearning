package gentypes

type Category struct {
	UUID  UUID
	Name  string
	Color string
}

type CreateCategoryInput struct {
	Name  string
	Color string `valid:"hexcolor"`
}
