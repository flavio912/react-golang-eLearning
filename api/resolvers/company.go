package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CompanyResolver struct {
	company gentypes.Company
}

type NewCompanyArgs struct {
	Company gentypes.Company
	UUID    string
}

func NewCompanyResolver(ctx context.Context, args NewCompanyArgs) (*CompanyResolver, error) {
	var (
		company gentypes.Company
		err     error
	)

	switch {
	case args.UUID != "":
		company, err = loader.LoadCompany(ctx, args.UUID)
	case args.Company.UUID.String() != "":
		company = args.Company
	default:
		err = &errors.ErrUnableToResolve
	}

	if err != nil {
		return &CompanyResolver{}, err
	}

	return &CompanyResolver{
		company: company,
	}, nil
}

func (r *CompanyResolver) Name() string { return r.company.Name }
func (r *CompanyResolver) UUID() string { return r.company.UUID.String() }
func (r *CompanyResolver) Managers(ctx context.Context, args struct {
	Page   *gentypes.Page
	Filter *gentypes.ManagersFilter
}) (*ManagerPageResolver, error) {
	managers, err := loader.LoadManagersFromCompany(ctx, r.company.UUID.String())
	if err != nil {
		return &ManagerPageResolver{}, err
	}
	var res = make([]*ManagerResolver, len(managers))
	for i, manager := range managers {
		res[i] = &ManagerResolver{
			manager: manager,
		}
	}

	return &ManagerPageResolver{
		edges: &res,
		pageInfo: &PageInfoResolver{
			pageInfo: &gentypes.PageInfo{
				Total:  0,
				Offset: 0,
				Limit:  0,
				Given:  0,
			},
		},
	}, nil
}

type CompanyPageResolver struct {
	edges    *[]*CompanyResolver
	pageInfo *PageInfoResolver
}

func (r *CompanyPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *CompanyPageResolver) Edges() *[]*CompanyResolver  { return r.edges }
