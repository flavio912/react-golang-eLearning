package migration

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// InitMigrations - Run the auto database migrations
func InitMigrations() {
	database.GormDB.AutoMigrate(&models.Admin{})
	database.GormDB.AutoMigrate(&models.Manager{})
	database.GormDB.AutoMigrate(&models.Delegate{})
}