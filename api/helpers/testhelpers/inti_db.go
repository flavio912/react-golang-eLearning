package testhelpers

import (
	"database/sql"
	"fmt"
	"path"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func SetupTestDatabase(logMode bool, dbName string) (*testfixtures.Loader, error) {
	if err := helpers.LoadConfig("../dev_env/test_config.yml"); err != nil {
		return nil, err
	}
	config := helpers.Config

	host := config.Database.Host
	port := config.Database.Port
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbName,
		dbName,
		dbName)

	gorm, errConnect := gorm.Open("postgres", connectionString)
	if errConnect != nil {
		fmt.Printf("Unable to connect to gorm: %s", errConnect.Error())
		return nil, errConnect
	}
	fmt.Printf("DB Connected: %s\n", connectionString)

	gorm.LogMode(logMode)
	database.GormDB = gorm

	migration.InitMigrations()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Unable to connect to test DB: %s", err.Error())
		return nil, err
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db), // Your database connection
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(path.Dir(filename)+"/fixtures"), // the directory containing the YAML files
	)

	if err != nil {
		fmt.Printf("Unable get fixtures: %s", err.Error())
		panic("Cannot get test fixtures")
	}

	return fixtures, nil
}
