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
		return nil, &errors.ErrUnableToResolve
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
