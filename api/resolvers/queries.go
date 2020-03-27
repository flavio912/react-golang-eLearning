package resolvers

func (r *RootResolver) Info() (string, error) {
	return "This is the TTC server api", nil
}
