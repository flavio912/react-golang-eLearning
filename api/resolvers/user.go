package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type UserResolver struct {
	user gentypes.User
}

type NewUserArgs struct {
	User            gentypes.User
	UUID            *gentypes.UUID
	CourseTakerUUID *gentypes.UUID
}

func NewUserResolver(ctx context.Context, args NewUserArgs) (*UserResolver, error) {
	if args.UUID != nil {
		app := auth.AppFromContext(ctx)
		user, err := app.UsersApp.GetCurrentUser() // TODO: Use Dataloaders
		if err != nil {
			return &UserResolver{}, err
		}

		return &UserResolver{user: user}, nil
	}
	if args.CourseTakerUUID != nil {
		app := auth.AppFromContext(ctx)
		usersMap, err := app.UsersApp.UsersFromTakers([]gentypes.UUID{*args.CourseTakerUUID})
		if err != nil {
			return &UserResolver{}, err
		}

		return &UserResolver{user: usersMap[*args.CourseTakerUUID]}, nil
	}

	return &UserResolver{user: args.User}, nil
}

func (u *UserResolver) Type() gentypes.UserType  { return u.user.Type }
func (u *UserResolver) Email() *string           { return u.user.Email }
func (u *UserResolver) FirstName() string        { return u.user.FirstName }
func (u *UserResolver) LastName() string         { return u.user.LastName }
func (u *UserResolver) Telephone() *string       { return u.user.Telephone }
func (u *UserResolver) JobTitle() *string        { return u.user.JobTitle }
func (u *UserResolver) LastLogin() string        { return u.user.LastLogin }
func (u *UserResolver) ProfileImageUrl() *string { return u.user.ProfileImageUrl }
func (u *UserResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: grant.Claims.Company.String(),
	})
}
func (u *UserResolver) Activity(ctx context.Context, args struct{ Page *gentypes.Page }) (*ActivityPageResolver, error) {
	return NewActivityPageResolver(ctx, NewActivityPageArgs{
		CourseTakerUUID: u.user.CourseTakerUUID,
	}, args.Page)
}
func (u *UserResolver) MyCourses(ctx context.Context) (*[]*MyCourseResolver, error) {
	return NewMyCoursesResolvers(ctx, NewMyCoursesArgs{
		TakerUUID: u.user.CourseTakerUUID,
	})
}
func (u *UserResolver) MyActiveCourse(ctx context.Context, args struct{ ID int32 }) (*MyCourseResolver, error) {
	app := auth.AppFromContext(ctx)

	if u.user.CourseTakerUUID == nil {
		return &MyCourseResolver{}, &errors.ErrNotFound
	}

	myCourse, err := app.UsersApp.TakerCourse(*u.user.CourseTakerUUID, uint(args.ID))

	return &MyCourseResolver{
		MyCourse: myCourse,
	}, err
}
