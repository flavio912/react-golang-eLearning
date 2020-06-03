package auth

import (
	"context"
	"net/http"
	"strings"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

type contextKey string

var (
	// AuthKey is the contextKey for the JWT Auth token
	AuthKey = contextKey("token")
	// GrantKey is the contextKey for getting a grant
	GrantKey = contextKey("grant")
	RespKey  = contextKey("resp")
)

// Handler handles creating a grant for authenticated users
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")

		if token == "" {
			_token, err := r.Cookie("auth")
			if err == nil {
				// If valid cookie given, check XSRF token is present
				var csrfToken = r.Header.Get("X-CSRF-TOKEN")
				if csrfToken == "" {
					w.Header().Set("Content-Type", "application/text")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid CSRF TOKEN"))
					return
				}

				// TODO: IMPORTANT - Carry out AUTH checks
				token = _token.Value
			}
		}

		// Attempt to get a grant
		grant, err := middleware.Authenticate(token)
		if err == nil {
			ctx = context.WithValue(ctx, GrantKey, grant)

			addSentryContext(r, grant)
		}

		ctx = context.WithValue(ctx, AuthKey, token)
		ctx = context.WithValue(ctx, RespKey, &w)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addSentryContext(r *http.Request, grant *middleware.Grant) {
	// Add sentry context
	if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
		hub.Scope().SetUser(sentry.User{
			ID: grant.Claims.UUID.String(),
		})
		hub.Scope().SetTag("role", auth.RoleToString(grant.Claims.Role))
		hub.Scope().SetTag("company", grant.Claims.Company.String())
	}
}

// GrantFromContext returns a grant if the context has one (i.e the user is authenticated)
// otherwise returns nil
func GrantFromContext(ctx context.Context) *middleware.Grant {
	val := ctx.Value(GrantKey)
	if val == nil {
		return nil
	}

	v := val.(*middleware.Grant)
	v.Logger = logging.GetLoggerFromCtx(ctx)
	return v
}

// ResponseFromContext returns a pointer to the response writer of the http request
func ResponseFromContext(ctx context.Context) *http.ResponseWriter {
	val := ctx.Value(RespKey)
	if val == nil {
		return nil
	}

	v := val.(*http.ResponseWriter)
	return v
}

// SetAuthCookie gets the responseWriter from context and uses it to set the jwt cookie
func SetAuthCookie(ctx context.Context, token string) {
	writer := ResponseFromContext(ctx)
	if writer == nil {
		glog.Warning("Unable to set auth cookie, no writer set")
		return
	}

	if helpers.Config.IsDev {
		// In dev, don't use a secure cookie as we aren't using tls in dev
		http.SetCookie(*writer, &http.Cookie{Name: "auth", Value: token, HttpOnly: true, Secure: false})
	} else {
		http.SetCookie(*writer, &http.Cookie{Name: "auth", Value: token, HttpOnly: true, Secure: true, Domain: "*.ttc-hub.com"})
	}
}
