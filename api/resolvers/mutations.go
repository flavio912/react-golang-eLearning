package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

// MutationResolver - Root resolver for mutations
type MutationResolver struct{}

// AuthToken - Type used for returning JWT's
type AuthToken struct {
	Token string
}

// AdminLogin - Resolver for getting an authToken
func (m *MutationResolver) AdminLogin(ctx context.Context, args struct{ Input gentypes.AdminLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetAdminAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}
	auth.SetAuthCookies(ctx, token)

	return &gentypes.AuthToken{Token: token}, nil
}

// ManagerLogin - Resolver for getting an authToken
func (m *MutationResolver) ManagerLogin(ctx context.Context, args struct{ Input gentypes.ManagerLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetManagerAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}

	auth.SetAuthCookies(ctx, token)

	// If NoResp given return a blank token in the response - @temmerson
	if args.Input.NoResp != nil && *args.Input.NoResp {
		return &gentypes.AuthToken{Token: ""}, nil
	}

	return &gentypes.AuthToken{Token: token}, nil
}

// CreateManager is for an admin to create new managers manually
func (m *MutationResolver) CreateManager(ctx context.Context, args struct{ Input gentypes.CreateManagerInput }) (*ManagerResolver, error) {
	// Validate the manager input
	if err := args.Input.Validate(); err != nil {
		return &ManagerResolver{}, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	manager, err := usersApp.CreateManager(args.Input)
	if err != nil {
		return &ManagerResolver{}, err
	}

	loadedManager, loadErr := loader.LoadManager(ctx, manager.UUID.String())

	return &ManagerResolver{
		manager: loadedManager,
	}, loadErr
}

// Allows a manager to update their details or an admin
func (m *MutationResolver) UpdateManager(ctx context.Context, args struct{ Input gentypes.UpdateManagerInput }) (*ManagerResolver, error) {
	// Validate the manager input
	if err := args.Input.Validate(); err != nil {
		return &ManagerResolver{}, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	manager, err := usersApp.UpdateManager(args.Input)
	if err != nil {
		return &ManagerResolver{}, err
	}

	return &ManagerResolver{
		manager: manager,
	}, nil
}

func (m *MutationResolver) DeleteManager(ctx context.Context, args struct{ Input gentypes.DeleteManagerInput }) (bool, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	success, err := usersApp.DeleteManager(args.Input.UUID)
	return success, err
}

func (m *MutationResolver) CreateAdmin(ctx context.Context, args struct{ Input gentypes.CreateAdminInput }) (*AdminResolver, error) {
	if err := args.Input.Validate(); err != nil {
		return nil, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return nil, &errors.ErrUnauthorized
	}
	adminFuncs := application.NewAdminApp(grant)

	admin, addErr := adminFuncs.CreateAdmin(args.Input)
	if addErr != nil {
		return nil, addErr
	}

	return &AdminResolver{
		admin: admin,
	}, nil
}

func (m *MutationResolver) UpdateAdmin(ctx context.Context, args struct{ Input gentypes.UpdateAdminInput }) (*AdminResolver, error) {
	if err := args.Input.Validate(); err != nil {
		return nil, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return nil, &errors.ErrUnauthorized
	}

	adminFuncs := application.NewAdminApp(grant)
	admin, err := adminFuncs.UpdateAdmin(args.Input)
	if err != nil {
		return nil, err
	}

	return &AdminResolver{
		admin: admin,
	}, nil
}

func (m *MutationResolver) DeleteAdmin(ctx context.Context, args struct{ Input gentypes.DeleteAdminInput }) (bool, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}
	adminFuncs := application.NewAdminApp(grant)
	success, err := adminFuncs.DeleteAdmin(args.Input.UUID)
	return success, err
}

func (m *MutationResolver) ProfileImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	url, successToken, err := usersApp.ProfileUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) UpdateManagerProfileImage(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileSuccess },
) (*ManagerResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	err := usersApp.ManagerProfileUploadSuccess(args.Input.SuccessToken)
	if err != nil {
		return &ManagerResolver{}, err
	}

	res, err := NewManagerResolver(ctx, NewManagerArgs{
		UUID: grant.Claims.UUID.String(),
	})

	if err != nil {
		return &ManagerResolver{}, err
	}

	return res, nil
}

