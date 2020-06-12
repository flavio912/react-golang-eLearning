package courses

import (
	"github.com/asaskevich/govalidator"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func PurchaseCourses(grant *middleware.Grant, input gentypes.PurchaseCoursesInput) (gentypes.PurchaseCoursesResponse, error) {
	// Validate input
	if ok, err := govalidator.ValidateStruct(input); !ok {
		return gentypes.PurchaseCoursesResponse{}, err
	}

	// Find courses
	courseModels, err := g.getCourseModels(helpers.Int32sToUints(input.Courses))
	if err != nil {
		return gentypes.PurchaseCoursesResponse{}, err
	}

	if !g.isAuthorizedToBook(courseModels) {
		return gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorizedToBook
	}

	// Check users exist and are valid

	//
}
