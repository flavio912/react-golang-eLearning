package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type IndividualResolver struct {
	Individual gentypes.Individual
}

func (i *IndividualResolver) UUID() gentypes.UUID { return i.Individual.UUID }
func (i *IndividualResolver) CreatedAt() *string  { return i.Individual.CreatedAt }
func (i *IndividualResolver) Email() string       { return i.Individual.Email }
func (i *IndividualResolver) FirstName() string   { return i.Individual.FirstName }
func (i *IndividualResolver) LastName() string    { return i.Individual.LastName }
func (i *IndividualResolver) JobTitle() *string   { return i.Individual.JobTitle }
func (i *IndividualResolver) Telephone() *string  { return i.Individual.Telephone }
func (i *IndividualResolver) LastLogin() string   { return i.Individual.LastLogin }
func (i *IndividualResolver) MyCourses(ctx context.Context) (*[]*MyCourseResolver, error) {
	return NewMyCoursesResolvers(ctx, NewMyCoursesArgs{
		TakerUUID: &i.Individual.CourseTakerUUID,
	})
}

type NewIndividualArgs struct {
	Individual     *gentypes.Individual
	IndividualUUID *gentypes.UUID
}

func NewIndividualResolver(ctx context.Context, args NewIndividualArgs) (*IndividualResolver, error) {
	switch {
	case args.IndividualUUID != nil:
		app := auth.AppFromContext(ctx)
		ind, err := app.UsersApp.Individual(*args.IndividualUUID)
		if err != nil {
			return &IndividualResolver{}, err
		}

		return &IndividualResolver{
			Individual: ind,
		}, nil
	case args.Individual != nil:
		return &IndividualResolver{
			Individual: *args.Individual,
		}, nil
	default:
		return &IndividualResolver{}, &errors.ErrUnableToResolve
	}
}

type IndividualPageResolver struct {
	edges    *[]*IndividualResolver
	pageInfo *PageInfoResolver
}

func (r *IndividualPageResolver) PageInfo() *PageInfoResolver   { return r.pageInfo }
func (r *IndividualPageResolver) Edges() *[]*IndividualResolver { return r.edges }

type NewIndividualPageArgs struct {
	PageInfo    gentypes.PageInfo
	Individuals *[]gentypes.Individual
}

func NewIndividualPageResolver(ctx context.Context, args NewIndividualPageArgs) (*IndividualPageResolver, error) {
	if args.Individuals == nil {
		return &IndividualPageResolver{}, &errors.ErrUnableToResolve
	}

	var resolvers []*IndividualResolver
	for _, ind := range *args.Individuals {
		res, err := NewIndividualResolver(ctx, NewIndividualArgs{
			Individual: &ind,
		})

		if err != nil {
			return &IndividualPageResolver{}, err
		}
		resolvers = append(resolvers, res)
	}

	return &IndividualPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}