func (m *MutationResolver) CreateCompany(ctx context.Context, args struct{ Input gentypes.CreateCompanyInput }) (*CompanyResolver, error) {
	if err := args.Input.Validate(); err != nil {
		return &CompanyResolver{}, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	company, err := usersApp.CreateCompany(args.Input)
	if err != nil {
		return &CompanyResolver{}, err
	}

	return NewCompanyResolver(ctx, NewCompanyArgs{Company: company})
}

func (m *MutationResolver) UpdateCompany(ctx context.Context, args struct{ Input gentypes.UpdateCompanyInput }) (*CompanyResolver, error) {
	// Validate the company input
	if err := args.Input.Validate(); err != nil {
		return &CompanyResolver{}, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	company, err := usersApp.UpdateCompany(args.Input)
	if err != nil {
		return &CompanyResolver{}, err
	}

	return NewCompanyResolver(ctx, NewCompanyArgs{Company: company})
}

type companyRequestInput struct {
	Company   gentypes.CreateCompanyInput
	Manager   gentypes.CreateManagerInput
	Recaptcha string
}

// CreateCompanyRequest is used to request that an admin allows you to create company
func (m *MutationResolver) CreateCompanyRequest(ctx context.Context, args companyRequestInput) (bool, error) {
	// TODO: Check recaptcha token
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	success, err := usersApp.CreateCompanyRequest(args.Company, args.Manager)
	return success, err
}

func (m *MutationResolver) ApproveCompany(ctx context.Context, args struct{ UUID gentypes.UUID }) (*CompanyResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	company, err := usersApp.ApproveCompany(args.UUID)
	if err != nil {
		return nil, err
	}

	return NewCompanyResolver(ctx, NewCompanyArgs{
		Company: company,
	})
}

func (m *MutationResolver) SaveOnlineCourse(
	ctx context.Context,
	args struct {
		Input gentypes.SaveOnlineCourseInput
	}) (*CourseResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CourseResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	course, err := courseFuncs.SaveOnlineCourse(args.Input)
	if err != nil {
		return &CourseResolver{}, err
	}

	return NewCourseResolver(ctx, NewCourseArgs{
		Course: &course,
	})
}

func (m *MutationResolver) SaveClassroomCourse(ctx context.Context, args struct {
	Input gentypes.SaveClassroomCourseInput
}) (*CourseResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CourseResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	course, err := courseFuncs.SaveClassroomCourse(args.Input)
	if err != nil {
		return &CourseResolver{}, err
	}

	return NewCourseResolver(ctx, NewCourseArgs{
		Course: &course,
	})
}

func (m *MutationResolver) CreateTag(ctx context.Context, args struct{ Input gentypes.CreateTagInput }) (*TagResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &TagResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	tag, err := courseFuncs.CreateTag(args.Input)
	return &TagResolver{
		Tag: tag,
	}, err
}

func (m *MutationResolver) CreateCategory(ctx context.Context, args struct{ Input gentypes.CreateCategoryInput }) (*CategoryResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CategoryResolver{}, &errors.ErrUnauthorized
	}

	category, err := grant.CreateCategory(ctx, args.Input)
	if err != nil {
		return &CategoryResolver{}, err
	}

	return NewCategoryResolver(ctx, NewCategoryResolverArgs{
		Category: category,
	})
}

func (m *MutationResolver) CreateLesson(ctx context.Context, args struct{ Input gentypes.CreateLessonInput }) (*LessonResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &LessonResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	lesson, err := courseFuncs.CreateLesson(args.Input)
	if err != nil {
		return &LessonResolver{}, err
	}

	return &LessonResolver{
		Lesson: lesson,
	}, err
}

func (m *MutationResolver) PurchaseCourses(ctx context.Context, args struct{ Input gentypes.PurchaseCoursesInput }) (*gentypes.PurchaseCoursesResponse, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.PurchaseCoursesResponse{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	return courseApp.PurchaseCourses(args.Input)
}

type CreateIndividualResponse struct {
	User *UserResolver
}

func (m *MutationResolver) CreateIndividual(ctx context.Context, args struct {
	Input gentypes.CreateIndividualInput
}) (*CreateIndividualResponse, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return nil, &errors.ErrUnauthorized
	}

	usersApp := users.NewUsersApp(grant)
	user, err := usersApp.CreateIndividual(args.Input)
	if err != nil {
		return &CreateIndividualResponse{}, err
	}

	res, err := NewUserResolver(ctx, NewUserArgs{
		User: user,
	})

	return &CreateIndividualResponse{
		User: res,
	}, err
}

func (m *MutationResolver) UpdateLesson(ctx context.Context, args struct{ Input gentypes.UpdateLessonInput }) (*LessonResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &LessonResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	lesson, err := courseFuncs.UpdateLesson(args.Input)
	if err != nil {
		return &LessonResolver{}, err
	}

	return &LessonResolver{
		Lesson: lesson,
	}, nil
}

func (m *MutationResolver) DeleteLesson(ctx context.Context, args struct{ Input gentypes.DeleteLessonInput }) (bool, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	b, err := courseFuncs.DeleteLesson(args.Input)
	if err != nil {
		return false, err
	}

	return b, nil
}

func (m *MutationResolver) CreateBlog(ctx context.Context, args struct{ Input gentypes.CreateBlogInput }) (*BlogResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	courseFuncs := course.NewCourseApp(grant)
	blog, err := courseFuncs.CreateBlog(args.Input)
	if err != nil {
		return &BlogResolver{}, err
	}

	return &BlogResolver{
		Blog: blog,
	}, nil
}

func (m *MutationResolver) BlogHeaderImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	url, successToken, err := courseApp.BlogHeaderImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) UpdateBlogHeaderImage(
	ctx context.Context,
	args struct {
		Input gentypes.UpdateBlogHeaderImageInput
	},
) (*BlogResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	err := courseApp.UpdateBlogHeaderImage(args.Input.BlogUUID, args.Input.FileSucess.SuccessToken)
	if err != nil {
		return &BlogResolver{}, err
	}

	return NewBlogResolver(ctx, NewBlogArgs{
		UUID: args.Input.BlogUUID.String(),
	})
}

func (m *MutationResolver) BlogBodyImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	url, successToken, err := courseApp.BlogBodyImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) UpdateBlog(ctx context.Context, args struct{ Input gentypes.UpdateBlogInput }) (*BlogResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	blog, err := courseApp.UpdateBlog(args.Input)
	return &BlogResolver{
		Blog: blog,
	}, err
}
