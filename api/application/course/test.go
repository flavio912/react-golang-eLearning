package course

import (
	"time"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *courseAppImpl) testToGentype(test models.Test) gentypes.Test {
	if c.grant.IsAdmin {
		return gentypes.Test{
			UUID:              test.UUID,
			Name:              test.Name,
			AttemptsAllowed:   &test.AttemptsAllowed,
			PassPercentage:    &test.PassPercentage,
			QuestionsToAnswer: &test.QuestionsToAnswer,
			RandomiseAnswers:  &test.RandomiseAnswers,
		}
	}
	if c.grant.IsDelegate || c.grant.IsIndividual {
		return gentypes.Test{
			UUID:              test.UUID,
			Name:              test.Name,
			AttemptsAllowed:   &test.AttemptsAllowed,
			QuestionsToAnswer: &test.QuestionsToAnswer,
		}
	}
	return gentypes.Test{}
}

func (c *courseAppImpl) testsToGentypes(tests []models.Test) []gentypes.Test {
	var modTests = make([]gentypes.Test, len(tests))
	for i, test := range tests {
		modTests[i] = c.testToGentype(test)
	}
	return modTests
}

func (c *courseAppImpl) CreateTest(input gentypes.CreateTestInput) (gentypes.Test, error) {
	if !c.grant.IsAdmin {
		return gentypes.Test{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Test{}, err
	}

	// Check all uploaded images

	createTest := course.CreateTestInput{
		Name:              input.Name,
		PassPercentage:    input.PassPercentage,
		RandomiseAnswers:  input.RandomiseAnswers,
		Questions:         input.Questions,
		QuestionsToAnswer: uint(input.QuestionsToAnswer),
		AttemptsAllowed:   uint(input.AttemptsAllowed),
	}

	test, err := c.coursesRepository.CreateTest(createTest)

	if err != nil {
		return gentypes.Test{}, err
	}

	return c.testToGentype(test), nil
}

func (c *courseAppImpl) UpdateTest(input gentypes.UpdateTestInput) (gentypes.Test, error) {
	if !c.grant.IsAdmin {
		return gentypes.Test{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Test{}, err
	}

	updateInput := course.UpdateTestInput{
		UUID:             input.UUID,
		Name:             input.Name,
		PassPercentage:   input.PassPercentage,
		RandomiseAnswers: input.RandomiseAnswers,
		Tags:             input.Tags,
		Questions:        input.Questions,
	}

	if input.QuestionsToAnswer != nil {
		updateInput.QuestionsToAnswer = helpers.UintPointer(uint(*input.QuestionsToAnswer))
	}
	if input.AttemptsAllowed != nil {
		updateInput.AttemptsAllowed = helpers.UintPointer(uint(*input.AttemptsAllowed))
	}

	test, err := c.coursesRepository.UpdateTest(updateInput)
	return c.testToGentype(test), err
}

// Test gets a test from the db and applies applicable access control
func (c *courseAppImpl) Test(testUUID gentypes.UUID) (gentypes.Test, error) {
	// TODO; Should be allowed if user is assigned the course (but only get specific info)
	if !c.grant.IsAdmin {
		return gentypes.Test{}, &errors.ErrUnauthorized
	}

	test, err := c.coursesRepository.Test(testUUID)
	return c.testToGentype(test), err
}

func (c *courseAppImpl) Tests(
	page *gentypes.Page,
	filter *gentypes.TestFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Test, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Test{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	tests, pageInfo, err := c.coursesRepository.Tests(page, filter, orderBy)
	if err != nil {
		return []gentypes.Test{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return c.testsToGentypes(tests), pageInfo, nil
}

func (c *courseAppImpl) grantCanViewSyllabusItems(courseElementUUIDs []gentypes.UUID, elementType gentypes.CourseElement) bool {
	if !c.grant.IsAdmin && !c.grant.IsDelegate && !c.grant.IsIndividual {
		return false
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		// Check user is taking a course with those tests in it
		var courseTakerID gentypes.UUID
		if c.grant.IsDelegate {
			delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
			courseTakerID = delegate.CourseTakerUUID
		}

		if c.grant.IsIndividual {
			individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
			courseTakerID = individual.CourseTakerUUID
		}

		activeCourses, err := c.usersRepository.TakerActiveCourses(courseTakerID)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "grantCanViewSyllabusItems: Unable to get taker active courses")
			return false
		}

		var courseIds = make([]uint, len(activeCourses))
		for i, activeCourse := range activeCourses {
			courseIds[i] = activeCourse.CourseID
		}

		areTestsInCourses, err := c.coursesRepository.AreInCourses(courseIds, courseElementUUIDs, elementType)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "grantCanViewSyllabusItems: AreInCourses error")
			return false
		}

		if !areTestsInCourses {
			return false
		}
	}

	return true
}

func (c *courseAppImpl) TestsByUUIDs(uuids []gentypes.UUID) ([]gentypes.Test, error) {
	if !c.grantCanViewSyllabusItems(uuids, gentypes.TestType) {
		return []gentypes.Test{}, &errors.ErrUnauthorized
	}

	tests, err := c.coursesRepository.TestsByUUIDs(uuids)
	return c.testsToGentypes(tests), err
}

func (c *courseAppImpl) takerTestMark(courseTaker gentypes.UUID, courseID uint, testUUID gentypes.UUID) (models.TestMark, error) {
	marks, err := c.usersRepository.TakerTestMarks(courseTaker, courseID)
	if err != nil {
		return models.TestMark{}, &errors.ErrWhileHandling
	}

	for _, mark := range marks {
		if mark.TestUUID == testUUID {
			return mark, nil
		}
	}

	return models.TestMark{}, &errors.ErrNotFound
}

// SubmitTest verifies answers given to a test
func (c *courseAppImpl) SubmitTest(input gentypes.SubmitTestInput) (bool, gentypes.CourseStatus, error) {
	// Only course takers can submit a test
	if !c.grant.IsDelegate && !c.grant.IsIndividual {
		return false, gentypes.CourseIncomplete, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return false, gentypes.CourseIncomplete, err
	}

	// Check taker can access this course
	var courseTakerUUID gentypes.UUID
	if c.grant.IsDelegate {
		delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
		courseTakerUUID = delegate.CourseTakerUUID
	}
	if c.grant.IsIndividual {
		individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
		courseTakerUUID = individual.CourseTakerUUID
	}

	courseID := uint(input.CourseID)

	activeCourse, err := c.usersRepository.TakerActiveCourse(courseTakerUUID, courseID)
	if err != nil {
		return false, gentypes.CourseIncomplete, &errors.ErrUnauthorized
	}

	//TODO: Check this test is part of this course

	test, err := c.coursesRepository.Test(input.TestUUID)
	if err != nil {
		return false, activeCourse.Status, &errors.ErrWhileHandling
	}

	// Check enough answers given to complete test
	questions, err := c.coursesRepository.TestQuestions(input.TestUUID)
	if err != nil {
		return false, activeCourse.Status, &errors.ErrWhileHandling
	}

	var acceptedQuestions = make(map[gentypes.UUID]gentypes.QuestionAnswer)
	for _, answer := range input.Answers {
		for _, question := range questions {
			if answer.QuestionUUID == question.UUID {
				acceptedQuestions[question.UUID] = answer
			}
		}
	}

	if len(acceptedQuestions) < int(test.QuestionsToAnswer) {
		return false, activeCourse.Status, &errors.ErrNotEnoughAnswersGiven
	}

	// Check how many questions the taker got right

	// Get answers for each question
	var acceptedQuestionUUIDs []gentypes.UUID
	for key := range acceptedQuestions {
		acceptedQuestionUUIDs = append(acceptedQuestionUUIDs, key)
	}

	questionsToAnswers, err := c.coursesRepository.ManyAnswers(acceptedQuestionUUIDs)
	if err != nil {
		return false, activeCourse.Status, &errors.ErrWhileHandling
	}

	var correct uint = 0
	for questionUUID, inputAnswer := range acceptedQuestions {
		for _, answer := range questionsToAnswers[questionUUID] {
			if inputAnswer.AnswerUUID == answer.UUID && answer.IsCorrect {
				correct = correct + 1
			}
		}
	}

	// Check if the user has passed the test or not
	percentCorrect := (float64(correct) / float64(test.QuestionsToAnswer)) * 100
	testPassed := percentCorrect > test.PassPercentage

	prevMark, err := c.takerTestMark(courseTakerUUID, courseID, input.TestUUID)

	if err != nil && err != &errors.ErrNotFound {
		c.grant.Logger.Log(sentry.LevelError, err, "SubmitTest: Unable to get taker test mark")
		return false, activeCourse.Status, &errors.ErrWhileHandling
	}

	currentAttempt := uint(1)
	if err != &errors.ErrNotFound {
		currentAttempt = prevMark.CurrentAttempt + 1
	}

	// If this is the last attempt check if it passed
	if test.AttemptsAllowed == 1 || (err != &errors.ErrNotFound && test.AttemptsAllowed == prevMark.CurrentAttempt) {
		if !testPassed {
			err := c.completeCourse(courseTakerUUID, courseID, activeCourse.MinutesTracked, false)
			if err != nil {
				c.grant.Logger.Log(sentry.LevelError, err, "Unable to complete course - fail")
			}
			return false, gentypes.CourseFailed, nil
		}
	}

	// Save marks into DB
	marks := models.TestMark{
		TestUUID:        test.UUID,
		CourseTakerUUID: courseTakerUUID,
		CourseID:        courseID,
		Passed:          testPassed,
		CurrentAttempt:  currentAttempt,
	}

	err = c.usersRepository.SaveTestMarks(marks)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to save marks for test")
		return false, activeCourse.Status, &errors.ErrWhileHandling
	}

	// If taker has completed the whole course
	completed, err := c.isOnlineCourseCompleted(courseTakerUUID, courseID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelFatal, err, "SubmitTest: UNABLE TO CHECK COURSE COMPLETION")
		return true, activeCourse.Status, &errors.ErrWhileHandling
	}

	if completed {
		err := c.completeCourse(courseTakerUUID, courseID, activeCourse.MinutesTracked, true)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "Unable to complete course - success")
		}
		return true, gentypes.CourseComplete, nil
	}

	return true, gentypes.CourseIncomplete, nil
}

