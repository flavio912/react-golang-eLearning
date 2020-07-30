package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
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

func (c *courseAppImpl) UpdateCategory(input gentypes.UpdateCategoryInput) (gentypes.Category, error) {
	if !c.grant.IsAdmin {
		return gentypes.Category{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Category{}, err
	}

	category, err := c.coursesRepository.UpdateCategory(input)
	return categoryToGentype(category), err
}

func (c *courseAppImpl) DeleteCategory(input gentypes.DeleteCategoryInput) error {
	if !c.grant.IsAdmin {
		return &errors.ErrUnauthorized
	}

	err := c.coursesRepository.DeleteCategory(input.UUID)
	if err != nil {
		return &errors.ErrDeleteFailed
	}

	return nil
}
