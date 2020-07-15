package course

import (
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

func (c *courseAppImpl) TestsByUUIDs(uuids []gentypes.UUID) ([]gentypes.Test, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Test{}, &errors.ErrUnauthorized
	}

	tests, err := c.coursesRepository.TestsByUUIDs(uuids)
	return c.testsToGentypes(tests), err
}

func (c *courseAppImpl) takerHasSubmittedTest(courseTaker gentypes.UUID, courseID uint, testUUID gentypes.UUID) (bool, error) {
	marks, err := c.usersRepository.TakerTestMarks(courseTaker, courseID)
	if err != nil {
		return false, &errors.ErrWhileHandling
	}

	for _, mark := range marks {
		if mark.TestUUID == testUUID {
			return true, nil
		}
	}
	return false, nil
}

// SubmitTest verifies answers given to a test
func (c *courseAppImpl) SubmitTest(input gentypes.SubmitTestInput) (bool, error) {
	// Only course takers can submit a test
	if !c.grant.IsDelegate && !c.grant.IsIndividual {
		return false, &errors.ErrUnauthorized
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

	success, _ := c.usersRepository.TakerHasActiveCourse(courseTakerUUID, input.CourseID)
	if !success {
		return false, &errors.ErrUnauthorized
	}

	// Check this test's answers haven't already been submitted by this user
	if taken, _ := c.takerHasSubmittedTest(courseTakerUUID, input.CourseID, input.TestUUID); taken {
		return false, &errors.ErrAlreadyTakenTest
	}

	//TODO: Check this test is part of this course

	test, err := c.coursesRepository.Test(input.TestUUID)
	if err != nil {
		return false, &errors.ErrWhileHandling
	}

	// Check enough answers given to complete test
	questions, err := c.coursesRepository.TestQuestions(input.TestUUID)
	if err != nil {
		return false, &errors.ErrWhileHandling
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
		return false, &errors.ErrNotEnoughAnswersGiven
	}

	// Check how many questions the taker got right

	// Get answers for each question
	var acceptedQuestionUUIDs []gentypes.UUID
	for key := range acceptedQuestions {
		acceptedQuestionUUIDs = append(acceptedQuestionUUIDs, key)
	}

	questionsToAnswers, err := c.coursesRepository.ManyAnswers(acceptedQuestionUUIDs)
	if err != nil {
		return false, &errors.ErrWhileHandling
	}

	var correct uint = 0
	for questionUUID, inputAnswer := range acceptedQuestions {
		for _, answer := range questionsToAnswers[questionUUID] {
			if inputAnswer.AnswerUUID == answer.UUID && answer.IsCorrect {
				correct = correct + 1
			}
		}
	}

	// Save marks into DB
	marks := models.TestMark{
		TestUUID:        test.UUID,
		CourseTakerUUID: courseTakerUUID,
		CourseID:        input.CourseID,
		NumCorrect:      correct,
		Total:           test.QuestionsToAnswer,
	}
	err = c.coursesRepository.CreateTestMarks(marks)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to save marks for test")
		return false, &errors.ErrWhileHandling
	}

	// TODO: Check if taker has completed course
	return true, nil
}

// course, err := c.coursesRepository.OnlineCourse(input.CourseID)
// if err != nil {
// 	if err == &errors.ErrNotFound {
// 		return false, &errors.ErrNotFound
// 	}
// 	return false, &errors.ErrWhileHandling
// }

// tests, err := c.coursesRepository.CourseTests(course.UUID)
// if err != nil {
// 	return false, &errors.ErrWhileHandling
// }

func (c *courseAppImpl) DeleteTest(input gentypes.DeleteTestInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	return c.coursesRepository.DeleteTest(input.UUID)
}
