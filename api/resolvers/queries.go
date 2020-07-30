package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
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
		return nil, &errors.ErrUnauthorized
	}

	adminFuncs := application.NewAdminApp(grant)
	admins, page, err := adminFuncs.PageAdmins(args.Page, nil)
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
func (q *QueryResolver) Admin(ctx context.Context, args struct{ UUID *gentypes.UUID }) (*AdminResolver, error) {
	app := auth.AppFromContext(ctx)

	admin, err := app.AdminApp.Admin(args.UUID)
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

	usersApp := users.NewUsersApp(grant)
	delegates, pageInfo, err := usersApp.GetDelegates(args.Page, args.Filter, args.OrderBy)
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

	usersApp := users.NewUsersApp(grant)
	managers, page, err := usersApp.GetManagers(args.Page, args.Filter, args.OrderBy)
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
	app := auth.AppFromContext(ctx)
	companies, page, err := app.UsersApp.GetCompanyUUIDs(args.Page, args.Filter, args.OrderBy)
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

func (q *QueryResolver) User(ctx context.Context) (*UserResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &UserResolver{}, &errors.ErrUnauthorized
	}
	return NewUserResolver(ctx, NewUserArgs{
		UUID: &grant.Claims.UUID,
	})
}

func (q *QueryResolver) Lesson(ctx context.Context, args struct{ UUID gentypes.UUID }) (*LessonResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &LessonResolver{}, &errors.ErrUnauthorized
	}

	return NewLessonResolver(ctx, NewLessonArgs{
		UUID: &args.UUID,
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

	courseFuncs := course.NewCourseApp(grant)
	lessons, page, err := courseFuncs.GetLessons(args.Page, args.Filter, args.OrderBy)
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

func (q *QueryResolver) Courses(ctx context.Context, args struct {
	Page    *gentypes.Page
	Filter  *gentypes.CourseFilter
	OrderBy *gentypes.OrderBy
}) (*CoursePageResolver, error) {
	app := auth.AppFromContext(ctx)
	courses, page, err := app.CourseApp.GetCourses(args.Page, args.Filter, args.OrderBy)

	var courseResolvers []*CourseResolver
	for _, course := range courses {
		courseResolvers = append(courseResolvers, &CourseResolver{
			Course: course,
		})
	}

	return &CoursePageResolver{
		edges: &courseResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}

func (q *QueryResolver) Course(ctx context.Context, args struct{ ID int32 }) (*CourseResolver, error) {
	id := uint(args.ID)
	return NewCourseResolver(ctx, NewCourseArgs{
		ID: &id,
	})
}

func (q *QueryResolver) Question(ctx context.Context, args struct{ UUID gentypes.UUID }) (*QuestionResolver, error) {
	return NewQuestionResolver(ctx, NewQuestionArgs{
		UUID: &args.UUID,
	})
}

func (q *QueryResolver) Questions(
	ctx context.Context,
	args struct {
		Page    *gentypes.Page
		Filter  *gentypes.QuestionFilter
		OrderBy *gentypes.OrderBy
	}) (*QuestionPageResolver, error) {
	app := auth.AppFromContext(ctx)
	questions, pageInfo, err := app.CourseApp.Questions(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &QuestionPageResolver{}, err
	}

	return NewQuestionPageResolver(ctx, NewQuestionPageArgs{
		PageInfo:  pageInfo,
		Questions: &questions,
	})
}

func (q *QueryResolver) CertificateInfo(ctx context.Context, args struct{ Token string }) (gentypes.CertficateInfo, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.CertificateInfo(args.Token)
}

func (q *QueryResolver) Test(ctx context.Context, args struct{ UUID gentypes.UUID }) (*TestResolver, error) {
	return NewTestResolver(ctx, NewTestArgs{
		TestUUID: &args.UUID,
	})
}

func (q *QueryResolver) Tests(
	ctx context.Context,
	args struct {
		Page    *gentypes.Page
		Filter  *gentypes.TestFilter
		OrderBy *gentypes.OrderBy
	}) (*TestPageResolver, error) {
	app := auth.AppFromContext(ctx)
	tests, pageInfo, err := app.CourseApp.Tests(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &TestPageResolver{}, err
	}

	return NewTestPageResolver(ctx, NewTestPageArgs{
		PageInfo: pageInfo,
		Tests:    &tests,
	})
}

func (q *QueryResolver) Module(ctx context.Context, args struct{ UUID gentypes.UUID }) (*ModuleResolver, error) {
	return NewModuleResolver(ctx, NewModuleArgs{
		ModuleUUID: &args.UUID,
	})
}

func (q *QueryResolver) Blog(ctx context.Context, args struct{ UUID string }) (*BlogResolver, error) {
	return NewBlogResolver(ctx, NewBlogArgs{
		UUID: args.UUID,
	})
}

func (q *QueryResolver) Blogs(ctx context.Context, args struct {
	Page    *gentypes.Page
	OrderBy *gentypes.OrderBy
}) (*BlogPageResolver, error) {
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogPageResolver{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	blogs, page, err := blogApp.GetBlogs(args.Page, args.OrderBy)
	var blogsResolvers []*BlogResolver
	for _, blog := range blogs {
		blogsResolvers = append(blogsResolvers, &BlogResolver{
			Blog: blog,
		})
	}

	return &BlogPageResolver{
		edges: &blogsResolvers,
		pageInfo: &PageInfoResolver{
			&page,
		},
	}, err
}

func (q *QueryResolver) Modules(
	ctx context.Context,
	args struct {
		Page    *gentypes.Page
		Filter  *gentypes.ModuleFilter
		OrderBy *gentypes.OrderBy
	}) (*ModulePageResolver, error) {
	app := auth.AppFromContext(ctx)
	modules, pageInfo, err := app.CourseApp.Modules(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &ModulePageResolver{}, err
	}

	return NewModulePageResolver(ctx, NewModulePageArgs{
		PageInfo: pageInfo,
		Modules:  &modules,
	})
}

func (q *QueryResolver) SearchSyllabus(
	ctx context.Context,
	args struct {
		Page   *gentypes.Page
		Filter *gentypes.SyllabusFilter
	}) (*SearchSyllabusResultResolver, error) {
	app := auth.AppFromContext(ctx)

	results, pageInfo, err := app.CourseApp.SearchSyllabus(args.Page, args.Filter)

	var syllabusResolvers []*SyllabusResolver

	for _, res := range results {
		switch res.Type {
		case gentypes.ModuleType:
			m, _ := NewModuleResolver(ctx, NewModuleArgs{
				ModuleUUID: &res.UUID,
			})
			syllabusResolvers = append(syllabusResolvers, &SyllabusResolver{m})
		case gentypes.LessonType:
			l, _ := NewLessonResolver(ctx, NewLessonArgs{
				UUID: &res.UUID,
			})
			syllabusResolvers = append(syllabusResolvers, &SyllabusResolver{l})
		case gentypes.TestType:
			t, _ := NewTestResolver(ctx, NewTestArgs{
				TestUUID: &res.UUID,
			})
			syllabusResolvers = append(syllabusResolvers, &SyllabusResolver{t})
		}
	}

	return &SearchSyllabusResultResolver{
		edges: &syllabusResolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		},
	}, err
}

func (q *QueryResolver) Categories(
	ctx context.Context,
	args struct {
		Page *gentypes.Page
		Text *string
	}) (*CategoryPageResolver, error) {
	app := auth.AppFromContext(ctx)
	categories, pageInfo, err := app.CourseApp.Categories(args.Page, args.Text)
	if err != nil {
		return &CategoryPageResolver{}, err
	}

	return NewCategoryPageResolver(ctx, NewCategoryPageArgs{
		PageInfo:   pageInfo,
		Categories: &categories,
	})
}

func (q *QueryResolver) CertificateTypes(
	ctx context.Context,
	args struct {
		Page   *gentypes.Page
		Filter *gentypes.CertificateTypeFilter
	}) (*CertificateTypePageResolver, error) {
	app := auth.AppFromContext(ctx)
	certTypes, pageInfo, err := app.CourseApp.CertificateTypes(args.Page, args.Filter)
	if err != nil {
		return &CertificateTypePageResolver{}, err
	}

	return NewCertificateTypePageResolver(ctx, NewCertificateTypePageArgs{
		CertificateTypes: &certTypes,
		PageInfo:         &pageInfo,
	})
}

func (q *QueryResolver) CertificateType(ctx context.Context, args struct{ UUID gentypes.UUID }) (*CertificateTypeResolver, error) {
	return NewCertificateTypeResolver(ctx, NewCertificateTypeArgs{CertificateTypeUUID: &args.UUID})
}

func (q *QueryResolver) CAANumbers(
	ctx context.Context,
	args struct {
		Page   *gentypes.Page
		Filter *gentypes.CAANumberFilter
	}) (*CAANumberPageResolver, error) {
	app := auth.AppFromContext(ctx)
	numbers, pageInfo, err := app.CourseApp.CAANumbers(args.Page, args.Filter)

	if err != nil {
		return &CAANumberPageResolver{}, err
	}

	var resolvers []*CAANumberResolver
	for _, no := range numbers {
		resolvers = append(resolvers, &CAANumberResolver{
			CAANumber: no,
		})
	}

	return &CAANumberPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &pageInfo,
		},
	}, nil
}

func (q *QueryResolver) Individual(ctx context.Context, args struct{ UUID gentypes.UUID }) (*IndividualResolver, error) {
	return NewIndividualResolver(ctx, NewIndividualArgs{
		IndividualUUID: &args.UUID,
	})
}

func (q *QueryResolver) Individuals(
	ctx context.Context,
	args struct {
		Page    *gentypes.Page
		Filter  *gentypes.IndividualFilter
		OrderBy *gentypes.OrderBy
	}) (*IndividualPageResolver, error) {
	app := auth.AppFromContext(ctx)
	inds, pageInfo, err := app.UsersApp.Individuals(args.Page, args.Filter, args.OrderBy)
	if err != nil {
		return &IndividualPageResolver{}, err
	}

	return NewIndividualPageResolver(ctx, NewIndividualPageArgs{
		Individuals: &inds,
		PageInfo:    pageInfo,
	})
}

func (q *QueryResolver) Tutor(ctx context.Context, args struct{ UUID gentypes.UUID }) (*TutorResolver, error) {
	return NewTutorResolver(ctx, NewTutorArgs{
		TutorUUID: &args.UUID,
	})
}

func (q *QueryResolver) Tutors(
	ctx context.Context,
	args struct {
		Page    *gentypes.Page
		Filter  *gentypes.TutorFilter
		OrderBy *gentypes.OrderBy
	}) (*TutorPageResolver, error) {
	app := auth.AppFromContext(ctx)
	tutors, pageInfo, err := app.CourseApp.Tutors(args.Page, args.Filter, args.OrderBy)

	if err != nil {
		return &TutorPageResolver{}, err
	}

	return NewTutorPageResolver(ctx, NewTutorPageArgs{
		Tutors:   &tutors,
		PageInfo: &pageInfo,
	})
}

func (q *QueryResolver) Category(
	ctx context.Context, args struct{ UUID gentypes.UUID }) (*CategoryResolver, error) {
	return NewCategoryResolver(ctx, NewCategoryResolverArgs{
		UUID: args.UUID,
	})
}
