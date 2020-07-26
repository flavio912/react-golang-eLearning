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

func (m *MutationResolver) CourseBannerImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	app := auth.AppFromContext(ctx)
	url, successToken, err := app.CourseApp.CourseBannerImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
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

func (m *MutationResolver) SubmitTest(ctx context.Context, args struct{ Input gentypes.SubmitTestInput }) (*gentypes.SubmitTestPayload, error) {
	app := auth.AppFromContext(ctx)

	passed, courseStatus, err := app.CourseApp.SubmitTest(args.Input)
	return &gentypes.SubmitTestPayload{
		Passed:       passed,
		CourseStatus: courseStatus,
	}, err
}

type CreateTestPayload struct {
	Test *TestResolver
}

func (m *MutationResolver) CreateTest(ctx context.Context, args struct{ Input gentypes.CreateTestInput }) (*CreateTestPayload, error) {
	app := auth.AppFromContext(ctx)

	test, err := app.CourseApp.CreateTest(args.Input)
	if err != nil {
		return nil, err
	}

	res, err := NewTestResolver(ctx, NewTestArgs{
		Test: &test,
	})
	return &CreateTestPayload{
		Test: res,
	}, err
}

type UpdateTestPayload struct {
	Test *TestResolver
}

func (m *MutationResolver) UpdateTest(ctx context.Context, args struct{ Input gentypes.UpdateTestInput }) (*UpdateTestPayload, error) {
	app := auth.AppFromContext(ctx)

	test, err := app.CourseApp.UpdateTest(args.Input)
	if err != nil {
		return nil, err
	}

	res, err := NewTestResolver(ctx, NewTestArgs{
		Test: &test,
	})
	return &UpdateTestPayload{
		Test: res,
	}, err
}

type CreateQuestionPayload struct {
	Question *QuestionResolver
}

func (m *MutationResolver) CreateQuestion(ctx context.Context, args struct{ Input gentypes.CreateQuestionInput }) (*CreateQuestionPayload, error) {
	app := auth.AppFromContext(ctx)
	question, err := app.CourseApp.CreateQuestion(args.Input)
	if err != nil {
		return &CreateQuestionPayload{}, err
	}

	res, err := NewQuestionResolver(ctx, NewQuestionArgs{
		Question: &question,
	})

	return &CreateQuestionPayload{
		Question: res,
	}, err
}

type UpdateQuestionPayload struct {
	Question *QuestionResolver
}

func (m *MutationResolver) UpdateQuestion(ctx context.Context, args struct{ Input gentypes.UpdateQuestionInput }) (*UpdateQuestionPayload, error) {
	app := auth.AppFromContext(ctx)
	question, err := app.CourseApp.UpdateQuestion(args.Input)
	if err != nil {
		return &UpdateQuestionPayload{}, err
	}

	res, err := NewQuestionResolver(ctx, NewQuestionArgs{
		Question: &question,
	})

	return &UpdateQuestionPayload{
		Question: res,
	}, err
}

type CreateModulePayload struct {
	Module *ModuleResolver
}

func (m *MutationResolver) CreateModule(ctx context.Context, args struct{ Input gentypes.CreateModuleInput }) (*CreateModulePayload, error) {
	app := auth.AppFromContext(ctx)
	module, err := app.CourseApp.CreateModule(args.Input)
	if err != nil {
		return &CreateModulePayload{}, err
	}

	res, err := NewModuleResolver(ctx, NewModuleArgs{
		Module: &module,
	})

	return &CreateModulePayload{
		Module: res,
	}, err
}

type UpdateModulePayload struct {
	Module *ModuleResolver
}

func (m *MutationResolver) UpdateModule(ctx context.Context, args struct{ Input gentypes.UpdateModuleInput }) (*UpdateModulePayload, error) {
	app := auth.AppFromContext(ctx)
	module, err := app.CourseApp.UpdateModule(args.Input)
	if err != nil {
		return &UpdateModulePayload{}, err
	}

	res, err := NewModuleResolver(ctx, NewModuleArgs{
		Module: &module,
	})

	return &UpdateModulePayload{
		Module: res,
	}, err
}

