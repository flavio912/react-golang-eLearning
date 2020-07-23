package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type QuestionResolver struct {
	question gentypes.Question
}

func (q *QuestionResolver) UUID() gentypes.UUID { return q.question.UUID }
func (q *QuestionResolver) Text() string {
	return q.question.Text
}
func (q *QuestionResolver) RandomiseAnswers() *bool {
	return q.question.RandomiseAnswers
}
func (q *QuestionResolver) QuestionType() gentypes.QuestionType {
	return q.question.QuestionType
}
func (q *QuestionResolver) Answers(ctx context.Context) (*[]*AnswerResolver, error) {
	app := auth.AppFromContext(ctx)
	answers, err := app.CourseApp.ManyAnswers([]gentypes.UUID{q.question.UUID})
	if err != nil {
		return nil, err
	}

	var res []*AnswerResolver
	for _, ans := range answers[q.question.UUID] {
		answerRes, _ := NewAnswerResolver(ctx, NewAnswerArgs{
			Answer: &ans,
		})
		res = append(res, answerRes)
	}
	return &res, nil
}

type NewQuestionArgs struct {
	UUID     *gentypes.UUID
	Question *gentypes.Question
}

func NewQuestionResolver(ctx context.Context, args NewQuestionArgs) (*QuestionResolver, error) {
	app := auth.AppFromContext(ctx)
	switch {
	case args.Question != nil:
		return &QuestionResolver{
			question: *args.Question,
		}, nil
	case args.UUID != nil:
		question, err := app.CourseApp.Question(*args.UUID)
		return &QuestionResolver{
			question: question,
		}, err
	default:
		return &QuestionResolver{}, &errors.ErrUnableToResolve
	}
}

type QuestionPageResolver struct {
	edges    *[]*QuestionResolver
	pageInfo *PageInfoResolver
}

func (r *QuestionPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *QuestionPageResolver) Edges() *[]*QuestionResolver { return r.edges }

type NewQuestionPageArgs struct {
	PageInfo      gentypes.PageInfo
	Questions     *[]gentypes.Question
	QuestionUUIDs *[]gentypes.UUID
}

func NewQuestionPageResolver(ctx context.Context, args NewQuestionPageArgs) (*QuestionPageResolver, error) {
	var resolvers []*QuestionResolver

	switch {
	case args.Questions != nil:
		for _, question := range *args.Questions {
			res, err := NewQuestionResolver(ctx, NewQuestionArgs{
				Question: &question,
			})

			if err != nil {
				return &QuestionPageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	case args.QuestionUUIDs != nil:
		for _, id := range *args.QuestionUUIDs {
			res, err := NewQuestionResolver(ctx, NewQuestionArgs{
				UUID: &id,
			})

			if err != nil {
				return &QuestionPageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	default:
		return &QuestionPageResolver{}, &errors.ErrUnableToResolve
	}

	return &QuestionPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}
