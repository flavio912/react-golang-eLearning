package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type TutorResolver struct {
	Tutor gentypes.Tutor
}

type NewTutorArgs struct {
	Tutor     *gentypes.Tutor
	TutorUUID *gentypes.UUID
}

func NewTutorResolver(ctx context.Context, args NewTutorArgs) (*TutorResolver, error) {
	switch {
	case args.TutorUUID != nil:
		app := auth.AppFromContext(ctx)
		tutor, err := app.CourseApp.Tutor(*args.TutorUUID)
		if err != nil {
			return &TutorResolver{}, err
		}

		return &TutorResolver{
			Tutor: tutor,
		}, nil
	case args.Tutor != nil:
		return &TutorResolver{
			Tutor: *args.Tutor,
		}, nil
	default:
		return &TutorResolver{}, &errors.ErrUnableToResolve
	}
}

func (t *TutorResolver) UUID() gentypes.UUID  { return t.Tutor.UUID }
func (t *TutorResolver) Name() string         { return t.Tutor.Name }
func (t *TutorResolver) CIN() string          { return t.Tutor.CIN }
func (t *TutorResolver) SignatureURL() string { return t.Tutor.SignatureURL }
