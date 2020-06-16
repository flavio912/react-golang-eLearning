package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
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

func (c *courseAppImpl) CreateTag(input gentypes.CreateTagInput) (gentypes.Tag, error) {
	if !c.grant.IsAdmin {
		return gentypes.Tag{}, &errors.ErrUnauthorized
	}

	tag, err := c.coursesRepository.CreateTag(input)
	return tagToGentype(tag), err
}

func (c *courseAppImpl) GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]gentypes.Tag, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Tag{}, &errors.ErrUnauthorized
	}

	tags, err := c.coursesRepository.GetTags(page, filter, orderBy)
	return tagsToGentypes(tags), err
}

func (c *courseAppImpl) GetTagsByCourseInfoIDs(ids []uint) (map[uint][]gentypes.Tag, error) {

	tags, err := c.coursesRepository.GetTagsByCourseInfoIDs(ids)

	var genTags = map[uint][]gentypes.Tag{}
	for key, element := range tags {
		genTags[key] = tagsToGentypes(element)
	}

	return genTags, err
}

func (c *courseAppImpl) GetTagsByLessonUUID(uuid string) ([]gentypes.Tag, error) {
	tags, err := c.coursesRepository.GetTagsByLessonUUID(uuid)
	return tagsToGentypes(tags), err
}
