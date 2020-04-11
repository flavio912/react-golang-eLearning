/*
Package database provides functions for saving and retrieving objects from the database.
*/
package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// GormDB The database object that can be used by middleware to get data
var GormDB *gorm.DB

//SetupDatabase - Connects the database
func SetupDatabase(logMode bool) error {
	config := helpers.Config

	// Configure database rds or local test
	host := os.Getenv("RDS_HOSTNAME")
	port := os.Getenv("RDS_PORT")
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	dbName := os.Getenv("RDS_DB_NAME")

	if host == "" {
		host = config.Database.Host
		port = config.Database.Port
		username = config.Database.User
		password = config.Database.Password
		dbName = config.Database.DatabaseName
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		dbName,
		password)
	fmt.Printf("DB Connected: %s\n", connectionString)
	db, errConnect := gorm.Open("postgres", connectionString)
	if errConnect != nil {
		return errConnect
	}

	// TODO: Create debug config option
	db.LogMode(logMode)
	GormDB = db
	return nil
}
