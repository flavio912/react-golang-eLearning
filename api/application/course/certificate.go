package course

import (
	"net/http"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func (c *courseAppImpl) generateCertificate(historicalCourseUUID gentypes.UUID) {
	token, err := auth.GenerateCertificateToken(historicalCourseUUID)

	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to generate cert token - UUID: "+historicalCourseUUID.String())
	}

	generatorURL := helpers.Config.PDF.ServerURL +
		"?url=" +
		helpers.Config.PDF.RequestURL +
		"%3Ftoken=" +
		token

	client := &http.Client{}

	req, err := http.NewRequest("POST", generatorURL, nil)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to create request to PDF generator - UUID: "+historicalCourseUUID.String())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to make request to PDF generator - UUID: "+historicalCourseUUID.String())
		return
	}

	key, err := uploads.UploadCertificate(resp.Body)

	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to upload certificate - UUID: "+historicalCourseUUID.String())
		return
	}

	err = c.usersRepository.UpdateHistoricalCourse(user.UpdateHistoricalCourseInput{
		UUID:           historicalCourseUUID,
		CertificateKey: helpers.StringPointer(key),
	})

	if err != nil {
		// TODO: Delete certificate created
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to update historical course with key - UUID: "+historicalCourseUUID.String())
		return
	}
}
