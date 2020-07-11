package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/dbutils"
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

func filterQuestion(query *gorm.DB, filter *gentypes.QuestionFilter) *gorm.DB {
	if filter != nil {
		if filter.UUID != nil {
			query = query.Where("uuid = ?", *filter.UUID)
		}

		if filter.Text != nil && *filter.Text != "" {
			query = query.Where("text ILIKE ?", "%%"+*filter.Text+"%%")
		}

		// TODO: Filter tags
		// if filter.Tags != nil && *filter.Tags != "" {
		// 	query = query.Where("text ILIKE ?", "%%"+*filter.Text+"%%")
		// }
	}

	return query
}

func (c *coursesRepoImpl) Questions(page *gentypes.Page, filter *gentypes.QuestionFilter, orderBy *gentypes.OrderBy) ([]models.Question, gentypes.PageInfo, error) {
	var questions []models.Question
	utils := dbutils.NewDBUtils(c.Logger)
	pageInfo, err := utils.GetPageOf(
		&models.Question{},
		&questions,
		page,
		orderBy,
		[]string{"created_at", "text", "randomise_answers"},
		"created_at DESC",
		func(db *gorm.DB) *gorm.DB {
			return filterQuestion(db, filter)
		},
	)
	pageInfo.Given = int32(len(questions))

	return questions, pageInfo, err
}

// ManyAnswers gets a mapping between questionUUIDs and their respective answers
func (c *coursesRepoImpl) ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.BasicAnswer, error) {
	var answers []models.BasicAnswer
	query := database.GormDB.Where("question_uuid IN (?)", questionUUIDs).Order("question_uuid, rank ASC").Find(&answers)
	if query.Error != nil && !query.RecordNotFound() {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get many answers")
		return map[gentypes.UUID][]models.BasicAnswer{}, &errors.ErrWhileHandling
	}

	var output = make(map[gentypes.UUID][]models.BasicAnswer)
	for _, answer := range answers {
		output[answer.QuestionUUID] = append(output[answer.QuestionUUID], answer)
	}

	return output, nil
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
	AnswerType gentypes.AnswerType
	UUID       *gentypes.UUID
	IsCorrect  *bool
	Text       *string
	ImageKey   *string
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

				switch ans.AnswerType {
				case gentypes.TextAnswer:
					updates["image_key"] = nil

					if ans.Text == nil {
						tx.Rollback()
						return models.Question{}, errors.ErrInputValidation("Answers", "Text answer has no text")
					}
					updates["text"] = *ans.Text
				case gentypes.ImageAnswer:
					updates["text"] = nil

					if ans.ImageKey == nil {
						break
					}
					updates["image_key"] = *ans.ImageKey
				case gentypes.TextImageAnswer:
					if ans.Text == nil {
						tx.Rollback()
						return models.Question{}, errors.ErrInputValidation("Answers", "Text + Image answer has no text")
					}
					updates["text"] = *ans.Text

					if ans.ImageKey != nil {
						updates["image_key"] = *ans.ImageKey
					}
				default:
					tx.Rollback()
					return models.Question{}, errors.ErrInputValidation("Answers", "Invalid answer type")
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

func (c *coursesRepoImpl) DeleteQuestion(input gentypes.UUID) (bool, error) {
	tx := database.GormDB.Begin()

	var test_question_link models.TestQuestionsLink
	if !tx.Model(&models.TestQuestionsLink{}).Where("question_uuid = ?", input).First(&test_question_link).RecordNotFound() {
		err := errors.ErrUnableToDelete("Cannot delete question that is part of a test")
		c.Logger.Log(sentry.LevelError, err, "Unable to delete question")
		return false, err
	}

	if err := tx.Delete(models.Question{}, "uuid = ?", input).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Unable to delete question: %s", input)
		return false, &errors.ErrDeleteFailed
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return false, &errors.ErrWhileHandling
	}

	return true, nil
}
