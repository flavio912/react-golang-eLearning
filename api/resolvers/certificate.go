package resolvers

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type CertificateTypeResolver struct {
	CertificateType gentypes.CertificateType
}

func (c *CertificateTypeResolver) UUID() gentypes.UUID { return c.CertificateType.UUID }
func (c *CertificateTypeResolver) Name() string        { return c.CertificateType.Name }
func (c *CertificateTypeResolver) CreatedAt() string   { return c.CertificateType.CreatedAt }
func (c *CertificateTypeResolver) CertificateBodyImageURL() *string {
	return c.CertificateType.CertificateBodyImageURL
}
func (c *CertificateTypeResolver) RegulationText() string { return c.CertificateType.RegulationText }
func (c *CertificateTypeResolver) RequiresCAANo() bool    { return c.CertificateType.RequiresCAANo }
func (c *CertificateTypeResolver) ShowTrainingSection() bool {
	return c.CertificateType.ShowTrainingSection
}
