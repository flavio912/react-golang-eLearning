package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
)

type LessonResolver struct {
	Lesson gentypes.Lesson
}

type NewLessonArgs struct {
	UUID   *gentypes.UUID
	Lesson gentypes.Lesson
}

type NewLessonsArgs struct {
	UUIDs []string
}

func NewLessonResolver(ctx context.Context, args NewLessonArgs) (*LessonResolver, error) {
	var (
		lesson gentypes.Lesson
		err    error
	)

	switch {
	case args.UUID != nil:
		lesson, err = loader.LoadLesson(ctx, *args.UUID)
	case args.Lesson.UUID.String() != "":
		lesson = args.Lesson
	default:
		err = &errors.ErrUnableToResolve
	}

	if err != nil {
		return &LessonResolver{}, err
	}

	return &LessonResolver{
		Lesson: lesson,
	}, nil
}

func NewLessonResolvers(ctx context.Context, args NewLessonsArgs) (*[]*LessonResolver, error) {
	results, err := loader.LoadLessons(ctx, args.UUIDs)
	if err != nil {
		return nil, err
	}

	var (
		lessons   = results
		resolvers = make([]*LessonResolver, 0, len(lessons))
	)

	for _, lesson := range lessons {
		if lesson.Error != nil {
			logging.Log(ctx, sentry.LevelWarning, "Lesson resolver error", lesson.Error)
		}

		resolver, err := NewLessonResolver(ctx, NewLessonArgs{Lesson: lesson.Lesson})
		if err != nil {
			glog.Error("Unable to create resolver")
		}

		resolvers = append(resolvers, resolver)
	}

	return &resolvers, nil
}

func (l *LessonResolver) UUID() gentypes.UUID { return l.Lesson.UUID }
func (l *LessonResolver) Name() string        { return l.Lesson.Name }
func (l *LessonResolver) Type() gentypes.CourseElement {
	return gentypes.LessonType
}
func (l *LessonResolver) Description() string     { return l.Lesson.Description }
func (l *LessonResolver) Complete() *bool         { return helpers.BoolPointer(false) } // TODO
func (l *LessonResolver) BannerImageURL() *string { return l.Lesson.BannerImageURL }
func (l *LessonResolver) VoiceoverURL() *string   { return l.Lesson.VoiceoverURL }
func (l *LessonResolver) Transcript() *string     { return l.Lesson.Transcript }
func (l *LessonResolver) Video() *gentypes.Video  { return l.Lesson.Video }

func (l *LessonResolver) Tags(ctx context.Context) (*[]*TagResolver, error) {
	app := auth.AppFromContext(ctx)
	tags, err := app.CourseApp.GetTagsByLessonUUID(l.UUID().String())
	if err != nil {
		glog.Info("Unable to resolve tags")
		return nil, err

	}
	var res []*TagResolver
	for _, tag := range tags {
		res = append(res, &TagResolver{
			Tag: tag,
		})
	}
	return &res, nil
}

type LessonPageResolver struct {
	edges    *[]*LessonResolver
	pageInfo *PageInfoResolver
}

func (r *LessonPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *LessonPageResolver) Edges() *[]*LessonResolver   { return r.edges }
