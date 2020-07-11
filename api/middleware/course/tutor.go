package course

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *coursesRepoImpl) Tutor(uuid gentypes.UUID) (models.Tutor, error) {
	var tutor models.Tutor
	query := database.GormDB.Where("uuid = ?", uuid).Find(&tutor)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Tutor{}, errors.ErrTutorDoesNotExist(uuid.String())
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to get tutor: %s", uuid)
		return models.Tutor{}, &errors.ErrWhileHandling
	}

	return tutor, nil
}

func (c *coursesRepoImpl) CreateTutor(details gentypes.CreateTutorInput) (models.Tutor, error) {
	tutor := models.Tutor{
		Name: details.Name,
		CIN:  uint(details.CIN),
	}

	query := database.GormDB.Create(&tutor)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create tutor")
		return models.Tutor{}, &errors.ErrWhileHandling
	}

	return tutor, nil
}

func (c *coursesRepoImpl) UpdateTutorSignature(tutorUUID gentypes.UUID, s3key string) error {
	query := database.GormDB.Model(&models.Tutor{}).Where("uuid = ?", tutorUUID).Update("signature_key", s3key)
	if query.Error != nil {
		if query.RecordNotFound() {
			return errors.ErrTutorDoesNotExist(tutorUUID.String())
		}

		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to update tutor signature: %s", tutorUUID)
		return &errors.ErrWhileHandling
	}

	return nil
}
