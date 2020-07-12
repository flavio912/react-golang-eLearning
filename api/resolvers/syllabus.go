package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type SyllabusResolver interface {
	Name() string
	UUID() gentypes.UUID
	Complete() *bool
}

type NewSyllabusArgs struct {
	CourseID   *uint
	ModuleUUID *gentypes.UUID
}

func NewSyllabusResolvers(ctx context.Context, args NewSyllabusArgs) (*[]SyllabusResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &[]SyllabusResolver{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	var resolvers []SyllabusResolver
	switch {
	case args.CourseID != nil:
		syllabus, err := courseApp.CourseSyllabus(*args.CourseID)
		if err != nil {
			return &[]SyllabusResolver{}, &errors.ErrUnableToResolve
		}

		for _, item := range syllabus {
			if item.Type == gentypes.TestType {
				res, err := NewTestResolver(ctx, NewTestArgs{TestUUID: &item.UUID})
				if err != nil {
					return &[]SyllabusResolver{}, &errors.ErrUnableToResolve
				}
				resolvers = append(resolvers, res)
			}
		}
	case args.ModuleUUID != nil:
		return &[]SyllabusResolver{}, &errors.ErrUnableToResolve
	default:
		return &[]SyllabusResolver{}, &errors.ErrUnableToResolve
	}

	return &resolvers, nil
}

type SearchSyllabusResultResolver struct {
	edges    *[]SyllabusResolver
	pageInfo *PageInfoResolver
}

func (r *SearchSyllabusResultResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *SearchSyllabusResultResolver) Edges() *[]SyllabusResolver  { return r.edges }
