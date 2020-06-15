package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func tagToGentype(tag models.Tag) gentypes.Tag {
	return gentypes.Tag{
		UUID:  tag.UUID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func tagsToGentypes(tags []models.Tag) []gentypes.Tag {
	var genTags = make([]gentypes.Tag, len(tags))
	for i, tag := range tags {
		genTags[i] = tagToGentype(tag)
	}
	return genTags
}

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
