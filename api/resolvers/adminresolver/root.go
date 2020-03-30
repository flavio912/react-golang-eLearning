package adminresolver

// RootResolver - Resolves all admin spefic routes
type RootResolver struct {
	QueryResolver
	MutationResolver
}
