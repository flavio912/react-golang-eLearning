package resolvers

// QueryResolver -
type QueryResolver struct{}

// Info -
func (q *QueryResolver) Info() (string, error) {
	return "This is the TTC server api", nil
}
