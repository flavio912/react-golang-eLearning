package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
)

func TestCreateTest(t *testing.T) {
	t.Run("Creates correctly", func(t *testing.T) {
		prepareTestDatabase()

		input := course.CreateTestInput{
			Name:                 "Awesome test",
			AttemptsAllowed:      helpers.IntPointer(12),
			PassPercentage:       23,
			MinQuestionsToAnswer: 12,
			RandomiseAnswers:     false,
		}
		test, err := courseRepo.CreateTest(input)
		assert.Nil(t, err)
		assert.Equal(t, "Awesome test", test.Name)
		assert.Equal(t, float32(23), test.PassPercentage)
		assert.Equal(t, 12, test.MinQuestionsToAnswer)
		assert.False(t, test.RandomiseAnswers)
	})
}
