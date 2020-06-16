package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CourseTaker struct {
	UUID           gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	MinutesTracked float64
	ActiveCourses  []ActiveCourse
	Activity       CourseTakerActivity
}

// ActiveCourse represents a course the taker is assigned and can complete
type ActiveCourse struct {
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID `gorm:"primary_key"`
	Course          Course
	CourseID        uint `gorm:"primary_key;auto_increment:false"` // FKEY
	CurrentAttempt  uint `gorm:"default:1"`
}

type CourseTakerActivity struct {
	UUID            gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time
	CourseTakerUUID gentypes.UUID
	ActivityType    gentypes.ActivityType // The type of activity
	CourseID        *uint
	Course          Course
}
