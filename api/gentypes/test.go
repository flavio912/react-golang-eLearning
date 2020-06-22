package gentypes

import (
	"github.com/asaskevich/govalidator"
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

type QuestionType string

// You can't change the actual text in prod without also
// updating the text saved in the DB - @temmerson
const (
	SingleAnswerType QuestionType = "singleAnswer"
)

type CreateBasicAnswerInput struct {
	Text       *string
	ImageToken *string
	IsCorrect  bool
}

func (c CreateBasicAnswerInput) Validate() error {
	if c.ImageToken != nil {
		if !govalidator.IsBase64(*c.ImageToken) {
			return errors.ErrInputValidation("ImageToken", "Not JWT string")
		}
	}
	if c.ImageToken == nil && c.Text == nil {
		return errors.ErrInputValidation("ImageToken, Text", "Answer type must have Text or ImageToken")
	}

	return nil
}

type CreateQuestionInput struct {
	Text             string
	RandomiseAnswers bool
	QuestionType     QuestionType
	Tags             []UUID
	Answers          []CreateBasicAnswerInput
}

func (c CreateQuestionInput) Validate() error {
	if len(c.Answers) == 0 {
		return errors.ErrInputValidation("Answers", "Question has no answers")
	}

	for _, ans := range c.Answers {
		if err := ans.Validate(); err != nil {
			return err
		}
	}
	for _, ans := range c.Answers {
		if ans.IsCorrect {
			return nil
		}
	}

	return errors.ErrInputValidation("Answers", "No answer is set as correct")
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
