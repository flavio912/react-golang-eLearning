package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type ClassroomCourseResolver struct {
	ClassroomCourse gentypes.ClassroomCourse
}

type NewClassroomCourseArgs struct {
	ClassroomCourse gentypes.ClassroomCourse
	UUID            gentypes.UUID
}

func NewClassroomCourseResolver(ctx context.Context, args NewClassroomCourseArgs) (*ClassroomCourseResolver, error) {
	var (
		course gentypes.ClassroomCourse
		err    error
	)

	switch {
	case args.UUID != gentypes.UUID{}:
		//TODO: Load course using loader
	case args.ClassroomCourse.UUID != gentypes.UUID{}:
		course = args.ClassroomCourse
	default:
		err = &errors.ErrUnableToResolve
	}

	if err != nil {
		return &ClassroomCourseResolver{}, err
	}

	return &ClassroomCourseResolver{
		ClassroomCourse: course,
	}, err
}

func (r *ClassroomCourseResolver) UUID() gentypes.UUID { return r.ClassroomCourse.UUID }
func (r *ClassroomCourseResolver) Info(ctx context.Context) (*CourseInfoResolver, error) {
	return NewCourseInfoResolver(ctx, NewCourseInfoArgs{
		ID: &r.ClassroomCourse.CourseInfoID,
	})
}

type ClassroomCoursePageResolver struct {
	edges    *[]*ClassroomCourseResolver
	pageInfo *PageInfoResolver
}

func (r *ClassroomCoursePageResolver) PageInfo() *PageInfoResolver        { return r.pageInfo }
func (r *ClassroomCoursePageResolver) Edges() *[]*ClassroomCourseResolver { return r.edges }
