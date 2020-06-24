package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *courseAppImpl) testToGentype(test models.Test) gentypes.Test {
	if c.grant.IsAdmin {
		return gentypes.Test{
			UUID:              test.UUID,
			Name:              test.Name,
			AttemptsAllowed:   test.AttemptsAllowed,
			PassPercentage:    &test.PassPercentage,
			QuestionsToAnswer: &test.QuestionsToAnswer,
			RandomiseAnswers:  &test.RandomiseAnswers,
		}
	}
	return gentypes.Test{}
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
		AttemptsAllowed:   input.AttemptsAllowed,
		PassPercentage:    input.PassPercentage,
		QuestionsToAnswer: input.QuestionsToAnswer,
		RandomiseAnswers:  input.RandomiseAnswers,
		Questions:         input.Questions,
	}

	test, err := c.coursesRepository.CreateTest(createTest)

	if err != nil {
		return gentypes.Test{}, err
	}

	return c.testToGentype(test), nil
}

// Test gets a test from the db and applies applicable access control
func (c *courseAppImpl) Test(testUUID gentypes.UUID) (gentypes.Test, error) {
	// TODO; Should be allowed if user is assigned the course (but only get specific info)
	if !c.grant.IsAdmin {
		return gentypes.Test{}, &errors.ErrUnauthorized
	}

	return gentypes.Test{
		UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
		Name: "Cheseecake",
	}, nil
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
		delegate, err := c.usersRepository.Delegate(c.grant.Claims.UUID)
		courseTakerUUID = delegate.CourseTakerUUID
	}
	if c.grant.IsIndividual {
		individual, err := c.usersRepository.Individual(c.grant.Claims.UUID)
		courseTakerUUID = individual.CourseTakerUUID
	}

	success, _ := c.usersRepository.TakerHasActiveCourse(courseTakerUUID, input.CourseID)
	if !success {
		return false, &errors.ErrUnauthorized
	}

	// Check this test's answers haven't already been submitted by this user
	if taken, _ := c.takerHasSubmittedTest(courseTakerUUID, input.CourseID, input.TestUUID); !taken {
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

	var acceptedQuestions map[gentypes.UUID]gentypes.QuestionAnswer
	for _, answer := range input.Answers {
		for _, question := range questions {
			if answer.QuestionUUID == question.UUID {
				acceptedQuestions[question.UUID] = answer
			}
		}
	}

	if len(acceptedQuestions) < test.QuestionsToAnswer {
		return false, &errors.ErrNotEnoughAnswersGiven
	}

	// Check how many questions the taker got right

	// Get answers for each question
	var acceptedQuestionUUIDs []gentypes.UUID
	for key, _ := range acceptedQuestions {
		acceptedQuestionUUIDs = append(acceptedQuestionUUIDs, key)
	}

	answers, err := c.coursesRepository.ManyAnswers(acceptedQuestionUUIDs)

	for key, answer := range answers {

	}
	// Count how many correct

	// Save marks into DB
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
