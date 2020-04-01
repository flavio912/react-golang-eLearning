package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

// QueryResolver -
type QueryResolver struct{}

// Info -
func (q *QueryResolver) Info() (string, error) {
	return "This is the TTC server api", nil
}

// Admins - Get a list of admins
func (q *QueryResolver) Admins(ctx context.Context, args gentypes.Page) (*AdminPageResolver, error) {
	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return &AdminPageResolver{}, err
	}
	admins, err := grant.GetAdmins(&args, nil)
	if err != nil {
		return nil, err
	}
	adminResolvers := []*AdminResolver{}
	for _, admin := range admins {
		adminResolvers = append(adminResolvers, &AdminResolver{
			admin: *admin,
		})
	}
	return &AdminPageResolver{
		edges: &adminResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &gentypes.PageInfo{
				PagesAfter: 0,
				Offset:     0,
				Limit:      0,
				Given:      int32(len(admins)),
			},
		},
	}, nil
}

// Admin gets a single admin
func (q *QueryResolver) Admin(ctx context.Context, args struct{ UUID string }) (*AdminResolver, error) {
	admin, err := loader.LoadAdmin(ctx, args.UUID)

	return &AdminResolver{admin: admin}, err
}

// Manager gets a single manager
func (q *QueryResolver) Manager(ctx context.Context, args struct{ UUID string }) (*ManagerResolver, error) {
	manager, err := loader.LoadManager(ctx, args.UUID)
	if err != nil {
		return &ManagerResolver{}, err
	}
	return &ManagerResolver{manager: manager}, nil
}
