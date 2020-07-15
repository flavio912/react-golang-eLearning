package users

import (
	"time"

	"github.com/getsentry/sentry-go"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func activeCourseToGentype(activeCourse models.ActiveCourse) gentypes.ActiveCourse {
	return gentypes.ActiveCourse{
		CourseID:       activeCourse.CourseID,
		MinutesTracked: activeCourse.MinutesTracked,
		Status:         activeCourse.Status,
	}
}

func activeCoursesToGentypes(activeCourses []models.ActiveCourse) []gentypes.ActiveCourse {
	var genCourses = make([]gentypes.ActiveCourse, len(activeCourses))
	for i, course := range activeCourses {
		genCourses[i] = activeCourseToGentype(course)
	}
	return genCourses
}

func (u *usersAppImpl) DelegateToUser(delegate models.Delegate) gentypes.User {

	var uploadUrl *string
	if delegate.ProfileKey != nil {
		uploadUrl = helpers.StringPointer(uploads.GetImgixURL(*delegate.ProfileKey))
	}

	return gentypes.User{
		UUID:            delegate.UUID,
		Type:            gentypes.DelegateType,
		Email:           delegate.Email,
		FirstName:       delegate.FirstName,
		LastName:        delegate.LastName,
		Telephone:       delegate.Telephone,
		JobTitle:        &delegate.JobTitle,
		LastLogin:       delegate.LastLogin.Format(time.RFC3339),
		ProfileImageUrl: uploadUrl,
		CourseTakerUUID: &delegate.CourseTakerUUID,
	}
}

func (u *usersAppImpl) ManagerToUser(manager models.Manager) gentypes.User {
	genManager := u.managerToGentype(manager)
	return gentypes.User{
		Type:            gentypes.ManagerType,
		Email:           &genManager.Email,
		FirstName:       genManager.FirstName,
		LastName:        genManager.LastName,
		Telephone:       &genManager.Telephone,
		JobTitle:        &genManager.JobTitle,
		LastLogin:       genManager.LastLogin,
		ProfileImageUrl: genManager.ProfileImageURL,
	}
}

func (u *usersAppImpl) IndividualToUser(individual models.Individual) gentypes.User {
	return gentypes.User{
		Type:            gentypes.IndividualType,
		Email:           &individual.Email,
		FirstName:       individual.FirstName,
		LastName:        individual.LastName,
		Telephone:       individual.Telephone,
		JobTitle:        individual.JobTitle,
		LastLogin:       individual.LastLogin.String(),
		CourseTakerUUID: &individual.CourseTakerUUID,
	}
}

func (u *usersAppImpl) GetCurrentUser() (gentypes.User, error) {
	if u.grant.IsDelegate {
		delegate, err := u.usersRepository.Delegate(u.grant.Claims.UUID)
		return u.DelegateToUser(delegate), err
	}

	if u.grant.IsManager {
		manager, err := u.usersRepository.Manager(u.grant.Claims.UUID)
		return u.ManagerToUser(manager), err
	}

	if u.grant.IsIndividual {
		individual, err := u.usersRepository.Individual(u.grant.Claims.UUID)
		return u.IndividualToUser(individual), err
	}

	return gentypes.User{}, &errors.ErrUnauthorized
}

// ActiveCourses gets the activeCourses for a course taker
func (u *usersAppImpl) ActiveCourses(takerUUID gentypes.UUID) ([]gentypes.ActiveCourse, error) {
	var authorized = false

	switch {
	case u.grant.IsAdmin:
		authorized = true
	case u.grant.IsManager:
		ok, err := u.usersRepository.CompanyManagesCourseTakers(u.grant.Claims.Company, []gentypes.UUID{takerUUID})
		if err != nil {
			u.grant.Logger.Log(sentry.LevelError, err, "ActiveCourses error checking if company manages takers")
			return []gentypes.ActiveCourse{}, &errors.ErrUnauthorized
		}

		if ok {
			authorized = true
		}
	case u.grant.IsDelegate:
		delegate, _ := u.usersRepository.Delegate(u.grant.Claims.UUID)
		if delegate.CourseTakerUUID == takerUUID {
			authorized = true
		}
	case u.grant.IsIndividual:
		individual, _ := u.usersRepository.Individual(u.grant.Claims.UUID)
		if individual.CourseTakerUUID == takerUUID {
			authorized = true
		}
	}

	if authorized {
		activeCourses, err := u.usersRepository.TakerActiveCourses(takerUUID)
		if err != nil {
			return []gentypes.ActiveCourse{}, &errors.ErrWhileHandling
		}

		return activeCoursesToGentypes(activeCourses), nil
	}

	return []gentypes.ActiveCourse{}, &errors.ErrUnauthorized
}

func (u *usersAppImpl) MyActiveCourses() ([]gentypes.ActiveCourse, error) {
	if !u.grant.IsDelegate && !u.grant.IsIndividual {
		return []gentypes.ActiveCourse{}, &errors.ErrUnauthorized
	}

	var takerUUID gentypes.UUID
	if u.grant.IsDelegate {
		delegate, _ := u.usersRepository.Delegate(u.grant.Claims.UUID)
		takerUUID = delegate.CourseTakerUUID
	}

	if u.grant.IsIndividual {
		individual, _ := u.usersRepository.Individual(u.grant.Claims.UUID)
		takerUUID = individual.CourseTakerUUID
	}

	if (takerUUID == gentypes.UUID{}) {
		return []gentypes.ActiveCourse{}, &errors.ErrNotFound
	}

	activeCourses, err := u.usersRepository.TakerActiveCourses(takerUUID)
	if err != nil {
		return []gentypes.ActiveCourse{}, &errors.ErrWhileHandling
	}

	return activeCoursesToGentypes(activeCourses), nil
}
