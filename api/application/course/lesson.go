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
		UUID:        lesson.UUID,
		Name:        lesson.Name,
		Tags:        tags,
		Description: lesson.Description,
	}
}

func (c *courseAppImpl) lessonsToGentype(lessons []models.Lesson) []gentypes.Lesson {
	var genLessons []gentypes.Lesson
	for _, lesson := range lessons {
		genLessons = append(genLessons, c.lessonToGentype(lesson))
	}
	return genLessons
}

func (c *courseAppImpl) GetLessonsByUUID(uuids []string) ([]gentypes.Lesson, error) {
	if !c.grant.IsAdmin && !c.grant.IsDelegate && !c.grant.IsIndividual {
		return []gentypes.Lesson{}, &errors.ErrUnauthorized
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		// Check user is taking a course with those lessons in it
		var courseTakerID gentypes.UUID
		if c.grant.IsDelegate {
			delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
			courseTakerID = delegate.CourseTakerUUID
		}

		if c.grant.IsIndividual {
			individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
			courseTakerID = individual.CourseTakerUUID
		}

		activeCourses, err := c.usersRepository.TakerActiveCourses(courseTakerID)
		if err != nil {
			return []gentypes.Lesson{}, &errors.ErrWhileHandling
		}

		var courseIds = make([]uint, len(activeCourses))
		for i, activeCourse := range activeCourses {
			courseIds[i] = activeCourse.CourseID
		}

		var gen_uuids = make([]gentypes.UUID, len(uuids))
		for i, uuid := range uuids {
			gen_uuids[i] = gentypes.MustParseToUUID(uuid)
		}

		areLessonsInCourses, err := c.coursesRepository.AreInCourses(courseIds, gen_uuids, gentypes.LessonType)
		if err != nil {
			return []gentypes.Lesson{}, &errors.ErrWhileHandling
		}

		if !areLessonsInCourses {
			return []gentypes.Lesson{}, &errors.ErrWhileHandling
		}
	}

	lessons, err := c.coursesRepository.GetLessonsByUUID(uuids)
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

	b, err := c.coursesRepository.DeleteLesson(input.UUID)
	return b, err
}
