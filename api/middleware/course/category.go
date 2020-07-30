package course

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/dbutils"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (c *coursesRepoImpl) Categories(page *gentypes.Page, text *string) ([]models.Category, gentypes.PageInfo, error) {
	var categories []models.Category
	utils := dbutils.NewDBUtils(c.Logger)

	pageInfo, err := utils.GetPageOf(
		&models.Category{},
		&categories,
		page,
		nil,
		[]string{},
		"created_at DESC",
		func(db *gorm.DB) *gorm.DB {
			if text == nil {
				return db
			}

			return db.Where("name ILIKE ?", text)
		},
	)

	pageInfo.Given = int32(len(categories))

	return categories, pageInfo, err
}

func (c *coursesRepoImpl) UpdateCategory(input gentypes.UpdateCategoryInput) (models.Category, error) {
	updates := make(map[string]interface{})

	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.Color != nil {
		updates["color"] = *input.Color
	}

	err := database.GormDB.Model(models.Category{}).Where("uuid = ?", input.UUID).Updates(updates).Error
	if err != nil {
		return models.Category{}, &errors.ErrWhileHandling
	}

	var category models.Category
	err = database.GormDB.Where("uuid = ?", input.UUID).Find(&category).Error
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "UpdateCategory: Cannot get category after update")
		return category, &errors.ErrWhileHandling
	}

	return category, nil
}

func (c *coursesRepoImpl) DeleteCategory(uuid gentypes.UUID) error {
	err := database.GormDB.Where("uuid = ?", uuid).Delete(models.Category{}).Error

	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "DeleteCategory: Unable to delete category")
		return &errors.ErrWhileHandling
	}

	return nil
}