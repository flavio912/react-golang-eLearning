package resolvers

import (
	"context"
	"strings"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func (m *MutationResolver) DelegateLogin(ctx context.Context, args struct{ Input gentypes.DelegateLoginInput }) (*gentypes.AuthToken, error) {
	var (
		token string
		err   error
	)

	if strings.Contains(args.Input.TTC_ID, "@") {
		token, err = middleware.GetIndividualAccessToken(args.Input.TTC_ID, args.Input.Password)
	} else {
		token, err = middleware.GetDelegateAccessToken(args.Input.TTC_ID, args.Input.Password)
	}

	if err != nil {
		return nil, err
	}

	auth.SetAuthCookies(ctx, token)

	// If NoResp given return a blank token in the response - @temmerson
	if args.Input.NoResp != nil && *args.Input.NoResp {
		return &gentypes.AuthToken{Token: ""}, nil
	}
	return &gentypes.AuthToken{Token: token}, nil
}

func (m *MutationResolver) CreateDelegate(ctx context.Context, args struct{ Input gentypes.CreateDelegateInput }) (*CreateDelegateResponse, error) {
	if err := args.Input.Validate(); err != nil {
		return nil, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CreateDelegateResponse{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	delegate, password, err := usersApp.CreateDelegate(args.Input)
	return &CreateDelegateResponse{delegate: delegate, generatedPassword: password}, err
}

func (m *MutationResolver) UpdateDelegate(ctx context.Context, args struct{ Input gentypes.UpdateDelegateInput }) (*DelegateResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &DelegateResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	delegate, err := usersApp.UpdateDelegate(args.Input)
	return &DelegateResolver{
		delegate: delegate,
	}, err
}
