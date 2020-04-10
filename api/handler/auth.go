package handler

import (
	"context"
	"net/http"
	"strings"
)

func AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// todo: Pass JWT to resolvers for now; this should be moved to middleware
		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")

		ctx = context.WithValue(ctx, "token", token)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
