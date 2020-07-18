package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type Syllabus interface {
	Name() string
	UUID() gentypes.UUID
	Type() gentypes.CourseElement
	Complete() *bool
}

type SyllabusResolver struct {
	Syllabus
}

func (s *SyllabusResolver) ToModule() (*ModuleResolver, bool) {
	c, ok := s.Syllabus.(*ModuleResolver)
	return c, ok
}

func (s *SyllabusResolver) ToLesson() (*LessonResolver, bool) {
	c, ok := s.Syllabus.(*LessonResolver)
	return c, ok
}

func (s *SyllabusResolver) ToTest() (*TestResolver, bool) {
	c, ok := s.Syllabus.(*TestResolver)
	return c, ok
}

type NewSyllabusArgs struct {
	CourseID   *uint
	ModuleUUID *gentypes.UUID
}

func NewSyllabusResolvers(ctx context.Context, args NewSyllabusArgs) (*[]*SyllabusResolver, error) {
	app := auth.AppFromContext(ctx)
	var resolvers []*SyllabusResolver
	switch {
	case args.CourseID != nil:
		syllabus, err := app.CourseApp.CourseSyllabus(*args.CourseID)
		if err != nil {
			return &[]*SyllabusResolver{}, err
		}

		for _, item := range syllabus {
			var (
				res Syllabus
				err error
			)

			switch item.Type {
			case gentypes.TestType:
				res, err = NewTestResolver(ctx, NewTestArgs{TestUUID: &item.UUID})
			case gentypes.ModuleType:
				res, err = NewModuleResolver(ctx, NewModuleArgs{ModuleUUID: &item.UUID})
			case gentypes.LessonType:
				uuid := item.UUID
				res, err = NewLessonResolver(ctx, NewLessonArgs{UUID: &uuid})
			}

			if err != nil {
				return &[]*SyllabusResolver{}, err
			}
			resolvers = append(resolvers, &SyllabusResolver{res})
		}
	case args.ModuleUUID != nil:
		syllabus, err := app.CourseApp.ModuleSyllabus(*args.ModuleUUID)
		if err != nil {
			return &[]*SyllabusResolver{}, err
		}

		for _, item := range syllabus {
			var (
				res Syllabus
				err error
			)

			switch item.Type {
			case gentypes.ModuleTest:
				res, err = NewTestResolver(ctx, NewTestArgs{TestUUID: &item.UUID})
			case gentypes.ModuleLesson:
				uuid := item.UUID
				res, err = NewLessonResolver(ctx, NewLessonArgs{UUID: &uuid})
			}

			if err != nil {
				return &[]*SyllabusResolver{}, err
			}
			resolvers = append(resolvers, &SyllabusResolver{res})
		}
	default:
		return &[]*SyllabusResolver{}, &errors.ErrUnableToResolve
	}

	return &resolvers, nil
}

type SearchSyllabusResultResolver struct {
	edges    *[]*SyllabusResolver
	pageInfo *PageInfoResolver
}

func (r *SearchSyllabusResultResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *SearchSyllabusResultResolver) Edges() *[]*SyllabusResolver { return r.edges }
