package middleware

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func (g *Grant) CreatePendingOrder(price float64, courseIds []uint, courseTakerIDs []uint, extraInvoiceEmail *string) (*stripe.PaymentIntent, error) {
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

func (g *Grant) FulfilPendingOrder() {

}
