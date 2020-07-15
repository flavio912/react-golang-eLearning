package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *courseAppImpl) generateCertificate(historicalCourse models.HistoricalCourse) {
	// - Get certificate from generator
	// - Save to S3
	// - Save S3 Key to given historical course
}
