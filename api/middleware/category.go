package middleware

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/lib/pq"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
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

func (g *Grant) CreateCategory(ctx context.Context, input gentypes.CreateCategoryInput) (gentypes.Category, error) {
	if !g.IsAdmin {
		return gentypes.Category{}, &errors.ErrUnauthorized
	}

	if ok, err := govalidator.ValidateStruct(input); !ok {
		return gentypes.Category{}, err
	}

	category := models.Category{
		Name:  input.Name,
		Color: input.Color,
	}

	if query := database.GormDB.Create(&category); query.Error != nil {
		if errors.CodeUniqueViolation == query.Error.(*pq.Error).Code {
			return gentypes.Category{}, &errors.ErrCategoryAlreadyExists
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Could not create category")
		return gentypes.Category{}, &errors.ErrWhileHandling
	}

	return categoryToGentype(category), nil
}

func (g *Grant) GetCategoryByUUID(uuid gentypes.UUID) (gentypes.Category, error) {
	var category models.Category
	if query := database.GormDB.Where("uuid = ?", uuid).First(&category); query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Category{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to get by UUID")
		return gentypes.Category{}, &errors.ErrWhileHandling
	}

	return categoryToGentype(category), nil
}
