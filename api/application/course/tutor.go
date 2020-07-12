package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) tutorToGentype(tutor models.Tutor) gentypes.Tutor {
	url := uploads.GetImgixURL(tutor.SignatureKey)

	return gentypes.Tutor{
		UUID:         tutor.UUID,
		Name:         tutor.Name,
		CIN:          int32(tutor.CIN),
		SignatureURL: url,
	}
}

func (c *courseAppImpl) CreateTutor(input gentypes.CreateTutorInput) (gentypes.Tutor, error) {
	if !c.grant.IsAdmin {
		return gentypes.Tutor{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Tutor{}, err
	}

	tutor, err := c.coursesRepository.CreateTutor(input)

	if input.SignatureToken != nil {
		key, err := c.UpdateTutorSignature(gentypes.UpdateTutorSignatureInput{
			FileSuccess: gentypes.UploadFileSuccess{
				SuccessToken: *input.SignatureToken,
			},
			TutorUUID: tutor.UUID,
		})
		if err != nil {
			return gentypes.Tutor{}, err
		}

		tutor.SignatureKey = key
	}

	return c.tutorToGentype(tutor), err
}

func (c *courseAppImpl) UpdateTutor(input gentypes.UpdateTutorInput) (gentypes.Tutor, error) {
	if !c.grant.IsAdmin {
		return gentypes.Tutor{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Tutor{}, err
	}

	tutor, err := c.coursesRepository.UpdateTutor(input)

	if input.SignatureToken != nil {
		key, err := c.UpdateTutorSignature(gentypes.UpdateTutorSignatureInput{
			FileSuccess: gentypes.UploadFileSuccess{
				SuccessToken: *input.SignatureToken,
			},
			TutorUUID: tutor.UUID,
		})
		if err != nil {
			return gentypes.Tutor{}, err
		}

		tutor.SignatureKey = key
	}

	return c.tutorToGentype(tutor), err
}

func (c *courseAppImpl) TutorSignatureImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,
		imageMeta.ContentLength,
		[]string{"png"},
		int32(5000000), // 5 Mb, how big a signature can be?
		"tutor_signature",
		"tutorSignature",
	)

	return url, successToken, err
}

func (c *courseAppImpl) UpdateTutorSignature(input gentypes.UpdateTutorSignatureInput) (string, error) {
	if !c.grant.IsAdmin {
		return "", &errors.ErrUnauthorized
	}

	s3key, err := uploads.VerifyUploadSuccess(input.FileSuccess.SuccessToken, "tutorSignature")
	if err != nil {
		return "", err
	}

	err = c.coursesRepository.UpdateTutorSignature(input.TutorUUID, s3key)
	if err != nil {
		return "", err
	}

	return s3key, err
}
