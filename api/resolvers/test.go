package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
)

type TestResolver struct {
	test gentypes.Test
}

func (t *TestResolver) Name() string        { return t.test.Name }
func (t *TestResolver) UUID() gentypes.UUID { return t.test.UUID }
func (t *TestResolver) Type() gentypes.CourseElement {
	return gentypes.TestType
}
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
		return helpers.Int32Pointer(int32(*t.test.QuestionsToAnswer))
	}
	return nil
}
func (t *TestResolver) RandomiseAnswers() *bool { return t.test.RandomiseAnswers }
func (t *TestResolver) Questions(ctx context.Context) (*[]*QuestionResolver, error) {
	app := auth.AppFromContext(ctx)
	questions, err := app.CourseApp.TestQuestions(t.test.UUID)
	if err != nil {
		return &([]*QuestionResolver{}), err
	}

	var resolvers = make([]*QuestionResolver, len(questions))
	for i, question := range questions {
		res, err := NewQuestionResolver(ctx, NewQuestionArgs{
			Question: &question,
		})
		if err != nil {
			return &([]*QuestionResolver{}), err
		}
		resolvers[i] = res
	}

	return &resolvers, nil
}

type NewTestArgs struct {
	Test     *gentypes.Test
	TestUUID *gentypes.UUID
}

func NewTestResolver(ctx context.Context, args NewTestArgs) (*TestResolver, error) {
	switch {
	case args.TestUUID != nil:
		test, err := loader.LoadTest(ctx, *args.TestUUID)
		if err != nil {
			return &TestResolver{}, err
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

type TestPageResolver struct {
	edges    *[]*TestResolver
	pageInfo *PageInfoResolver
}

func (r *TestPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *TestPageResolver) Edges() *[]*TestResolver     { return r.edges }

type NewTestPageArgs struct {
	PageInfo gentypes.PageInfo
	Tests    *[]gentypes.Test
}

func NewTestPageResolver(ctx context.Context, args NewTestPageArgs) (*TestPageResolver, error) {
	var resolvers []*TestResolver

	switch {
	case args.Tests != nil:
		for _, test := range *args.Tests {
			res, err := NewTestResolver(ctx, NewTestArgs{
				Test: &test,
			})

			if err != nil {
				return &TestPageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	default:
		return &TestPageResolver{}, &errors.ErrUnableToResolve
	}

	return &TestPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}
