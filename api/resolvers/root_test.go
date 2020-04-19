package resolvers_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

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
	db             *sql.DB
	fixtures       *testfixtures.Loader
	schema         *graphql.Schema = helpers.ParseSchema(s.String(), &resolvers.RootResolver{})
	defaultContext context.Context = context.Background()
	adminContext   context.Context = context.Background()
	managerContext context.Context = context.Background()
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

	loaders := loader.Init()
	defaultContext = loaders.Attach(defaultContext)

	adminContext = loaders.Attach(adminContext)
	adminContext, err = addAdminCreds(adminContext)
	if err != nil {
		return
	}

	managerContext = loaders.Attach(managerContext)
	managerContext, err = addManagerCreds(managerContext)
	if err != nil {
		return
	}

	os.Exit(m.Run())
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
	}
}

func accessTest(
	t *testing.T, schema *graphql.Schema, query string,
	expectedErrors []gqltest.TestQueryError,
	mustAuth, adminAllowed, managerAllowed, delegateAllowed bool,
) {
	checkAccess := func(t *testing.T, ctx context.Context, allowed bool) {
		r := schema.Exec(ctx, query, "", nil)
		if allowed {
			gqltest.CheckErrors(t, []gqltest.TestQueryError{}, r.Errors)
		} else {
			gqltest.CheckErrors(t, expectedErrors, r.Errors)
		}
	}

	t.Run(fmt.Sprintf("Must be authed:%v", mustAuth), func(t *testing.T) {
		checkAccess(t, defaultContext, !mustAuth)
	})

	t.Run(fmt.Sprintf("Admin Allowed:%v", adminAllowed), func(t *testing.T) {
		checkAccess(t, adminContext, adminAllowed)
	})

	t.Run(fmt.Sprintf("Manager Allowed:%v", managerAllowed), func(t *testing.T) {
		checkAccess(t, adminContext, adminAllowed)
	})

	// there's no delegate context yet ...
}
