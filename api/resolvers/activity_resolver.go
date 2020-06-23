package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type ActivityResolver struct {
	activity gentypes.Activity
}

func (a *ActivityResolver) UUID() gentypes.UUID {
	return a.activity.UUID
}

func (a *ActivityResolver) CreatedAt() string {
	return a.activity.CreatedAt
}

func (a *ActivityResolver) Type() gentypes.ActivityType {
	return a.activity.ActivityType
}

func (a *ActivityResolver) Course(ctx context.Context) (*CourseResolver, error) {
	if a.activity.CourseID == nil {
		return nil, nil
	}

	return NewCourseResolver(ctx, NewCourseArgs{
		ID: a.activity.CourseID,
	})
}

type NewActivityPageArgs struct {
	ActivityItems   *[]gentypes.Activity
	CourseTakerUUID *gentypes.UUID
}

type ActivityPageResolver struct {
	edges    *[]*ActivityResolver
	pageInfo *PageInfoResolver
}

func NewActivityPageResolver(ctx context.Context, args NewActivityPageArgs, page *gentypes.Page) (*ActivityPageResolver, error) {
	if args.CourseTakerUUID != nil {
		grant := auth.GrantFromContext(ctx)
		if grant == nil {
			return &ActivityPageResolver{}, &errors.ErrUnauthorized
		}

		usersApp := users.NewUsersApp(grant)
		activityItems, pageInfo, err := usersApp.TakerActivity(*args.CourseTakerUUID, page)
		if err != nil {
			return &ActivityPageResolver{}, err
		}

		var resolvers = make([]*ActivityResolver, len(activityItems))
		for i, activity := range activityItems {
			resolvers[i] = &ActivityResolver{activity: activity}
		}

		return &ActivityPageResolver{
			edges: &resolvers,
			pageInfo: &PageInfoResolver{
				pageInfo: &pageInfo,
			},
		}, nil
	}

	return &ActivityPageResolver{}, &errors.ErrUnableToResolve
}

func (r *ActivityPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *ActivityPageResolver) Edges() *[]*ActivityResolver { return r.edges }
