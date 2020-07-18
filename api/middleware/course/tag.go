package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/lib/pq"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

/* Tags CRUD */

// CheckTagsExist returns a slice of tags if all the given tag uuids are in the database
// If *any* are not found it returns an error and no tags
func (c *coursesRepoImpl) CheckTagsExist(tags []gentypes.UUID) ([]models.Tag, error) {
	var tagModels []models.Tag
	query := database.GormDB.Where("uuid IN (?)", tags).Find(&tagModels)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Error while checking tags exist")
		return []models.Tag{}, &errors.ErrWhileHandling
	}

	if len(tagModels) != len(tags) {
		return []models.Tag{}, &errors.ErrTagsNotFound
	}
	return tagModels, nil
}

// CreateTag makes a new course/module tag
func (c *coursesRepoImpl) CreateTag(input gentypes.CreateTagInput) (models.Tag, error) {
	if _, err := govalidator.ValidateStruct(input); err != nil {
		return models.Tag{}, err
	}

	tag := models.Tag{
		Name:  input.Name,
		Color: input.Color,
	}
	if query := database.GormDB.Create(&tag); query.Error != nil {
		if errors.CodeUniqueViolation == query.Error.(*pq.Error).Code {
			return models.Tag{}, &errors.ErrTagAlreadyExists
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Could not create tag")
		return models.Tag{}, &errors.ErrWhileHandling
	}

	return tag, nil
}

// ManyCourseTags takes a list of courseInfo ids and returns a mapping
// of a courseInfo Id to a slice of tags
func (c *coursesRepoImpl) ManyCourseTags(ids []uint) (map[uint][]models.Tag, error) {
	// TODO: Check if user has access to this particular course

	// Find the table links
	var links []models.CourseTagsLink
	query := database.GormDB.Where("course_id IN (?)", ids).Find(&links)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get course tags links")
		return map[uint][]models.Tag{}, &errors.ErrWhileHandling
	}

	var tagUUIDs []gentypes.UUID
	for _, i := range links {
		tagUUIDs = append(tagUUIDs, i.TagUUID)
	}

	// Get all tags
	var tags []models.Tag
	query = database.GormDB.Where("uuid IN (?)", tagUUIDs).Find(&tags)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get course tags")
		return map[uint][]models.Tag{}, &errors.ErrWhileHandling
	}

	var tagsMap = make(map[gentypes.UUID]models.Tag)
	for _, tag := range tags {
		tagsMap[tag.UUID] = tag
	}

	// Put tags into map: courseIDs > models.Tag
	var courseInfoIdsToTags = make(map[uint][]models.Tag, len(tagUUIDs))
	for _, link := range links {
		id := link.CourseID
		if _, ok := courseInfoIdsToTags[id]; ok {
			courseInfoIdsToTags[id] = append(courseInfoIdsToTags[id], tagsMap[link.TagUUID])
		} else {
			courseInfoIdsToTags[id] = []models.Tag{tagsMap[link.TagUUID]}
		}
	}
	return courseInfoIdsToTags, nil
}

// TODO: Finish func
// GetTags returns a slice of tags
func (c *coursesRepoImpl) GetTags(page gentypes.Page, filter gentypes.GetTagsFilter, orderBy gentypes.OrderBy) ([]models.Tag, error) {
	return []models.Tag{}, nil
}

// GetTagsByLessonUUID returns a slice of tags associated with a given lesson
func (c *coursesRepoImpl) GetTagsByLessonUUID(uuid string) ([]models.Tag, error) {
	var tags []models.Tag
	query := database.GormDB.Table("tags").
		Joins("JOIN lesson_tags_link ON lesson_tags_link.tag_uuid = tags.uuid AND lesson_tags_link.lesson_uuid = ?", uuid).
		Order("name DESC").
		Find(&tags)

	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Tag{}, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to finds tags associated with lesson "+uuid)
		return []models.Tag{}, &errors.ErrWhileHandling
	}

	return tags, nil
}
