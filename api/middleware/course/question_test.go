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
			Answers: []course.CreateAnswerArgs{
				course.CreateAnswerArgs{
					IsCorrect: false,
					Text:      helpers.StringPointer("Cheesecake"),
				},
				course.CreateAnswerArgs{
					IsCorrect: true,
					Text:      helpers.StringPointer("Liz"),
				},
				course.CreateAnswerArgs{
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

		ansId := gentypes.MustParseToUUID("8901853c-e73e-46f9-b7a6-7cd6e714dbe0")
		ansId1 := gentypes.MustParseToUUID("8deab18c-824f-4a45-9b65-d533833d80bf")
		answerArgs := []course.UpdateAnswerArgs{
			course.UpdateAnswerArgs{
				UUID:       &ansId,
				IsCorrect:  helpers.BoolPointer(false),
				Text:       helpers.StringPointer("Cheesecake"),
				AnswerType: gentypes.TextAnswer,
			},
			course.UpdateAnswerArgs{
				IsCorrect:  helpers.BoolPointer(true),
				Text:       helpers.StringPointer("Liz"),
				AnswerType: gentypes.TextAnswer,
			},
			course.UpdateAnswerArgs{
				UUID:       &ansId1,
				IsCorrect:  helpers.BoolPointer(false),
				Text:       helpers.StringPointer("Some cool text"),
				AnswerType: gentypes.TextAnswer,
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
		assert.Equal(t, *answerArgs[0].IsCorrect, answers[0].IsCorrect)

		assert.Equal(t, *answerArgs[1].Text, *answers[1].Text)
		assert.Equal(t, *answerArgs[1].IsCorrect, answers[1].IsCorrect)

		assert.Equal(t, *answerArgs[2].Text, *answers[2].Text)
		assert.Equal(t, *answerArgs[2].IsCorrect, answers[2].IsCorrect)
	})
}

func TestDeleteQuestion(t *testing.T) {
	t.Run("Should not delete question that is part of a test", func(t *testing.T) {
		prepareTestDatabase()

		uuid := gentypes.MustParseToUUID("d8ff8501-4381-4217-a332-8e87a64b968c")
		b, err := courseRepo.DeleteQuestion(uuid)

		assert.NotNil(t, err)
		assert.False(t, b)
	})

	t.Run("Deletes existing question", func(t *testing.T) {
		prepareTestDatabase()

		uuid := gentypes.MustParseToUUID("ba070bfb-d3d0-4ff7-a35d-6263180a43f9")
		b, err := courseRepo.DeleteQuestion(uuid)

		assert.Nil(t, err)
		assert.True(t, b)

		var answers []models.BasicAnswer
		database.GormDB.Table("basic_answers").Where("question_uuid = ?", uuid).Find(&answers)

		assert.Len(t, answers, 0)
	})
}
