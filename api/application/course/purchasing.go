package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type CourseApp interface {
	PurchaseCourses(input gentypes.PurchaseCoursesInput) (*gentypes.PurchaseCoursesResponse, error)

	Course(courseID uint) (gentypes.Course, error)
	Courses(courseIDs []uint) ([]gentypes.Course, error)
	GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error)

	SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error)
	SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error)

	CreateTag(input gentypes.CreateTagInput) (gentypes.Tag, error)
	GetTagsByCourseInfoIDs(ids []uint) (map[uint][]gentypes.Tag, error)
	GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]gentypes.Tag, error)
	GetTagsByLessonUUID(uuid string) ([]gentypes.Tag, error)

	GetLessonsByUUID(uuid []string) ([]gentypes.Lesson, error)
	CreateLesson(lesson gentypes.CreateLessonInput) (gentypes.Lesson, error)
	GetLessons(
		page *gentypes.Page,
		filter *gentypes.LessonFilter,
		orderBy *gentypes.OrderBy,
	) ([]gentypes.Lesson, gentypes.PageInfo, error)
}

type courseAppImpl struct {
	grant             *middleware.Grant
	ordersRepository  middleware.OrdersRepository
	coursesRepository course.CoursesRepository
}

func NewCourseApp(grant *middleware.Grant) CourseApp {
	return &courseAppImpl{
		grant:             grant,
		ordersRepository:  middleware.NewOrdersRepository(&grant.Logger),
		coursesRepository: course.NewCoursesRepository(&grant.Logger),
	}
}

func (c *courseAppImpl) courseToGentype(courseInfo models.Course) gentypes.Course {

	// Get bullet points
	requirementModels, err := c.coursesRepository.RequirementBullets(courseInfo.ID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to get bullets")
		return gentypes.Course{}
	}

	var requirementBullets []string
	for _, bullet := range requirementModels {
		requirementBullets = append(requirementBullets, bullet.Text)
	}

	// Get WhatYouLearn bullet points
	learnModels, err := c.coursesRepository.LearnBullets(courseInfo.ID)
	if err != nil {
		return gentypes.Course{}
	}

	var learnBullets []string
	for _, bullet := range learnModels {
		learnBullets = append(learnBullets, bullet.Text)
	}

	var allowedToBuy = true
	if courseInfo.AccessType == gentypes.Restricted {
		allowedToBuy = c.grant.IsFullyApproved()
	}

	// TODO: Check if user has access to this course
	return gentypes.Course{
		ID:              courseInfo.ID,
		Name:            courseInfo.Name,
		AccessType:      courseInfo.AccessType,
		BackgroundCheck: courseInfo.BackgroundCheck,
		Price:           courseInfo.Price,
		Color:           courseInfo.Color,
		Introduction:    courseInfo.Introduction,
		HowToComplete:   courseInfo.HowToComplete,
		HoursToComplete: courseInfo.HoursToComplete,
		WhatYouLearn:    learnBullets,
		Requirements:    requirementBullets,
		Excerpt:         courseInfo.Excerpt,
		SpecificTerms:   courseInfo.SpecificTerms,
		CategoryUUID:    courseInfo.CategoryUUID,
		AllowedToBuy:    allowedToBuy,
		CourseType:      courseInfo.CourseType,
	}
}

// TODO: Bulk load rather than making a million db calls
func (c *courseAppImpl) coursesToGentypes(courses []models.Course) []gentypes.Course {
	var genCourses []gentypes.Course
	for _, course := range courses {
		genCourses = append(genCourses, c.courseToGentype(course))
	}
	return genCourses
}

func (c *courseAppImpl) PurchaseCourses(input gentypes.PurchaseCoursesInput) (*gentypes.PurchaseCoursesResponse, error) {
	// Validate input
	if ok, err := govalidator.ValidateStruct(input); !ok {
		return &gentypes.PurchaseCoursesResponse{}, err
	}

	courseModels, err := c.coursesRepository.Courses(helpers.Int32sToUints(input.Courses))
	if err != nil {
		return &gentypes.PurchaseCoursesResponse{}, err
	}

	// Calculate total price in pounds
	var price float64
	for _, course := range courseModels {
		price = price + course.Price
	}

	if !c.grant.IsAuthorizedToBook(courseModels) {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorizedToBook
	}

	var courseTakerIDs []uint

	//	Individual can only book courses for themselves
	if c.grant.IsIndividual {
		ind, err := c.grant.Individual(c.grant.Claims.UUID)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "Unable to get current user")
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}
		courseTakerIDs = []uint{ind.CourseTakerID}
	}

	// Managers can only purchase for users that exist and that they are manager of
	if c.grant.IsManager {
		for _, uuid := range input.Users {
			delegate, err := c.grant.Delegate(uuid)
			if err != nil {
				return &gentypes.PurchaseCoursesResponse{}, errors.ErrDelegateDoesNotExist(uuid.String())
			}

			courseTakerIDs = append(courseTakerIDs, delegate.CourseTakerID)
		}
	}

	// Create paymentIntent
	pennyPrice := int64(price * 100) // This will discard any digit after two decimal places

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(pennyPrice), // Convert to pence
		Currency: stripe.String(string(stripe.CurrencyGBP)),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to create payment intent")
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
	}

	// Create a pending order
	err = c.ordersRepository.CreatePendingOrder(intent.ClientSecret, helpers.Int32sToUints(input.Courses), courseTakerIDs, input.ExtraInvoiceEmail)
	if err != nil {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
	}

	// If manager is part of a contract company don't charge them and fulfil immediately
	if c.grant.IsManager {
		manager, err := c.grant.Manager(c.grant.Claims.UUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		company, err := c.grant.Company(manager.CompanyUUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		if company.IsContract {
			err := c.ordersRepository.FulfilPendingOrder(intent.ClientSecret)
			if err != nil {
				c.grant.Logger.Log(sentry.LevelError, err, "Unable to fulfil contract order")
				return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
			}

			return &gentypes.PurchaseCoursesResponse{
				StripeClientSecret:  nil,
				TransactionComplete: true, // As customer doesn't need to pay
			}, nil
		}
	}

	// If normal purchasing applies
	return &gentypes.PurchaseCoursesResponse{
		StripeClientSecret:  &intent.ClientSecret,
		TransactionComplete: false, // As user still needs to pay
	}, nil
}

func (c *courseAppImpl) Course(courseID uint) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	course, err := c.coursesRepository.Course(courseID)
	return c.courseToGentype(course), err
}

