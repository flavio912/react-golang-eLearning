package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers/adminresolver"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers/managerresolver"
)

type RootResolver struct {
	QueryResolver
	MutationResolver
	adminresolver.AdminRoot
	managerresolver.ManagerRoot
}
