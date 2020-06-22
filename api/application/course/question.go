package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) questionToGentype(question models.Question) gentypes.Question {
	if c.grant.IsAdmin {
		return gentypes.Question{
			UUID:             question.UUID,
			Text:             question.Text,
			RandomiseAnswers: &question.RandomiseAnswers,
			QuestionType:     question.QuestionType,
		}
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		return gentypes.Question{
			UUID:         question.UUID,
			Text:         question.Text,
			QuestionType: question.QuestionType,
		}
	}

	return gentypes.Question{}
}

func (c *courseAppImpl) CreateQuestion(input gentypes.CreateQuestionInput) (gentypes.Question, error) {
	if !c.grant.IsAdmin {
		return gentypes.Question{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Question{}, err
	}

	var answerArgs []course.CreateAnswerArgs
	for _, answer := range input.Answers {
		var imageKey *string
		if answer.ImageToken != nil {
			key, err := uploads.VerifyUploadSuccess(*answer.ImageToken, "answerImages")
			if err != nil {
				return gentypes.Question{}, err
			}
			imageKey = &key
		}
		answerArgs = append(answerArgs, course.CreateAnswerArgs{
			Text:      answer.Text,
			ImageKey:  imageKey,
			IsCorrect: answer.IsCorrect,
		})
	}

	question, err := c.coursesRepository.CreateQuestion(course.CreateQuestionArgs{
		Text:             input.Text,
		RandomiseAnswers: input.RandomiseAnswers,
		QuestionType:     input.QuestionType,
		Answers:          answerArgs,
		Tags:             input.Tags,
	})
	return c.questionToGentype(question), err
}

func (c *courseAppImpl) UpdateQuestion(input gentypes.UpdateQuestionInput) (gentypes.Question, error) {
	return gentypes.Question{}, nil
}

// func (c *courseAppImpl) Question(uuid gentypes.UUID) (gentypes.Question, err) {
// 	// Check user is assigned course with this question in
// }
