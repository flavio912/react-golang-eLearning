// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	gentypes "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	course "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"

	gorm "github.com/jinzhu/gorm"

	mock "github.com/stretchr/testify/mock"

	models "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// CoursesRepository is an autogenerated mock type for the CoursesRepository type
type CoursesRepository struct {
	mock.Mock
}

// AreInCourses provides a mock function with given fields: courseIDs, uuids, courseElement
func (_m *CoursesRepository) AreInCourses(courseIDs []uint, uuids []gentypes.UUID, courseElement gentypes.CourseElement) (bool, error) {
	ret := _m.Called(courseIDs, uuids, courseElement)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]uint, []gentypes.UUID, gentypes.CourseElement) bool); ok {
		r0 = rf(courseIDs, uuids, courseElement)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint, []gentypes.UUID, gentypes.CourseElement) error); ok {
		r1 = rf(courseIDs, uuids, courseElement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CAANumber provides a mock function with given fields: uuid
func (_m *CoursesRepository) CAANumber(uuid gentypes.UUID) (models.CAANumber, error) {
	ret := _m.Called(uuid)

	var r0 models.CAANumber
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.CAANumber); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(models.CAANumber)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CAANumbers provides a mock function with given fields: page, filter
func (_m *CoursesRepository) CAANumbers(page *gentypes.Page, filter *gentypes.CAANumberFilter) ([]models.CAANumber, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter)

	var r0 []models.CAANumber
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.CAANumberFilter) []models.CAANumber); ok {
		r0 = rf(page, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.CAANumber)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.CAANumberFilter) gentypes.PageInfo); ok {
		r1 = rf(page, filter)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.CAANumberFilter) error); ok {
		r2 = rf(page, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Categories provides a mock function with given fields: page, text
func (_m *CoursesRepository) Categories(page *gentypes.Page, text *string) ([]models.Category, gentypes.PageInfo, error) {
	ret := _m.Called(page, text)

	var r0 []models.Category
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *string) []models.Category); ok {
		r0 = rf(page, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Category)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *string) gentypes.PageInfo); ok {
		r1 = rf(page, text)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *string) error); ok {
		r2 = rf(page, text)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CertificateType provides a mock function with given fields: uuid
func (_m *CoursesRepository) CertificateType(uuid gentypes.UUID) (models.CertificateType, error) {
	ret := _m.Called(uuid)

	var r0 models.CertificateType
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.CertificateType); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(models.CertificateType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateTypes provides a mock function with given fields: page, filter
func (_m *CoursesRepository) CertificateTypes(page *gentypes.Page, filter *gentypes.CertificateTypeFilter) ([]models.CertificateType, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter)

	var r0 []models.CertificateType
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.CertificateTypeFilter) []models.CertificateType); ok {
		r0 = rf(page, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.CertificateType)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.CertificateTypeFilter) gentypes.PageInfo); ok {
		r1 = rf(page, filter)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.CertificateTypeFilter) error); ok {
		r2 = rf(page, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CheckTagsExist provides a mock function with given fields: tags
func (_m *CoursesRepository) CheckTagsExist(tags []gentypes.UUID) ([]models.Tag, error) {
	ret := _m.Called(tags)

	var r0 []models.Tag
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) []models.Tag); ok {
		r0 = rf(tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComposeCourse provides a mock function with given fields: courseInfo
func (_m *CoursesRepository) ComposeCourse(courseInfo course.CourseInput) (models.Course, error) {
	ret := _m.Called(courseInfo)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(course.CourseInput) models.Course); ok {
		r0 = rf(courseInfo)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.CourseInput) error); ok {
		r1 = rf(courseInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Course provides a mock function with given fields: courseID
func (_m *CoursesRepository) Course(courseID uint) (models.Course, error) {
	ret := _m.Called(courseID)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(uint) models.Course); ok {
		r0 = rf(courseID)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CourseTests provides a mock function with given fields: onlineCourseUUID
func (_m *CoursesRepository) CourseTests(onlineCourseUUID gentypes.UUID) ([]models.Test, error) {
	ret := _m.Called(onlineCourseUUID)

	var r0 []models.Test
	if rf, ok := ret.Get(0).(func(gentypes.UUID) []models.Test); ok {
		r0 = rf(onlineCourseUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Test)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(onlineCourseUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Courses provides a mock function with given fields: courseIDs, showUnpublished
func (_m *CoursesRepository) Courses(courseIDs []uint, showUnpublished bool) ([]models.Course, error) {
	ret := _m.Called(courseIDs, showUnpublished)

	var r0 []models.Course
	if rf, ok := ret.Get(0).(func([]uint, bool) []models.Course); ok {
		r0 = rf(courseIDs, showUnpublished)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint, bool) error); ok {
		r1 = rf(courseIDs, showUnpublished)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCAANumber provides a mock function with given fields: identifier
func (_m *CoursesRepository) CreateCAANumber(identifier string) (models.CAANumber, error) {
	ret := _m.Called(identifier)

	var r0 models.CAANumber
	if rf, ok := ret.Get(0).(func(string) models.CAANumber); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Get(0).(models.CAANumber)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCertificateType provides a mock function with given fields: input
func (_m *CoursesRepository) CreateCertificateType(input gentypes.CreateCertificateTypeInput) (models.CertificateType, error) {
	ret := _m.Called(input)

	var r0 models.CertificateType
	if rf, ok := ret.Get(0).(func(gentypes.CreateCertificateTypeInput) models.CertificateType); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.CertificateType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.CreateCertificateTypeInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateClassroomCourse provides a mock function with given fields: courseInfo
func (_m *CoursesRepository) CreateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error) {
	ret := _m.Called(courseInfo)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(gentypes.SaveClassroomCourseInput) models.Course); ok {
		r0 = rf(courseInfo)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.SaveClassroomCourseInput) error); ok {
		r1 = rf(courseInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLesson provides a mock function with given fields: lesson
func (_m *CoursesRepository) CreateLesson(lesson course.CreateLessonInput) (models.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 models.Lesson
	if rf, ok := ret.Get(0).(func(course.CreateLessonInput) models.Lesson); ok {
		r0 = rf(lesson)
	} else {
		r0 = ret.Get(0).(models.Lesson)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.CreateLessonInput) error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateModule provides a mock function with given fields: input
func (_m *CoursesRepository) CreateModule(input course.CreateModuleInput) (models.Module, error) {
	ret := _m.Called(input)

	var r0 models.Module
	if rf, ok := ret.Get(0).(func(course.CreateModuleInput) models.Module); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Module)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.CreateModuleInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOnlineCourse provides a mock function with given fields: courseInfo
func (_m *CoursesRepository) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error) {
	ret := _m.Called(courseInfo)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(gentypes.SaveOnlineCourseInput) models.Course); ok {
		r0 = rf(courseInfo)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.SaveOnlineCourseInput) error); ok {
		r1 = rf(courseInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQuestion provides a mock function with given fields: input
func (_m *CoursesRepository) CreateQuestion(input course.CreateQuestionArgs) (models.Question, error) {
	ret := _m.Called(input)

	var r0 models.Question
	if rf, ok := ret.Get(0).(func(course.CreateQuestionArgs) models.Question); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Question)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.CreateQuestionArgs) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTag provides a mock function with given fields: input
func (_m *CoursesRepository) CreateTag(input gentypes.CreateTagInput) (models.Tag, error) {
	ret := _m.Called(input)

	var r0 models.Tag
	if rf, ok := ret.Get(0).(func(gentypes.CreateTagInput) models.Tag); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.CreateTagInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTest provides a mock function with given fields: input
func (_m *CoursesRepository) CreateTest(input course.CreateTestInput) (models.Test, error) {
	ret := _m.Called(input)

	var r0 models.Test
	if rf, ok := ret.Get(0).(func(course.CreateTestInput) models.Test); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Test)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.CreateTestInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTutor provides a mock function with given fields: details
func (_m *CoursesRepository) CreateTutor(details gentypes.CreateTutorInput) (models.Tutor, error) {
	ret := _m.Called(details)

	var r0 models.Tutor
	if rf, ok := ret.Get(0).(func(gentypes.CreateTutorInput) models.Tutor); ok {
		r0 = rf(details)
	} else {
		r0 = ret.Get(0).(models.Tutor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.CreateTutorInput) error); ok {
		r1 = rf(details)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCourse provides a mock function with given fields: ID
func (_m *CoursesRepository) DeleteCourse(ID uint) (bool, error) {
	ret := _m.Called(ID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint) bool); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLesson provides a mock function with given fields: uuid
func (_m *CoursesRepository) DeleteLesson(uuid gentypes.UUID) (bool, error) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(gentypes.UUID) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteModule provides a mock function with given fields: uuid
func (_m *CoursesRepository) DeleteModule(uuid gentypes.UUID) (bool, error) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(gentypes.UUID) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteQuestion provides a mock function with given fields: input
func (_m *CoursesRepository) DeleteQuestion(input gentypes.UUID) (bool, error) {
	ret := _m.Called(input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(gentypes.UUID) bool); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTest provides a mock function with given fields: uuid
func (_m *CoursesRepository) DeleteTest(uuid gentypes.UUID) (bool, error) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(gentypes.UUID) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourses provides a mock function with given fields: page, filter, orderBy, fullyApproved, showPublished
func (_m *CoursesRepository) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy, fullyApproved bool, showPublished bool) ([]models.Course, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter, orderBy, fullyApproved, showPublished)

	var r0 []models.Course
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.CourseFilter, *gentypes.OrderBy, bool, bool) []models.Course); ok {
		r0 = rf(page, filter, orderBy, fullyApproved, showPublished)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Course)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.CourseFilter, *gentypes.OrderBy, bool, bool) gentypes.PageInfo); ok {
		r1 = rf(page, filter, orderBy, fullyApproved, showPublished)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.CourseFilter, *gentypes.OrderBy, bool, bool) error); ok {
		r2 = rf(page, filter, orderBy, fullyApproved, showPublished)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLessonByUUID provides a mock function with given fields: UUID
func (_m *CoursesRepository) GetLessonByUUID(UUID gentypes.UUID) (models.Lesson, error) {
	ret := _m.Called(UUID)

	var r0 models.Lesson
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.Lesson); ok {
		r0 = rf(UUID)
	} else {
		r0 = ret.Get(0).(models.Lesson)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLessons provides a mock function with given fields: page, filter, orderBy
func (_m *CoursesRepository) GetLessons(page *gentypes.Page, filter *gentypes.LessonFilter, orderBy *gentypes.OrderBy) ([]models.Lesson, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter, orderBy)

	var r0 []models.Lesson
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.LessonFilter, *gentypes.OrderBy) []models.Lesson); ok {
		r0 = rf(page, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Lesson)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.LessonFilter, *gentypes.OrderBy) gentypes.PageInfo); ok {
		r1 = rf(page, filter, orderBy)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.LessonFilter, *gentypes.OrderBy) error); ok {
		r2 = rf(page, filter, orderBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLessonsByUUID provides a mock function with given fields: uuids
func (_m *CoursesRepository) GetLessonsByUUID(uuids []string) ([]models.Lesson, error) {
	ret := _m.Called(uuids)

	var r0 []models.Lesson
	if rf, ok := ret.Get(0).(func([]string) []models.Lesson); ok {
		r0 = rf(uuids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(uuids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModuleByUUID provides a mock function with given fields: moduleUUID
func (_m *CoursesRepository) GetModuleByUUID(moduleUUID gentypes.UUID) (models.Module, error) {
	ret := _m.Called(moduleUUID)

	var r0 models.Module
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.Module); ok {
		r0 = rf(moduleUUID)
	} else {
		r0 = ret.Get(0).(models.Module)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(moduleUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModuleStructure provides a mock function with given fields: moduleUUID
func (_m *CoursesRepository) GetModuleStructure(moduleUUID gentypes.UUID) ([]gentypes.ModuleItem, error) {
	ret := _m.Called(moduleUUID)

	var r0 []gentypes.ModuleItem
	if rf, ok := ret.Get(0).(func(gentypes.UUID) []gentypes.ModuleItem); ok {
		r0 = rf(moduleUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gentypes.ModuleItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(moduleUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: page, filter, orderBy
func (_m *CoursesRepository) GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]models.Tag, error) {
	ret := _m.Called(page, filter, orderBy)

	var r0 []models.Tag
	if rf, ok := ret.Get(0).(func(gentypes.Page, gentypes.GetTagsFilter, gentypes.OrderBy) []models.Tag); ok {
		r0 = rf(page, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.Page, gentypes.GetTagsFilter, gentypes.OrderBy) error); ok {
		r1 = rf(page, filter, orderBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagsByLessonUUID provides a mock function with given fields: uuid
func (_m *CoursesRepository) GetTagsByLessonUUID(uuid string) ([]models.Tag, error) {
	ret := _m.Called(uuid)

	var r0 []models.Tag
	if rf, ok := ret.Get(0).(func(string) []models.Tag); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsModuleInCourses provides a mock function with given fields: courseIDs, moduleUUID
func (_m *CoursesRepository) IsModuleInCourses(courseIDs []uint, moduleUUID gentypes.UUID) (bool, error) {
	ret := _m.Called(courseIDs, moduleUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]uint, gentypes.UUID) bool); ok {
		r0 = rf(courseIDs, moduleUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint, gentypes.UUID) error); ok {
		r1 = rf(courseIDs, moduleUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LearnBullets provides a mock function with given fields: courseID
func (_m *CoursesRepository) LearnBullets(courseID uint) ([]models.WhatYouLearnBullet, error) {
	ret := _m.Called(courseID)

	var r0 []models.WhatYouLearnBullet
	if rf, ok := ret.Get(0).(func(uint) []models.WhatYouLearnBullet); ok {
		r0 = rf(courseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.WhatYouLearnBullet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManyAnswers provides a mock function with given fields: questionUUIDs
func (_m *CoursesRepository) ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.BasicAnswer, error) {
	ret := _m.Called(questionUUIDs)

	var r0 map[gentypes.UUID][]models.BasicAnswer
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) map[gentypes.UUID][]models.BasicAnswer); ok {
		r0 = rf(questionUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[gentypes.UUID][]models.BasicAnswer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(questionUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManyCourseTags provides a mock function with given fields: ids
func (_m *CoursesRepository) ManyCourseTags(ids []uint) (map[uint][]models.Tag, error) {
	ret := _m.Called(ids)

	var r0 map[uint][]models.Tag
	if rf, ok := ret.Get(0).(func([]uint) map[uint][]models.Tag); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint][]models.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManyModuleTags provides a mock function with given fields: moduleUUIDs
func (_m *CoursesRepository) ManyModuleTags(moduleUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.Tag, error) {
	ret := _m.Called(moduleUUIDs)

	var r0 map[gentypes.UUID][]models.Tag
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) map[gentypes.UUID][]models.Tag); ok {
		r0 = rf(moduleUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[gentypes.UUID][]models.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(moduleUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManyOnlineCourseStructures provides a mock function with given fields: onlineCourseUUIDs
func (_m *CoursesRepository) ManyOnlineCourseStructures(onlineCourseUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.CourseStructure, error) {
	ret := _m.Called(onlineCourseUUIDs)

	var r0 map[gentypes.UUID][]models.CourseStructure
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) map[gentypes.UUID][]models.CourseStructure); ok {
		r0 = rf(onlineCourseUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[gentypes.UUID][]models.CourseStructure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(onlineCourseUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManyTests provides a mock function with given fields: testUUIDs
func (_m *CoursesRepository) ManyTests(testUUIDs []gentypes.UUID) (map[gentypes.UUID]models.Test, error) {
	ret := _m.Called(testUUIDs)

	var r0 map[gentypes.UUID]models.Test
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) map[gentypes.UUID]models.Test); ok {
		r0 = rf(testUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[gentypes.UUID]models.Test)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(testUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modules provides a mock function with given fields: page, filter, orderBy
func (_m *CoursesRepository) Modules(page *gentypes.Page, filter *gentypes.ModuleFilter, orderBy *gentypes.OrderBy) ([]models.Module, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter, orderBy)

	var r0 []models.Module
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.ModuleFilter, *gentypes.OrderBy) []models.Module); ok {
		r0 = rf(page, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Module)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.ModuleFilter, *gentypes.OrderBy) gentypes.PageInfo); ok {
		r1 = rf(page, filter, orderBy)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.ModuleFilter, *gentypes.OrderBy) error); ok {
		r2 = rf(page, filter, orderBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ModulesByUUIDs provides a mock function with given fields: uuids
func (_m *CoursesRepository) ModulesByUUIDs(uuids []gentypes.UUID) ([]models.Module, error) {
	ret := _m.Called(uuids)

	var r0 []models.Module
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) []models.Module); ok {
		r0 = rf(uuids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Module)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(uuids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnlineCourse provides a mock function with given fields: courseID
func (_m *CoursesRepository) OnlineCourse(courseID uint) (models.OnlineCourse, error) {
	ret := _m.Called(courseID)

	var r0 models.OnlineCourse
	if rf, ok := ret.Get(0).(func(uint) models.OnlineCourse); ok {
		r0 = rf(courseID)
	} else {
		r0 = ret.Get(0).(models.OnlineCourse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnlineCourseStructure provides a mock function with given fields: onlineCourseUUID
func (_m *CoursesRepository) OnlineCourseStructure(onlineCourseUUID gentypes.UUID) ([]models.CourseStructure, error) {
	ret := _m.Called(onlineCourseUUID)

	var r0 []models.CourseStructure
	if rf, ok := ret.Get(0).(func(gentypes.UUID) []models.CourseStructure); ok {
		r0 = rf(onlineCourseUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.CourseStructure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(onlineCourseUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Question provides a mock function with given fields: uuid
func (_m *CoursesRepository) Question(uuid gentypes.UUID) (models.Question, error) {
	ret := _m.Called(uuid)

	var r0 models.Question
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.Question); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(models.Question)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Questions provides a mock function with given fields: page, filter, orderBy
func (_m *CoursesRepository) Questions(page *gentypes.Page, filter *gentypes.QuestionFilter, orderBy *gentypes.OrderBy) ([]models.Question, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter, orderBy)

	var r0 []models.Question
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.QuestionFilter, *gentypes.OrderBy) []models.Question); ok {
		r0 = rf(page, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Question)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.QuestionFilter, *gentypes.OrderBy) gentypes.PageInfo); ok {
		r1 = rf(page, filter, orderBy)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.QuestionFilter, *gentypes.OrderBy) error); ok {
		r2 = rf(page, filter, orderBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RequirementBullets provides a mock function with given fields: courseID
func (_m *CoursesRepository) RequirementBullets(courseID uint) ([]models.RequirementBullet, error) {
	ret := _m.Called(courseID)

	var r0 []models.RequirementBullet
	if rf, ok := ret.Get(0).(func(uint) []models.RequirementBullet); ok {
		r0 = rf(courseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.RequirementBullet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchSyllabus provides a mock function with given fields: page, filter
func (_m *CoursesRepository) SearchSyllabus(page *gentypes.Page, filter *gentypes.SyllabusFilter) ([]gentypes.CourseItem, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter)

	var r0 []gentypes.CourseItem
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.SyllabusFilter) []gentypes.CourseItem); ok {
		r0 = rf(page, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gentypes.CourseItem)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.SyllabusFilter) gentypes.PageInfo); ok {
		r1 = rf(page, filter)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.SyllabusFilter) error); ok {
		r2 = rf(page, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Test provides a mock function with given fields: testUUID
func (_m *CoursesRepository) Test(testUUID gentypes.UUID) (models.Test, error) {
	ret := _m.Called(testUUID)

	var r0 models.Test
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.Test); ok {
		r0 = rf(testUUID)
	} else {
		r0 = ret.Get(0).(models.Test)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(testUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestQuestions provides a mock function with given fields: testUUID
func (_m *CoursesRepository) TestQuestions(testUUID gentypes.UUID) ([]models.Question, error) {
	ret := _m.Called(testUUID)

	var r0 []models.Question
	if rf, ok := ret.Get(0).(func(gentypes.UUID) []models.Question); ok {
		r0 = rf(testUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(testUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tests provides a mock function with given fields: page, filter, orderBy
func (_m *CoursesRepository) Tests(page *gentypes.Page, filter *gentypes.TestFilter, orderBy *gentypes.OrderBy) ([]models.Test, gentypes.PageInfo, error) {
	ret := _m.Called(page, filter, orderBy)

	var r0 []models.Test
	if rf, ok := ret.Get(0).(func(*gentypes.Page, *gentypes.TestFilter, *gentypes.OrderBy) []models.Test); ok {
		r0 = rf(page, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Test)
		}
	}

	var r1 gentypes.PageInfo
	if rf, ok := ret.Get(1).(func(*gentypes.Page, *gentypes.TestFilter, *gentypes.OrderBy) gentypes.PageInfo); ok {
		r1 = rf(page, filter, orderBy)
	} else {
		r1 = ret.Get(1).(gentypes.PageInfo)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gentypes.Page, *gentypes.TestFilter, *gentypes.OrderBy) error); ok {
		r2 = rf(page, filter, orderBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TestsByUUIDs provides a mock function with given fields: testUUIDs
func (_m *CoursesRepository) TestsByUUIDs(testUUIDs []gentypes.UUID) ([]models.Test, error) {
	ret := _m.Called(testUUIDs)

	var r0 []models.Test
	if rf, ok := ret.Get(0).(func([]gentypes.UUID) []models.Test); ok {
		r0 = rf(testUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Test)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]gentypes.UUID) error); ok {
		r1 = rf(testUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tutor provides a mock function with given fields: uuid
func (_m *CoursesRepository) Tutor(uuid gentypes.UUID) (models.Tutor, error) {
	ret := _m.Called(uuid)

	var r0 models.Tutor
	if rf, ok := ret.Get(0).(func(gentypes.UUID) models.Tutor); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(models.Tutor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UUID) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCAANumber provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateCAANumber(input gentypes.UpdateCAANumberInput) (models.CAANumber, error) {
	ret := _m.Called(input)

	var r0 models.CAANumber
	if rf, ok := ret.Get(0).(func(gentypes.UpdateCAANumberInput) models.CAANumber); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.CAANumber)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UpdateCAANumberInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCertificateType provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateCertificateType(input gentypes.UpdateCertificateTypeInput) (models.CertificateType, error) {
	ret := _m.Called(input)

	var r0 models.CertificateType
	if rf, ok := ret.Get(0).(func(gentypes.UpdateCertificateTypeInput) models.CertificateType); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.CertificateType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UpdateCertificateTypeInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClassroomCourse provides a mock function with given fields: courseInfo
func (_m *CoursesRepository) UpdateClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (models.Course, error) {
	ret := _m.Called(courseInfo)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(gentypes.SaveClassroomCourseInput) models.Course); ok {
		r0 = rf(courseInfo)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.SaveClassroomCourseInput) error); ok {
		r1 = rf(courseInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCourse provides a mock function with given fields: courseID, infoChanges
func (_m *CoursesRepository) UpdateCourse(courseID uint, infoChanges course.CourseInput) (models.Course, error) {
	ret := _m.Called(courseID, infoChanges)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(uint, course.CourseInput) models.Course); ok {
		r0 = rf(courseID, infoChanges)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, course.CourseInput) error); ok {
		r1 = rf(courseID, infoChanges)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLesson provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateLesson(input gentypes.UpdateLessonInput) (models.Lesson, error) {
	ret := _m.Called(input)

	var r0 models.Lesson
	if rf, ok := ret.Get(0).(func(gentypes.UpdateLessonInput) models.Lesson); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Lesson)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UpdateLessonInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateModule provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateModule(input course.UpdateModuleInput) (models.Module, error) {
	ret := _m.Called(input)

	var r0 models.Module
	if rf, ok := ret.Get(0).(func(course.UpdateModuleInput) models.Module); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Module)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.UpdateModuleInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateModuleStructure provides a mock function with given fields: tx, moduleUUID, moduleStructure
func (_m *CoursesRepository) UpdateModuleStructure(tx *gorm.DB, moduleUUID gentypes.UUID, moduleStructure []gentypes.ModuleItem) (models.Module, error) {
	ret := _m.Called(tx, moduleUUID, moduleStructure)

	var r0 models.Module
	if rf, ok := ret.Get(0).(func(*gorm.DB, gentypes.UUID, []gentypes.ModuleItem) models.Module); ok {
		r0 = rf(tx, moduleUUID, moduleStructure)
	} else {
		r0 = ret.Get(0).(models.Module)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, gentypes.UUID, []gentypes.ModuleItem) error); ok {
		r1 = rf(tx, moduleUUID, moduleStructure)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOnlineCourse provides a mock function with given fields: courseInfo
func (_m *CoursesRepository) UpdateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (models.Course, error) {
	ret := _m.Called(courseInfo)

	var r0 models.Course
	if rf, ok := ret.Get(0).(func(gentypes.SaveOnlineCourseInput) models.Course); ok {
		r0 = rf(courseInfo)
	} else {
		r0 = ret.Get(0).(models.Course)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.SaveOnlineCourseInput) error); ok {
		r1 = rf(courseInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuestion provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateQuestion(input course.UpdateQuestionArgs) (models.Question, error) {
	ret := _m.Called(input)

	var r0 models.Question
	if rf, ok := ret.Get(0).(func(course.UpdateQuestionArgs) models.Question); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Question)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.UpdateQuestionArgs) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTest provides a mock function with given fields: input
func (_m *CoursesRepository) UpdateTest(input course.UpdateTestInput) (models.Test, error) {
	ret := _m.Called(input)

	var r0 models.Test
	if rf, ok := ret.Get(0).(func(course.UpdateTestInput) models.Test); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.Test)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(course.UpdateTestInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTutor provides a mock function with given fields: details
func (_m *CoursesRepository) UpdateTutor(details gentypes.UpdateTutorInput) (models.Tutor, error) {
	ret := _m.Called(details)

	var r0 models.Tutor
	if rf, ok := ret.Get(0).(func(gentypes.UpdateTutorInput) models.Tutor); ok {
		r0 = rf(details)
	} else {
		r0 = ret.Get(0).(models.Tutor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gentypes.UpdateTutorInput) error); ok {
		r1 = rf(details)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTutorSignature provides a mock function with given fields: tutorUUID, s3key
func (_m *CoursesRepository) UpdateTutorSignature(tutorUUID gentypes.UUID, s3key string) error {
	ret := _m.Called(tutorUUID, s3key)

	var r0 error
	if rf, ok := ret.Get(0).(func(gentypes.UUID, string) error); ok {
		r0 = rf(tutorUUID, s3key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
