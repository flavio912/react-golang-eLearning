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
	if !c.grant.IsAdmin {
		return gentypes.Question{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Question{}, err
	}

	// Validate image tokens if given
	var ans *[]course.UpdateAnswerArgs
	if input.Answers != nil {
		answers := []course.UpdateAnswerArgs{}
		for _, ans := range *input.Answers {
			var key *string
			if ans.ImageToken != nil {
				imgKey, err := uploads.VerifyUploadSuccess(*ans.ImageToken, "questionImages")
				if err != nil {
					return gentypes.Question{}, err
				}

				key = &imgKey
			}

			answers = append(answers, course.UpdateAnswerArgs{
				UUID:      ans.UUID,
				IsCorrect: ans.IsCorrect,
				Text:      ans.Text,
				ImageKey:  key,
			})

		}

		ans = &answers
	}

	updateArgs := course.UpdateQuestionArgs{
		UUID:             input.UUID,
		Text:             input.Text,
		RandomiseAnswers: input.RandomiseAnswers,
		QuestionType:     input.QuestionType,
		Answers:          ans,
		Tags:             input.Tags,
	}

	question, err := c.coursesRepository.UpdateQuestion(updateArgs)
	return c.questionToGentype(question), err
}

// func (c *courseAppImpl) Question(uuid gentypes.UUID) (gentypes.Question, err) {
// 	// Check user is assigned course with this question in
// }
