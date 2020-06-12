package migration

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// InitMigrations - Run the auto database migrations
func InitMigrations() {
	// Users
	database.GormDB.AutoMigrate(&models.Admin{})
	database.GormDB.AutoMigrate(&models.Manager{})
	database.GormDB.AutoMigrate(&models.Company{})
	database.GormDB.AutoMigrate(&models.Address{})
	database.GormDB.AutoMigrate(&models.Delegate{})
	database.GormDB.AutoMigrate(&models.CourseTaker{})

	// Courses
	database.GormDB.AutoMigrate(&models.Course{})
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

	// If course is deleted, delete the requirements and what you learn too
	database.GormDB.Model(&models.RequirementBullet{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "RESTRICT")
	database.GormDB.Model(&models.WhatYouLearnBullet{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "RESTRICT")

	// Orders
	database.GormDB.AutoMigrate(&models.PendingOrder{})
}
