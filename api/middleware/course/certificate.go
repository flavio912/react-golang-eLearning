package course

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *coursesRepoImpl) CertificateType(uuid gentypes.UUID) (models.CertificateType, error) {
	var certType models.CertificateType
	query := database.GormDB.Where("uuid = ?", uuid).First(&certType)
	if query.Error != nil {
		if query.RecordNotFound() {
			return certType, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get certType")
		return certType, &errors.ErrWhileHandling
	}
	return certType, nil
}

func (c *coursesRepoImpl) CreateCertificateType(input gentypes.CreateCertificateTypeInput) (models.CertificateType, error) {
	tx := database.GormDB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.Logger.LogMessage(sentry.LevelFatal, "CreateCertificateType: Forced to recover")
		}
	}()

	certType := models.CertificateType{
		Name:                input.Name,
		RegulationText:      input.RegulationText,
		RequiresCAANo:       input.RequiresCAANo != nil && *input.RequiresCAANo,
		ShowTrainingSection: input.ShowTrainingSection != nil && *input.ShowTrainingSection,
	}

	if err := tx.Create(&certType).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create certificate type")
		tx.Rollback()
		return models.CertificateType{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		tx.Rollback()
		return models.CertificateType{}, &errors.ErrWhileHandling
	}

	return certType, nil
}

func (c *coursesRepoImpl) CreateCAANumber(identifier string) (models.CAANumber, error) {
	no := models.CAANumber{
		Identifier: identifier,
		Used:       false,
	}

	query := database.GormDB.Create(&no)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to create CAANumber")
		return models.CAANumber{}, &errors.ErrWhileHandling
	}

	return no, nil
}
