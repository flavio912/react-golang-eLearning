package application

import (
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
