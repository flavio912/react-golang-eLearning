package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *coursesRepoImpl) Question(uuid gentypes.UUID) (models.Question, error) {
	var question models.Question
	if err := database.GormDB.Where("uuid = ?", uuid).Find(&question).Error; err != nil {
		c.Logger.Log(sentry.LevelWarning, err, "Unable to get question")
		return models.Question{}, &errors.ErrWhileHandling
	}

	return question, nil
}

type AnswerArgs struct {
	IsCorrect bool
	Text      *string
	ImageKey  *string
}

type CreateQuestionArgs struct {
	Text             string
	RandomiseAnswers bool
	QuestionType     gentypes.QuestionType
	Answers          []AnswerArgs
	Tags             []gentypes.UUID
}

func answerArgsToModels(answers []AnswerArgs) []models.BasicAnswer {
	answerMods := []models.BasicAnswer{}
	for i, answer := range answers {
		answerMods = append(answerMods, models.BasicAnswer{
			IsCorrect: answer.IsCorrect,
			Text:      answer.Text,
			ImageKey:  answer.ImageKey,
			Rank:      strconv.Itoa(i),
		})
	}
	return answerMods
}

func (c *coursesRepoImpl) CreateQuestion(input CreateQuestionArgs) (models.Question, error) {

	tags, err := c.CheckTagsExist(input.Tags)
	if err != nil {
		return models.Question{}, &errors.ErrTagsNotFound
	}

	answers := answerArgsToModels(input.Answers)

	question := models.Question{
		Text:             input.Text,
		RandomiseAnswers: input.RandomiseAnswers,
		QuestionType:     input.QuestionType,
		Answers:          answers,
		Tags:             tags,
	}

	if err := database.GormDB.Create(&question).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create question")
		return models.Question{}, &errors.ErrWhileHandling
	}

	return question, nil
}

type UpdateQuestionArgs struct {
	UUID             gentypes.UUID
	Text             *string
	RandomiseAnswers *bool
	QuestionType     *gentypes.QuestionType
	Answers          *[]AnswerArgs
	Tags             *[]gentypes.UUID
}

func (c *coursesRepoImpl) UpdateQuestion(input UpdateQuestionArgs) (models.Question, error) {
	// Get Question
	question, err := c.Question(input.UUID)
	if err != nil {
		c.Logger.Log(sentry.LevelWarning, err, "Unable to get question")
		return models.Question{}, &errors.ErrNotFound
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete all previous answers and replace with new ones
	if input.Answers != nil {
		err := tx.Where("question_uuid = ?", input.UUID).Delete(&models.BasicAnswer{}).Error
		if err != nil {
			c.Logger.Log(sentry.LevelError, err, "Unable to delete answers")
			tx.Rollback()
			return models.Question{}, &errors.ErrWhileHandling
		}

		answers := answerArgsToModels(*input.Answers)
		tx.Model(&question).Association("Answers").Replace(answers)
	}

	// Update tags
	if input.Tags != nil {
		tags, err := c.CheckTagsExist(*input.Tags)
		if err != nil {
			c.Logger.Log(sentry.LevelError, err, "Unable to get tags")
			return models.Question{}, &errors.ErrTagsNotFound
		}
		tx.Model(&question).Association("Tags").Replace(tags)
	}

	if input.Text != nil && *input.Text != question.Text {
		question.Text = *input.Text
	}
	if input.RandomiseAnswers != nil && *input.RandomiseAnswers != question.RandomiseAnswers {
		question.RandomiseAnswers = *input.RandomiseAnswers
	}
	if input.QuestionType != nil && *input.QuestionType != question.QuestionType {
		question.QuestionType = *input.QuestionType
	}

	if err := tx.Save(&question).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to update question")
		tx.Rollback()
		return models.Question{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit question")
		tx.Rollback()
		return models.Question{}, &errors.ErrWhileHandling
	}

	return question, nil
}
