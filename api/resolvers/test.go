package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

type TestResolver struct {
	test gentypes.Test
}

func (t *TestResolver) Title() string       { return t.test.Name }
func (t *TestResolver) UUID() gentypes.UUID { return t.test.UUID }
func (t *TestResolver) Complete() *bool     { return helpers.BoolPointer(false) }

type NewTestArgs struct {
	TestUUID *gentypes.UUID
}

func NewTestResolver(ctx context.Context, args NewTestArgs) (SyllabusResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return &TestResolver{}, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)

	switch {
	case args.TestUUID != nil:
		test, err := courseApp.Test(*args.TestUUID)
		if err != nil {
			return &TestResolver{}, &errors.ErrUnableToResolve
		}
		return &TestResolver{
			test: test,
		}, nil
	default:
		return &TestResolver{}, &errors.ErrUnableToResolve
	}
}
