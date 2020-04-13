package resolvers

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/golang/glog"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
)

type ManagerResolver struct {
	manager gentypes.Manager
}

type NewManagerArgs struct {
	UUID    string
	Manager gentypes.Manager
}

type NewManagersArgs struct {
	UUIDs []string
}

func NewManagerResolver(ctx context.Context, args NewManagerArgs) (*ManagerResolver, error) {
	var (
		manager gentypes.Manager
		err     error
	)

	switch {
	case args.UUID != "":
		manager, err = loader.LoadManager(ctx, args.UUID)
	case args.Manager.UUID.String() != "":
		manager = args.Manager
	default:
		err = &errors.ErrUnableToResolve
	}

	if err != nil {
		return &ManagerResolver{}, err
	}

	return &ManagerResolver{
		manager: manager,
	}, nil
}

func NewManagerResolvers(ctx context.Context, args NewManagersArgs) (*[]*ManagerResolver, error) {
	results, err := loader.LoadManagers(ctx, args.UUIDs)
	if err != nil {
		return nil, err
	}

	var (
		managers  = results
		resolvers = make([]*ManagerResolver, 0, len(managers))
	)

	// var errs []error
	for _, manager := range managers {
		if manager.Error != nil {
			logging.Log(ctx, sentry.LevelWarning, "Manager resolver error", manager.Error)
			// TODO: Add errors to each item correctly
			// errs = append(errs, errors.WithIndex(err, i))
		}

		resolver, err := NewManagerResolver(ctx, NewManagerArgs{Manager: manager.Manager})
		if err != nil {
			glog.Error("Unable to create resolver")
		}

		resolvers = append(resolvers, resolver)
	}
	return &resolvers, nil
}

func (m *ManagerResolver) UUID() string             { return m.manager.UUID.String() }
func (m *ManagerResolver) CreatedAt() *string       { return m.manager.CreatedAt }
func (m *ManagerResolver) Email() string            { return m.manager.Email }
func (m *ManagerResolver) FirstName() string        { return m.manager.FirstName }
func (m *ManagerResolver) LastName() string         { return m.manager.LastName }
func (m *ManagerResolver) Telephone() string        { return m.manager.Telephone }
func (m *ManagerResolver) JobTitle() string         { return m.manager.JobTitle }
func (m *ManagerResolver) LastLogin() string        { return m.manager.LastLogin }
func (m *ManagerResolver) ProfileImageURL() *string { return m.manager.ProfileImageURL }
func (m *ManagerResolver) Company(ctx context.Context) (*CompanyResolver, error) {
	return NewCompanyResolver(ctx, NewCompanyArgs{
		UUID: m.manager.CompanyID.String(),
	})
}

type ManagerPageResolver struct {
	edges    *[]*ManagerResolver
	pageInfo *PageInfoResolver
}

func (r *ManagerPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *ManagerPageResolver) Edges() *[]*ManagerResolver  { return r.edges }
