package course_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
)

func TestCreateTest(t *testing.T) {
	t.Run("Creates correctly", func(t *testing.T) {
		prepareTestDatabase()

		input := course.CreateTestInput{
			Name:              "Awesome test",
			AttemptsAllowed:   12,
			PassPercentage:    23,
			QuestionsToAnswer: 12,
			RandomiseAnswers:  false,
		}
		test, err := courseRepo.CreateTest(input)
		assert.Nil(t, err)
		assert.Equal(t, "Awesome test", test.Name)
		assert.Equal(t, float64(23), test.PassPercentage)
		assert.Equal(t, uint(12), test.QuestionsToAnswer)
		assert.False(t, test.RandomiseAnswers)
	})
}

func TestUpdateTest(t *testing.T) {
	inputs := []struct {
		Name    string
		Input   course.UpdateTestInput
		WantErr error
	}{
		{
			Name: "Updates everything",
			Input: course.UpdateTestInput{
				UUID:              gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9"),
				Name:              helpers.StringPointer("I like lion cakes"),
				AttemptsAllowed:   helpers.UintPointer(12),
				PassPercentage:    helpers.FloatPointer(33.4),
				QuestionsToAnswer: helpers.UintPointer(2),
				RandomiseAnswers:  helpers.BoolPointer(false),
				Questions: &[]gentypes.UUID{
					gentypes.MustParseToUUID("797efc50-f980-42a2-a008-2991a1162631"),
					gentypes.MustParseToUUID("18cce315-ef3c-4597-aca8-f3d06e9347b1"),
				},
				Tags: &[]gentypes.UUID{
					gentypes.MustParseToUUID("1894c148-04dc-4166-ae8f-571b106c2835"),
					gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003"),
				},
			},
			WantErr: nil,
		},
		{
			Name: "Updates name and attempts",
			Input: course.UpdateTestInput{
				UUID:            gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9"),
				Name:            helpers.StringPointer("I like cheese"),
				AttemptsAllowed: helpers.UintPointer(17),
			},
			WantErr: nil,
		},
		{
			Name: "Gives invalid UUID",
			Input: course.UpdateTestInput{
				UUID:            gentypes.MustParseToUUID("9bbaf248-b832-42f9-bd0a-49e4b72d5e7d"),
				Name:            helpers.StringPointer("I like yellow"),
				AttemptsAllowed: helpers.UintPointer(34),
			},
			WantErr: &errors.ErrNotFound,
		},
	}

	assert := assert.New(t)
	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			prepareTestDatabase()

			prevTest, _ := courseRepo.Test(input.Input.UUID)

			test, err := courseRepo.UpdateTest(input.Input)
			assert.Equal(err, input.WantErr)

			if input.Input.Name != nil && input.WantErr == nil {
				assert.Equal(*input.Input.Name, test.Name)
			} else {
				assert.Equal(prevTest.Name, test.Name)
			}

			if input.Input.AttemptsAllowed != nil && input.WantErr == nil {
				assert.Equal(*input.Input.AttemptsAllowed, test.AttemptsAllowed)
			} else {
				assert.Equal(prevTest.AttemptsAllowed, test.AttemptsAllowed)
			}

			if input.Input.PassPercentage != nil && input.WantErr == nil {
				assert.Equal(*input.Input.PassPercentage, test.PassPercentage)
			} else {
				assert.Equal(prevTest.PassPercentage, test.PassPercentage)
			}

			if input.Input.QuestionsToAnswer != nil && input.WantErr == nil {
				assert.Equal(*input.Input.QuestionsToAnswer, test.QuestionsToAnswer)
			} else {
				assert.Equal(prevTest.QuestionsToAnswer, test.QuestionsToAnswer)
			}

			if input.Input.RandomiseAnswers != nil && input.WantErr == nil {
				assert.Equal(*input.Input.RandomiseAnswers, test.RandomiseAnswers)
			} else {
				assert.Equal(prevTest.RandomiseAnswers, test.RandomiseAnswers)
			}

			// Get tests
			questions, err := courseRepo.TestQuestions(test.UUID)
			assert.Nil(err)

			var uuids = make([]gentypes.UUID, len(questions))
			for i, q := range questions {
				uuids[i] = q.UUID
			}

			if input.Input.Questions != nil {
				assert.Equal(len(*input.Input.Questions), len(questions))
				for _, qid := range *input.Input.Questions {
					assert.Contains(uuids, qid)
				}
			}

			// Check tags
			// var tags []models.Tag
			// database.GormDB.Model(test).Association("Tags").Find(&tags)

			// var tagids []gentypes.UUID
			// for _, tag := range tags {

			// }
		})
	}
}

