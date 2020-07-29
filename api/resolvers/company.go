package resolvers

import (
	"context"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

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
	}, err
}

type NewCompanyPageArgs struct {
	UUIDs     []string
	Companies []gentypes.Company
}

func NewCompanyPageResolver(ctx context.Context, args NewCompanyPageArgs, _pageInfo gentypes.PageInfo) (*CompanyPageResolver, error) {
	var resolvers []*CompanyResolver

	if len(args.UUIDs) > 0 {
		for _, uuid := range args.UUIDs {
			resolver, err := NewCompanyResolver(ctx, NewCompanyArgs{
				UUID: uuid,
			})
			if err != nil {
				glog.Errorf("Unable to resolve: %s", uuid)
				resolvers = append(resolvers, &CompanyResolver{})
			}
			resolvers = append(resolvers, resolver)
		}
	} else {
		for _, comp := range args.Companies {
			resolver, err := NewCompanyResolver(ctx, NewCompanyArgs{
				Company: comp,
			})
			if err != nil {
				glog.Errorf("Unable to resolve: %s", comp.UUID.String())
				resolvers = append(resolvers, &CompanyResolver{})
			}
			resolvers = append(resolvers, resolver)
		}
	}

	return &CompanyPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &_pageInfo,
		},
	}, nil
}

func uuidsToStrings(uuids []gentypes.UUID) []string {
	var strings = make([]string, len(uuids))
	for i, uuid := range uuids {
		strings[i] = uuid.String()
	}
	return strings
}

func (r *CompanyResolver) Name() string         { return r.company.Name }
func (r *CompanyResolver) CreatedAt() *string   { return r.company.CreatedAt }
func (r *CompanyResolver) UUID() gentypes.UUID  { return r.company.UUID }
func (r *CompanyResolver) ContactEmail() string { return r.company.ContactEmail }
func (r *CompanyResolver) Approved(ctx context.Context) *bool {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return nil
	}

	// TODO: Add a key onto the grant 'CanGetApproved' to check if user can see if approved (or something similar)
	if grant.IsAdmin {
		return r.company.Approved
	}
	return nil
}
func (r *CompanyResolver) IsContract() bool { return r.company.IsContract }
func (r *CompanyResolver) Address(ctx context.Context) (gentypes.Address, error) {
	return loader.LoadAddress(ctx, r.company.AddressID)
}
func (r *CompanyResolver) Managers(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.ManagersFilter
	OrderBy *gentypes.OrderBy
}) (*ManagerPageResolver, error) {
	// TODO: N+1 problem - get it to use dataloaders
	app := auth.AppFromContext(ctx)
	managers, pageInfo, _ := app.UsersApp.GetManagerIDsByCompany(r.company.UUID, args.Page, args.Filter, args.OrderBy)
	resolver, err := NewManagerResolvers(ctx, NewManagersArgs{UUIDs: uuidsToStrings(managers)})
	if err != nil {
		glog.Info("Unable to resolve a manager: ")
		return &ManagerPageResolver{}, err
	}

	return &ManagerPageResolver{
		edges: resolver,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		},
	}, nil
}
func (r *CompanyResolver) Delegates(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.DelegatesFilter
	OrderBy *gentypes.OrderBy
}) (*DelegatePageResolver, error) {
	app := auth.AppFromContext(ctx)
	delegates, pageInfo, err := app.UsersApp.GetDelegates(args.Page, args.Filter, args.OrderBy)

	if err != nil {
		return &DelegatePageResolver{}, err
	}
	return NewDelegatePageResolver(ctx, NewDelegatePageArgs{
		Delegates: &delegates,
		PageInfo:  pageInfo,
	})
}

type CompanyPageResolver struct {
	edges    *[]*CompanyResolver
	pageInfo *PageInfoResolver
}

func (r *CompanyPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *CompanyPageResolver) Edges() *[]*CompanyResolver  { return r.edges }
