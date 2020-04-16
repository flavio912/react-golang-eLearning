package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type OnlineCourseResolver struct {
	OnlineCourse gentypes.OnlineCourse
}

type NewOnlineCourseArgs struct {
	OnlineCourse gentypes.OnlineCourse
	UUID         gentypes.UUID
}

func NewOnlineCourseResolver(ctx context.Context, args NewOnlineCourseArgs) (*OnlineCourseResolver, error) {
	var (
		course gentypes.OnlineCourse
		err    error
	)

	switch {
	case args.UUID != gentypes.UUID{}:
		//TODO: Load course using loader
	case args.OnlineCourse.UUID != gentypes.UUID{}:
		course = args.OnlineCourse
	default:
		err = &errors.ErrUnableToResolve
	}

	if err != nil {
		return &OnlineCourseResolver{}, err
	}

	return &OnlineCourseResolver{
		OnlineCourse: course,
	}, err
}

func (r *OnlineCourseResolver) UUID() gentypes.UUID { return r.OnlineCourse.UUID }
func (r *OnlineCourseResolver) Info(ctx context.Context) (*CourseInfoResolver, error) {
	return NewCourseInfoResolver(ctx, NewCourseInfoArgs{
		ID: &r.OnlineCourse.CourseInfoID,
	})
}
