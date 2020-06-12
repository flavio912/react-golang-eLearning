package middleware_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/testhelpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader

	adminGrant    = middleware.Grant{auth.UserClaims{}, true, false, false, false, false, logging.Logger{}}
	nonAdminGrant = middleware.Grant{auth.UserClaims{}, false, true, true, false, false, logging.Logger{}}
	managerGrant  = middleware.Grant{auth.UserClaims{
		UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		Role:    auth.ManagerRole,
	}, false, true, false, false, false, logging.Logger{}}
	delegateGrant = middleware.Grant{auth.UserClaims{
		UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
		Role:    auth.DelegateRole,
	}, false, false, true, false, false, logging.Logger{}}
	publicGrant = middleware.Grant{IsPublic: true}
	uuidZero    = gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")
)

func TestMain(m *testing.M) {
	var err error
	fixtures, err = testhelpers.SetupTestDatabase(true, "middleware_test")
	if err != nil {
		panic("Failed to init test db")
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Unable to load fixtures for test: %s", err.Error())
		panic("cannot load test fixtures")
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
			assert.Equal(t, grant, &middleware.Grant{IsPublic: true})
		})
	}
}
