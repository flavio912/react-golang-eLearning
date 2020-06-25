package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// duplicateModule copys a module and its stucture
// TODO: This really isn't nice, make more efficient
func (c *coursesRepoImpl) duplicateModule(tx *gorm.DB, module models.Module, template bool, duplicateStructure bool) (models.Module, error) {
	newModule := models.Module{
		Template:   template,
		TemplateID: &module.UUID,
	}

	query := tx.Create(&newModule)
	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to save duplicated module")
		return models.Module{}, &errors.ErrWhileHandling
	}

	if !duplicateStructure {
		return newModule, nil
	}

	for _, item := range module.Structure {
		structure := models.ModuleStructure{
			ModuleUUID: newModule.UUID,
			LessonUUID: item.LessonUUID,
			TestUUID:   item.TestUUID,
			Rank:       item.Rank,
		}
		query := tx.Create(&structure)
		if query.Error != nil {
			c.Logger.Log(sentry.LevelError, query.Error, "Unable to save module structure while duplicating")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return newModule, nil
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

// GetModuleStructure builds the structure of the module into gentype form
func (c *coursesRepoImpl) GetModuleStructure(moduleUUID gentypes.UUID) (gentypes.CourseItem, error) {
	var moduleChildren []models.ModuleStructure
	query := database.GormDB.Where("module_uuid = ?", moduleUUID).
		Order("rank DESC").
		Find(&moduleChildren)

	if query.Error != nil {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get module structure")
		return gentypes.CourseItem{}, &errors.ErrWhileHandling
	}

	var structure []gentypes.ModuleItem
	for _, child := range moduleChildren {
		if child.LessonUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.LessonType,
				UUID: *child.LessonUUID,
			})
		} else if child.TestUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.TestType,
				UUID: *child.TestUUID,
			})
		} else {
			c.Logger.LogMessage(sentry.LevelError, "Blank Module structure item found")
		}
	}

	return gentypes.CourseItem{
		Type:  gentypes.ModuleType,
		UUID:  moduleUUID,
		Items: structure,
	}, nil
}

// UpdateModuleStructure takes in a transaction, its your job to rollback that transaction if this function returns an error
// or panics
func (c *coursesRepoImpl) UpdateModuleStructure(tx *gorm.DB, moduleItem gentypes.CourseItem, duplicateTemplates bool) (models.Module, error) {

	// Check it is actually a module
	if moduleItem.Type != gentypes.ModuleType {
		c.Logger.LogMessage(sentry.LevelWarning, "UpdateModule given type other than ModuleItem")
		return models.Module{}, &errors.ErrWhileHandling
	}

	var moduleModel models.Module
	query := tx.Where("uuid = ?", moduleItem.UUID).First(&moduleModel)
	if query.Error != nil {
		if query.RecordNotFound() {
			glog.Infof("Could not find uuid: %s", moduleItem.UUID)
			return models.Module{}, &errors.ErrNotFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get module")
		return models.Module{}, &errors.ErrWhileHandling
	}

	// Module templates should be duplicated
	if duplicateTemplates && moduleModel.Template {
		newmod, err := c.duplicateModule(tx, moduleModel, false, false)
		if err != nil {
			return models.Module{}, err
		}
		moduleModel = newmod
		moduleItem.UUID = moduleModel.UUID
	} else {
		query = tx.Where("module_uuid = ?", moduleItem.UUID).Delete(models.ModuleStructure{})
		if query.Error != nil {
			c.Logger.Log(sentry.LevelError, query.Error, "Error deleting old module")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	for _, item := range moduleItem.Items {

		// TODO check if lessons + tests exist
		structureItem := models.ModuleStructure{
			ModuleUUID: moduleItem.UUID,
		}
		if item.Type == gentypes.LessonType {
			structureItem.LessonUUID = &item.UUID
		}
		if item.Type == gentypes.TestType {
			structureItem.TestUUID = &item.UUID
		}

		if err := tx.Save(&structureItem).Error; err != nil {
			c.Logger.Log(sentry.LevelError, err, "Unable to save structure item")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return moduleModel, nil
}