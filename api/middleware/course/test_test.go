package course_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
)

func TestCreateTest(t *testing.T) {
	t.Run("Creates correctly", func(t *testing.T) {
		prepareTestDatabase()

		input := course.CreateTestInput{
			Name:              "Awesome test",
			AttemptsAllowed:   helpers.UintPointer(12),
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

// Note to self - stop calling functions Test...
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
