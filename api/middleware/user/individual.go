package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
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

func filterIndividual(query *gorm.DB, filter *gentypes.IndividualFilter) *gorm.DB {
	if filter != nil {
		query = middleware.FilterUser(query, &filter.UserFilter)

		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
	}

	return query
}

func (u *usersRepoImpl) Individuals(page *gentypes.Page, filter *gentypes.IndividualFilter, orderBy *gentypes.OrderBy) ([]models.Individual, gentypes.PageInfo, error) {
	var individuals []models.Individual
	query := filterIndividual(database.GormDB, filter)

	var count int32
	if err := query.Model(&models.Individual{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to count individuals")
		return individuals, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title"}, "created_at DESC")
	if orderErr != nil {
		return individuals, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&individuals)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Individual{}, gentypes.PageInfo{}, &errors.ErrNotAllFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to find individuals")
		return []models.Individual{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return individuals, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(individuals)),
	}, nil
}

// CreateIndividual - PUBLIC
func (u *usersRepoImpl) CreateIndividual(input gentypes.CreateIndividualInput) (models.Individual, error) {
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
		u.Logger.Log(sentry.LevelError, err, "Unable to create individual")
		return models.Individual{}, &errors.ErrWhileHandling
	}

	return newIndividual, nil
}

func (u *usersRepoImpl) UpdateIndividual(input gentypes.UpdateIndividualInput) (models.Individual, error) {
	if ok, err := govalidator.ValidateStruct(input); !ok || err != nil {
		return models.Individual{}, err
	}

	updates := make(map[string]interface{})

	if input.FirstName != nil {
		updates["first_name"] = *input.FirstName
	}
	if input.LastName != nil {
		updates["last_name"] = *input.LastName
	}
	if input.JobTitle != nil {
		updates["job_title"] = *input.JobTitle
	}
	if input.Telephone != nil {
		updates["telephone"] = *input.Telephone
	}
	if input.Email != nil {
		updates["email"] = *input.Email
	}
	if input.Password != nil {
		updates["password"] = *input.Password
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			u.Logger.LogMessage(sentry.LevelFatal, "UpdateIndividual: Forced to recover")
		}
	}()

	if err := tx.Model(&models.Individual{}).Where("uuid = ?", input.UUID).Updates(updates).Error; err != nil {
		u.Logger.Logf(sentry.LevelError, err, "Unable to update individual: %s", input.UUID)
		tx.Rollback()
		return models.Individual{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		tx.Rollback()
		return models.Individual{}, &errors.ErrWhileHandling
	}

	ind, err := u.Individual(input.UUID)
	if err != nil {
		return models.Individual{}, err
	}

	return ind, nil
}
