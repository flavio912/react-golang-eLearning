package course

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *coursesRepoImpl) CreateTutor(details gentypes.CreateTutorInput, s3key *string) (models.Tutor, error) {
	tutor := models.Tutor{
		Name:         details.Name,
		CIN:          uint(details.CIN),
		SignatureKey: s3key,
	}

	query := database.GormDB.Create(&tutor)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create tutor")
		return models.Tutor{}, &errors.ErrWhileHandling
	}

	return tutor, nil
}
