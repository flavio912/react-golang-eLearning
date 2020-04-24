package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CourseInfoResolver struct {
	CourseInfo gentypes.CourseInfo
}

type NewCourseInfoArgs struct {
	ID *uint
}

func NewCourseInfoResolver(ctx context.Context, args NewCourseInfoArgs) (*CourseInfoResolver, error) {
	if args.ID != nil {
		// TODO: Use loader
		grant := auth.GrantFromContext(ctx)
		if grant == nil {
			return &CourseInfoResolver{}, &errors.ErrUnauthorized
		}

		info, err := grant.GetCourseInfoFromID(*args.ID)
		if err != nil {
			return &CourseInfoResolver{}, err
		}

		return &CourseInfoResolver{
			CourseInfo: info,
		}, nil
	}

	return &CourseInfoResolver{}, &errors.ErrUnableToResolve
}

func (r *CourseInfoResolver) Name() *string { return helpers.StringPointer(r.CourseInfo.Name) }