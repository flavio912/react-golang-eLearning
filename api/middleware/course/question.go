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

type CreateAnswerArgs struct {
	IsCorrect bool
	Text      *string
	ImageKey  *string
}

type CreateQuestionArgs struct {
	Text             string
	RandomiseAnswers bool
	QuestionType     gentypes.QuestionType
	Answers          []CreateAnswerArgs
	Tags             []gentypes.UUID
}

func answerArgsToModels(answers []CreateAnswerArgs) []models.BasicAnswer {
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

	answers := []models.BasicAnswer{}
	for i, answer := range input.Answers {
		answers = append(answers, models.BasicAnswer{
			IsCorrect: answer.IsCorrect,
			Text:      answer.Text,
			ImageKey:  answer.ImageKey,
			Rank:      strconv.Itoa(i),
		})
	}

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

type UpdateAnswerArgs struct {
	UUID      *gentypes.UUID
	IsCorrect *bool
	Text      *string
	ImageKey  *string
}

type UpdateQuestionArgs struct {
	UUID             gentypes.UUID
	Text             *string
	RandomiseAnswers *bool
	QuestionType     *gentypes.QuestionType
	Answers          *[]UpdateAnswerArgs
	Tags             *[]gentypes.UUID
}

func (u UpdateQuestionArgs) Validate() error {
	if u.Answers != nil {
		for _, ans := range *u.Answers {
			if ans.UUID == nil && ans.IsCorrect == nil {
				return errors.ErrInputValidation("Answers", "Cannot have undefined uuid and isCorrect")
			}
		}
	}
	return nil
}

func (c *coursesRepoImpl) UpdateQuestion(input UpdateQuestionArgs) (models.Question, error) {
	if err := input.Validate(); err != nil {
		return models.Question{}, err
	}

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

	if input.Answers != nil {
		// Get all current answers
		var currentAnswers []models.BasicAnswer
		if err := tx.Model(question).Association("Answers").Find(&currentAnswers).Error; err != nil {
			c.Logger.Log(sentry.LevelWarning, err, "Unable to get current answers")
			tx.Rollback()
			return models.Question{}, &errors.ErrWhileHandling
		}

		// We need to remove the deleted answers manually from the DB
		var toDelete []gentypes.UUID
		for _, currentAns := range currentAnswers {
			var found = false
			for _, updatedAns := range *input.Answers {
				if updatedAns.UUID != nil && currentAns.UUID == *updatedAns.UUID {
					found = true
					break
				}
			}
			if !found {
				toDelete = append(toDelete, currentAns.UUID)
			}
		}

		var updatedAnswers []models.BasicAnswer
		for i, ans := range *input.Answers {
			if ans.UUID != nil {
				updates := map[string]interface{}{
					"rank": strconv.Itoa(i),
				}
				if ans.Text != nil {
					updates["text"] = *ans.Text
				}
				if ans.ImageKey != nil {
					updates["image_key"] = *ans.ImageKey
				}
				if ans.IsCorrect != nil {
					updates["is_correct"] = *ans.IsCorrect
				}
				tx.Model(&models.BasicAnswer{}).Where("uuid = ?", *ans.UUID).Updates(updates)
				updatedAnswers = append(updatedAnswers, models.BasicAnswer{UUID: *ans.UUID})
			} else {
				newAns := models.BasicAnswer{
					Text:      ans.Text,
					ImageKey:  ans.ImageKey,
					IsCorrect: *ans.IsCorrect, // Checked in validation
					Rank:      strconv.Itoa(i),
				}
				updatedAnswers = append(updatedAnswers, newAns)
			}
		}

		err := tx.Model(&question).Association("Answers").Replace(updatedAnswers).Error
		if err != nil {
			c.Logger.Log(sentry.LevelWarning, err, "Unable replace answer associations")
			tx.Rollback()
			return models.Question{}, &errors.ErrWhileHandling
		}

		// Remove dangling answers
		if err := tx.Where("uuid IN (?)", toDelete).Delete(&models.BasicAnswer{}).Error; err != nil {
			c.Logger.Log(sentry.LevelWarning, err, "Unable delete dangling answers")
			tx.Rollback()
			return models.Question{}, &errors.ErrWhileHandling
		}
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
