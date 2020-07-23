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
