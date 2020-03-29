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
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// GormDB The database object that can be used by middleware to get data
var GormDB *gorm.DB

//SetupDatabase - Connects the database
func SetupDatabase() error {
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
	fmt.Print(connectionString)
	db, errConnect := gorm.Open("postgres", connectionString)
	if errConnect != nil {
		return errConnect
	}

	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Manager{})
	db.AutoMigrate(&models.Delegate{})

	GormDB = db
	return nil
}
