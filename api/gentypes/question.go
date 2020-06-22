package gentypes

import (
	"github.com/asaskevich/govalidator"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Question struct {
	UUID             UUID
	Text             string
	RandomiseAnswers *bool
	QuestionType     QuestionType
}

type QuestionType string

// You can't change this text in prod without also
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

type UpdateQuestionInput struct {
	
}