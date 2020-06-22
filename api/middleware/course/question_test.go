package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
)

func TestCreateQuestion(t *testing.T) {
	t.Run("Create valid question with tags and answers", func(t *testing.T) {
		prepareTestDatabase()

		question, err := courseRepo.CreateQuestion(course.CreateQuestionArgs{
			Text:             "Who is the queen of England?",
			RandomiseAnswers: true,
			QuestionType:     gentypes.SingleAnswerType,
			Answers: []course.AnswerArgs{
				course.AnswerArgs{
					IsCorrect: false,
					Text:      helpers.StringPointer("Cheesecake"),
				},
				course.AnswerArgs{
					IsCorrect: true,
					Text:      helpers.StringPointer("Liz"),
				},
				course.AnswerArgs{
					IsCorrect: true,
					Text:      helpers.StringPointer("Tom Riddle"),
				},
			},
			Tags: []gentypes.UUID{
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, "Who is the queen of England?", question.Text)
		assert.Equal(t, true, question.RandomiseAnswers)
		assert.Equal(t, gentypes.SingleAnswerType, question.QuestionType)

		assert.Equal(t, gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"), question.Tags[0].UUID)

		// Get answer links
		var answers []models.BasicAnswer
		database.GormDB.Where("question_uuid = ?", question.UUID).Order("rank ASC").Find(&answers)
		assert.Equal(t, 3, len(answers))
		assert.Equal(t, "0", answers[0].Rank)
		assert.Equal(t, "Cheesecake", *answers[0].Text)
		assert.Nil(t, answers[0].ImageKey)
		assert.False(t, answers[0].IsCorrect)

		assert.Equal(t, "1", answers[1].Rank)
		assert.True(t, answers[1].IsCorrect)

		assert.Equal(t, "2", answers[2].Rank)
	})
}

func TestUpdateQuestion(t *testing.T) {
	t.Run("Update existing question", func(t *testing.T) {
		prepareTestDatabase()

		answerArgs := []course.AnswerArgs{
			course.AnswerArgs{
				IsCorrect: false,
				Text:      helpers.StringPointer("Cheesecake"),
			},
			course.AnswerArgs{
				IsCorrect: true,
				Text:      helpers.StringPointer("Liz"),
			},
			course.AnswerArgs{
				IsCorrect: true,
				Text:      helpers.StringPointer("Tom Riddle"),
			},
		}

		args := course.UpdateQuestionArgs{
			UUID:             gentypes.MustParseToUUID("d8ff8501-4381-4217-a332-8e87a64b968c"),
			Text:             helpers.StringPointer("New question"),
			RandomiseAnswers: helpers.BoolPointer(true),
			Answers:          &answerArgs,
			Tags: &[]gentypes.UUID{
				gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000002"),
			},
		}

		question, err := courseRepo.UpdateQuestion(args)
		assert.Nil(t, err)
		assert.Equal(t, *args.Text, question.Text)
		assert.Equal(t, *args.RandomiseAnswers, question.RandomiseAnswers)

		var answers []models.BasicAnswer
		database.GormDB.Where("question_uuid = ?", args.UUID).Order("rank ASC").Find(&answers)
		assert.Equal(t, 3, len(answers))

		assert.Equal(t, *answerArgs[0].Text, *answers[0].Text)
		assert.Equal(t, answerArgs[0].IsCorrect, answers[0].IsCorrect)

		assert.Equal(t, *answerArgs[1].Text, *answers[1].Text)
		assert.Equal(t, answerArgs[1].IsCorrect, answers[1].IsCorrect)

		assert.Equal(t, *answerArgs[2].Text, *answers[2].Text)
		assert.Equal(t, answerArgs[2].IsCorrect, answers[2].IsCorrect)
	})
}
