package user_test

import (
	"fmt"
	"os"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"

	"github.com/go-testfixtures/testfixtures/v3"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/testhelpers"
)

var fixtures *testfixtures.Loader
var usersRepo = user.NewUsersRepository(&logging.Logger{})

func TestMain(m *testing.M) {
	var err error
	fixtures, err = testhelpers.SetupTestDatabase(false, "middleware_test")
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
