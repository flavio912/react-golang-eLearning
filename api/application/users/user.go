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

func activeCourseToMyCourse(activeCourse models.ActiveCourse, upTo *gentypes.UUID, progress *gentypes.Progress) gentypes.MyCourse {
	return gentypes.MyCourse{
		CourseID:       activeCourse.CourseID,
		MinutesTracked: activeCourse.MinutesTracked,
		Status:         gentypes.CourseIncomplete,
		CreatedAt:      activeCourse.CreatedAt.Format(time.RFC3339),
		UpTo:           upTo,
		Progress:       progress,
	}
}

func historicalCourseToMyCourse(historicalCourse models.HistoricalCourse) gentypes.MyCourse {
	status := gentypes.CourseFailed
	if historicalCourse.Passed {
		status = gentypes.CourseComplete
	}
	return gentypes.MyCourse{
		CourseID:       historicalCourse.CourseID,
		MinutesTracked: historicalCourse.MinutesTracked,
		Status:         status,
	}
}

func (u *usersAppImpl) userCoursesToGentypes(activeCourses []models.ActiveCourse, historicalCourses []models.HistoricalCourse) []gentypes.MyCourse {
	var genCourses = make([]gentypes.MyCourse, len(activeCourses)+len(historicalCourses))
	for i, course := range activeCourses {
		progress, lastTest, _ := u.activeCourseProgress(course)
		genCourses[i] = activeCourseToMyCourse(course, lastTest, progress)
	}
	for i, course := range historicalCourses {
		genCourses[len(activeCourses)+i] = historicalCourseToMyCourse(course)
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

// TakerCourses gets the courses for a course taker
func (u *usersAppImpl) TakerCourses(takerUUID gentypes.UUID, showHistorical bool) ([]gentypes.MyCourse, error) {
	var authorized = false

	switch {
	case u.grant.IsAdmin:
		authorized = true
	case u.grant.IsManager:
		ok, err := u.usersRepository.CompanyManagesCourseTakers(u.grant.Claims.Company, []gentypes.UUID{takerUUID})
		if err != nil {
			u.grant.Logger.Log(sentry.LevelError, err, "MyCourses error checking if company manages takers")
			return []gentypes.MyCourse{}, &errors.ErrUnauthorized
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
			return []gentypes.MyCourse{}, &errors.ErrWhileHandling
		}

		if !showHistorical {
			return u.userCoursesToGentypes(activeCourses, []models.HistoricalCourse{}), nil
		}

		historicalCourses, err := u.usersRepository.TakerHistoricalCourses(takerUUID)
		if err != nil || err == &errors.ErrNotFound {
			return []gentypes.MyCourse{}, &errors.ErrWhileHandling
		}

		return u.userCoursesToGentypes(activeCourses, historicalCourses), nil
	}

	return []gentypes.MyCourse{}, &errors.ErrUnauthorized
}

// TakerCourse gets an active course from the courseID
func (u *usersAppImpl) TakerCourse(takerUUID gentypes.UUID, courseID uint) (gentypes.MyCourse, error) {
	if !(u.grant.IsDelegate || u.grant.IsIndividual) {
		return gentypes.MyCourse{}, &errors.ErrUnauthorized
	}
	//TODO: Auth checks

	activeCourse, err := u.usersRepository.TakerActiveCourse(takerUUID, courseID)
	if err != nil {
		return gentypes.MyCourse{}, &errors.ErrWhileHandling
	}

	progress, latestTest, _ := u.activeCourseProgress(activeCourse)

	return activeCourseToMyCourse(activeCourse, latestTest, progress), nil
}

func (u *usersAppImpl) activeCourseProgress(activeCourse models.ActiveCourse) (*gentypes.Progress, *gentypes.UUID, error) {

	onlineCourse, err := u.coursesRepository.OnlineCourse(activeCourse.CourseID)
	if err != nil {
		if err == &errors.ErrNotFound {
			u.grant.Logger.Log(sentry.LevelInfo, err, "activeCourseProgress: onlineCourse not found")
			return nil, nil, nil
		}
		u.grant.Logger.Log(sentry.LevelError, err, "activeCourseProgress: Unable to get online course")
		return nil, nil, &errors.ErrWhileHandling
	}

	testUUIDs, err := u.coursesRepository.CourseTestUUIDs(onlineCourse.UUID)
	if err != nil {
		u.grant.Logger.Log(sentry.LevelWarning, err, "activeCourseProgress: Unable to get course tests")
		return nil, nil, err
	}

	// Get the last tests taken by finding the most recently taken testMark
	var latestTest *gentypes.UUID
	var latestDate = time.Time{}
	marks, err := u.usersRepository.TakerTestMarks(activeCourse.CourseTakerUUID, activeCourse.CourseID) // TODO: Inefficiency
	if err != nil {
		u.grant.Logger.Log(sentry.LevelWarning, err, "TakerCourse: Unable to get test marks")
	} else {
		for _, mark := range marks {
			if mark.CreatedAt.After(latestDate) && mark.Passed {
				latestDate = mark.CreatedAt
				latestTest = &mark.TestUUID
			}
		}
	}

	if latestTest == nil {
		return &gentypes.Progress{
			Total:     len(testUUIDs),
			Completed: 0,
		}, nil, nil
	}

	for i, uuid := range testUUIDs {
		if uuid == *latestTest {
			return &gentypes.Progress{
				Total:     len(testUUIDs),
				Completed: i,
			}, latestTest, nil
		}
	}

	u.grant.Logger.LogMessage(sentry.LevelError, "Unable to get testUUID for user")
	return nil, latestTest, &errors.ErrWhileHandling
}
