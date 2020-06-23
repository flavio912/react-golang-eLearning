package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"
)

type CourseApp interface {
	PurchaseCourses(input gentypes.PurchaseCoursesInput) (*gentypes.PurchaseCoursesResponse, error)

	Course(courseID uint) (gentypes.Course, error)
	Courses(courseIDs []uint) ([]gentypes.Course, error)
	GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error)
	CourseSyllabus(courseID uint) ([]gentypes.CourseItem, error)

	SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error)
	SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error)

	CreateTag(input gentypes.CreateTagInput) (gentypes.Tag, error)
	ManyCourseTags(ids []uint) (map[uint][]gentypes.Tag, error)
	GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]gentypes.Tag, error)
	GetTagsByLessonUUID(uuid string) ([]gentypes.Tag, error)

	GetLessonsByUUID(uuid []string) ([]gentypes.Lesson, error)
	CreateLesson(lesson gentypes.CreateLessonInput) (gentypes.Lesson, error)
	GetLessons(
		page *gentypes.Page,
		filter *gentypes.LessonFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.Lesson, gentypes.PageInfo, error)

	Test(testUUID gentypes.UUID) (gentypes.Test, error)
}

type courseAppImpl struct {
	grant             *middleware.Grant
	ordersRepository  middleware.OrdersRepository
	coursesRepository course.CoursesRepository
	usersRepository   user.UsersRepository
}

func NewCourseApp(grant *middleware.Grant) CourseApp {
	return &courseAppImpl{
		grant:             grant,
		ordersRepository:  middleware.NewOrdersRepository(&grant.Logger),
		coursesRepository: course.NewCoursesRepository(&grant.Logger),
		usersRepository:   user.NewUsersRepository(&grant.Logger),
	}
}
