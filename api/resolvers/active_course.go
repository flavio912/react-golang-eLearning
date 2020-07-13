package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type ActiveCourseResolver struct {
	ActiveCourse gentypes.ActiveCourse
}

func (a *ActiveCourseResolver) Course(ctx context.Context) (*CourseResolver, error) {
	return NewCourseResolver(ctx, NewCourseArgs{
		ID: &a.ActiveCourse.CourseID,
	})
}
func (a *ActiveCourseResolver) CurrentAttempt() int32 {
	return int32(a.ActiveCourse.CurrentAttempt)
}
func (a *ActiveCourseResolver) MinutesTracked() float64 {
	return a.ActiveCourse.MinutesTracked
}

type NewActiveCourseArgs struct {
	TakerUUID *gentypes.UUID
}

func NewActiveCoursesResolvers(ctx context.Context, args NewActiveCourseArgs) (*[]*ActiveCourseResolver, error) {
	app := auth.AppFromContext(ctx)

	var resolvers []*ActiveCourseResolver
	switch {
	case args.TakerUUID != nil:
		activeCourses, err := app.UsersApp.MyActiveCourses()
		if err != nil {
			return &resolvers, err
		}
		for _, course := range activeCourses {
			resolvers = append(resolvers, &ActiveCourseResolver{
				ActiveCourse: course,
			})
		}
	default:
		return &resolvers, &errors.ErrUnableToResolve
	}

	return &resolvers, nil
}
