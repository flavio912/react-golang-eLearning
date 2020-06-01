package resolvers

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
)

type LessonResolver struct {
	Lesson gentypes.Lesson
}

type NewLessonArgs struct {
	UUID   string
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
	case args.UUID != "":
		lesson, err = loader.LoadLesson(ctx, args.UUID)
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
func (l *LessonResolver) Title() string       { return l.Lesson.Title }
func (l *LessonResolver) Text() string        { return l.Lesson.Text }
func (l *LessonResolver) Tags() []*TagResolver {
	var res []*TagResolver
	for _, tag := range l.Lesson.Tags {
		res = append(res, &TagResolver{
			Tag: tag,
		})
	}
	return res
}

type LessonPageResolver struct {
	edges    *[]*LessonResolver
	pageInfo *PageInfoResolver
}

func (r *LessonPageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *LessonPageResolver) Edges() *[]*LessonResolver   { return r.edges }
