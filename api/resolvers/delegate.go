package resolvers

import (
	"context"

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

		delegate, err := grant.GetDelegateFromUUID(*args.UUID)
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
func (d *DelegateResolver) Email() string            { return d.delegate.Email }
func (d *DelegateResolver) FirstName() string        { return d.delegate.FirstName }
func (d *DelegateResolver) LastName() string         { return d.delegate.LastName }
func (d *DelegateResolver) Telephone() string        { return d.delegate.Telephone }
func (d *DelegateResolver) JobTitle() string         { return d.delegate.JobTitle }
func (d *DelegateResolver) LastLogin() string        { return d.delegate.LastLogin }
func (d *DelegateResolver) ProfileImageURL() *string { return d.delegate.ProfileImageURL }

func (d *DelegateResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: d.delegate.CompanyUUID.String(),
	})
}
