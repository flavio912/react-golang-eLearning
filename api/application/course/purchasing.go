package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

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

	if !application.IsAuthorizedToBook(&c.usersRepository, c.grant, courseModels) {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorizedToBook
	}

	var courseTakerIDs []gentypes.UUID

	//	Individual can only book courses for themselves
	if c.grant.IsIndividual {
		ind, err := c.usersRepository.Individual(c.grant.Claims.UUID)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "Unable to get current user")
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}
		courseTakerIDs = []gentypes.UUID{ind.CourseTakerUUID}
	}

	// Managers can only purchase for users that exist and that they are manager of
	if c.grant.IsManager {
		for _, uuid := range input.Users {
			delegate, err := c.usersRepository.Delegate(uuid)
			if err != nil {
				return &gentypes.PurchaseCoursesResponse{}, errors.ErrDelegateDoesNotExist(uuid.String())
			}

			courseTakerIDs = append(courseTakerIDs, delegate.CourseTakerUUID)
		}
	}

	// Add VAT on top of prices
	pennyPrice := int64((price * 100) + (price * 0.2)) // This will discard any digit after two decimal places

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
		manager, err := c.usersRepository.Manager(c.grant.Claims.UUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		company, err := c.usersRepository.Company(manager.CompanyUUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		if company.IsContract {
			success, err := c.FulfilPendingOrder(intent.ClientSecret)
			if err != nil || !success {
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

func (c *courseAppImpl) FulfilPendingOrder(clientSecret string) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	activeCourses, err := c.ordersRepository.FulfilPendingOrder(clientSecret)
	if err != nil {
		return false, err
	}

	for _, activeCourse := range activeCourses {
		_, err := c.usersRepository.CreateTakerActivity(activeCourse.CourseTakerUUID, gentypes.ActivityNewCourse, &activeCourse.CourseID)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelWarning, err, "FulfilPendingOrder: Unable to create taker activity")
		}
	}

	return true, nil
}

func (c *courseAppImpl) CancelPendingOrder(clientSecret string) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	err := c.ordersRepository.CancelPendingOrder(clientSecret)
	if err != nil {
		return false, err
	}

	return true, nil
}
