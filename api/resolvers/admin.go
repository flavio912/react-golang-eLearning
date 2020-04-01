package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type AdminResolver struct {
	admin gentypes.Admin
}

func (r *AdminResolver) FirstName() string { return r.admin.FirstName }
func (r *AdminResolver) LastName() string  { return r.admin.LastName }
func (r *AdminResolver) UUID() string      { return r.admin.UUID }
func (r *AdminResolver) Email() string     { return r.admin.Email }

type AdminPageResolver struct {
	edges    *[]*AdminResolver
	pageInfo *PageInfoResolver
}

func (r *AdminPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *AdminPageResolver) Edges() *[]*AdminResolver    { return r.edges }
