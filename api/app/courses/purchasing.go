package courses

import (
	"github.com/asaskevich/govalidator"
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

	// TODO: If you are an individual you can only purchase for yourself so ignore users
	if grant.IsIndividual {
		// courseTakerIDs = grant.Individual(grant.Claims.UUID).CourseTakerID
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorizedToBook
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
		//TODO: check if contract customer
	}

	// If normal purchasing applies
	return &gentypes.PurchaseCoursesResponse{
		StripeClientSecret:  &intent.ClientSecret,
		TransactionComplete: false, // As user still needs to pay
	}, nil
}
