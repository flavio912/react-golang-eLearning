package gentypes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

type Test struct {
	UUID              UUID
	CreatedAt         time.Time
	Name              string
	AttemptsAllowed   *uint
	PassPercentage    *float64
	QuestionsToAnswer *uint
	RandomiseAnswers  *bool
}

type TestFilter struct {
	UUID *UUID
	Name *string
}

type CreateTestInput struct {
	Name              string
	Tags              *[]UUID
	AttemptsAllowed   int32
	PassPercentage    float64
	QuestionsToAnswer int32
	RandomiseAnswers  bool
	Questions         []UUID
}

func (c CreateTestInput) Validate() error {
	if c.AttemptsAllowed <= 0 {
		return errors.ErrInputValidation("AttemptsAllowed", "Not greater than 0")
	}

	if c.QuestionsToAnswer <= 0 {
		return errors.ErrInputValidation("QuestionsToAnswer", "Not greater than 0")
	}

	return nil
}

type UpdateTestInput struct {
	UUID              UUID
	Name              *string
	AttemptsAllowed   *int32
	PassPercentage    *float64
	QuestionsToAnswer *int32
	RandomiseAnswers  *bool
	Tags              *[]UUID
	Questions         *[]UUID
}

func (c UpdateTestInput) Validate() error {
	if c.AttemptsAllowed != nil && *c.AttemptsAllowed <= 0 {
		return errors.ErrInputValidation("AttemptsAllowed", "Not greater than 0")
	}

	if c.QuestionsToAnswer != nil && *c.QuestionsToAnswer <= 0 {
		return errors.ErrInputValidation("QuestionsToAnswer", "Not greater than 0")
	}

	return nil
}

type QuestionAnswer struct {
	QuestionUUID UUID
	AnswerUUID   UUID
}

type SubmitTestInput struct {
	CourseID int32
	TestUUID UUID
	Answers  []QuestionAnswer
}

func (s SubmitTestInput) Validate() error {
	if s.CourseID < 0 {
		return errors.ErrInputValidation("CourseID", "must be greater than 0")
	}
	return nil
}

type SubmitTestPayload struct {
	CourseStatus CourseStatus
	Passed       bool
}

type DeleteTestInput struct {
	UUID UUID `valid:"required"`
}

func (d *DeleteTestInput) Validate() error {
	_, err := govalidator.ValidateStruct(d)
	return err
}