// Note to self - stop calling functions Test...
// H.Haikal: it your fault. Should've called the model Quiz or Exam
func TestTest(t *testing.T) {
	prepareTestDatabase()

	t.Run("Gets existing test", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9")
		test, err := courseRepo.Test(uuid)
		assert.Nil(t, err)
		assert.Equal(t, uuid, test.UUID)
		assert.Equal(t, "How to fibbonacci", test.Name)
		assert.Equal(t, float64(70), test.PassPercentage)
	})

	t.Run("Fails to get non-existant test", func(t *testing.T) {
		unusedUUID := gentypes.MustParseToUUID("e480c157-d758-4073-8093-ee0a4d30b4e3")
		test, err := courseRepo.Test(unusedUUID)
		assert.Equal(t, &errors.ErrNotFound, err)
		assert.Equal(t, models.Test{}, test)
	})
}

func TestManyTests(t *testing.T) {
	prepareTestDatabase()

	t.Run("Gets correctly", func(t *testing.T) {
		coolTest := gentypes.MustParseToUUID("2a7e551a-0291-422d-8508-c0ee8ff4c67e")
		fibTest := gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9")

		testMap, err := courseRepo.ManyTests([]gentypes.UUID{coolTest, fibTest})
		assert.Nil(t, err)
		assert.Equal(t, coolTest, testMap[coolTest].UUID)
		assert.Equal(t, fibTest, testMap[fibTest].UUID)

		assert.Equal(t, "Cool test name", testMap[coolTest].Name)
		assert.Equal(t, "How to fibbonacci", testMap[fibTest].Name)
	})

	t.Run("Returns errors when one not found", func(t *testing.T) {
		coolTest := gentypes.MustParseToUUID("2a7e551a-0291-422d-8508-c0ee8ff4c67e")

		// A uuid not in the database
		unusedUUID := gentypes.MustParseToUUID("c53113c5-4155-4589-8076-33f7745c6ea0")

		testMap, err := courseRepo.ManyTests([]gentypes.UUID{coolTest, unusedUUID})
		assert.Equal(t, &errors.ErrNotAllFound, err)
		assert.Equal(t, coolTest, testMap[coolTest].UUID)

		// Check that the unused uuid hasn't been added to the map
		_, ok := testMap[unusedUUID]
		assert.False(t, ok)
	})

	t.Run("Returns empty map when no items given", func(t *testing.T) {
		testMap, err := courseRepo.ManyTests([]gentypes.UUID{})
		assert.Nil(t, err)
		assert.Equal(t, map[gentypes.UUID]models.Test{}, testMap)
	})
}

