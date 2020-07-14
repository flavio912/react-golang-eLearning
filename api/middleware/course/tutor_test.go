package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func TestCreateTutor(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must create a tutor", func(t *testing.T) {
		input := gentypes.CreateTutorInput{
			Name: "Richard Feynman",
			CIN:  "69",
		}

		tutor, err := courseRepo.CreateTutor(input)

		assert.Nil(t, err)
		assert.Equal(t, input.Name, tutor.Name)
		assert.Equal(t, input.CIN, tutor.CIN)
	})
}

func TestTutor(t *testing.T) {
	prepareTestDatabase()

	t.Run("Gets existing tutor", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("386bd256-82e0-4d8a-91af-b4a117e0eda8")
		tutor, err := courseRepo.Tutor(uuid)

		assert.Nil(t, err)
		assert.Equal(t, "Mohammed Rashwan", tutor.Name)
		assert.Equal(t, "100", tutor.CIN)
	})

	t.Run("Must fail to get non-existant", func(t *testing.T) {
		uuid := gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000")

		tutor, err := courseRepo.Tutor(uuid)

		assert.Equal(t, errors.ErrTutorDoesNotExist(uuid.String()), err)
		assert.Equal(t, models.Tutor{}, tutor)
	})
}

func TestUpdateTutor(t *testing.T) {
	prepareTestDatabase()

	t.Run("Cannot update non-existant tutor", func(t *testing.T) {
		input := gentypes.UpdateTutorInput{
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"),
		}

		tutor, err := courseRepo.UpdateTutor(input)

		assert.Equal(t, errors.ErrTutorDoesNotExist(input.UUID.String()), err)
		assert.Equal(t, models.Tutor{}, tutor)
	})

	t.Run("Update some fields of tutor", func(t *testing.T) {
		input := gentypes.UpdateTutorInput{
			UUID: gentypes.MustParseToUUID("386bd256-82e0-4d8a-91af-b4a117e0eda8"),
			Name: helpers.StringPointer("Walter White"),
			CIN:  helpers.StringPointer("69"),
		}

		tutor, err := courseRepo.UpdateTutor(input)

		assert.Nil(t, err)
		assert.Equal(t, *input.Name, tutor.Name)
		assert.Equal(t, *input.CIN, tutor.CIN)
	})
}
