package models

import (
	"time"

	"github.com/google/uuid"
)

type CourseInfo struct {
	ID              uint // PRIM, is uint as no need to ever go clientside
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string    // The course name/title
	Price           float64   // 0.00 if free course
	Color           string    // The primary color for the course
	Category        Category  // The category that the course belongs to
	CategoryID      uuid.UUID // FKEY
	Tags            []Tag     `gorm:"many2many:course_tags_link;"`
	Excerpt         string    `sql:"json"` // Excert quill json
	Introduction    string    `sql:"json"` // Introduction quill json
	AccessType      int       // Restricted or Open Access
	ImageKey        string    // S3 Key for the course image
	BackgroundCheck bool      // Is a background check required
	SpecificTerms   string    `sql:"json"` // Terms specific to this course in qull json
}

type Category struct {
	UUID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string
}

type Tag struct {
	UUID  uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name  string    // UNIQUE
	Color string    // A hex color for the tag
}

type OnlineCourse struct {
	UUID         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseInfo   CourseInfo
	CourseInfoID uint // FKEY
	Structure    []CourseStructure
}

type CourseStructure struct {
	OnlineCourseID uuid.UUID
	ModuleID       *uuid.UUID
	LessonID       *uuid.UUID
	TestID         *uuid.UUID
	Rank           string
}

type ClassroomCourse struct {
	UUID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseInfo CourseInfo
	//Tutor      Tutor // The tutor user running this course
	StartDate time.Time
	EndDate   time.Time
	Location  string // e.g The Ritz, London.

	// Classroom courses can require you to take some online courses first
	OnlineCourses []OnlineCourse `gorm:"many2many:online_classroom_link;"`
}
