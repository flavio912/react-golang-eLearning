package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) tutorToGentype(tutor models.Tutor) gentypes.Tutor {
	url := uploads.GetImgixURL(*tutor.SignatureKey)

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

	var s3key *string

	if input.SignatureToken != nil {
		// get key
	}

	tutor, err := c.coursesRepository.CreateTutor(input, s3key)
	return c.tutorToGentype(tutor), err
}
