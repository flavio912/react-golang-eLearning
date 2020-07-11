package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateTutor(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must create a tutor", func(t *testing.T) {
		input := gentypes.CreateTutorInput{
			Name: "Richard Feynman",
			CIN:  69,
		}

		tutor, err := courseRepo.CreateTutor(input)

		assert.Nil(t, err)
		assert.Equal(t, input.Name, tutor.Name)
		assert.Equal(t, uint(input.CIN), tutor.CIN)
	})
}
