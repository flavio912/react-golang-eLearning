package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"github.com/golang/glog"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
)

// QueryResolver -
type QueryResolver struct{}

// Info -
func (q *QueryResolver) Info(ctx context.Context) (string, error) {
	glog.Info(auth.GrantFromContext(ctx))
	return "This is the TTC server api", nil
}

// Admins - Get a list of admins
func (q *QueryResolver) Admins(ctx context.Context, args struct{ Page *gentypes.Page }) (*AdminPageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &AdminPageResolver{}, &errors.ErrUnauthorized
	}
	admins, err := grant.GetAdmins(args.Page, nil)
	if err != nil {
		return nil, err
	}
	adminResolvers := []*AdminResolver{}
	for _, admin := range admins {
		adminResolvers = append(adminResolvers, &AdminResolver{
			admin: admin,
		})
	}
	return &AdminPageResolver{
		edges: &adminResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &gentypes.PageInfo{
				Total:  0,
				Offset: 0,
				Limit:  0,
				Given:  int32(len(admins)),
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

func (q *QueryResolver) Managers(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.ManagersFilter
	OrderBy *gentypes.OrderBy
}) (*ManagerPageResolver, error) {

	if args.Filter != nil {
		err := (*args.Filter).Validate()
		if err != nil {
			return &ManagerPageResolver{}, err
		}
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerPageResolver{}, &errors.ErrUnauthorized
	}

	managers, page, err := grant.GetManagers(args.Page, args.Filter, args.OrderBy)
	var managerResolvers []*ManagerResolver
	for _, manager := range managers {
		managerResolvers = append(managerResolvers, &ManagerResolver{
			manager: manager,
		})
	}

	return &ManagerPageResolver{
		edges: &managerResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}

func (q *QueryResolver) Companies(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.CompanyFilter
	OrderBy *gentypes.OrderBy
}) (*CompanyPageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyPageResolver{}, &errors.ErrUnauthorized
	}

	companies, page, err := grant.GetCompanyUUIDs(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &CompanyPageResolver{}, err
	}
	return NewCompanyPageResolver(ctx, NewCompanyPageArgs{
		UUIDs: companies,
	}, page)
}

func (q *QueryResolver) Company(ctx context.Context, args struct{ UUID string }) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: args.UUID,
	})
}
