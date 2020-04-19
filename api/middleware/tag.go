package middleware

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Tags CRUD */

// CheckTagsExist returns nil if all the given tag uuids are in the database
func CheckTagsExist(tags []gentypes.UUID) ([]models.Tag, error) {
	var tagModels []models.Tag
	query := database.GormDB.Where("uuid IN (?)", tags).Find(&tagModels)
	if query.Error != nil {
		glog.Errorf("Error while checking tags exist: %s", query.Error.Error())
		return []models.Tag{}, &errors.ErrWhileHandling
	}

	if len(tagModels) != len(tags) {
		return []models.Tag{}, &errors.ErrTagsNotFound
	}
	return tagModels, nil
}

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

// CreateTag makes a new course/module tag
func (g *Grant) CreateTag(input gentypes.CreateTagInput) (gentypes.Tag, error) {
	if !g.IsAdmin {
		return gentypes.Tag{}, &errors.ErrUnauthorized
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return gentypes.Tag{}, err
	}

	tag := models.Tag{
		Name:  input.Name,
		Color: input.Color,
	}
	if query := database.GormDB.Create(&tag); query.Error != nil {
		if errors.CodeUniqueViolation == query.Error.(*pq.Error).Code {
			return gentypes.Tag{}, &errors.ErrTagAlreadyExists
		}
		glog.Errorf("Could not create tag: %s", query.Error.Error())
		return gentypes.Tag{}, &errors.ErrWhileHandling
	}

	return tagToGentype(tag), nil
}

// GetTagsByCourseInfoIDs takes a list of courseInfo ids and returns a mapping
// of courseInfo ids to their respective Tags
func (g *Grant) GetTagsByCourseInfoIDs(ids []uint) (map[uint][]gentypes.Tag, error) {
	// TODO: Check if user has access to this particular course
	if !g.IsAdmin {
		return map[uint][]gentypes.Tag{}, &errors.ErrUnauthorized
	}

	// Find the table links
	var links []models.CourseTagsLink
	query := database.GormDB.Where("course_info_id IN (?)", ids).Find(&links)
	if query.Error != nil {
		glog.Errorf("Unable to get course tags links: %s", query.Error.Error())
		return map[uint][]gentypes.Tag{}, &errors.ErrWhileHandling
	}

	// Put into a map for easy searching later
	var tagUUIDsToIds = make(map[gentypes.UUID]uint, len(links))
	var tagUUIDs []gentypes.UUID
	for _, i := range links {
		tagUUIDsToIds[i.TagUUID] = i.CourseInfoID
		tagUUIDs = append(tagUUIDs, i.TagUUID)
	}

	// Get all tags
	var tags []models.Tag
	query = database.GormDB.Where("uuid IN (?)", tagUUIDs).Find(&tags)
	if query.Error != nil {
		glog.Errorf("Unable to get course tags: %s", query.Error.Error())
		return map[uint][]gentypes.Tag{}, &errors.ErrWhileHandling
	}

	var tagsMap = make(map[gentypes.UUID]models.Tag)
	for _, tag := range tags {
		tagsMap[tag.UUID] = tag
	}

	// Put tags into map: courseIDs > gentypes.Tag
	var courseInfoIdsToTags = make(map[uint][]gentypes.Tag, len(tagUUIDs))
	for _, link := range links {
		id := link.CourseInfoID
		if _, ok := courseInfoIdsToTags[id]; ok {
			courseInfoIdsToTags[id] = append(courseInfoIdsToTags[id], tagToGentype(tagsMap[link.TagUUID]))
		} else {
			courseInfoIdsToTags[id] = []gentypes.Tag{tagToGentype(tagsMap[link.TagUUID])}
		}
	}
	return courseInfoIdsToTags, nil
}
