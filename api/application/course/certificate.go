package course

import (
	"net/http"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

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

// CertificateInfo gets the information for the pdf generator to build a cert
func (c *courseAppImpl) CertificateInfo(token string) (gentypes.CertficateInfo, error) {
	histCourseUUID, err := auth.ValidateCertificateToken(token)
	if err != nil {
		return gentypes.CertficateInfo{}, &errors.ErrUnauthorized
	}

	historicalCourse, err := c.usersRepository.HistoricalCourse(histCourseUUID)
	if err != nil {
		return gentypes.CertficateInfo{}, err
	}

	if !historicalCourse.Passed || historicalCourse.ExpirationDate == nil {
		c.grant.Logger.LogMessage(sentry.LevelError, "Tried to create a cert for an invalid hist course - UUID: "+histCourseUUID.String())
	}

	course, err := c.coursesRepository.Course(historicalCourse.CourseID)
	if err != nil {
		return gentypes.CertficateInfo{}, err
	}

	var (
		firstName      string
		lastName       string
		companyName    *string
		caaNo          *string
		regulationText string
		certTitle      string
	)
	delegate, individual := c.usersRepository.UserFromCourseTaker(historicalCourse.CourseTakerUUID)
	switch {
	case delegate != nil:
		firstName = delegate.FirstName
		lastName = delegate.LastName
		comp, _ := c.usersRepository.Company(delegate.CompanyUUID)
		companyName = helpers.StringPointer(comp.Name)
	case individual != nil:
		firstName = individual.FirstName
		lastName = individual.LastName
	default:
		c.grant.Logger.LogMessage(sentry.LevelError, "No delegate or individual associdate with taker - UUID:"+histCourseUUID.String())
		return gentypes.CertficateInfo{}, &errors.ErrUserNotFound
	}

	if course.CertificateTypeUUID != nil {
		certType, err := c.coursesRepository.CertificateType(*course.CertificateTypeUUID)
		if err != nil {
			c.grant.Logger.Log(sentry.LevelError, err, "CertInfo: Unable to get certtype")
			return gentypes.CertficateInfo{}, &errors.ErrWhileHandling
		}

		regulationText = certType.RegulationText
		certTitle = certType.Name
	}

	// TODO get caaNo

	return gentypes.CertficateInfo{
		CourseTitle:            course.Name,
		ExpiryDate:             (*historicalCourse.ExpirationDate).Format(time.RFC3339),
		CompletionDate:         historicalCourse.CreatedAt.Format(time.RFC3339),
		CompanyName:            companyName,
		TakerFirstName:         firstName,
		TakerLastName:          lastName,
		CertificateBodyURL:     nil,
		RegulationText:         regulationText,
		CAANo:                  caaNo,
		Title:                  certTitle,
		InstructorName:         "",
		InstructorCIN:          "",
		InstructorSignatureURL: nil,
	}, nil
}
