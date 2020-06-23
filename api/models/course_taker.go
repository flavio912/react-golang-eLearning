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
	UUID            gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID `gorm:"primary_key"`
	Course          Course
	CourseID        uint `gorm:"primary_key;auto_increment:false"` // FKEY
	CurrentAttempt  uint `gorm:"default:1"`                        // Starts at 1
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

type TestMark struct {
	TestUUID   gentypes.UUID `gorm:"type:uuid;"` //FKEY
	NumCorrect uint
	// Total keeps its own record of total marks in case the
	// test changes in between a user completing this test and completing the course
	Total uint
}
