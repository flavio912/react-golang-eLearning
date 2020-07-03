package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type AnswerResolver struct {
	answer gentypes.Answer
}

func (a *AnswerResolver) UUID() gentypes.UUID { return a.answer.UUID }
func (a *AnswerResolver) Text() *string       { return a.answer.Text }
func (a *AnswerResolver) ImageURL() *string   { return a.answer.ImageURL }
func (a *AnswerResolver) IsCorrect() *bool    { return a.answer.IsCorrect }

type NewAnswerArgs struct {
	Answer     *gentypes.Answer
	AnswerUUID *gentypes.UUID
}

func NewAnswerResolver(ctx context.Context, args NewAnswerArgs) (*AnswerResolver, error) {
	// app := auth.AppFromContext(ctx)

	switch {
	case args.Answer != nil:
		return &AnswerResolver{
			answer: *args.Answer,
		}, nil
	default:
		return &AnswerResolver{}, &errors.ErrUnableToResolve
	}
}
