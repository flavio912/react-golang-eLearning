package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersRepoImpl) Individual(uuid gentypes.UUID) (models.Individual, error) {
	// Only individuals themselves and admins can get an individual
	// if !g.IsIndividual && !g.IsAdmin {
	// 	return models.Individual{}, &errors.ErrUnauthorized
	// }

	// if g.IsIndividual && uuid != g.Claims.UUID {
	// 	return models.Individual{}, &errors.ErrUnauthorized
	// }

	var individual models.Individual

	query := database.GormDB.Where("uuid = ?", uuid).Find(&individual)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Individual{}, &errors.ErrNotFound
		}
		// If some other error occurs log it
		u.Logger.Logf(sentry.LevelError, query.Error, "Unable to find admin for UUID: %s", uuid)
		return models.Individual{}, &errors.ErrWhileHandling
	}

	return individual, nil
}

// CreateIndividual
func (u *usersRepoImpl) CreateIndividual(input gentypes.CreateIndividualInput) (models.Individual, error) {
	if ok, err := govalidator.ValidateStruct(input); !ok {
		return models.Individual{}, err
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	taker, err := u.createCourseTaker(tx)
	if err != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, err, "CreateIndividual: Unable to create course taker")
		return models.Individual{}, &errors.ErrWhileHandling
	}

	newIndividual := models.Individual{
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		JobTitle:        input.JobTitle,
		Telephone:       input.Telephone,
		Password:        input.Password,
		Email:           input.Email,
		CourseTakerUUID: taker.UUID,
	}

	if err := tx.Create(&newIndividual).Error; err != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, err, "Unable to create individual")
		return models.Individual{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "CreateIndividual: Unable to commit")
		return models.Individual{}, &errors.ErrWhileHandling
	}

	return newIndividual, nil
}
