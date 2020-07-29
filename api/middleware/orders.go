package middleware

import (
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type OrdersRepository interface {
	CreatePendingOrder(clientSecret string, courseIds []uint, courseTakerIds []gentypes.UUID, extraInvoiceEmail *string) error
	FulfilPendingOrder(clientSecret string) ([]models.ActiveCourse, error)
	CancelPendingOrder(clientSecret string) error
}

type ordersRepositoryImpl struct {
	Logger *logging.Logger
}

func NewOrdersRepository(logger *logging.Logger) OrdersRepository {
	return &ordersRepositoryImpl{
		Logger: logger,
	}
}

func (o *ordersRepositoryImpl) CreatePendingOrder(clientSecret string, courseIDs []uint, courseTakerUUIDs []gentypes.UUID, extraInvoiceEmail *string) error {
	// Input validation
	if len(courseIDs) == 0 || len(courseTakerUUIDs) == 0 {
		o.Logger.LogMessage(sentry.LevelWarning, "CourseIDs or takers empty")
		return &errors.ErrNotFound
	}

	var numFoundCourses int
	if err := database.GormDB.Model(models.Course{}).Where("id IN (?)", courseIDs).Count(&numFoundCourses).Error; err != nil {
		o.Logger.Log(sentry.LevelWarning, err, "Unable to get courses for pending order")
		return err
	}
	if numFoundCourses != len(courseIDs) {
		return &errors.ErrNotAllFound
	}

	var numFoundTakers int
	if err := database.GormDB.Model(models.CourseTaker{}).Where("uuid IN (?)", courseTakerUUIDs).Count(&numFoundTakers).Error; err != nil {
		o.Logger.Log(sentry.LevelInfo, err, "Unable to get course takers for pending order")
		return err
	}
	if numFoundTakers != len(courseTakerUUIDs) {
		return &errors.ErrNotAllFound
	}

	courses := make([]models.Course, len(courseIDs))
	for index, id := range courseIDs {
		courses[index] = models.Course{ID: id}
	}

	takers := make([]models.CourseTaker, len(courseTakerUUIDs))
	for index, id := range courseTakerUUIDs {
		takers[index] = models.CourseTaker{UUID: id}
	}

	pendingOrder := models.PendingOrder{
		StripeClientSecret: clientSecret,
		Courses:            courses,
		CourseTakers:       takers,
		ExtraInvoiceEmail:  extraInvoiceEmail,
	}

	if err := database.GormDB.Create(&pendingOrder).Error; err != nil {
		o.Logger.Log(sentry.LevelError, err, "Unable to create pending order")
		return &errors.ErrWhileHandling
	}

	return nil
}

// FulfilPendingOrder gives the users the courses they purchased, should only be run after payment confirmation
// returns the activeCourses created
func (o *ordersRepositoryImpl) FulfilPendingOrder(clientSecret string) ([]models.ActiveCourse, error) {

	var pendingOrder models.PendingOrder
	query := database.GormDB.
		Preload("Courses").
		Preload("CourseTakers").
		Where("stripe_client_secret = ?", clientSecret).
		Find(&pendingOrder)

	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.ActiveCourse{}, &errors.ErrNotFound
		}

		sentry.CaptureException(query.Error)
		return []models.ActiveCourse{}, &errors.ErrWhileHandling
	}

	var createdActive []models.ActiveCourse
	for _, courseTaker := range pendingOrder.CourseTakers {
		var activeCourses []models.ActiveCourse

		for _, course := range pendingOrder.Courses {
			newCourse := models.ActiveCourse{
				CourseTakerUUID: courseTaker.UUID,
				CourseID:        course.ID,
				Status:          gentypes.CourseIncomplete,
				MinutesTracked:  0,
			}
			activeCourses = append(activeCourses, newCourse)
			createdActive = append(createdActive, newCourse)
		}

		courseTaker.ActiveCourses = activeCourses
		if err := database.GormDB.Save(&courseTaker).Error; err != nil {
			sentry.CaptureException(err)
			sentry.CaptureMessage(fmt.Sprintf("Unable to fulfil pending order: %s", clientSecret))
			glog.Errorf("Unable to fufil pending order: %s : %s", err.Error(), clientSecret)
			return []models.ActiveCourse{}, &errors.ErrWhileHandling
		}
	}

	err := database.GormDB.Delete(&pendingOrder).Error
	if err != nil {
		o.Logger.Log(sentry.LevelError, err, "Unable to delete pending order")
		return []models.ActiveCourse{}, &errors.ErrWhileHandling
	}

	return createdActive, nil
}

// CancelPendingOrder deletes a pending order from the DB. Usually after stripe
// has confirmed the payment has failed for some reason
func (o *ordersRepositoryImpl) CancelPendingOrder(clientSecret string) error {
	query := database.GormDB.Where("stripe_client_secret = ?", clientSecret).Delete(&models.PendingOrder{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return &errors.ErrNotFound
		}

		sentry.CaptureException(query.Error)
		return &errors.ErrWhileHandling
	}

	return nil
}
