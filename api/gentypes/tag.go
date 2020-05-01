package gentypes

type Tag struct {
	UUID  UUID
	Name  string
	Color string
}

type CreateTagInput struct {
	Name  string
	Color string `valid:"hexcolor"`
}

type GetTagsFilter struct {
	Name string
}
