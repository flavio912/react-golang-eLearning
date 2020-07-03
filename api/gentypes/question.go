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

type QuestionFilter struct {
	UUID *UUID
	Text *string
	Tags *[]UUID
}

func (q QuestionFilter) Validate() error {
	ok, err := govalidator.ValidateStruct(q)
	if !ok {
		return err
	}

	return nil
}

type QuestionType string

// You can't change this text in prod without also
// updating the text saved in the DB - @temmerson
const (
	SingleAnswerType QuestionType = "singleAnswer"
)

type Answer struct {
	UUID      UUID
	IsCorrect *bool
	Text      *string
	ImageURL  *string
}

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

type UpdateBasicAnswerInput struct {
	UUID       *UUID
	Text       *string
	ImageToken *string
	IsCorrect  *bool
}

func (u UpdateBasicAnswerInput) Validate() error {
	if u.UUID == nil && u.IsCorrect == nil {
		return errors.ErrInputValidation("IsCorrect", "IsCorrect must be given if no UUID")
	}

	if u.ImageToken != nil {
		if !govalidator.IsBase64(*u.ImageToken) {
			return errors.ErrInputValidation("ImageToken", "Not JWT string")
		}
	}
	if u.UUID == nil && u.ImageToken == nil && u.Text == nil {
		return errors.ErrInputValidation("ImageToken, Text", "Answer type must have Text or ImageToken of no UUID given")
	}

	return nil
}

type UpdateQuestionInput struct {
	UUID             UUID
	Text             *string
	RandomiseAnswers *bool
	QuestionType     *QuestionType
	Tags             *[]UUID
	Answers          *[]UpdateBasicAnswerInput
}

func (u UpdateQuestionInput) Validate() error {
	if u.Answers != nil {
		for _, ans := range *u.Answers {
			if err := ans.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

type DeleteQuestionInput struct {
	UUID UUID `valid:"required"`
}

func (d *DeleteQuestionInput) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}
