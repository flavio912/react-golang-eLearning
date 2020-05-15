package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

func (m *MutationResolver) CreateDelegate(ctx context.Context, args struct{ Input gentypes.CreateDelegateInput }) (*DelegateResolver, error) {
	if err := args.Input.Validate(); err != nil {
		return nil, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &DelegateResolver{}, &errors.ErrUnauthorized
	}

	delegate, err := grant.CreateDelegate(args.Input)
	return &DelegateResolver{delegate}, err
}
