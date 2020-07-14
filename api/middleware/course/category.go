package course

import (
	"github.com/jinzhu/gorm"
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
		&gentypes.OrderBy{},
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
