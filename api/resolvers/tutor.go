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

type TutorPageResolver struct {
	edges    *[]*TutorResolver
	pageInfo *PageInfoResolver
}

type NewTutorPageArgs struct {
	Tutors   *[]gentypes.Tutor
	PageInfo *gentypes.PageInfo
}

func NewTutorPageResolver(ctx context.Context, args NewTutorPageArgs) (*TutorPageResolver, error) {
	var resolvers []*TutorResolver

	switch {
	case args.Tutors != nil:
		for _, t := range *args.Tutors {
			res, err := NewTutorResolver(ctx, NewTutorArgs{
				Tutor: &t,
			})

			if err != nil {
				return &TutorPageResolver{}, err
			}

			resolvers = append(resolvers, res)
		}
	default:
		return &TutorPageResolver{}, &errors.ErrUnableToResolve
	}

	return &TutorPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: args.PageInfo,
		},
	}, nil
}

func (r *TutorPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *TutorPageResolver) Edges() *[]*TutorResolver    { return r.edges }
