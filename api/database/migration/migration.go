package migration

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// InitMigrations - Run the auto database migrations
func InitMigrations() {
	database.GormDB.AutoMigrate(&models.Admin{})
	database.GormDB.AutoMigrate(&models.Manager{})
	database.GormDB.AutoMigrate(&models.Company{})
	database.GormDB.AutoMigrate(&models.Address{})
	database.GormDB.AutoMigrate(&models.Delegate{})

	// Courses
	database.GormDB.AutoMigrate(&models.CourseInfo{})
	database.GormDB.AutoMigrate(&models.CourseStructure{})
	database.GormDB.AutoMigrate(&models.Category{})
	database.GormDB.AutoMigrate(&models.Tag{})
	database.GormDB.AutoMigrate(&models.OnlineCourse{})
	database.GormDB.AutoMigrate(&models.ClassroomCourse{})
	database.GormDB.AutoMigrate(&models.Module{})
	database.GormDB.AutoMigrate(&models.ModuleStructure{})
	database.GormDB.AutoMigrate(&models.Lesson{})
	database.GormDB.AutoMigrate(&models.Test{})
	database.GormDB.AutoMigrate(&models.WhatYouLearnBullet{})
	database.GormDB.AutoMigrate(&models.RequirementBullet{})
}
