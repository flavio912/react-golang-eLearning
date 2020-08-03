package users

import (
	"time"

	"github.com/getsentry/sentry-go"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func activityToGentype(activity models.CourseTakerActivity) gentypes.Activity {
	return gentypes.Activity{
		UUID:            activity.UUID,
		CourseTakerUUID: activity.CourseTakerUUID,
		ActivityType:    activity.ActivityType,
		CourseID:        activity.CourseID,
		CreatedAt:       activity.CreatedAt.Format(time.RFC3339),
	}
}

func activityToGentypes(activityItems []models.CourseTakerActivity) []gentypes.Activity {
	items := make([]gentypes.Activity, len(activityItems))
	for i, activity := range activityItems {
		items[i] = activityToGentype(activity)
	}
	return items
}

func (u *usersAppImpl) TakerActivity(courseTakerUUID gentypes.UUID, page *gentypes.Page) ([]gentypes.Activity, gentypes.PageInfo, error) {
	// Enforce max items at a time
	if page == nil {
		plimit := middleware.MaxPageLimit
		page = &gentypes.Page{
			Limit: &plimit,
		}
	}

	// Admin does what they want
	if u.grant.IsAdmin {
		activityItems, pageInfo, err := u.usersRepository.TakerActivity(courseTakerUUID, page)
		return activityToGentypes(activityItems), pageInfo, err
	}

	// Check manager is asking about one of their companys delegates
	if u.grant.IsManager {
		ok, err := u.usersRepository.CompanyManagesCourseTakers(u.grant.Claims.Company, []gentypes.UUID{courseTakerUUID})
		if err == nil && ok {
			activityItems, pageInfo, err := u.usersRepository.TakerActivity(courseTakerUUID, page)
			return activityToGentypes(activityItems), pageInfo, err
		}
	}

	// Check delegate is asking about themselves
	if u.grant.IsDelegate {
		delegate, err := u.usersRepository.Delegate(u.grant.Claims.UUID)
		if err != nil {
			u.grant.Logger.Log(sentry.LevelInfo, err, "Unable to get delegate")
			return []gentypes.Activity{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
		}

		if delegate.CourseTakerUUID == courseTakerUUID {
			activityItems, pageInfo, err := u.usersRepository.TakerActivity(courseTakerUUID, page)
			return activityToGentypes(activityItems), pageInfo, err
		}
	}

	// Check individual is asking about themselves
	if u.grant.IsIndividual {
		individual, err := u.usersRepository.Individual(u.grant.Claims.UUID)
		if err != nil {
			return []gentypes.Activity{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
		}

		if individual.CourseTakerUUID == courseTakerUUID {
			activityItems, pageInfo, err := u.usersRepository.TakerActivity(courseTakerUUID, page)
			return activityToGentypes(activityItems), pageInfo, err
		}
	}

	return []gentypes.Activity{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
}

func (u *usersAppImpl) CompanyActivity(companyUUID gentypes.UUID, page *gentypes.Page) ([]gentypes.Activity, gentypes.PageInfo, error) {
	if !u.grant.ManagesCompany(companyUUID) {
		return []gentypes.Activity{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	// Get company activity
	activity, pageInfo, err := u.usersRepository.CompanyActivity(companyUUID, page)
	if err != nil {
		u.grant.Logger.Log(sentry.LevelError, err, "CompanyActivity: Unable to get company activity")
		return activityToGentypes(activity), pageInfo, &errors.ErrWhileHandling
	}

	return activityToGentypes(activity), pageInfo, nil
}
