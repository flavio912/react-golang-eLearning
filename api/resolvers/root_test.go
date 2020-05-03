package resolvers_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/go-testfixtures/testfixtures/v3"
	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers"
	s "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/schema"
)

var (
	db                 *sql.DB
	fixtures           *testfixtures.Loader
	schema             *graphql.Schema = helpers.ParseSchema(s.String(), &resolvers.RootResolver{})
	baseAdminContext   context.Context
	baseManagerContext context.Context
)

func TestMain(m *testing.M) {
	var err error

	// Load in the config.yaml file
	if err := helpers.LoadConfig("../dev_env/test_config.yml"); err != nil {
		panic(err)
	}
	errDb := database.SetupDatabase(false)
	if errDb != nil {
		panic(errDb)
	}
	migration.InitMigrations()

	db, err = sql.Open("postgres", "host=test_db port=5432 user=test dbname=testdb password=test sslmode=disable")
	if err != nil {
		fmt.Printf("Unable to connect to test DB: %s", err.Error())
		return
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db), // Your database connection
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../middleware/fixtures"), // the directory containing the YAML files
	)

	if err != nil {
		fmt.Printf("Unable get fixtures: %s", err.Error())
		return
	}

	prepareTestDatabase()

	baseAdminContext, err = addAdminCreds(context.Background())
	if err != nil {
		return
	}
	baseManagerContext, err = addManagerCreds(context.Background())
	if err != nil {
		return
	}

	os.Exit(m.Run())
}

func attachLoaders(ctx context.Context) context.Context {
	loaders := loader.Init()
	ctx = loaders.Attach(ctx)
	return ctx
}

func defaultContext() context.Context {
	return attachLoaders(context.Background())
}

func adminContext() context.Context {
	return attachLoaders(baseAdminContext)
}

func managerContext() context.Context {
	return attachLoaders(baseManagerContext)
}

// adds admin 1 to context
func addAdminCreds(ctx context.Context) (context.Context, error) {
	adminToken, err := middleware.GetAdminAccessToken("test123@test.com", "iamasuperadmin")
	if err != nil {
		fmt.Printf("Unable to get admin token! %#v", err)
		return nil, err
	}
	ctx = context.WithValue(ctx, auth.AuthKey, adminToken)

	grant, err := middleware.Authenticate(adminToken)
	if err != nil {
		fmt.Printf("Unable auth admin! %#v", err)
		return nil, err
	}

	return context.WithValue(ctx, auth.GrantKey, grant), nil
}

// adds manager 1 to the context
func addManagerCreds(ctx context.Context) (context.Context, error) {
	managerToken, err := middleware.GetManagerAccessToken("man@managers.com", "iamamanager")
	if err != nil {
		fmt.Printf("Unable to get manager token! %#v", err)
		return nil, err
	}
	ctx = context.WithValue(ctx, auth.AuthKey, managerToken)

	grant, err := middleware.Authenticate(managerToken)
	if err != nil {
		fmt.Printf("Unable auth manager! %#v", err)
		return nil, err
	}

	return context.WithValue(ctx, auth.GrantKey, grant), nil
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Unable to load fixtures for test: %s", err.Error())
		panic(err)
	}
}

type accessTestOpts struct {
	Query           string
	Path            []interface{}
	MustAuth        bool
	AdminAllowed    bool
	ManagerAllowed  bool
	DelegateAllowed bool
	CleanDB         bool
}

func accessTest(t *testing.T, schema *graphql.Schema, opts accessTestOpts) {
	checkAccess := func(t *testing.T, ctx context.Context, allowed bool) {
		if opts.CleanDB {
			prepareTestDatabase()
		}

		r := schema.Exec(ctx, opts.Query, "", nil)
		if allowed {
			gqltest.CheckErrors(t, []gqltest.TestQueryError{}, r.Errors)
		} else {
			gqltest.CheckErrors(t, []gqltest.TestQueryError{{
				Path:          opts.Path,
				ResolverError: &errors.ErrUnauthorized,
			}}, r.Errors)
		}
	}

	t.Run(fmt.Sprintf("Must be authed:%v", opts.MustAuth), func(t *testing.T) {
		checkAccess(t, defaultContext(), !opts.MustAuth)
	})

	t.Run(fmt.Sprintf("Admin Allowed:%v", opts.AdminAllowed), func(t *testing.T) {
		checkAccess(t, adminContext(), opts.AdminAllowed)
	})

	t.Run(fmt.Sprintf("Manager Allowed:%v", opts.ManagerAllowed), func(t *testing.T) {
		checkAccess(t, managerContext(), opts.ManagerAllowed)
	})

	// there's no delegate context yet ...
}
