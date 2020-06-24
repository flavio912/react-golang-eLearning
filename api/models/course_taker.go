package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CourseTaker struct {
	UUID          gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ActiveCourses []ActiveCourse
	Activity      CourseTakerActivity
}

// ActiveCourse represents a course the taker is assigned and can complete
type ActiveCourse struct {
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID `gorm:"primary_key;type:uuid;"`
	Course          Course
	CourseID        uint `gorm:"primary_key;type:uuid;"` // FKEY
	CurrentAttempt  uint `gorm:"default:1"`              // Starts at 1
	MinutesTracked  float64
}

type CourseTakerActivity struct {
	UUID            gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time
	CourseTakerUUID gentypes.UUID         `gorm:"type:uuid;"`
	ActivityType    gentypes.ActivityType // The type of activity
	CourseID        *uint
	Course          Course
}

// TestMark Stores a mark for a user on a test
type TestMark struct {
	TestUUID        gentypes.UUID `gorm:"primary_key;type:uuid;"`
	CourseTakerUUID gentypes.UUID `gorm:"primary_key;type:uuid;"`
	CourseID        uint          `gorm:"primary_key;auto_increment:false"`
	NumCorrect      uint
	// Total keeps its own record of total marks in case the
	// test changes in between a user completing this test and completing the course
	Total uint
}
