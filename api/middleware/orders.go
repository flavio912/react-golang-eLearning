package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type OrdersRepository interface {
	CreatePendingOrder(clientSecret string, courseIds []uint, courseTakerIds []uint, extraInvoiceEmail *string) error
	FulfilPendingOrder(clientSecret string) error
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

func (o *ordersRepositoryImpl) CreatePendingOrder(clientSecret string, courseIDs []uint, courseTakerIDs []uint, extraInvoiceEmail *string) error {
	// Input validation
	if len(courseIDs) == 0 || len(courseTakerIDs) == 0 {
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
	if err := database.GormDB.Model(models.CourseTaker{}).Where("id IN (?)", courseTakerIDs).Count(&numFoundTakers).Error; err != nil {
		o.Logger.Log(sentry.LevelInfo, err, "Unable to get course takers for pending order")
		return err
	}
	if numFoundTakers != len(courseTakerIDs) {
		return &errors.ErrNotAllFound
	}

	courses := make([]models.Course, len(courseIDs))
	for index, id := range courseIDs {
		courses[index] = models.Course{ID: id}
	}

	takers := make([]models.CourseTaker, len(courseTakerIDs))
	for index, id := range courseIDs {
		takers[index] = models.CourseTaker{ID: id}
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
func (o *ordersRepositoryImpl) FulfilPendingOrder(clientSecret string) error {

	var pendingOrder models.PendingOrder
	query := database.GormDB.
		Preload("Courses").
		Preload("CourseTakers").
		Where("stripe_client_secret = ?", clientSecret).
		Find(&pendingOrder)

	if query.Error != nil {
		if query.RecordNotFound() {
			return &errors.ErrNotFound
		}

		sentry.CaptureException(query.Error)
		return &errors.ErrWhileHandling
	}

	for _, courseTaker := range pendingOrder.CourseTakers {
		var activeCourses []models.ActiveCourse

		for _, course := range pendingOrder.Courses {
			activeCourses = append(activeCourses, models.ActiveCourse{
				CourseTakerID: courseTaker.ID,
				CourseID:      course.ID,
			})
		}

		courseTaker.ActiveCourses = activeCourses
		if err := database.GormDB.Save(&courseTaker).Error; err != nil {
			sentry.CaptureException(err)
			sentry.CaptureMessage(fmt.Sprintf("Unable to fulfil pending order: %s", clientSecret))
			glog.Errorf("Unable to fufil pending order: %s : %s", err.Error(), clientSecret)
			return &errors.ErrWhileHandling
		}
	}

	return nil
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
