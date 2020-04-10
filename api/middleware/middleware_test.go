package middleware_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// func TestMain(m *testing.M) {
// 	db, err := gorm.Open("postgres", "host=test_db port=5432 user=test dbname=testdb password=test sslmode=disable")
// 	if err != nil {
// 		fmt.Println("Failed to connected to db")
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	database.GormDB = db
// 	cleaner := DeleteCreatedEntities(db)

// 	migration.InitMigrations()

// 	exitVal := m.Run()

// 	cleaner()
// 	db.Close()
// 	os.Exit(exitVal)
// }
var (
	db       *sql.DB
	fixtures *testfixtures.Loader
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
