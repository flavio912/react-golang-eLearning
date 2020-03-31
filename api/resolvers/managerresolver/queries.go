package managerresolver

// QueryResolver -
type QueryResolver struct {
}

// func (q *QueryResolver) Manager(ctx context.Context, args struct{ UUID string }) (*ManagerResolver, error) {
// 	manager, err := middleware.GetManager(ctx.Value("token").(string), args.UUID)
// 	if err != nil {
// 		return &ManagerResolver{}, &middleware.ErrUnauthorized
// 	}
// 	return &ManagerResolver{
// 		manager: manager,
// 	}, nil
// }
