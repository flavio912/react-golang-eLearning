package application

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func IsFullyApproved(usersRepo *user.UsersRepository, grant *middleware.Grant) bool {
	if grant.IsAdmin {
		return true
	}

	if grant.IsManager {
		// Check company is approved
		if usersRepo == nil {
			return false
		}

		company, err := (*usersRepo).Company(grant.Claims.Company)
		if err != nil {
			return false
		}
		if company.Approved {
			return true
		}
	}

	return false
}

func IsAuthorizedToBook(usersRepo *user.UsersRepository, grant *middleware.Grant, courses []models.Course) bool {
	if !grant.IsManager && !grant.IsIndividual {
		return false
	}

	var canAccessRestricted = false

	if grant.IsManager {
		// Check company is approved
		if usersRepo == nil {
			return false
		}
		company, err := (*usersRepo).Company(grant.Claims.Company)
		if err != nil {
			return false
		}
		if company.Approved {
			canAccessRestricted = true
		}
	}

	for _, course := range courses {
		if course.AccessType == gentypes.Restricted {
			if !canAccessRestricted {
				return false
			}
		}
	}
	return true

}

// GrantCanViewSyllabus checks to see if a user can view a courses syllabus
func GrantCanViewSyllabus(usersRepo *user.UsersRepository, grant *middleware.Grant, courseID uint) bool {
	switch {
	case grant.IsAdmin:
		return true
	case grant.IsManager:
		return false
	case grant.IsIndividual || grant.IsDelegate:
		var takerUUID gentypes.UUID
		if grant.IsIndividual {
			individual, _ := (*usersRepo).Individual(grant.Claims.UUID)
			takerUUID = individual.CourseTakerUUID
		}
		if grant.IsDelegate {
			delegate, _ := (*usersRepo).Delegate(grant.Claims.UUID)
			takerUUID = delegate.CourseTakerUUID
		}
		success, err := (*usersRepo).TakerHasActiveCourse(takerUUID, courseID)
		if err != nil {
			grant.Logger.LogMessage(sentry.LevelError, "Unable to check if grant can view syllabus")
			return false
		}

		return success
	}

	return false
}
