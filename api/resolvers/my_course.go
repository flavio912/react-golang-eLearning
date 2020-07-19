package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type MyCourseResolver struct {
	MyCourse gentypes.MyCourse
}

func (a *MyCourseResolver) Course(ctx context.Context) (*CourseResolver, error) {
	return NewCourseResolver(ctx, NewCourseArgs{
		ID: &a.MyCourse.CourseID,
	})
}
func (a *MyCourseResolver) Status(ctx context.Context) gentypes.CourseStatus {
	return a.MyCourse.Status
}
func (a *MyCourseResolver) MinutesTracked() float64 {
	return a.MyCourse.MinutesTracked
}
func (a *MyCourseResolver) EnrolledAt() string {
	return a.MyCourse.CreatedAt
}

type NewMyCoursesArgs struct {
	TakerUUID *gentypes.UUID
}

func NewMyCoursesResolvers(ctx context.Context, args NewMyCoursesArgs) (*[]*MyCourseResolver, error) {
	app := auth.AppFromContext(ctx)

	var resolvers []*MyCourseResolver
	switch {
	case args.TakerUUID != nil:
		myCourses, err := app.UsersApp.TakerCourses(*args.TakerUUID, true)
		if err != nil {
			return &resolvers, err
		}
		for _, course := range myCourses {
			resolvers = append(resolvers, &MyCourseResolver{
				MyCourse: course,
			})
		}
	default:
		return &resolvers, &errors.ErrUnableToResolve
	}

	return &resolvers, nil
}
