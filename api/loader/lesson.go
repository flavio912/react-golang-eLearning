package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"

	"github.com/graph-gophers/dataloader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type lessonLoader struct {
}

func sortLessons(lessons []gentypes.Lesson, keys dataloader.Keys) []gentypes.Lesson {
	var (
		k         = keys.Keys()
		lessonMap = map[string]gentypes.Lesson{}
		sorted    = make([]gentypes.Lesson, len(k))
	)

	for _, lesson := range lessons {
		lessonMap[lesson.UUID.String()] = lesson
	}

	for i, key := range keys {
		sorted[i] = lessonMap[key.String()]
	}

	return sorted
}

// loadBatch loads a batch of lessons via dataloader
func (l *lessonLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return loadBatchError(&errors.ErrUnauthorized, n)
	}

	courseFuncs := course.NewCourseApp(grant)
	lessons, err := courseFuncs.GetLessonsByUUID(keys.Keys())
	if err != nil {
		return loadBatchError(err, n)
	}

	lessons = sortLessons(lessons, keys)
	res := make([]*dataloader.Result, n)
	for i, lesson := range lessons {
		res[i] = &dataloader.Result{Data: lesson}
	}
	return res
}

// LoadLesson loads Lesson via dataloader
func LoadLesson(ctx context.Context, uuid gentypes.UUID) (gentypes.Lesson, error) {
	var lesson gentypes.Lesson
	data, err := extractAndLoad(ctx, lessonLoaderKey, uuid.String())
	if err != nil {
		return lesson, err
	}

	lesson, ok := data.(gentypes.Lesson)
	if !ok {
		return lesson, fmt.Errorf("Wrong type: %T", data)
	}

	return lesson, nil
}

type LessonResult struct {
	Lesson gentypes.Lesson
	Error  error
}

func LoadLessons(ctx context.Context, uuids []string) ([]LessonResult, error) {
	ldr, err := extract(ctx, lessonLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(uuids))()

	results := make([]LessonResult, 0, len(uuids))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		lesson, ok := d.(gentypes.Lesson)
		if !ok && e == nil {
			e = fmt.Errorf("Wrong type: %T", lesson)
		}

		results = append(results, LessonResult{Lesson: lesson, Error: e})
	}

	return results, nil
}