func (c *courseAppImpl) Courses(courseIDs []uint) ([]gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Course{}, &errors.ErrUnauthorized
	}

	courses, err := c.coursesRepository.Courses(courseIDs)
	return c.coursesToGentypes(courses), err
}

func (c *courseAppImpl) GetCourses(page *gentypes.Page, filter *gentypes.CourseFilter, orderBy *gentypes.OrderBy) ([]gentypes.Course, gentypes.PageInfo, error) {
	courses, pageInfo, err := c.coursesRepository.GetCourses(page, filter, orderBy, c.grant.IsFullyApproved())

	return c.coursesToGentypes(courses), pageInfo, err
}

func (c *courseAppImpl) CreateOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	course, err := c.coursesRepository.CreateOnlineCourse(courseInfo)
	return c.courseToGentype(course), err
}

func (c *courseAppImpl) SaveOnlineCourse(courseInfo gentypes.SaveOnlineCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	_, err := govalidator.ValidateStruct(courseInfo)
	if err != nil {
		return gentypes.Course{}, err
	}

	var course models.Course
	// If courseUUID given, update
	if courseInfo.ID != nil {
		// Update CourseInfo
		course, err = c.coursesRepository.UpdateOnlineCourse(courseInfo)
	} else {
		course, err = c.coursesRepository.CreateOnlineCourse(courseInfo)
	}

	return c.courseToGentype(course), err
}

// SaveClassroomCourse is a wrapper around CreateClassroomCourse and UpdateClassroomCourse to
// update the course if a uuid is provided, otherwise it will create a new one
func (c *courseAppImpl) SaveClassroomCourse(courseInfo gentypes.SaveClassroomCourseInput) (gentypes.Course, error) {
	if !c.grant.IsAdmin {
		return gentypes.Course{}, &errors.ErrUnauthorized
	}

	var course models.Course
	var err error
	if courseInfo.ID != nil {
		course, err = c.coursesRepository.UpdateClassroomCourse(courseInfo)
	} else {
		course, err = c.coursesRepository.CreateClassroomCourse(courseInfo)
	}

	return c.courseToGentype(course), err
}

func (c *courseAppImpl) CreateTag(input gentypes.CreateTagInput) (gentypes.Tag, error) {
	if !c.grant.IsAdmin {
		return gentypes.Tag{}, &errors.ErrUnauthorized
	}

	tag, err := c.coursesRepository.CreateTag(input)
	return tagToGentype(tag), err
}

func (c *courseAppImpl) GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]gentypes.Tag, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Tag{}, &errors.ErrUnauthorized
	}

	tags, err := c.coursesRepository.GetTags(page, filter, orderBy)
	return tagsToGentypes(tags), err
}

func (c *courseAppImpl) GetTagsByCourseInfoIDs(ids []uint) (map[uint][]gentypes.Tag, error) {

	tags, err := c.coursesRepository.GetTagsByCourseInfoIDs(ids)

	var genTags = map[uint][]gentypes.Tag{}
	for key, element := range tags {
		genTags[key] = tagsToGentypes(element)
	}

	return genTags, err
}

func (c *courseAppImpl) GetTagsByLessonUUID(uuid string) ([]gentypes.Tag, error) {
	tags, err := c.coursesRepository.GetTagsByLessonUUID(uuid)
	return tagsToGentypes(tags), err
}

func (c *courseAppImpl) GetLessonsByUUID(uuid []string) ([]gentypes.Lesson, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	lessons, err := c.coursesRepository.GetLessonsByUUID(uuid)
	return c.lessonsToGentype(lessons), err
}

func (c *courseAppImpl) CreateLesson(lesson gentypes.CreateLessonInput) (gentypes.Lesson, error) {
	if !c.grant.IsAdmin {
		return gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	lessonMod, err := c.coursesRepository.CreateLesson(lesson)
	return c.lessonToGentype(lessonMod), err
}

func (c *courseAppImpl) GetLessons(
	page *gentypes.Page,
	filter *gentypes.LessonFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Lesson, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Lesson{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	lessons, pageInfo, err := c.coursesRepository.GetLessons(page, filter, orderBy)
	return c.lessonsToGentype(lessons), pageInfo, err
}
