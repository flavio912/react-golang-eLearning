package gentypes

type OnlineCourse struct {
	Course
}

type SaveOnlineCourseInput struct {
	CourseInput
	Structure *[]CourseItem
}
