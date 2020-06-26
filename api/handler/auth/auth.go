package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/domain"

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
	AppKey  = contextKey("application")
	// GrantKey is the contextKey for getting a grant
	GrantKey = contextKey("grant")
	RespKey  = contextKey("resp")
)

// Handler handles creating a grant for authenticated users
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")
		cookieUsed := false
		if token == "" {
			_token, err := r.Cookie("auth")
			if err == nil {
				token = _token.Value
				cookieUsed = true
			}
		}

		// Attempt to get a grant
		grant, _ := middleware.Authenticate(token)

		// Check CSRF if cookie was used for authentication and disallow if CSRF is invalid
		// This is happening after authetication so the user can still make unauthenticated
		// requests (like logging in) even if CSRF fails
		if grant != nil {

			var csrfHeader = r.Header.Get("X-CSRF-TOKEN")
			csrfCookie, _ := r.Cookie("csrf")

			var allowRequest = false
			if cookieUsed && csrfCookie != nil {
				if fmt.Sprintf("csrf=%s", csrfHeader) == csrfCookie.String() {
					allowRequest = true
				} else {
					w.Header().Set("CSRF-FAIL", "true") // Useful for diagnosis
					glog.Warningf("CSRF Tokens don't match: IP - %s", r.RemoteAddr)
					allowRequest = false
				}
			}

			if !cookieUsed {
				allowRequest = true
			}

			// Public grants don't need to worry about CSRF tokens
			if allowRequest || grant.IsPublic {
				ctx = context.WithValue(ctx, GrantKey, grant)
				ctx = context.WithValue(ctx, AuthKey, token)
				ctx = AddAppContext(ctx, grant)
				addSentryContext(r, grant)
			}
		}

		ctx = context.WithValue(ctx, RespKey, &w)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddAppContext(ctx context.Context, grant *middleware.Grant) context.Context {
	app := domain.Application{
		CourseApp: course.NewCourseApp(grant),
		UsersApp:  users.NewUsersApp(grant),
		AdminApp:  application.NewAdminApp(grant),
	}

	ctx = context.WithValue(ctx, AppKey, app)
	return ctx
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

// AppFromContext gets an application from the context
func AppFromContext(ctx context.Context) domain.Application {
	val := ctx.Value(AppKey)
	if val == nil {
		return domain.Application{}
	}

	app := val.(domain.Application)
	return app
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

// SetAuthCookies gets the responseWriter from context and uses it to set the jwt cookie
func SetAuthCookies(ctx context.Context, authToken string) {

	writer := ResponseFromContext(ctx)
	if writer == nil {
		glog.Warning("Unable to set auth cookie, no writer set")
		return
	}

	grant, err := middleware.Authenticate(authToken)
	if err != nil {
		grant.Logger.LogMessage(sentry.LevelError, "Unable to authenticate with token, auth cookies not set")
		return
	}

	csrfToken, err := grant.GenerateCSRFToken()
	if err != nil {
		grant.Logger.Log(sentry.LevelError, err, "Unable to generate CSRF token")
		return
	}

	var expirationTime = time.Now().Add(time.Hour * time.Duration(helpers.Config.Jwt.TokenExpirationHours))

	var authCookie = http.Cookie{Name: "auth", Value: authToken, HttpOnly: true, Domain: helpers.Config.CookieDomain, Secure: true, Expires: expirationTime}
	var csrfCookie = http.Cookie{Name: "csrf", Value: csrfToken, HttpOnly: false, Domain: helpers.Config.CookieDomain, Secure: true, Expires: expirationTime}

	if helpers.Config.IsDev {
		authCookie.Domain = ""
		csrfCookie.Domain = ""
		authCookie.Secure = false
		csrfCookie.Secure = false
	}

	http.SetCookie(*writer, &authCookie)
	http.SetCookie(*writer, &csrfCookie)
}
