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

type CategoryPageResolver struct {
	edges    *[]*CategoryResolver
	pageInfo *PageInfoResolver
}

func (r *CategoryPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *CategoryPageResolver) Edges() *[]*CategoryResolver { return r.edges }

type NewCategoryPageArgs struct {
	Categories *[]gentypes.Category
	PageInfo   gentypes.PageInfo
}

func NewCategoryPageResolver(ctx context.Context, args NewCategoryPageArgs) (*CategoryPageResolver, error) {
	var resolvers []*CategoryResolver

	switch {
	case args.Categories != nil:
		for _, category := range *args.Categories {
			res, err := NewCategoryResolver(ctx, NewCategoryResolverArgs{
				Category: category,
			})

			if err != nil {
				return &CategoryPageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	default:
		return &CategoryPageResolver{}, &errors.ErrUnableToResolve
	}

	return &CategoryPageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}
