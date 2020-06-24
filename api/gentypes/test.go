package gentypes

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Test struct {
	UUID              UUID
	Name              string
	AttemptsAllowed   *int
	PassPercentage    *float32
	QuestionsToAnswer *int
	RandomiseAnswers  *bool
}

type CreateTestInput struct {
	Name              string
	AttemptsAllowed   *int
	PassPercentage    float32
	QuestionsToAnswer int
	RandomiseAnswers  bool
	Questions         []UUID
}

func (c CreateTestInput) Validate() error {
	if c.AttemptsAllowed != nil {
		if *c.AttemptsAllowed <= 0 {
			return errors.ErrInputValidation("AttemptsAllowed", "Not greater than 0")
		}
	}

	if c.QuestionsToAnswer <= 0 {
		return errors.ErrInputValidation("QuestionsToAnswer", "Not greater than 0")
	}

	return nil
}

type QuestionAnswer struct {
	QuestionUUID UUID
	AnswerUUID   UUID
}

type SubmitTestInput struct {
	CourseID uint
	TestUUID UUID
	Answers  []QuestionAnswer
}
