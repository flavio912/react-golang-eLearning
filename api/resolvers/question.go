package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type QuestionResolver struct {
	question gentypes.Question
}

func (q *QuestionResolver) Text() string {
	return q.question.Text
}
func (q *QuestionResolver) RandomiseAnswers() *bool {
	return q.question.RandomiseAnswers
}
func (q *QuestionResolver) QuestionType() gentypes.QuestionType {
	return q.question.QuestionType
}
func (q *QuestionResolver) Answers() *[]*AnswerResolver {
	return nil
}

type NewQuestionArgs struct {
	UUID     *gentypes.UUID
	Question *gentypes.Question
}

func NewQuestionResolver(ctx context.Context, args NewQuestionArgs) (*QuestionResolver, error) {
	switch {
	case args.Question != nil:
		return &QuestionResolver{
			question: *args.Question,
		}, nil
	default:
		return &QuestionResolver{}, &errors.ErrUnableToResolve
	}
}
