package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type ManagerResolver struct {
	manager gentypes.Manager
}

func (m *ManagerResolver) UUID() string      { return m.manager.UUID.String() }
func (m *ManagerResolver) Email() string     { return m.manager.Email }
func (m *ManagerResolver) FirstName() string { return m.manager.FirstName }
func (m *ManagerResolver) LastName() string  { return m.manager.LastName }
func (m *ManagerResolver) Telephone() string { return m.manager.Telephone }
func (m *ManagerResolver) JobTitle() string  { return m.manager.JobTitle }
func (m *ManagerResolver) LastLogin() string { return m.manager.LastLogin }
func (m *ManagerResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: m.manager.CompanyID.String(),
	})
}

type ManagerPageResolver struct {
	edges    *[]*ManagerResolver
	pageInfo *PageInfoResolver
}

func (r *ManagerPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *ManagerPageResolver) Edges() *[]*ManagerResolver  { return r.edges }
