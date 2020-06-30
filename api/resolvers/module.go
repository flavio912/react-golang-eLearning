package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

type ModuleResolver struct {
	Module gentypes.Module
}

func (m *ModuleResolver) UUID() gentypes.UUID {
	return m.Module.UUID
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
func (m *ModuleResolver) Syllabus(ctx context.Context) (*[]SyllabusResolver, error) {
	return NewSyllabusResolvers(ctx, NewSyllabusArgs{
		ModuleUUID: &m.Module.UUID,
	})
}

type NewModuleArgs struct {
	Module     *gentypes.Module
	ModuleUUID *gentypes.UUID
}

func NewModuleResolver(ctx context.Context, args NewModuleArgs) (*ModuleResolver, error) {
	app := auth.AppFromContext(ctx)

	switch {
	case args.Module != nil:
		return &ModuleResolver{
			Module: *args.Module,
		}, nil
	case args.ModuleUUID != nil:
		module, err := app.CourseApp.Module(*args.ModuleUUID)
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
