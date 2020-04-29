package gentypes

type ClassroomCourse struct {
	Course
	StartDate       Time
	EndDate         Time
	Location        string
	MaxParticipants int
}

type ClassroomCourseFilter struct {
	CourseInfo *CourseInfoFilter
}

type SaveClassroomCourseInput struct {
	CourseInput
	TutorUUID       *UUID
	MaxParticipants *int
	StartDate       *Time
	EndDate         *Time
	Location        *string
}
