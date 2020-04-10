package auth

import (
	"context"
	"net/http"
	"strings"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

type contextKey string

var (
	// AuthKey is the contextKey for the JWT Auth token
	AuthKey = contextKey("token")
	// GrantKey is the contextKey for getting a grant
	GrantKey = contextKey("grant")
)

// Handler handles creating a grant for authenticated users
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// todo: Pass JWT to resolvers for now; this should be moved to middleware
		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")

		// Attempt to get a grant
		grant, err := middleware.Authenticate(token)
		if err == nil {
			ctx = context.WithValue(ctx, GrantKey, grant)
		}

		ctx = context.WithValue(ctx, AuthKey, token)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GrantFromContext returns a grant if the context has one (i.e the user is authenticated)
// otherwise returns nil
func GrantFromContext(ctx context.Context) *middleware.Grant {
	val := ctx.Value(GrantKey)
	if val == nil {
		return nil
	}

	v := val.(*middleware.Grant)
	return v
}
