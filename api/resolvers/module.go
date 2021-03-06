package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
)

type ModuleResolver struct {
	Module gentypes.Module
}

func (m *ModuleResolver) UUID() gentypes.UUID {
	return m.Module.UUID
}
func (m *ModuleResolver) Type() gentypes.CourseElement {
	return gentypes.ModuleType
}
func (m *ModuleResolver) Name() string {
	return m.Module.Name
}
func (m *ModuleResolver) BannerImageURL() *string {
	return m.Module.BannerImageURL
}
func (m *ModuleResolver) Description() string {
	return m.Module.Description
}
func (m *ModuleResolver) Transcript() string {
	return m.Module.Transcript
}
func (m *ModuleResolver) VoiceoverURL() *string {
	return m.Module.VoiceoverURL
}
func (m *ModuleResolver) Video() *gentypes.Video {
	return m.Module.Video
}
func (m *ModuleResolver) Complete() *bool {
	return helpers.BoolPointer(false)
}
func (m *ModuleResolver) Syllabus(ctx context.Context) (*[]*SyllabusResolver, error) {
	return NewSyllabusResolvers(ctx, NewSyllabusArgs{
		ModuleUUID: &m.Module.UUID,
	})
}
func (m *ModuleResolver) Tags(ctx context.Context) ([]*TagResolver, error) {
	app := auth.AppFromContext(ctx)
	modulesToTags, err := app.CourseApp.ManyModuleTags([]gentypes.UUID{m.Module.UUID})
	if err != nil {
		return NewTagsResolver([]gentypes.Tag{}), err
	}
	tags := modulesToTags[m.Module.UUID]

	return NewTagsResolver(tags), nil
}

type NewModuleArgs struct {
	Module     *gentypes.Module
	ModuleUUID *gentypes.UUID
}

func NewModuleResolver(ctx context.Context, args NewModuleArgs) (*ModuleResolver, error) {
	switch {
	case args.Module != nil:
		return &ModuleResolver{
			Module: *args.Module,
		}, nil
	case args.ModuleUUID != nil:
		module, err := loader.LoadModule(ctx, *args.ModuleUUID)
		if err != nil {
			return &ModuleResolver{}, err
		}
		return &ModuleResolver{
			Module: module,
		}, nil
	default:
		return &ModuleResolver{}, &errors.ErrUnableToResolve
	}
}

type ModulePageResolver struct {
	edges    *[]*ModuleResolver
	pageInfo *PageInfoResolver
}

func (r *ModulePageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *ModulePageResolver) Edges() *[]*ModuleResolver   { return r.edges }

type NewModulePageArgs struct {
	PageInfo    gentypes.PageInfo
	Modules     *[]gentypes.Module
	ModuleUUIDs *[]gentypes.UUID
}

func NewModulePageResolver(ctx context.Context, args NewModulePageArgs) (*ModulePageResolver, error) {
	var resolvers []*ModuleResolver

	switch {
	case args.Modules != nil:
		for _, module := range *args.Modules {
			res, err := NewModuleResolver(ctx, NewModuleArgs{
				Module: &module,
			})

			if err != nil {
				return &ModulePageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	case args.ModuleUUIDs != nil:
		for _, id := range *args.ModuleUUIDs {
			res, err := NewModuleResolver(ctx, NewModuleArgs{
				ModuleUUID: &id,
			})

			if err != nil {
				return &ModulePageResolver{}, err
			}
			resolvers = append(resolvers, res)
		}
	default:
		return &ModulePageResolver{}, &errors.ErrUnableToResolve
	}

	return &ModulePageResolver{
		edges: &resolvers,
		pageInfo: &PageInfoResolver{
			pageInfo: &args.PageInfo,
		},
	}, nil
}
