package gentypes

import "github.com/asaskevich/govalidator"

type CertficateInfo struct {
	CourseTitle            string
	ExpiryDate             string
	CompletionDate         string
	CompanyName            *string
	TakerFirstName         string
	TakerLastName          string
	CertificateBodyURL     *string
	RegulationText         string
	CAANo                  *string
	Title                  string
	InstructorName         string
	InstructorCIN          string
	InstructorSignatureURL *string
}

type CertificateType struct {
	UUID                    UUID
	Name                    string
	CreatedAt               string
	CertificateBodyImageURL *string
	RegulationText          string
	RequiresCAANo           bool
	ShowTrainingSection     bool
}

type CreateCertificateTypeInput struct {
	Name                 string `valid:"required"`
	RegulationText       string `valid:"required"`
	RequiresCAANo        *bool
	ShowTrainingSection  *bool
	CertificateBodyToken *string
}

func (c *CreateCertificateTypeInput) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
