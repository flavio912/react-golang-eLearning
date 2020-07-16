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

// TakerHasActiveCourse returns true if the given courseTaker has the given courseID as an active course
func (u *usersRepoImpl) TakerHasActiveCourse(courseTaker gentypes.UUID, courseID uint) (bool, error) {
	var items int
	query := database.GormDB.
		Model(&models.ActiveCourse{}).
		Where("course_taker_uuid = ? AND course_id = ?", courseTaker, courseID).
		Count(&items)

	if query.Error != nil {
		if query.RecordNotFound() {
			return false, nil
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable check if taker has active course")
		return false, &errors.ErrWhileHandling
	}

	if items > 0 {
		return true, nil
	}

	return false, nil
}

func (u *usersRepoImpl) TakerActiveCourse(courseTaker gentypes.UUID, courseID uint) (models.ActiveCourse, error) {
	var activeCourse models.ActiveCourse
	query := database.GormDB.
		Model(&models.ActiveCourse{}).
		Where("course_taker_uuid = ? AND course_id = ?", courseTaker, courseID).
		Find(&activeCourse)

	if query.Error != nil {
		if query.RecordNotFound() {
			return models.ActiveCourse{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable get active course")
		return models.ActiveCourse{}, &errors.ErrWhileHandling
	}

	return activeCourse, nil
}

func (u *usersRepoImpl) TakerActiveCourses(courseTaker gentypes.UUID) ([]models.ActiveCourse, error) {
	var activeCourses []models.ActiveCourse
	query := database.GormDB.Where("course_taker_uuid = ?", courseTaker).Find(&activeCourses)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.ActiveCourse{}, nil
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to get taker active courses")
		return []models.ActiveCourse{}, &errors.ErrWhileHandling
	}

	return activeCourses, nil
}

// TakerHasSubmittedTest gets the testMarks for a particular course and taker
func (u *usersRepoImpl) TakerTestMarks(courseTaker gentypes.UUID, courseID uint) ([]models.TestMark, error) {
	var marks []models.TestMark
	query := database.GormDB.Where("course_taker_uuid = ? AND course_id = ?", courseTaker, courseID).Find(&marks)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.TestMark{}, nil
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to get taker course marks")
		return []models.TestMark{}, &errors.ErrWhileHandling
	}

	return marks, nil
}

func (u *usersRepoImpl) CreateHistoricalCourse(course models.HistoricalCourse) (models.HistoricalCourse, error) {
	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Remove any test marks for this user
	query := tx.
		Where("course_taker_uuid = ? AND course_id = ", course.CourseTakerUUID, course.CourseID).
		Delete(models.TestMark{})

	if query.Error != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, query.Error, "CreateHistoricalCourse: Unable to delete testmarks")
		return models.HistoricalCourse{}, &errors.ErrWhileHandling
	}

	if err := tx.Create(&course).Error; err != nil {
		tx.Rollback()
		u.Logger.Log(sentry.LevelError, err, "CreateHistoricalCourse: Unable to create historical course")
		return models.HistoricalCourse{}, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "CreateHistoricalCourse: Unable to commit historical course")
		return models.HistoricalCourse{}, &errors.ErrWhileHandling
	}

	return course, nil
}

type UpdateHistoricalCourseInput struct {
	UUID           gentypes.UUID
	CertificateKey *string
}

func (u *usersRepoImpl) UpdateHistoricalCourse(input UpdateHistoricalCourseInput) error {
	updates := make(map[string]interface{})

	if input.CertificateKey != nil {
		updates["certificate_key"] = input.CertificateKey
	}

	if err := database.GormDB.Where("uuid = ?", input.UUID).Updates(updates).Error; err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to update historical course")
		return &errors.ErrWhileHandling
	}

	return nil
}

func (u *usersRepoImpl) SaveTestMarks(mark models.TestMark) error {
	err := database.GormDB.Save(&mark).Error
	if err != nil {
		u.Logger.Log(sentry.LevelError, err, "Unable to save test marks")
		return &errors.ErrWhileHandling
	}

	return nil
}
