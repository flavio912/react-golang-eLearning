package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type UserResolver struct {
	user gentypes.User
}

type NewUserArgs struct {
	User gentypes.User
	UUID *gentypes.UUID
}

func NewUserResolver(ctx context.Context, args NewUserArgs) (*UserResolver, error) {
	if args.UUID != nil {
		grant := auth.GrantFromContext(ctx)
		if grant == nil {
			return &UserResolver{}, &errors.ErrUnauthorized
		}

		usersApp := users.NewUsersApp(grant)
		user, err := usersApp.GetCurrentUser() // TODO: Use Dataloaders
		if err != nil {
			return &UserResolver{}, err
		}

		return &UserResolver{user: user}, nil
	}

	return &UserResolver{user: args.User}, nil
}

func (u *UserResolver) Type() gentypes.UserType  { return u.user.Type }
func (u *UserResolver) Email() *string           { return u.user.Email }
func (u *UserResolver) FirstName() string        { return u.user.FirstName }
func (u *UserResolver) LastName() string         { return u.user.LastName }
func (u *UserResolver) Telephone() *string       { return u.user.Telephone }
func (u *UserResolver) JobTitle() *string        { return u.user.JobTitle }
func (u *UserResolver) LastLogin() string        { return u.user.LastLogin }
func (u *UserResolver) ProfileImageUrl() *string { return u.user.ProfileImageUrl }
func (u *UserResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: grant.Claims.Company.String(),
	})
}
