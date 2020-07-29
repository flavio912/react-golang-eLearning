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
	CertificateNumber      string
}

type CAANumber struct {
	UUID       UUID
	CreatedAt  string
	Identifier string
	Used       bool
}

type CreateCAANumberInput struct {
	Identifier string `valid:"required"`
}

type UpdateCAANumberInput struct {
	UUID       UUID `valid:"required"`
	Identifier *string
	Used       *bool
}

type CAANumberFilter struct {
	Identifier *string
	Used       *bool
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

type CertificateTypeFilter struct {
	Name                *string
	RegulationText      *string
	RequiresCAANo       *bool
	ShowTrainingSection *bool
}

type UpdateCertificateTypeInput struct {
	UUID                 UUID `valid:"required"`
	Name                 *string
	RegulationText       *string
	RequiresCAANo        *bool
	ShowTrainingSection  *bool
	CertificateBodyToken *string
}
