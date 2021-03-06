package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"
)

type CourseApp interface {
	SetOrdersRepository(r middleware.OrdersRepository)
	SetUsersRepository(r user.UsersRepository)
	SetCoursesRepository(r course.CoursesRepository)

	PurchaseCourses(input gentypes.PurchaseCoursesInput) (*gentypes.PurchaseCoursesResponse, error)
	FulfilPendingOrder(clientSecret string) (bool, error)
	CancelPendingOrder(clientSecret string) (bool, error)
	DeleteCourse(input gentypes.DeleteCourseInput) (bool, error)

	Course(courseID uint) (gentypes.Course, error)
	Courses(courseIDs []uint) ([]gentypes.Course, error)
	GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error)
	CourseSyllabus(courseID uint) ([]gentypes.CourseItem, error)

	SetCoursePublished(courseID uint, published bool) error
	SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error)
	SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error)
	CourseBannerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)

	Categories(page *gentypes.Page, text *string) ([]gentypes.Category, gentypes.PageInfo, error)
	UpdateCategory(input gentypes.UpdateCategoryInput) (gentypes.Category, error)
	DeleteCategory(input gentypes.DeleteCategoryInput) error

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
	UpdateLesson(input gentypes.UpdateLessonInput) (gentypes.Lesson, error)
	DeleteLesson(input gentypes.DeleteLessonInput) (bool, error)
	LessonBannerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)

	Test(testUUID gentypes.UUID) (gentypes.Test, error)
	Tests(
		page *gentypes.Page,
		filter *gentypes.TestFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.Test, gentypes.PageInfo, error)
	TestsByUUIDs(uuids []gentypes.UUID) ([]gentypes.Test, error)
	CreateTest(input gentypes.CreateTestInput) (gentypes.Test, error)
	UpdateTest(input gentypes.UpdateTestInput) (gentypes.Test, error)
	SubmitTest(input gentypes.SubmitTestInput) (bool, gentypes.CourseStatus, error)
	DeleteTest(input gentypes.DeleteTestInput) (bool, error)
	TestQuestions(testUUID gentypes.UUID) ([]gentypes.Question, error)

	Module(uuid gentypes.UUID) (gentypes.Module, error)
	ModulesByUUIDs(uuids []gentypes.UUID) ([]gentypes.Module, error)
	Modules(
		page *gentypes.Page,
		filter *gentypes.ModuleFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.Module, gentypes.PageInfo, error)
	CreateModule(input gentypes.CreateModuleInput) (gentypes.Module, error)
	UpdateModule(input gentypes.UpdateModuleInput) (gentypes.Module, error)
	DeleteModule(input gentypes.DeleteModuleInput) (bool, error)
	ModuleBannerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	ModuleSyllabus(uuid gentypes.UUID) ([]gentypes.ModuleItem, error)
	ManyModuleTags(moduleUUIDs []gentypes.UUID) (map[gentypes.UUID][]gentypes.Tag, error)

	VoiceoverUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)

	SearchSyllabus(
		page *gentypes.Page,
		filter *gentypes.SyllabusFilter,
	) ([]gentypes.CourseItem, gentypes.PageInfo, error)

	Question(uuid gentypes.UUID) (gentypes.Question, error)
	Questions(
		page *gentypes.Page,
		filter *gentypes.QuestionFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.Question, gentypes.PageInfo, error)
	CreateQuestion(input gentypes.CreateQuestionInput) (gentypes.Question, error)
	UpdateQuestion(input gentypes.UpdateQuestionInput) (gentypes.Question, error)
	DeleteQuestion(input gentypes.DeleteQuestionInput) (bool, error)
	AnswerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]gentypes.Answer, error)

	CertificateType(uuid gentypes.UUID) (gentypes.CertificateType, error)
	CertificateTypes(
		page *gentypes.Page,
		filter *gentypes.CertificateTypeFilter) ([]gentypes.CertificateType, gentypes.PageInfo, error)
	CreateCertificateType(input gentypes.CreateCertificateTypeInput) (gentypes.CertificateType, error)
	UpdateCertificateType(input gentypes.UpdateCertificateTypeInput) (gentypes.CertificateType, error)
	CAANumbers(
		page *gentypes.Page,
		filter *gentypes.CAANumberFilter) ([]gentypes.CAANumber, gentypes.PageInfo, error)
	CreateCAANumber(input gentypes.CreateCAANumberInput) (gentypes.CAANumber, error)
	UpdateCAANumber(input gentypes.UpdateCAANumberInput) (gentypes.CAANumber, error)
	CertificateBodyImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)

	Tutor(uuid gentypes.UUID) (gentypes.Tutor, error)
	Tutors(
		page *gentypes.Page,
		filter *gentypes.TutorFilter,
		order *gentypes.OrderBy) ([]gentypes.Tutor, gentypes.PageInfo, error)
	CreateTutor(input gentypes.CreateTutorInput) (gentypes.Tutor, error)
	UpdateTutor(input gentypes.UpdateTutorInput) (gentypes.Tutor, error)
	TutorSignatureImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error)
	UpdateTutorSignature(input gentypes.UpdateTutorSignatureInput) (string, error)

	CertificateInfo(token string) (gentypes.CertficateInfo, error)
	RegenerateCertificate(historicalCourseUUID gentypes.UUID) error
}

type courseAppImpl struct {
	grant             *middleware.Grant
	ordersRepository  middleware.OrdersRepository
	coursesRepository course.CoursesRepository
	usersRepository   user.UsersRepository
}

func (c *courseAppImpl) SetOrdersRepository(r middleware.OrdersRepository) {
	c.ordersRepository = r
}

func (c *courseAppImpl) SetUsersRepository(r user.UsersRepository) {
	c.usersRepository = r
}

func (c *courseAppImpl) SetCoursesRepository(r course.CoursesRepository) {
	c.coursesRepository = r
}

func NewCourseApp(grant *middleware.Grant) CourseApp {
	return &courseAppImpl{
		grant:             grant,
		ordersRepository:  middleware.NewOrdersRepository(&grant.Logger),
		coursesRepository: course.NewCoursesRepository(&grant.Logger),
		usersRepository:   user.NewUsersRepository(&grant.Logger),
	}
}
