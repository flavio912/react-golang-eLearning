package user_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateTakerActivity(t *testing.T) {
	prepareTestDatabase()

	validCourseId := uint(4)
	invalidCourseId := uint(3432)

	t.Run("Creates valid activity item with course", func(t *testing.T) {
		takerUUID := gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539")
		activityItem, err := usersRepo.CreateTakerActivity(takerUUID, gentypes.ActivityCompleted, &validCourseId)
		assert.Nil(t, err)

		assert.Equal(t, gentypes.ActivityCompleted, activityItem.ActivityType)
		assert.Equal(t, uint(4), *activityItem.CourseID)
		assert.Equal(t, takerUUID, activityItem.CourseTakerUUID)
	})

	t.Run("Can add nil courseID", func(t *testing.T) {
		takerUUID := gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539")
		activityItem, err := usersRepo.CreateTakerActivity(takerUUID, gentypes.ActivityCompleted, nil)
		assert.Nil(t, err)

		assert.Equal(t, gentypes.ActivityCompleted, activityItem.ActivityType)
		assert.Nil(t, activityItem.CourseID)
		assert.Equal(t, takerUUID, activityItem.CourseTakerUUID)
	})

	t.Run("Fails if attempt to add non-existant takerID", func(t *testing.T) {
		takerUUID := gentypes.MustParseToUUID("6a7dd024-0a48-4cab-b270-e429c7b92204")
		activityItem, err := usersRepo.CreateTakerActivity(takerUUID, gentypes.ActivityCompleted, &validCourseId)
		assert.Equal(t, &errors.ErrWhileHandling, err)
		assert.Equal(t, models.CourseTakerActivity{}, activityItem)
	})

	t.Run("Fails if attempt to add non-existant courseID", func(t *testing.T) {
		takerUUID := gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539")
		activityItem, err := usersRepo.CreateTakerActivity(takerUUID, gentypes.ActivityCompleted, &invalidCourseId)
		assert.Equal(t, &errors.ErrWhileHandling, err)
		assert.Equal(t, models.CourseTakerActivity{}, activityItem)
	})
}

func TestDeleteTakerActivity(t *testing.T) {
	prepareTestDatabase()
	t.Run("Can delete activity", func(t *testing.T) {
		err := usersRepo.DeleteTakerActivity(gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539"))
		assert.Nil(t, err)
	})
}

func TestTakerActivity(t *testing.T) {
	prepareTestDatabase()
	t.Run("Gets a takers activity", func(t *testing.T) {
		activityItems, pageInfo, err := usersRepo.TakerActivity(gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539"), nil)
		assert.Nil(t, err)
		assert.Len(t, activityItems, 2)
		assert.Equal(t, gentypes.PageInfo{
			Total:  2,
			Offset: 0,
			Limit:  2,
			Given:  2,
		}, pageInfo)
	})
}

func TestTakerActivitys(t *testing.T) {
	prepareTestDatabase()
	t.Run("Gets multiple takers activityItems", func(t *testing.T) {
		activityItems, pageInfo, err := usersRepo.TakerActivitys([]gentypes.UUID{
			gentypes.MustParseToUUID("f6d7978d-5e1a-45e9-ae8d-fad5d9c10539"),
			gentypes.MustParseToUUID("2c90d590-dbe1-425b-ad5c-eebf7abcf9da"),
		}, nil)
		assert.Nil(t, err)
		assert.Len(t, activityItems, 3)
		assert.Equal(t, gentypes.PageInfo{
			Total:  3,
			Offset: 0,
			Limit:  3,
			Given:  3,
		}, pageInfo)
	})
}
