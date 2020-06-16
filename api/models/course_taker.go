package models

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CourseTaker struct {
	ID             uint
	MinutesTracked float64
	ActiveCourses  []ActiveCourse
	Activity       CourseTakerActivity
}

// ActiveCourse represents a course the taker is assigned and can complete
type ActiveCourse struct {
	CourseTaker    CourseTaker
	CourseTakerID  uint `gorm:"primary_key;auto_increment:false"`
	Course         Course
	CourseID       uint `gorm:"primary_key;auto_increment:false"`
	CurrentAttempt uint `gorm:"default:1"`
}

type CourseTakerActivity struct {
	CourseTakerID uint
	ActivityType  gentypes.ActivityType // The type of activity
	CourseID      *uint
	Course        Course
}
