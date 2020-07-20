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
	Status          gentypes.CourseStatus
	CreatedAt       time.Time
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID `gorm:"primary_key;type:uuid;"`
	Course          Course
	CourseID        uint `gorm:"primary_key;"` // FKEY
	MinutesTracked  float64
}

type HistoricalCourse struct {
	Status          gentypes.CourseStatus
	UUID            gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time
	CourseTaker     CourseTaker
	CourseTakerUUID gentypes.UUID
	Course          Course
	CourseID        uint
	Passed          bool
	MinutesTracked  float64
	ExpirationDate  *time.Time
	CertificateKey  *string
}

type CourseTakerActivity struct {
	UUID            gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time
	CourseTakerUUID gentypes.UUID         `gorm:"type:uuid;"`
	ActivityType    gentypes.ActivityType // The type of activity
	CourseID        *uint
	Course          Course
}

// TestMark stores a mark for a user on a test
type TestMark struct {
	CreatedAt       time.Time
	TestUUID        gentypes.UUID `gorm:"primary_key;type:uuid;"`
	CourseTakerUUID gentypes.UUID `gorm:"primary_key;type:uuid;"`
	CourseID        uint          `gorm:"primary_key;auto_increment:false"`
	Passed          bool
	CurrentAttempt  uint
}
