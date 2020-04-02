package resolvers

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type PageInfoResolver struct {
	pageInfo *gentypes.PageInfo
}

func (r *PageInfoResolver) Total() int32  { return r.pageInfo.Total }
func (r *PageInfoResolver) Offset() int32 { return r.pageInfo.Offset }
func (r *PageInfoResolver) Limit() int32  { return r.pageInfo.Limit }
func (r *PageInfoResolver) Given() int32  { return r.pageInfo.Given }
