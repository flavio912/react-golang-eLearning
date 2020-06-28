package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type VideoInput struct {
	Type gentypes.VideoType
	URL  string
}

type CreateModuleInput struct {
	Name         string
	Description  string
	Transcript   string
	VoiceoverKey *string
	BannerKey    *string
	Video        *VideoInput
	Tags         *[]gentypes.UUID
	Syllabus     *[]gentypes.ModuleItem
}

func (c *coursesRepoImpl) CreateModule(input CreateModuleInput) (models.Module, error) {

	mod := models.Module{
		Name:         input.Name,
		Description:  input.Description,
		Transcript:   input.Transcript,
		BannerKey:    input.BannerKey,
		VoiceoverKey: input.VoiceoverKey,
	}

	if input.Video != nil {
		mod.VideoType = &(*input.Video).Type
		mod.VideoURL = &(*input.Video).URL
	}

	if input.Tags != nil {
		// Check tags exist
		tags, err := c.CheckTagsExist(*input.Tags)
		if err != nil {
			return models.Module{}, err
		}
		mod.Tags = tags
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&mod).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create module")
		return models.Module{}, &errors.ErrWhileHandling
	}

	if input.Syllabus != nil {
		module, err := c.UpdateModuleStructure(tx, mod.UUID, *input.Syllabus)
		if err := tx.Commit().Error; err != nil {
			c.Logger.Log(sentry.LevelError, err, "Unable to commit module")
			return module, &errors.ErrWhileHandling
		}

		return module, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit module")
		return mod, &errors.ErrWhileHandling
	}

	return mod, nil
}

type UpdateModuleInput struct {
	UUID         gentypes.UUID
	Name         *string
	Description  *string
	Transcript   *string
	VoiceoverKey *string
	BannerKey    *string
	Video        *VideoInput
	Tags         *[]gentypes.UUID
	Syllabus     *[]gentypes.ModuleItem
}

func (c *coursesRepoImpl) UpdateModule(input UpdateModuleInput) (models.Module, error) {
	module, err := c.GetModuleByUUID(input.UUID)
	if err != nil {
		return models.Module{}, &errors.ErrNotFound
	}

	updates := make(map[string]interface{})
	if input.Name != nil && *input.Name != module.Name {
		updates["name"] = *input.Name
	}
	if input.Description != nil && *input.Description != module.Description {
		updates["description"] = *input.Description
	}
	if input.Transcript != nil && *input.Transcript != module.Transcript {
		updates["transcript"] = *input.Transcript
	}
	if input.VoiceoverKey != nil {
		updates["voiceover_key"] = *input.VoiceoverKey
	}
	if input.BannerKey != nil {
		updates["banner_key"] = *input.BannerKey
	}
	if input.Video != nil {
		updates["video_url"] = (*input.Video).URL
		updates["video_type"] = (*input.Video).Type
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&module).Updates(updates).Error; err != nil {
		tx.Rollback()
		c.Logger.Log(sentry.LevelError, err, "Unable to update module")
		return module, &errors.ErrWhileHandling
	}

	if input.Tags != nil {
		tags, err := c.CheckTagsExist(*input.Tags)
		if err != nil {
			return module, err
		}
		if err := tx.Model(&module).Association("Tags").Replace(tags).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to replace tags")
			return module, &errors.ErrWhileHandling
		}
	}

	if input.Syllabus != nil {
		_, err := c.UpdateModuleStructure(tx, module.UUID, *input.Syllabus)
		if err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelWarning, err, "Unable to update module structure")
			return module, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit module update")
		return module, &errors.ErrWhileHandling
	}

	return c.GetModuleByUUID(module.UUID)
}

// GetModuleByUUID gets a module by its UUID
func (c *coursesRepoImpl) GetModuleByUUID(moduleUUID gentypes.UUID) (models.Module, error) {
	var module models.Module
	query := database.GormDB.Where("uuid = ?", moduleUUID).Find(&module)
	if query.Error != nil {
		if query.RecordNotFound() {
			return module, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Error getting module by UUID")
		return module, &errors.ErrWhileHandling
	}

	return module, nil
}

// ManyModuleItems
func (c *coursesRepoImpl) ManyModuleItems(moduleUUIDs []gentypes.UUID) (map[gentypes.UUID][]gentypes.ModuleItem, error) {
	var structures []models.ModuleStructure
	query := database.GormDB.Where("module_uuid IN (?)", moduleUUIDs).
		Order("module_uuid, rank ASC").
		Find(&structures)

	if query.Error != nil {
		if query.RecordNotFound() {
			return map[gentypes.UUID][]gentypes.ModuleItem{}, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get module structures")
		return map[gentypes.UUID][]gentypes.ModuleItem{}, &errors.ErrWhileHandling
	}

	var uuidToModules = make(map[gentypes.UUID][]gentypes.ModuleItem)
	for _, structure := range structures {
		var item gentypes.ModuleItem
		if structure.LessonUUID != nil {
			item = gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: *structure.LessonUUID,
			}
		} else if structure.TestUUID != nil {
			item = gentypes.ModuleItem{
				Type: gentypes.ModuleTest,
				UUID: *structure.TestUUID,
			}
		} else {
			c.Logger.LogMessage(sentry.LevelError, "Blank Module structure item found")
		}
		uuidToModules[structure.ModuleUUID] = append(uuidToModules[structure.ModuleUUID], item)
	}

	return uuidToModules, nil
}

// GetModuleStructure builds the structure of the module into gentype form
func (c *coursesRepoImpl) GetModuleStructure(moduleUUID gentypes.UUID) ([]gentypes.ModuleItem, error) {
	var moduleChildren []models.ModuleStructure
	query := database.GormDB.Where("module_uuid = ?", moduleUUID).
		Order("rank ASC").
		Find(&moduleChildren)

	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get module structure")
		return []gentypes.ModuleItem{}, &errors.ErrWhileHandling
	}

	var structure []gentypes.ModuleItem
	for _, child := range moduleChildren {
		if child.LessonUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.ModuleLesson,
				UUID: *child.LessonUUID,
			})
		} else if child.TestUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.ModuleTest,
				UUID: *child.TestUUID,
			})
		} else {
			c.Logger.LogMessage(sentry.LevelError, "Blank Module structure item found")
		}
	}

	return structure, nil
}

// UpdateModuleStructure takes in a transaction, its your job to rollback that transaction if this function returns an error
// or panics
func (c *coursesRepoImpl) UpdateModuleStructure(tx *gorm.DB, moduleUUID gentypes.UUID, moduleStructure []gentypes.ModuleItem) (models.Module, error) {

	var moduleModel models.Module
	query := tx.Where("uuid = ?", moduleUUID).First(&moduleModel)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Module{}, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get module")
		return models.Module{}, &errors.ErrWhileHandling
	}

	query = tx.Where("module_uuid = ?", moduleUUID).Delete(models.ModuleStructure{})
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Error deleting previous module structure items")
		return models.Module{}, &errors.ErrWhileHandling
	}

	for i, item := range moduleStructure {

		// TODO check if lessons + tests exist
		structureItem := models.ModuleStructure{
			ModuleUUID: moduleUUID,
			Rank:       strconv.Itoa(i),
		}
		if item.Type == gentypes.ModuleLesson {
			structureItem.LessonUUID = &item.UUID
		}
		if item.Type == gentypes.ModuleTest {
			structureItem.TestUUID = &item.UUID
		}

		if err := tx.Save(&structureItem).Error; err != nil {
			c.Logger.Log(sentry.LevelError, err, "Unable to save structure item")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return moduleModel, nil
}
