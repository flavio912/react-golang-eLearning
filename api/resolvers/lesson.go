package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type LessonResolver struct {
	Lesson gentypes.Lesson
}

func (l *LessonResolver) UUID() gentypes.UUID  { return l.Lesson.UUID }
func (l *LessonResolver) Title() string        { return l.Lesson.Title }
func (l *LessonResolver) Text() string         { return l.Lesson.Text }
func (l *LessonResolver) Tags() []gentypes.Tag { return l.Lesson.Tags }
