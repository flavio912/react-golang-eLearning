package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type DelegateResolver struct {
	delegate gentypes.Delegate
}

type NewDelegateArgs struct {
	Delegate gentypes.Delegate
	UUID     *gentypes.UUID
}

func NewDelegateResolver(ctx context.Context, args NewDelegateArgs) (*DelegateResolver, error) {
	if args.UUID != nil {
		grant := auth.GrantFromContext(ctx)
		if grant == nil {
			return &DelegateResolver{}, &errors.ErrUnauthorized
		}

		usersApp := users.NewUsersApp(grant)
		delegate, err := usersApp.Delegate(*args.UUID) // TODO: Use Dataloaders
		if err != nil {
			return &DelegateResolver{}, err
		}

		return &DelegateResolver{delegate}, nil
	}

	return &DelegateResolver{args.Delegate}, nil
}

func (d *DelegateResolver) UUID() gentypes.UUID      { return d.delegate.UUID }
func (d *DelegateResolver) TTC_ID() string           { return d.delegate.TTC_ID }
func (d *DelegateResolver) CreatedAt() *string       { return d.delegate.CreatedAt }
func (d *DelegateResolver) Email() *string           { return d.delegate.Email }
func (d *DelegateResolver) FirstName() string        { return d.delegate.FirstName }
func (d *DelegateResolver) LastName() string         { return d.delegate.LastName }
func (d *DelegateResolver) Telephone() *string       { return d.delegate.Telephone }
func (d *DelegateResolver) JobTitle() string         { return d.delegate.JobTitle }
func (d *DelegateResolver) LastLogin() string        { return d.delegate.LastLogin }
func (d *DelegateResolver) ProfileImageURL() *string { return d.delegate.ProfileImageURL }
func (d *DelegateResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: d.delegate.CompanyUUID.String(),
	})
}
func (d *DelegateResolver) Activity(ctx context.Context, args struct{ Page *gentypes.Page }) (*ActivityPageResolver, error) {
	return NewActivityPageResolver(ctx, NewActivityPageArgs{
		CourseTakerUUID: &d.delegate.CourseTakerUUID,
	}, args.Page)
}
func (d *DelegateResolver) ActiveCourses(ctx context.Context) (*[]*ActiveCourseResolver, error) {
	return NewActiveCoursesResolvers(ctx, NewActiveCourseArgs{
		TakerUUID: &d.delegate.CourseTakerUUID,
	})
}

type DelegatePageResolver struct {
	edges    *[]*DelegateResolver
	pageInfo *PageInfoResolver
}

func (r *DelegatePageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *DelegatePageResolver) Edges() *[]*DelegateResolver { return r.edges }

type NewDelegatePageArgs struct {
	Delegates *[]gentypes.Delegate
	PageInfo  gentypes.PageInfo
}

func NewDelegatePageResolver(ctx context.Context, args NewDelegatePageArgs) (*DelegatePageResolver, error) {
	var resolvers []*DelegateResolver
	switch {
	case args.Delegates != nil:
		for _, delegate := range *args.Delegates {
			res, err := NewDelegateResolver(ctx, NewDelegateArgs{
				Delegate: delegate,
			})

			if err != nil {
				return nil, err
			}

			resolvers = append(resolvers, res)
		}
	default:
		return &DelegatePageResolver{}, &errors.ErrUnableToResolve
	}

	return &DelegatePageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}

type CreateDelegateResponse struct {
	delegate          gentypes.Delegate
	generatedPassword *string
}

func (r *CreateDelegateResponse) Delegate(ctx context.Context) (*DelegateResolver, error) {
	return NewDelegateResolver(ctx, NewDelegateArgs{
		Delegate: r.delegate,
	})
}
func (r *CreateDelegateResponse) GeneratedPassword() *string {
	return r.generatedPassword
}
