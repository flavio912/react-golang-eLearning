package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"github.com/golang/glog"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
)

// QueryResolver -
type QueryResolver struct{}

// Info -
func (q *QueryResolver) Info(ctx context.Context) (string, error) {
	glog.Info(auth.GrantFromContext(ctx))
	return "This is the TTC server api", nil
}

// Admins - Get a list of admins
func (q *QueryResolver) Admins(ctx context.Context, args struct{ Page *gentypes.Page }) (*AdminPageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &AdminPageResolver{}, &errors.ErrUnauthorized
	}

	admins, page, err := grant.GetAdmins(args.Page, nil)
	adminResolvers := []*AdminResolver{}
	for _, admin := range admins {
		adminResolvers = append(adminResolvers, &AdminResolver{
			admin: admin,
		})
	}

	return &AdminPageResolver{
		edges: &adminResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}

// Admin gets a single admin
func (q *QueryResolver) Admin(ctx context.Context, args struct{ UUID gentypes.UUID }) (*AdminResolver, error) {
	admin, err := loader.LoadAdmin(ctx, args.UUID.String())

	return &AdminResolver{admin: admin}, err
}

func (q *QueryResolver) Delegate(ctx context.Context, args struct{ UUID gentypes.UUID }) (*DelegateResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &DelegateResolver{}, &errors.ErrUnauthorized
	}

	res, err := NewDelegateResolver(ctx, NewDelegateArgs{UUID: &args.UUID})

	return res, err
}

func (q *QueryResolver) Delegates(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.DelegatesFilter
	OrderBy *gentypes.OrderBy
}) (*DelegatePageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &DelegatePageResolver{}, &errors.ErrUnauthorized
	}

	if args.Filter != nil {
		err := (*args.Filter).Validate()
		if err != nil {
			return &DelegatePageResolver{}, err
		}
	}

	delegates, pageInfo, err := grant.GetDelegates(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &DelegatePageResolver{}, err
	}

	var delegateResolvers []*DelegateResolver
	for _, delegate := range delegates {
		resolver, err := NewDelegateResolver(ctx, NewDelegateArgs{Delegate: delegate})
		if err != nil {
			return &DelegatePageResolver{}, err
		}
		delegateResolvers = append(delegateResolvers, resolver)
	}

	return &DelegatePageResolver{
		edges: &delegateResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		},
	}, nil
}

func (q *QueryResolver) Manager(ctx context.Context, args struct{ UUID *gentypes.UUID }) (*ManagerResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerResolver{}, &errors.ErrUnauthorized
	}

	var managerUUID gentypes.UUID
	if args.UUID != nil {
		managerUUID = *args.UUID
	} else {
		managerUUID = grant.Claims.UUID
	}
	manager, err := loader.LoadManager(ctx, managerUUID.String())
	if err != nil {
		return &ManagerResolver{}, err
	}
	return &ManagerResolver{manager: manager}, nil
}

func (q *QueryResolver) Managers(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.ManagersFilter
	OrderBy *gentypes.OrderBy
}) (*ManagerPageResolver, error) {
	if args.Filter != nil {
		err := (*args.Filter).Validate()
		if err != nil {
			return &ManagerPageResolver{}, err
		}
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerPageResolver{}, &errors.ErrUnauthorized
	}

	managers, page, err := grant.GetManagers(args.Page, args.Filter, args.OrderBy)
	var managerResolvers []*ManagerResolver
	for _, manager := range managers {
		managerResolvers = append(managerResolvers, &ManagerResolver{
			manager: manager,
		})
	}

	return &ManagerPageResolver{
		edges: &managerResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}

func (q *QueryResolver) Companies(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.CompanyFilter
	OrderBy *gentypes.OrderBy
}) (*CompanyPageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyPageResolver{}, &errors.ErrUnauthorized
	}

	companies, page, err := grant.GetCompanyUUIDs(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &CompanyPageResolver{}, err
	}
	return NewCompanyPageResolver(ctx, NewCompanyPageArgs{
		UUIDs: uuidsToStrings(companies),
	}, page)
}

func (q *QueryResolver) Company(ctx context.Context, args struct{ UUID string }) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: args.UUID,
	})
}

func (q *QueryResolver) OnlineCourses(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.OnlineCourseFilter
	OrderBy *gentypes.OrderBy
}) (*OnlineCoursePageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &OnlineCoursePageResolver{}, &errors.ErrUnauthorized
	}

	courses, pageInfo, err := grant.GetOnlineCourses(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &OnlineCoursePageResolver{}, err
	}

	var courseResolvers []*OnlineCourseResolver
	for _, course := range courses {
		resolver, err := NewOnlineCourseResolver(ctx, NewOnlineCourseArgs{OnlineCourse: course})
		if err != nil {
			return &OnlineCoursePageResolver{}, err
		}
		courseResolvers = append(courseResolvers, resolver)
	}

	return &OnlineCoursePageResolver{
		edges: &courseResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		}}, nil

}

func (q *QueryResolver) ClassroomCourses(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.ClassroomCourseFilter
	OrderBy *gentypes.OrderBy
}) (*ClassroomCoursePageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ClassroomCoursePageResolver{}, &errors.ErrUnauthorized
	}

	courses, pageInfo, err := grant.GetClassroomCourses(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &ClassroomCoursePageResolver{}, err
	}

	var courseResolvers []*ClassroomCourseResolver
	for _, course := range courses {
		resolver, err := NewClassroomCourseResolver(ctx, NewClassroomCourseArgs{ClassroomCourse: course})
		if err != nil {
			return &ClassroomCoursePageResolver{}, err
		}
		courseResolvers = append(courseResolvers, resolver)
	}

	return &ClassroomCoursePageResolver{
		edges: &courseResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		}}, nil
}

func (q *QueryResolver) Lesson(ctx context.Context, args struct{ UUID string }) (*LessonResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &LessonResolver{}, &errors.ErrUnauthorized
	}

	return NewLessonResolver(ctx, NewLessonArgs{
		UUID: args.UUID,
	})
}

func (q *QueryResolver) Lessons(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.LessonFilter
	OrderBy *gentypes.OrderBy
}) (*LessonPageResolver, error) {
	if args.Filter != nil {
		err := (*args.Filter).Validate()
		if err != nil {
			return &LessonPageResolver{}, err
		}
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &LessonPageResolver{}, &errors.ErrUnauthorized
	}

	lessons, page, err := grant.GetLessons(args.Page, args.Filter, args.OrderBy)
	var lessonResolvers []*LessonResolver
	for _, lesson := range lessons {
		lessonResolvers = append(lessonResolvers, &LessonResolver{
			Lesson: lesson,
		})
	}

	return &LessonPageResolver{
		edges: &lessonResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}
