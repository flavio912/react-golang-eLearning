package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type LessonResolver struct {
	Lesson gentypes.Lesson
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
