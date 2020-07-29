package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func categoryToGentype(category models.Category) gentypes.Category {
	return gentypes.Category{
		UUID:  category.UUID,
		Name:  category.Name,
		Color: category.Color,
	}
}

func categoriesToGentypes(categories []models.Category) []gentypes.Category {
	var genCat = make([]gentypes.Category, len(categories))
	for i, category := range categories {
		genCat[i] = categoryToGentype(category)
	}

	return genCat
}

func (c *courseAppImpl) Categories(page *gentypes.Page, text *string) ([]gentypes.Category, gentypes.PageInfo, error) {
	// Public function

	categories, pageInfo, err := c.coursesRepository.Categories(page, text)
	return categoriesToGentypes(categories), pageInfo, err
}
