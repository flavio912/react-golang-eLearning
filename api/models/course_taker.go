package models

type CourseTaker struct {
	ID             uint
	MinutesTracked float64
	ActiveCourses  []ActiveCourse
}

// ActiveCourse represents a course the taker is assigned and can complete
type ActiveCourse struct {
	CourseTaker    CourseTaker
	CourseTakerID  uint `gorm:"primary_key;auto_increment:false"`
	Course         Course
	CourseID       uint `gorm:"primary_key;auto_increment:false"`
	CurrentAttempt uint `gorm:"default:1"`
}
