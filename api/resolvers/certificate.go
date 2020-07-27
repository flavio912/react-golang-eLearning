package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type CertificateTypeResolver struct {
	CertificateType gentypes.CertificateType
}

type NewCertificateTypeArgs struct {
	CertificateTypeUUID *gentypes.UUID
	CertificateType     *gentypes.CertificateType
}

func NewCertificateTypeResolver(ctx context.Context, args NewCertificateTypeArgs) (*CertificateTypeResolver, error) {
	switch {
	case args.CertificateTypeUUID != nil:
		app := auth.AppFromContext(ctx)
		certType, err := app.CourseApp.CertificateType(*args.CertificateTypeUUID)

		if err != nil {
			return &CertificateTypeResolver{}, err
		}
		return &CertificateTypeResolver{
			CertificateType: certType,
		}, nil

	case args.CertificateType != nil:
		return &CertificateTypeResolver{
			CertificateType: *args.CertificateType,
		}, nil
	default:
		return &CertificateTypeResolver{}, &errors.ErrUnableToResolve
	}
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

type CertificateTypePageResolver struct {
	edges    *[]*CertificateTypeResolver
	pageInfo *PageInfoResolver
}

type NewCertificateTypePageArgs struct {
	CertificateTypes *[]gentypes.CertificateType
	PageInfo         *gentypes.PageInfo
}

func NewCertificateTypePageResolver(ctx context.Context, args NewCertificateTypePageArgs) (*CertificateTypePageResolver, error) {
	var resolvers []*CertificateTypeResolver

	switch {
	case args.CertificateTypes != nil:
		for _, certType := range *args.CertificateTypes {
			res, err := NewCertificateTypeResolver(ctx, NewCertificateTypeArgs{
				CertificateType: &certType,
			})

			if err != nil {
				return &CertificateTypePageResolver{}, err
			}

			resolvers = append(resolvers, res)
		}
	}

	return &CertificateTypePageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: args.PageInfo,
		},
	}, nil
}

func (r *CertificateTypePageResolver) PageInfo() *PageInfoResolver        { return r.pageInfo }
func (r *CertificateTypePageResolver) Edges() *[]*CertificateTypeResolver { return r.edges }

type CAANumberResolver struct {
	CAANumber gentypes.CAANumber
}

func (c *CAANumberResolver) UUID() gentypes.UUID { return c.CAANumber.UUID }
func (c *CAANumberResolver) CreatedAt() string   { return c.CAANumber.CreatedAt }
func (c *CAANumberResolver) Identifier() string  { return c.CAANumber.Identifier }
func (c *CAANumberResolver) Used() bool          { return c.CAANumber.Used }

type CAANumberPageResolver struct {
	edges    *[]*CAANumberResolver
	pageInfo *PageInfoResolver
}

func (r *CAANumberPageResolver) PageInfo() *PageInfoResolver  { return r.pageInfo }
func (r *CAANumberPageResolver) Edges() *[]*CAANumberResolver { return r.edges }
