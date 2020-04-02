package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

// MutationResolver - Root resolver for mutations
type MutationResolver struct{}

// AuthToken - Type used for returning JWT's
type AuthToken struct {
	Token string
}

// AdminLogin - Resolver for getting an authToken
func (m *MutationResolver) AdminLogin(args struct{ Input gentypes.AdminLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetAdminAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}
	return &gentypes.AuthToken{Token: token}, nil
}

// ManagerLogin - Resolver for getting an authToken
func (m *MutationResolver) ManagerLogin(args struct{ Input gentypes.ManagerLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetManagerAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}
	return &gentypes.AuthToken{Token: token}, nil
}

// AddManager is for an admin to create new managers manually
func (m *MutationResolver) AddManager(ctx context.Context, args struct{ Input gentypes.AddManagerInput }) (*ManagerResolver, error) {
	// Validate the manager input
	if err := args.Input.Validate(); err != nil {
		return &ManagerResolver{}, err
	}

	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return &ManagerResolver{}, err
	}

	manager, err := grant.AddManager(args.Input)
	if err != nil {
		return &ManagerResolver{}, err
	}

	loadedManager, loadErr := loader.LoadManager(ctx, manager.UUID)

	return &ManagerResolver{
		manager: loadedManager,
	}, loadErr
}