func TestTestQuestions(t *testing.T) {
	prepareTestDatabase()

	t.Run("Gets all questions for a test", func(t *testing.T) {
		testWithQuestions := gentypes.MustParseToUUID("c212859c-ddd3-433c-9bf5-15cdd1db32f9")
		questions, err := courseRepo.TestQuestions(testWithQuestions)
		assert.Nil(t, err)

		assert.Len(t, questions, 3)

		// Check they're in rank order
		assert.Equal(t, gentypes.MustParseToUUID("c3c751ad-dc2d-4bf2-8a63-f1d59818aada"), questions[0].UUID)
		assert.Equal(t, gentypes.MustParseToUUID("d8ff8501-4381-4217-a332-8e87a64b968c"), questions[1].UUID)
		assert.Equal(t, gentypes.MustParseToUUID("18cce315-ef3c-4597-aca8-f3d06e9347b1"), questions[2].UUID)
	})

	t.Run("Gets empty list for tests with no questions", func(t *testing.T) {
		testWithoutQuestions := gentypes.MustParseToUUID("2a7e551a-0291-422d-8508-c0ee8ff4c67e")
		questions, err := courseRepo.TestQuestions(testWithoutQuestions)
		assert.Nil(t, err)
		assert.Len(t, questions, 0)
	})
}

func TestManyAnswers(t *testing.T) {
	prepareTestDatabase()

	t.Run("Gets all answers", func(t *testing.T) {
		cheddarCalories := gentypes.MustParseToUUID("d8ff8501-4381-4217-a332-8e87a64b968c")
		wally := gentypes.MustParseToUUID("c3c751ad-dc2d-4bf2-8a63-f1d59818aada")

		ansMap, err := courseRepo.ManyAnswers([]gentypes.UUID{wally, cheddarCalories})
		assert.Nil(t, err)

		assert.Len(t, ansMap[cheddarCalories], 2)
		assert.Len(t, ansMap[wally], 2)

		items := []struct {
			mapIndex      gentypes.UUID
			expectedUUIDs []string
		}{
			{
				mapIndex: cheddarCalories,
				expectedUUIDs: []string{
					"8deab18c-824f-4a45-9b65-d533833d80bf",
					"8901853c-e73e-46f9-b7a6-7cd6e714dbe0",
				},
			},
			{
				mapIndex: wally,
				expectedUUIDs: []string{
					"5ca49c84-8ff5-4fc6-8fb4-5b6406043427",
					"02e9d22d-af63-4f7d-af58-01c87101af19",
				},
			},
		}

		for _, item := range items {
			answers := ansMap[item.mapIndex]
			assert.Len(t, answers, len(item.expectedUUIDs))
			for i, uuid := range item.expectedUUIDs {
				assert.Equal(t, uuid, answers[i].UUID.String())
			}
		}
	})
}

func TestCourseTests(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should get all tests + nested tests", func(t *testing.T) {
		onlineCourse := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000003")

		tests, err := courseRepo.CourseTests(onlineCourse)
		assert.Nil(t, err)
		assert.Len(t, tests, 2)

		var testUUIDs = make([]string, len(tests))
		for i, test := range tests {
			testUUIDs[i] = test.UUID.String()
		}

		assert.Contains(t, testUUIDs, "2a7e551a-0291-422d-8508-c0ee8ff4c67e")
		assert.Contains(t, testUUIDs, "c212859c-ddd3-433c-9bf5-15cdd1db32f9")
	})
}

func TestDeleteTest(t *testing.T) {
	t.Run("Should not delete test that is part of a course", func(t *testing.T) {
		prepareTestDatabase()

		uuid := gentypes.MustParseToUUID("2a7e551a-0291-422d-8508-c0ee8ff4c67e")

		b, err := courseRepo.DeleteTest(uuid)

		assert.Equal(t, &errors.ErrWhileHandling, err)
		assert.False(t, b)
	})

	t.Run("Deletes existing test", func(t *testing.T) {
		prepareTestDatabase()

		uuid := gentypes.MustParseToUUID("2a56f8a8-1cd3-4e7b-bd10-c489b519828d")

		b, err := courseRepo.DeleteTest(uuid)

		assert.Nil(t, err)
		assert.True(t, b)

		questions, q_err := courseRepo.TestQuestions(uuid)

		assert.Nil(t, q_err)
		assert.Len(t, questions, 0)
	})
}
