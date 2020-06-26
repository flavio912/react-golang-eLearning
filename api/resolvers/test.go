package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

type TestResolver struct {
	test gentypes.Test
}

func (t *TestResolver) Title() string         { return t.test.Name }
func (t *TestResolver) UUID() gentypes.UUID   { return t.test.UUID }
func (t *TestResolver) Complete() *bool       { return helpers.BoolPointer(false) }
func (t *TestResolver) Tags() *[]*TagResolver { return nil }
func (t *TestResolver) AttemptsAllowed() *int32 {
	if t.test.AttemptsAllowed != nil {
		return helpers.Int32Pointer(int32(*t.test.AttemptsAllowed))
	}
	return nil
}
func (t *TestResolver) PassPercentage() *float64 { return t.test.PassPercentage }
func (t *TestResolver) QuestionsToAnswer() *int32 {
	if t.test.QuestionsToAnswer != nil {
		return helpers.Int32Pointer(int32(*t.test.AttemptsAllowed))
	}
	return nil
}
func (t *TestResolver) RandomiseAnswers() *bool         { return t.test.RandomiseAnswers }
func (t *TestResolver) Questions() *[]*QuestionResolver { return nil }

type NewTestArgs struct {
	Test     *gentypes.Test
	TestUUID *gentypes.UUID
}

func NewTestResolver(ctx context.Context, args NewTestArgs) (*TestResolver, error) {
	app := auth.AppFromContext(ctx)

	switch {
	case args.TestUUID != nil:
		test, err := app.CourseApp.Test(*args.TestUUID)
		if err != nil {
			return &TestResolver{}, &errors.ErrUnableToResolve
		}
		return &TestResolver{
			test: test,
		}, nil
	case args.Test != nil:
		return &TestResolver{
			test: *args.Test,
		}, nil
	default:
		return &TestResolver{}, &errors.ErrUnableToResolve
	}
}
