package gentypes

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Test struct {
	UUID                 UUID
	Name                 string
	AttemptsAllowed      *int
	PassPercentage       *float32
	MinQuestionsToAnswer *int
	RandomiseAnswers     *bool
}

type CreateTestInput struct {
	Name                 string
	AttemptsAllowed      *int
	PassPercentage       float32
	MinQuestionsToAnswer int
	RandomiseAnswers     bool
	Questions            []UUID
}

func (c CreateTestInput) Validate() error {
	if c.AttemptsAllowed != nil {
		if *c.AttemptsAllowed <= 0 {
			return errors.ErrInputValidation("AttemptsAllowed", "Not greater than 0")
		}
	}

	if c.MinQuestionsToAnswer <= 0 {
		return errors.ErrInputValidation("MinQuestionsToAnswer", "Not greater than 0")
	}

	return nil
}
