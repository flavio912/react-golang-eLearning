package middleware

import (
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// duplicateModule copys a module and its stucture
// TODO: This really isn't nice, make more efficient
func duplicateModule(tx *gorm.DB, module models.Module, template bool, duplicateStructure bool) (models.Module, error) {
	newModule := models.Module{
		Template:   template,
		TemplateID: &module.UUID,
	}

	query := tx.Create(&newModule)
	if query.Error != nil {
		glog.Error("Unable to save duplicated module")
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
			glog.Error("Unable to save module structure while duplicating")
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return newModule, nil
}

// GetModuleByUUID gets a module by its UUID
func (g *Grant) GetModuleByUUID(moduleUUID string) (models.Module, error) {
	var module models.Module
	query := database.GormDB.Where("uuid = ?", moduleUUID).Find(&module)
	if query.Error != nil {
		if query.RecordNotFound() {
			return module, &errors.ErrNotFound
		}
		glog.Errorf("Error getting module by UUID: %s", query.Error.Error())
		return module, &errors.ErrWhileHandling
	}

	return module, nil
}

// GetModuleStructure builds the structure of the module into gentype form
func (g *Grant) GetModuleStructure(moduleUUID string) (gentypes.CourseItem, error) {
	var moduleChildren []models.ModuleStructure
	query := database.GormDB.Where("module_uuid = ?", moduleUUID).
		Order("rank DESC").
		Find(&moduleChildren)

	if query.Error != nil {
		glog.Errorf("Unable to get module structure: %s", query.Error.Error())
		return gentypes.CourseItem{}, &errors.ErrWhileHandling
	}

	var structure []gentypes.ModuleItem
	for _, child := range moduleChildren {
		if child.LessonUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.LessonType,
				UUID: (*child.LessonUUID).String(),
			})
		} else if child.TestUUID != nil {
			structure = append(structure, gentypes.ModuleItem{
				Type: gentypes.TestType,
				UUID: (*child.TestUUID).String(),
			})
		} else {
			glog.Error("Blank Module structure item found")
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
func (g *Grant) UpdateModuleStructure(tx *gorm.DB, moduleItem gentypes.CourseItem, duplicateTemplates bool) (models.Module, error) {
	if !g.IsAdmin {
		return models.Module{}, &errors.ErrUnauthorized
	}

	// Check it is actually a module
	if moduleItem.Type != gentypes.ModuleType {
		glog.Warning("UpdateModule given type other than ModuleItem")
		return models.Module{}, &errors.ErrWhileHandling
	}

	var moduleModel models.Module
	query := tx.Where("uuid = ?", moduleItem.UUID).First(&moduleModel)
	if query.Error != nil {
		if query.RecordNotFound() {
			glog.Infof("Could not find uuid: %s", moduleItem.UUID)
			return models.Module{}, &errors.ErrNotFound
		}
		glog.Errorf("Unable to get module: %s", moduleItem.UUID)
		return models.Module{}, &errors.ErrWhileHandling
	}

	// Module templates should be duplicated
	if duplicateTemplates && moduleModel.Template {
		newmod, err := duplicateModule(tx, moduleModel, false, false)
		if err != nil {
			return models.Module{}, err
		}
		moduleModel = newmod
		moduleItem.UUID = moduleModel.UUID.String()
	} else {
		query = tx.Where("module_uuid = ?", moduleItem.UUID).Delete(models.ModuleStructure{})
		if query.Error != nil {
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	// ModuleUUID to real uuid
	modUUID, err := uuid.Parse(moduleItem.UUID)
	if err != nil {
		glog.Errorf("ModuleUUID is invalid: %s", err.Error())
		return models.Module{}, &errors.ErrUUIDInvalid
	}

	for _, item := range moduleItem.Items {
		itemUUID, err := uuid.Parse(item.UUID)
		if err != nil {
			glog.Errorf("ItemUUID is invalid: %s", err.Error())
		}

		// TODO check if lessons + tests exist
		structureItem := models.ModuleStructure{
			ModuleUUID: modUUID,
		}
		if item.Type == gentypes.LessonType {
			structureItem.LessonUUID = &itemUUID
		}
		if item.Type == gentypes.TestType {
			structureItem.TestUUID = &itemUUID
		}

		if err := tx.Save(&structureItem).Error; err != nil {
			glog.Errorf("Unable to save structure item: %s", err)
			return models.Module{}, &errors.ErrWhileHandling
		}
	}

	return moduleModel, nil
}
