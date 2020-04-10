package handler

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

type contextKey string

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
