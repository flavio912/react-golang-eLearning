package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func (g *Grant) CreatePendingOrder(price float64, courseIds []uint, courseTakerIDs []uint, extraInvoiceEmail *string) (*stripe.PaymentIntent, error) {
	// Input validation
	if len(courseIds) == 0 || len(courseTakerIDs) == 0 {
		g.Logger.LogMessage(sentry.LevelWarning, "CourseIds or takers empty")
		return &stripe.PaymentIntent{}, &errors.ErrNotFound
	}

	pennyPrice := int64(price * 100) // This will discard any digit after two decimal places

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(pennyPrice), // Convert to pence
		Currency: stripe.String(string(stripe.CurrencyGBP)),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to create payment intent")
		return &stripe.PaymentIntent{}, &errors.ErrWhileHandling
	}

	courses, err := g.Courses(courseIds)
	if err != nil {
		g.Logger.Log(sentry.LevelInfo, err, "Unable to get courses for pending order")
		return &stripe.PaymentIntent{}, err
	}

	takers, err := g.CourseTakers(courseTakerIDs)
	if err != nil {
		g.Logger.Log(sentry.LevelInfo, err, "Unable to get course takers for pending order")
		return &stripe.PaymentIntent{}, err
	}

	pendingOrder := models.PendingOrder{
		StripeClientSecret: intent.ClientSecret,
		Courses:            courses,
		CourseTakers:       takers,
		ExtraInvoiceEmail:  extraInvoiceEmail,
	}

	if err := database.GormDB.Create(&pendingOrder).Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to create pending order")
		return &stripe.PaymentIntent{}, &errors.ErrWhileHandling
	}

	return intent, nil
}

// FulfilPendingOrder gives the users the courses they purchased, should only be run after payment confirmation
func FulfilPendingOrder(clientSecret string) error {

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
func CancelPendingOrder(clientSecret string) error {
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
