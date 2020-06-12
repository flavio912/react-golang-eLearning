package middleware

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) CourseTakers(ids []uint) ([]models.CourseTaker, error) {
	var courseTakers []models.CourseTaker
	if err := database.GormDB.Where("id IN (?)", ids).Find(&courseTakers).Error; err != nil {
		return []models.CourseTaker{}, &errors.ErrWhileHandling
	}

	if len(courseTakers) != len(ids) {
		return courseTakers, &errors.ErrNotAllFound
	}
	// TODO: Check if the current user is allowed to get these courseTakers

	return courseTakers, nil
}
