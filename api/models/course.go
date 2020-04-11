package models

import (
	"time"

	"github.com/google/uuid"
)

type CourseInfo struct {
	ID              uint // PRIM, is uint as no need to ever go clientside
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	Price           float64
	Tags            []Tag
	Excerpt         string `sql:"json"`
	Introduction    string `sql:"json"`
	AccessType      int    // Restricted or Open Access
	Image           string // Course banner image
	BackgroundCheck bool
	SpecificTerms   string     `sql:"json"`
	Participants    []Delegate // The delegates that have/are done/doing the course
}

type CourseTagLink struct {
	CourseID uint
	TagID    uint
}

type Tag struct {
	UUID  uuid.UUID // PRIM
	Name  string    // UNIQUE
	Color string    // A hex color for the tag
}

type OnlineCourse struct {
	UUID         uuid.UUID // PRIM
	CourseInfo   CourseInfo
	CourseInfoID uint
	Structure    []CourseStructure
}

type ClassroomCourse struct {
	UUID       uuid.UUID // PRIM
	CourseInfo CourseInfo
	Tutor      Tutor // The tutor user running this course
	StartDate  time.Time
	EndDate    time.Time
	Location   string

	// Classroom courses can require you to take some online courses first
	OnlineCourses []OnlineCourse `gorm:"many2many:online_classroom_link;"`
}

type Module struct {
	Structure []ModuleStructure
}

type Lesson struct {
}

type Test struct {
}

type ModuleStructure struct {
	ModuleID uint
	LessonID *uint
	TestID   *uint
	Rank     string
}

type CourseStructure struct {
	OnlineCourseID uint
	ModuleID       *uint
	LessonID       *uint
	TestID         *uint
	Rank           string
}
