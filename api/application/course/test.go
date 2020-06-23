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
			UUID:                 test.UUID,
			Name:                 test.Name,
			AttemptsAllowed:      test.AttemptsAllowed,
			PassPercentage:       &test.PassPercentage,
			QuestionsToAnswer: &test.QuestionsToAnswer,
			RandomiseAnswers:     &test.RandomiseAnswers,
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
		Name:                 input.Name,
		AttemptsAllowed:      input.AttemptsAllowed,
		PassPercentage:       input.PassPercentage,
		QuestionsToAnswer: input.QuestionsToAnswer,
		RandomiseAnswers:     input.RandomiseAnswers,
		Questions:            input.Questions,
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
