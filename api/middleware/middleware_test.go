package middleware_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
)

func TestMain(m *testing.M) {
	db, err := gorm.Open("postgres", "host=test_db port=5432 user=test dbname=testdb password=test sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connected to db")
		fmt.Println(err)
		os.Exit(1)
	}
	database.GormDB = db
	defer db.Close()
	migration.InitMigrations()

	m.Run()

	os.Exit(0)
}
