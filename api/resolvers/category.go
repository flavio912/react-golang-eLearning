package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CategoryResolver struct {
	Category gentypes.Category
}

type NewCategoryResolverArgs struct {
	UUID     gentypes.UUID
	Category gentypes.Category
}

func NewCategoryResolver(ctx context.Context, args NewCategoryResolverArgs) (*CategoryResolver, error) {
	var category gentypes.Category
	var err error
	switch {
	case args.UUID != gentypes.UUID{}:
		//TODO: Load category using loader
		grant := auth.GrantFromContext(ctx)
		if grant == nil {
			err = &errors.ErrUnauthorized
		}

		cat, caterr := grant.GetCategoryByUUID(args.UUID)
		if caterr != nil {
			err = caterr
		}

		category = cat
	case args.Category.UUID != gentypes.UUID{}:
		category = args.Category
	default:
		err = &errors.ErrUnableToResolve
	}

	return &CategoryResolver{
		Category: category,
	}, err
}

func (t *CategoryResolver) UUID() *gentypes.UUID { return helpers.UUIDPointer(t.Category.UUID) }
func (t *CategoryResolver) Name() string         { return t.Category.Name }
func (t *CategoryResolver) Color() string        { return t.Category.Color }
