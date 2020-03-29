package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

// MutationResolver - Root resolver for mutations
type MutationResolver struct{}
type adminLoginInput struct {
	Email    string
	Password string
}

// AuthToken - Type used for returning JWT's
type AuthToken struct {
	Token string
}

// AdminLogin - Resolver for getting an authToken
func (m *MutationResolver) AdminLogin(args struct{ Input adminLoginInput }) (*AuthToken, error) {
	token, err := middleware.GetAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}
	return &AuthToken{Token: token}, nil
}
