package resolvers

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers/adminresolver"

type RootResolver struct {
	QueryResolver
	MutationResolver
	adminresolver.RootResolver
}
