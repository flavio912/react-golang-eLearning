package middleware_test

import (
	"database/sql"
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
	cleaner := DeleteCreatedEntities(db)

	migration.InitMigrations()

	exitVal := m.Run()

	cleaner()
	db.Close()
	os.Exit(exitVal)
}

func DeleteCreatedEntities(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		// fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			// fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}
