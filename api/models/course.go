package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type Course struct {
	ID                   uint
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 string         // The course name/title
	Price                float64        // 0.00 if free course
	Color                string         // The primary color for the course
	Category             Category       // The category that the course belongs to
	CategoryUUID         *gentypes.UUID // FKEY
	Tags                 []Tag          `gorm:"many2many:course_tags_link;"`
	Excerpt              string
	Introduction         string
	HowToComplete        string
	HoursToComplete      float64
	WhatYouLearn         []WhatYouLearnBullet
	Requirements         []RequirementBullet
	AccessType           gentypes.AccessType // Restricted or Open Access
	ImageKey             *string             // S3 Key for the course image
	BackgroundCheck      bool                // Is a background check required
	SpecificTerms        string              // Terms specific to this course
	Published            bool                // If not published users can't see this course
	CourseType           gentypes.CourseType // classroom or online course
	ExpiresInMonths      uint                // The number of months before the certificate for the course
	ExpirationToEndMonth bool                // If true the expiration date is calculated starting at the end of this month
	CertificateTypeUUID  *gentypes.UUID      // The type of certificate to use
	OnlineCourse         *OnlineCourse
	ClassroomCourse      *ClassroomCourse
}

type RequirementBullet struct {
	ID       uint
	OrderID  int // The precedence of the bullet point in the list
	CourseID uint
	Text     string
}

type WhatYouLearnBullet struct {
	ID       uint
	OrderID  int // The precedence of the bullet point in the list
	CourseID uint
	Text     string
}

// CourseTagsLink is not needed to create the table, but
// is used to extract information about the course_tags_link table
type CourseTagsLink struct {
	CourseID uint
	TagUUID  gentypes.UUID
}

func (CourseTagsLink) TableName() string {
	return "course_tags_link"
}

type Category struct {
	Base
	Name  string `gorm:"unique"`
	Color string
}

type Tag struct {
	Base
	Name  string `gorm:"unique"`
	Color string // A hex color for the tag
}

type OnlineCourse struct {
	Base
	CourseID  uint // FKEY
	Structure []CourseStructure
}

type CourseStructure struct {
	OnlineCourseUUID gentypes.UUID `gorm:"primary_key;"`
	ModuleUUID       *gentypes.UUID
	LessonUUID       *gentypes.UUID
	TestUUID         *gentypes.UUID
	Rank             string `gorm:"primary_key;"`
}

type ClassroomCourse struct {
	Base
	CourseID uint
	//Tutor      Tutor // The tutor user running this course
	StartDate       gentypes.Time
	EndDate         gentypes.Time
	Location        string // e.g The Ritz, London.
	MaxParticipants int
	// Classroom courses can require you to take some online courses first
	OnlineCourses []OnlineCourse `gorm:"many2many:online_classroom_link;"`
}
