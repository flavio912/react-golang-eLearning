package middleware

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) Individual(uuid gentypes.UUID) (models.Individual, error) {
	// Only individuals themselves and admins can get an individual
	if !g.IsIndividual && !g.IsAdmin {
		return models.Individual{}, &errors.ErrUnauthorized
	}

	if g.IsIndividual && uuid != g.Claims.UUID {
		return models.Individual{}, &errors.ErrUnauthorized
	}

	var individual models.Individual

	query := database.GormDB.Where("uuid = ?", uuid).Find(&individual)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Individual{}, &errors.ErrNotFound
		}
		// If some other error occurs log it
		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin for UUID: %s", uuid)
		return models.Individual{}, &errors.ErrWhileHandling
	}

	return individual, nil
}

// CreateIndividual - PUBLIC
func (g *Grant) CreateIndividual(input gentypes.CreateIndividualInput) (models.Individual, error) {
	if ok, err := govalidator.ValidateStruct(input); !ok {
		return models.Individual{}, err
	}

	newIndividual := models.Individual{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		JobTitle:  input.JobTitle,
		Telephone: input.Telephone,
		Password:  input.Password,
		Email:     input.Email,
	}

	if err := database.GormDB.Create(&newIndividual).Error; err != nil {
		g.Logger.Log(sentry.LevelError, err, "Unable to create individual")
		return models.Individual{}, &errors.ErrWhileHandling
	}

	return newIndividual, nil
}
