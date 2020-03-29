package resolvers

type QueryResolver struct{}

func (q *QueryResolver) Info() (string, error) {
	return "This is the TTC server api", nil
}
