package gentypes

type OnlineCourse struct {
	Course
}

type OnlineCourseFilter struct {
	CourseInfo *CourseInfoFilter
}

type SaveOnlineCourseInput struct {
	CourseInput
	Structure *[]CourseItem
}
