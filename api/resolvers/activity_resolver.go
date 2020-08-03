package resolvers

import (
	"context"

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
func (a *ActivityResolver) User(ctx context.Context) (*UserResolver, error) {
	return NewUserResolver(ctx, NewUserArgs{
		CourseTakerUUID: &a.activity.CourseTakerUUID,
	})
}

type NewActivityPageArgs struct {
	ActivityItems   *[]gentypes.Activity
	CourseTakerUUID *gentypes.UUID
	CompanyUUID     *gentypes.UUID
}

type ActivityPageResolver struct {
	edges    *[]*ActivityResolver
	pageInfo *PageInfoResolver
}

func NewActivityPageResolver(ctx context.Context, args NewActivityPageArgs, page *gentypes.Page) (*ActivityPageResolver, error) {
	var (
		activityItems []gentypes.Activity
		pageInfo      gentypes.PageInfo
	)

	app := auth.AppFromContext(ctx)

	switch {
	case args.CourseTakerUUID != nil:
		takerActivity, takerPageInfo, err := app.UsersApp.TakerActivity(*args.CourseTakerUUID, page)
		if err != nil {
			return &ActivityPageResolver{}, err
		}
		activityItems = takerActivity
		pageInfo = takerPageInfo
	case args.CompanyUUID != nil:
		companyActivity, companyPageInfo, err := app.UsersApp.CompanyActivity(*args.CompanyUUID, page)
		if err != nil {
			return &ActivityPageResolver{}, err
		}

		activityItems = companyActivity
		pageInfo = companyPageInfo
	default:
		return &ActivityPageResolver{}, &errors.ErrUnableToResolve
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

func (r *ActivityPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *ActivityPageResolver) Edges() *[]*ActivityResolver { return r.edges }
