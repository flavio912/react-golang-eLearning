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