// completeOnlineCourse completes a course and moves it to historical + generates a certificate if the user passed
func (c *courseAppImpl) completeCourse(takerUUID gentypes.UUID, courseID uint, minutesTracked float64, passed bool) error {
	// Calculate expiration date
	course, err := c.coursesRepository.Course(courseID)
	if err != nil {
		return err
	}

	var expirationDate *time.Time
	if passed {
		expDate := time.Now()
		if course.ExpirationToEndMonth {
			currentYear, currentMonth, _ := expDate.Date()
			currentLocation := expDate.Location()

			firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
			expDate = firstOfMonth.AddDate(0, 1, -1)
		}

		expDate = expDate.AddDate(0, int(course.ExpiresInMonths), 0)
		expirationDate = &expDate
	}

	// Create historical course
	histCourse, err := c.usersRepository.CreateHistoricalCourse(models.HistoricalCourse{
		CourseTakerUUID: takerUUID,
		CourseID:        courseID,
		MinutesTracked:  minutesTracked,
		Passed:          passed,
		CertificateKey:  nil, // No certificate initially while it is generated
		ExpirationDate:  expirationDate,
	})

	if err != nil {
		return err
	}

	// Generate activity
	activityType := gentypes.ActivityFailed
	if passed {
		activityType = gentypes.ActivityCompleted
	}
	_, err = c.usersRepository.CreateTakerActivity(takerUUID, activityType, &courseID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelWarning, err, "completeCourse: Unable to create activity")
	}

	if passed {
		go c.generateCertificate(histCourse.UUID)
	}

	return nil
}

func (c *courseAppImpl) isOnlineCourseCompleted(takerUUID gentypes.UUID, courseID uint) (bool, error) {
	onlineCourse, err := c.coursesRepository.OnlineCourse(courseID)
	if err != nil {
		if err == &errors.ErrNotFound {
			return false, &errors.ErrNotFound
		}
		return false, &errors.ErrWhileHandling
	}

	tests, err := c.coursesRepository.CourseTests(onlineCourse.UUID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to get course tests")
		return false, &errors.ErrWhileHandling
	}

	// Get all test marks for a course
	marks, err := c.usersRepository.TakerTestMarks(takerUUID, courseID)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to get test taker marks")
		return false, &errors.ErrWhileHandling
	}

	// Check Taker has taken enough tests to complete course
	if len(marks) < len(tests) {
		return false, nil
	}

	// Check all tests have been passed
	for _, mark := range marks {
		if !mark.Passed {
			return false, nil
		}
	}

	return true, nil
}

func (c *courseAppImpl) DeleteTest(input gentypes.DeleteTestInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	return c.coursesRepository.DeleteTest(input.UUID)
}
