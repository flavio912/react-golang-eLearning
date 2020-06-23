package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type CreateTestInput struct {
	Name                 string
	AttemptsAllowed      *int
	PassPercentage       float32
	QuestionsToAnswer int
	RandomiseAnswers     bool
	Questions            []gentypes.UUID
}

func (c *coursesRepoImpl) CreateTest(input CreateTestInput) (models.Test, error) {

	test := models.Test{
		Name:                 input.Name,
		AttemptsAllowed:      input.AttemptsAllowed,
		PassPercentage:       input.PassPercentage,
		QuestionsToAnswer: input.QuestionsToAnswer,
		RandomiseAnswers:     input.RandomiseAnswers,
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&test).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create test")
		tx.Rollback()
		return models.Test{}, &errors.ErrWhileHandling
	}

	// Assert that test uuid is valid
	if test.UUID == (gentypes.UUID{}) {
		tx.Rollback()
		c.Logger.LogMessage(sentry.LevelError, "Create test uuid blank")
		return models.Test{}, &errors.ErrWhileHandling
	}

	// Create question links
	for i, uuid := range input.Questions {
		link := models.TestQuestionsLink{
			TestUUID:     test.UUID,
			QuestionUUID: uuid,
			Rank:         strconv.Itoa(i),
		}
		if err := tx.Create(&link).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelWarning, err, "Unable to create test question link")
			return models.Test{}, &errors.ErrWhileHandling
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit test")
		return models.Test{}, &errors.ErrWhileHandling
	}

	return test, nil
}
