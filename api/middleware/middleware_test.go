package middleware_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

var adminGrant = middleware.Grant{auth.UserClaims{}, true, false, false}
var managerGrant = middleware.Grant{auth.UserClaims{
	UUID:    "00000000-0000-0000-0000-000000000001",
	Company: "00000000-0000-0000-0000-000000000001",
	Role:    auth.ManagerRole,
}, false, true, false}
var delegateGrant = middleware.Grant{auth.UserClaims{}, false, false, true}
var uuidZero = "00000000-0000-0000-0000-000000000000"

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
		testfixtures.Directory("fixtures"), // the directory containing the YAML files
	)

	if err != nil {
		fmt.Printf("Unable get fixtures: %s", err.Error())
		return
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Unable to load fixtures for test: %s", err.Error())
	}
}

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{"fake token bad secret", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoidGhpcyBpcyBub3QgYSByZWFsIHRva2VuIn0.Dv7R7hZtHJA8uyRSSZXe53cEDTHJvDh5OD1efdBpvp0"},
		{"expired token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY3MTc3MTIsImlhdCI6MTU1NjYzMTMxMiwiY2xhaW1zIjp7IlVVSUQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDEiLCJDb21wYW55IjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiUm9sZSI6Im1hbmFnZXIifX0.TQ_jxVTmQ5we-SOpYHGkmBdeOshHkl6mt_4ckUnJ82Y"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grant, err := middleware.Authenticate(test.token)
			assert.Equal(t, err, &errors.ErrTokenInvalid)
			assert.Equal(t, grant, &middleware.Grant{})
		})
	}
}
