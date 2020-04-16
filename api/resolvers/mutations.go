package resolvers

import (
	"context"

	"github.com/asaskevich/govalidator"
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
func (m *MutationResolver) AdminLogin(args struct{ Input gentypes.AdminLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetAdminAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
	}
	return &gentypes.AuthToken{Token: token}, nil
}

// ManagerLogin - Resolver for getting an authToken
func (m *MutationResolver) ManagerLogin(args struct{ Input gentypes.ManagerLoginInput }) (*gentypes.AuthToken, error) {
	token, err := middleware.GetManagerAccessToken(args.Input.Email, args.Input.Password)
	if err != nil {
		return nil, err
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

	manager, err := grant.CreateManager(args.Input)
	if err != nil {
		return &ManagerResolver{}, err
	}

	loadedManager, loadErr := loader.LoadManager(ctx, manager.UUID.String())

	return &ManagerResolver{
		manager: loadedManager,
	}, loadErr
}

func (m *MutationResolver) DeleteManager(ctx context.Context, args struct{ Input gentypes.DeleteManagerInput }) (bool, error) {
	if err := args.Input.Validate(); err != nil {
		return false, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}

	success, err := grant.DeleteManager(args.Input.UUID)
	return success, err
}

func (m *MutationResolver) CreateAdmin(ctx context.Context, args struct{ Input gentypes.CreateAdminInput }) (*AdminResolver, error) {
	if err := args.Input.Validate(); err != nil {
		return nil, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &AdminResolver{}, &errors.ErrUnauthorized
	}

	admin, addErr := grant.CreateAdmin(args.Input)
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

	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return nil, err
	}

	admin, err := grant.UpdateAdmin(args.Input)
	if err != nil {
		return nil, err
	}

	return &AdminResolver{
		admin: admin,
	}, nil
}

func (m *MutationResolver) DeleteAdmin(ctx context.Context, args struct{ Input gentypes.RemoveAdminInput }) (bool, error) {
	// TODO: Move grant + validate boilerplate stuff further up the request
	if err := args.Input.Validate(); err != nil {
		return false, err
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return false, &errors.ErrUnauthorized
	}

	success, err := grant.DeleteAdmin(args.Input.UUID)
	return success, err
}

func (m *MutationResolver) ManagerProfileUploadRequest(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileMeta },
) (*gentypes.UploadFileResp, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &gentypes.UploadFileResp{}, &errors.ErrUnauthorized
	}

	url, successToken, err := grant.ManagerProfileUploadRequest(args.Input)
	return &gentypes.UploadFileResp{
		URL:          url,
		SuccessToken: successToken,
	}, err
}

func (m *MutationResolver) ManagerProfileUploadSuccess(
	ctx context.Context,
	args struct{ Input gentypes.UploadFileSuccess },
) (*ManagerResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &ManagerResolver{}, &errors.ErrUnauthorized
	}

	err := grant.ManagerProfileUploadSuccess(args.Input.SuccessToken)
	if err != nil {
		return &ManagerResolver{}, err
	}

	res, err := NewManagerResolver(ctx, NewManagerArgs{
		UUID: grant.Claims.UUID,
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

	company, err := grant.CreateCompany(args.Input)
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

	err := middleware.CreateCompanyRequest(args.Company, args.Manager)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MutationResolver) ApproveCompany(ctx context.Context, args struct{ UUID string }) (*CompanyResolver, error) {
	if !govalidator.IsUUIDv4(args.UUID) {
		return &CompanyResolver{}, &errors.ErrUUIDInvalid
	}

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &CompanyResolver{}, &errors.ErrUnauthorized
	}

	company, err := grant.ApproveCompany(args.UUID)
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
	}) (*OnlineCourseResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &OnlineCourseResolver{}, &errors.ErrUnauthorized
	}

	course, err := grant.SaveOnlineCourse(args.Input)
	if err != nil {
		return &OnlineCourseResolver{}, err
	}

	return NewOnlineCourseResolver(ctx, NewOnlineCourseArgs{
		OnlineCourse: course,
	})
}

func (m *MutationResolver) SaveClassroomCourse(ctx context.Context, args struct {
	Input gentypes.SaveClassroomCourseInput
}) (*string, error) {
	return nil, nil
}