func (m *MutationResolver) ModuleBannerImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	app := auth.AppFromContext(ctx)
	url, successToken, err := app.CourseApp.ModuleBannerImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) VoiceoverUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	app := auth.AppFromContext(ctx)
	url, successToken, err := app.CourseApp.VoiceoverUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) FulfilPendingOrder(ctx context.Context, args struct{ ClientSecret string }) (bool, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.FulfilPendingOrder(args.ClientSecret)
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
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	blog, err := blogApp.CreateBlog(args.Input)
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
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	url, successToken, err := blogApp.BlogHeaderImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) DeleteTest(ctx context.Context, args struct{ Input gentypes.DeleteTestInput }) (bool, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.DeleteTest(args.Input)
}

func (m *MutationResolver) DeleteQuestion(ctx context.Context, args struct{ Input gentypes.DeleteQuestionInput }) (bool, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.DeleteQuestion(args.Input)
}

func (m *MutationResolver) AnswerImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	app := auth.AppFromContext(ctx)
	url, successToken, err := app.CourseApp.AnswerImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) CreateTutor(ctx context.Context, args struct{ Input gentypes.CreateTutorInput }) (*TutorResolver, error) {
	app := auth.AppFromContext(ctx)

	tutor, err := app.CourseApp.CreateTutor(args.Input)

	return &TutorResolver{
		Tutor: tutor,
	}, err
}

func (m *MutationResolver) UpdateTutor(ctx context.Context, args struct{ Input gentypes.UpdateTutorInput }) (*TutorResolver, error) {
	app := auth.AppFromContext(ctx)

	tutor, err := app.CourseApp.UpdateTutor(args.Input)

	return &TutorResolver{
		Tutor: tutor,
	}, err
}

func (m *MutationResolver) TutorSignatureImageUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	app := auth.AppFromContext(ctx)

	url, successToken, err := app.CourseApp.TutorSignatureImageUploadRequest(args.Input)
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
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	_, err := blogApp.UpdateBlogHeaderImage(args.Input.BlogUUID, args.Input.FileSucess.SuccessToken)
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
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	url, successToken, err := blogApp.BlogBodyImageUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) UpdateTutorSignature(
	ctx context.Context,
	args struct {
		Input gentypes.UpdateTutorSignatureInput
	},
) (*TutorResolver, error) {
	app := auth.AppFromContext(ctx)

	_, err := app.CourseApp.UpdateTutorSignature(args.Input)
	return &TutorResolver{}, err
}

func (m *MutationResolver) UpdateBlog(ctx context.Context, args struct{ Input gentypes.UpdateBlogInput }) (*BlogResolver, error) {
	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &BlogResolver{}, &errors.ErrUnauthorized
	}

	blogApp := application.NewBlogApp(grant)
	blog, err := blogApp.UpdateBlog(args.Input)
	return &BlogResolver{
		Blog: blog,
	}, err
}

func (m *MutationResolver) DeleteModule(ctx context.Context, args struct{ Input gentypes.DeleteModuleInput }) (bool, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.DeleteModule(args.Input)
}

func (m *MutationResolver) DeleteCourse(ctx context.Context, args struct{ Input gentypes.DeleteCourseInput }) (bool, error) {
	app := auth.AppFromContext(ctx)
	return app.CourseApp.DeleteCourse(args.Input)
}

func (m *MutationResolver) RegenerateCertificate(ctx context.Context, args struct {
	Input struct{ HistoricalCourseUUID gentypes.UUID }
}) (bool, error) {
	app := auth.AppFromContext(ctx)
	err := app.CourseApp.RegenerateCertificate(args.Input.HistoricalCourseUUID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MutationResolver) SetCoursePublished(ctx context.Context, args struct {
	CourseID  int32
	Published *bool
}) (bool, error) {
	app := auth.AppFromContext(ctx)
	pub := true
	if args.Published != nil {
		pub = *args.Published
	}

	err := app.CourseApp.SetCoursePublished(uint(args.CourseID), pub)
	if err != nil {
		return false, err
	}
	return true, err
}
