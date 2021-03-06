package course

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func (c *courseAppImpl) certificateTypeToGentype(certType models.CertificateType) gentypes.CertificateType {
	var cert_url *string
	if certType.CertificateBodyImageKey != nil {
		url := uploads.GetImgixURL(*certType.CertificateBodyImageKey)
		cert_url = &url
	}

	createdAt := certType.CreatedAt.Format(time.RFC3339)
	return gentypes.CertificateType{
		UUID:                    certType.UUID,
		Name:                    certType.Name,
		CreatedAt:               createdAt,
		RegulationText:          certType.RegulationText,
		RequiresCAANo:           certType.RequiresCAANo,
		ShowTrainingSection:     certType.ShowTrainingSection,
		CertificateBodyImageURL: cert_url,
	}
}

func (c *courseAppImpl) certificatesTypeToGentype(certTypes []models.CertificateType) []gentypes.CertificateType {
	gens := make([]gentypes.CertificateType, len(certTypes))
	for i, cert := range certTypes {
		gens[i] = c.certificateTypeToGentype(cert)
	}
	return gens
}

func (c *courseAppImpl) caaNumberToGentype(no models.CAANumber) gentypes.CAANumber {
	createdAt := no.CreatedAt.Format(time.RFC3339)
	return gentypes.CAANumber{
		UUID:       no.UUID,
		CreatedAt:  createdAt,
		Identifier: no.Identifier,
		Used:       no.Used,
	}
}

func (c *courseAppImpl) caaNumbersToGentypes(numbers []models.CAANumber) []gentypes.CAANumber {
	gens := make([]gentypes.CAANumber, len(numbers))
	for i, no := range numbers {
		gens[i] = c.caaNumberToGentype(no)
	}
	return gens
}

func GeneratePdfFromURL(url string) (io.Reader, error) {
	client := &http.Client{}

	// TODO: Make key an envvar
	var payload = []byte(
		fmt.Sprintf(`{"url":"%s", "direct": true, "key": "dea9fe3bc2b5981dbb46001bac2e7aa324971384235861a0d9666c70110a016224971384235861a0d9666c70110a0162"}`, url),
	)

	req, err := http.NewRequest("POST", helpers.Config.PDF.ServerURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &errors.ErrWhileHandling
	}

	return resp.Body, nil
}

func (c *courseAppImpl) RegenerateCertificate(historicalCourseUUID gentypes.UUID) error {
	if !c.grant.IsAdmin {
		return &errors.ErrUnauthorized
	}

	go c.generateCertificate(historicalCourseUUID)
	return nil
}

// generateCertificate updates a historicalCourse with its generated cert
func (c *courseAppImpl) generateCertificate(historicalCourseUUID gentypes.UUID) {
	token, err := auth.GenerateCertificateToken(historicalCourseUUID)
	glog.Error(token)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to generate cert token - UUID: "+historicalCourseUUID.String())
	}

	htmlPageURL := helpers.Config.PDF.RequestURL + "%3Ftoken=" + token

	pdf, err := GeneratePdfFromURL(htmlPageURL)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelError, err, "Unable to make request to PDF generator - UUID: "+historicalCourseUUID.String())
		return
	}

	key, err := uploads.UploadCertificate(pdf)

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
		InstructorName:         "Michelle Waddilove",
		InstructorCIN:          "0123445",
		InstructorSignatureURL: nil,
		CertificateNumber:      "0001",
	}, nil
}

func (c *courseAppImpl) CreateCertificateType(input gentypes.CreateCertificateTypeInput) (gentypes.CertificateType, error) {
	if !c.grant.IsAdmin {
		return gentypes.CertificateType{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.CertificateType{}, err
	}

	certType, err := c.coursesRepository.CreateCertificateType(input)
	return c.certificateTypeToGentype(certType), err
}

func (c *courseAppImpl) UpdateCertificateType(input gentypes.UpdateCertificateTypeInput) (gentypes.CertificateType, error) {
	if !c.grant.IsAdmin {
		return gentypes.CertificateType{}, &errors.ErrUnauthorized
	}

	if ok, err := govalidator.ValidateStruct(input); !ok || err != nil {
		return gentypes.CertificateType{}, err
	}

	certType, err := c.coursesRepository.UpdateCertificateType(input)
	return c.certificateTypeToGentype(certType), err
}

func (c *courseAppImpl) CreateCAANumber(input gentypes.CreateCAANumberInput) (gentypes.CAANumber, error) {
	if !c.grant.IsAdmin {
		return gentypes.CAANumber{}, &errors.ErrUnauthorized
	}

	if ok, err := govalidator.ValidateStruct(input); !ok || err != nil {
		return gentypes.CAANumber{}, err
	}

	number, err := c.coursesRepository.CreateCAANumber(input.Identifier)
	return c.caaNumberToGentype(number), err
}

func (c *courseAppImpl) UpdateCAANumber(input gentypes.UpdateCAANumberInput) (gentypes.CAANumber, error) {
	if !c.grant.IsAdmin {
		return gentypes.CAANumber{}, &errors.ErrUnauthorized
	}

	if ok, err := govalidator.ValidateStruct(input); !ok || err != nil {
		return gentypes.CAANumber{}, err
	}

	number, err := c.coursesRepository.UpdateCAANumber(input)
	return c.caaNumberToGentype(number), err
}

func (c *courseAppImpl) CertificateType(uuid gentypes.UUID) (gentypes.CertificateType, error) {
	if !c.grant.IsAdmin {
		return gentypes.CertificateType{}, &errors.ErrUnauthorized
	}

	certType, err := c.coursesRepository.CertificateType(uuid)
	return c.certificateTypeToGentype(certType), err
}

func (c *courseAppImpl) CertificateTypes(
	page *gentypes.Page,
	filter *gentypes.CertificateTypeFilter) ([]gentypes.CertificateType, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.CertificateType{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	certTypes, pageInfo, err := c.coursesRepository.CertificateTypes(page, filter)
	return c.certificatesTypeToGentype(certTypes), pageInfo, err
}

func (c *courseAppImpl) CAANumbers(
	page *gentypes.Page,
	filter *gentypes.CAANumberFilter) ([]gentypes.CAANumber, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.CAANumber{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	numbers, pageInfo, err := c.coursesRepository.CAANumbers(page, filter)
	return c.caaNumbersToGentypes(numbers), pageInfo, err
}

func (c *courseAppImpl) CertificateBodyImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,
		imageMeta.ContentLength,
		[]string{"jpg", "png"},
		int32(1000000),
		"certificate_body_image",
		"certificateBodyImage",
	)

	return url, successToken, err
}
