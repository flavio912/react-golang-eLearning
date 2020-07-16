package gentypes

import "time"

type CertficateInfo struct {
	CourseTitle            string
	ExpiryDate             time.Time
	CompletionDate         time.Time
	CompanyName            *string
	TakerFirstName         string
	TakerLastName          string
	CertificateBodyURL     *string
	RegulationText         string
	CAANo                  *string
	Title                  string
	InstructorName         string
	InstructorCIN          string
	InstructorSignatureURL string
}
