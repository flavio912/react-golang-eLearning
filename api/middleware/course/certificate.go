package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
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

func (c *coursesRepoImpl) CAANumber(uuid gentypes.UUID) (models.CAANumber, error) {
	var no models.CAANumber
	query := database.GormDB.Model(&models.CAANumber{}).Where("uuid = ?", uuid).Find(&no)
	if query.Error != nil {
		c.Logger.Logf(sentry.LevelError, query.Error, "Unable to find CAANumber: %s", uuid)
		return models.CAANumber{}, &errors.ErrNotFound
	}

	return no, nil
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

func filterCertificateTypes(query *gorm.DB, filter *gentypes.CertificateTypeFilter) *gorm.DB {
	if filter != nil {
		if helpers.StringNotNilOrEmpty(filter.Name) {
			query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if helpers.StringNotNilOrEmpty(filter.RegulationText) {
			query = query.Where("regulation_text ILIKE ?", "%%"+*filter.RegulationText+"%%")
		}
		if filter.RequiresCAANo != nil {
			query = query.Where("requires_caa_no = ?", *filter.RequiresCAANo)
		}
		if filter.ShowTrainingSection != nil {
			query = query.Where("show_training_section = ?", *filter.ShowTrainingSection)
		}
	}

	return query
}

func (c *coursesRepoImpl) CertificateTypes(
	page *gentypes.Page,
	filter *gentypes.CertificateTypeFilter) ([]models.CertificateType, gentypes.PageInfo, error) {
	var certTypes []models.CertificateType

	query := filterCertificateTypes(database.GormDB, filter)

	var count int32
	countErr := query.Model(&models.CertificateType{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		c.Logger.Log(sentry.LevelWarning, countErr, "Unable to count certificate types")
		return certTypes, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&certTypes)

	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get certificate types")
		return []models.CertificateType{}, gentypes.PageInfo{}, &errors.ErrNotAllFound
	}

	return certTypes, gentypes.PageInfo{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Given:  int32(len(certTypes)),
	}, nil
}

func filterCAANumbers(query *gorm.DB, filter *gentypes.CAANumberFilter) *gorm.DB {
	if filter != nil {
		if helpers.StringNotNilOrEmpty(filter.Identifier) {
			query = query.Where("identifier ILIKE ?", "%%"+*filter.Identifier+"%%")
		}
		if filter.Used != nil {
			query = query.Where("used = ?", *filter.Used)
		}
	}

	return query
}

func (c *coursesRepoImpl) CAANumbers(
	page *gentypes.Page,
	filter *gentypes.CAANumberFilter) ([]models.CAANumber, gentypes.PageInfo, error) {
	var numbers []models.CAANumber

	query := filterCAANumbers(database.GormDB, filter)

	var count int32
	countErr := query.Model(&models.CAANumber{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		c.Logger.Log(sentry.LevelWarning, countErr, "Unable to count CAANumbers")
		return numbers, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&numbers)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to find CAANumbers")
		return []models.CAANumber{}, gentypes.PageInfo{}, &errors.ErrNotAllFound
	}

	return numbers, gentypes.PageInfo{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Given:  int32(len(numbers)),
	}, nil
}
