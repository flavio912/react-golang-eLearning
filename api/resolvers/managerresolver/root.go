package managerresolver

// ManagerRoot - Resolves all admin spefic routes
type ManagerRoot struct {
	QueryResolver
	MutationResolver
}
