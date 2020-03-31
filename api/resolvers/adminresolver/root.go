package adminresolver

// AdminRoot - Resolves all admin spefic querys and mutations
type AdminRoot struct {
	QueryResolver
	MutationResolver
}
