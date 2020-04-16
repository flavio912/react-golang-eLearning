package resolvers_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
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
	adminContext = loaders.Attach(adminContext)

	adminToken, err := middleware.GetAdminAccessToken("test123@test.com", "iamasuperadmin")
	if err != nil {
		fmt.Printf("Unable to get admin token! %#v", err)
		return
	}
	adminContext = context.WithValue(adminContext, auth.AuthKey, adminToken)

	grant, err := middleware.Authenticate(adminToken)
	if err != nil {
		fmt.Printf("Unable auth admin token! %#v", err)
		return
	}
	adminContext = context.WithValue(adminContext, auth.GrantKey, grant)

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Unable to load fixtures for test: %s", err.Error())
	}
}
