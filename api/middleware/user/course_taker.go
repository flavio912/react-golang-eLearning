package user

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (u *usersRepoImpl) CourseTakers(uuids []gentypes.UUID) ([]models.CourseTaker, error) {
	var courseTakers []models.CourseTaker
	if err := database.GormDB.Where("uuid IN (?)", uuids).Find(&courseTakers).Error; err != nil {
		return []models.CourseTaker{}, &errors.ErrWhileHandling
	}

	if len(courseTakers) != len(uuids) {
		return courseTakers, &errors.ErrNotAllFound
	}
	return courseTakers, nil
}

func (u *usersRepoImpl) TakerActivity(courseTaker gentypes.UUID, page *gentypes.Page) ([]models.CourseTakerActivity, gentypes.PageInfo, error) {
	activityItems, pageInfo, err := u.TakerActivitys([]gentypes.UUID{courseTaker}, page)
	if err != nil && err != &errors.ErrNotFound {
		return []models.CourseTakerActivity{}, gentypes.PageInfo{}, err
	}

	return activityItems, pageInfo, nil
}

func (u *usersRepoImpl) TakerActivitys(courseTakers []gentypes.UUID, page *gentypes.Page) ([]models.CourseTakerActivity, gentypes.PageInfo, error) {

	query := database.GormDB.Where("course_taker_uuid IN (?)", courseTakers)

	// Get total before paging
	var count int32
	if err := query.Model(models.CourseTakerActivity{}).Count(&count).Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to count takers activity")
		return []models.CourseTakerActivity{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	var limit = count
	var offset int32
	if page != nil {
		query, limit, offset = middleware.GetPage(query, page)
	}

	var activityItems []models.CourseTakerActivity
	err := query.Find(&activityItems).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []models.CourseTakerActivity{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, err, "Unable to get takers activity")
		return []models.CourseTakerActivity{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return activityItems, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(activityItems)),
	}, nil
}

func (u *usersRepoImpl) CreateTakerActivity(courseTaker gentypes.UUID, activityType gentypes.ActivityType, relatedCourseID *uint) (models.CourseTakerActivity, error) {
	activityItem := models.CourseTakerActivity{
		CourseTakerUUID: courseTaker,
		ActivityType:    activityType,
		CourseID:        relatedCourseID,
	}

	err := database.GormDB.Create(&activityItem).Error
	if err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to create taker activity")
		return models.CourseTakerActivity{}, &errors.ErrWhileHandling
	}

	return activityItem, nil
}

func (u *usersRepoImpl) DeleteTakerActivity(activityUUID gentypes.UUID) error {
	err := database.GormDB.Where("uuid = ?", activityUUID).Delete(&models.CourseTakerActivity{}).Error
	if err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to delete taker activity")
		return &errors.ErrWhileHandling
	}

	return nil
}
