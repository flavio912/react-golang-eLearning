package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/dbutils"
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
		CIN:  details.CIN,
	}

	query := database.GormDB.Create(&tutor)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create tutor")
		return models.Tutor{}, &errors.ErrWhileHandling
	}

	return tutor, nil
}

func (c *coursesRepoImpl) UpdateTutor(details gentypes.UpdateTutorInput) (models.Tutor, error) {
	tutor, err := c.Tutor(details.UUID)
	if err != nil {
		return tutor, err
	}

	if details.Name != nil {
		tutor.Name = *details.Name
	}
	if details.CIN != nil {
		tutor.CIN = *details.CIN
	}

	if err := database.GormDB.Save(&tutor).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Unable to update tutor: %s", details.UUID)
		return models.Tutor{}, &errors.ErrSaveFail
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

func (c *coursesRepoImpl) Tutors(
	page *gentypes.Page,
	filter *gentypes.TutorFilter,
	order *gentypes.OrderBy) ([]models.Tutor, gentypes.PageInfo, error) {
	var tutors []models.Tutor
	utils := dbutils.NewDBUtils(c.Logger)

	pageInfo, err := utils.GetPageOf(
		&models.Tutor{},
		&tutors,
		page,
		order,
		[]string{"name", "cin"},
		"name ASC",
		func(query *gorm.DB) *gorm.DB {
			if filter != nil {
				if filter.Name != nil {
					query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
				}
				if filter.CIN != nil {
					query = query.Where("cin ILIKE ?", "%%"+*filter.CIN+"%%")
				}
			}
			return query
		},
	)

	pageInfo.Given = int32(len(tutors))
	return tutors, pageInfo, err
}
