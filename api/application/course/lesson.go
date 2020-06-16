package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *courseAppImpl) lessonToGentype(lesson models.Lesson) gentypes.Lesson {
	var tags []gentypes.Tag
	if lesson.Tags != nil {
		tags = tagsToGentypes(lesson.Tags)
	}
	return gentypes.Lesson{
		UUID:  lesson.UUID,
		Title: lesson.Title,
		Tags:  tags,
		Text:  lesson.Text,
	}
}

func (c *courseAppImpl) lessonsToGentype(lessons []models.Lesson) []gentypes.Lesson {
	var genLessons []gentypes.Lesson
	for _, lesson := range lessons {
		genLessons = append(genLessons, c.lessonToGentype(lesson))
	}
	return genLessons
}

func (c *courseAppImpl) GetLessonsByUUID(uuid []string) ([]gentypes.Lesson, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	lessons, err := c.coursesRepository.GetLessonsByUUID(uuid)
	return c.lessonsToGentype(lessons), err
}

func (c *courseAppImpl) CreateLesson(lesson gentypes.CreateLessonInput) (gentypes.Lesson, error) {
	if !c.grant.IsAdmin {
		return gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	lessonMod, err := c.coursesRepository.CreateLesson(lesson)
	return c.lessonToGentype(lessonMod), err
}

func (c *courseAppImpl) GetLessons(
	page *gentypes.Page,
	filter *gentypes.LessonFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Lesson, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Lesson{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	lessons, pageInfo, err := c.coursesRepository.GetLessons(page, filter, orderBy)
	return c.lessonsToGentype(lessons), pageInfo, err
}

func (c *courseAppImpl) UpdateLesson(input gentypes.UpdateLessonInput) (gentypes.Lesson, error) {
	if !c.grant.IsAdmin {
		return gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	lesson, err := c.coursesRepository.UpdateLesson(input)
	return c.lessonToGentype(lesson), err
}

func (c *courseAppImpl) DeleteLesson(input gentypes.DeleteLessonInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	b, err := c.coursesRepository.DeleteLesson(input)
	return b, err
}
