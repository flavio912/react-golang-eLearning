package courses

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func PurchaseCourses(grant *middleware.Grant, input gentypes.PurchaseCoursesInput) (*gentypes.PurchaseCoursesResponse, error) {
	// Validate input
	if ok, err := govalidator.ValidateStruct(input); !ok {
		return &gentypes.PurchaseCoursesResponse{}, err
	}

	courseModels, err := grant.Courses(helpers.Int32sToUints(input.Courses))
	if err != nil {
		return &gentypes.PurchaseCoursesResponse{}, err
	}

	// Calculate total price in pounds
	var price float64
	for _, course := range courseModels {
		price = price + course.Price
	}

	if !grant.IsAuthorizedToBook(courseModels) {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorizedToBook
	}

	var courseTakerIDs []uint

	//	Individual can only book courses for themselves
	if grant.IsIndividual {
		ind, err := grant.Individual(grant.Claims.UUID)
		if err != nil {
			grant.Logger.Log(sentry.LevelError, err, "Unable to get current user")
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}
		courseTakerIDs = []uint{ind.CourseTakerID}
	}

	// Managers can only purchase for users that exist and that they are manager of
	if grant.IsManager {
		for _, uuid := range input.Users {
			delegate, err := grant.Delegate(uuid)
			if err != nil {
				return &gentypes.PurchaseCoursesResponse{}, errors.ErrDelegateDoesNotExist(uuid.String())
			}

			courseTakerIDs = append(courseTakerIDs, delegate.CourseTakerID)
		}
	}

	// Create a pending order
	intent, err := grant.CreatePendingOrder(price, helpers.Int32sToUints(input.Courses), courseTakerIDs, input.ExtraInvoiceEmail)
	if err != nil {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
	}

	// If manager is part of a contract company don't charge them and fulfil immediately
	if grant.IsManager {
		manager, err := grant.Manager(grant.Claims.UUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		company, err := grant.Company(manager.CompanyUUID)
		if err != nil {
			return &gentypes.PurchaseCoursesResponse{}, &errors.ErrWhileHandling
		}

		if company.IsContract {
			err := middleware.FulfilPendingOrder(intent.ClientSecret)
			if err != nil {
				grant.Logger.Log(sentry.LevelError, err, "Unable to fulfil contract order")
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
