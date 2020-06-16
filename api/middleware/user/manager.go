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

func (u *usersRepoImpl) managerEmailExists(email string) bool {
	query := database.GormDB.Where("email = ?", email).First(&models.Manager{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		u.Logger.Logf(sentry.LevelError, query.Error, "Unable to find manager for Email: %s", email)
		return false
	}
	return true
}

func (u *usersRepoImpl) GetManagersByUUID(uuids []gentypes.UUID) ([]models.Manager, error) {
	var managers []models.Manager
	// if !g.IsAdmin && !g.IsManager {
	// 	return managers, &errors.ErrUnauthorized
	// }

	// var allowedUUIDs []string
	// // Managers can only get their own info
	// if g.IsManager {
	// 	for _, uuid := range uuids {
	// 		if g.Claims.UUID.String() == uuid {
	// 			allowedUUIDs = append(allowedUUIDs, uuid)
	// 		}
	// 	}
	// 	if len(uuids) > 0 && len(allowedUUIDs) == 0 {
	// 		return managers, &errors.ErrUnauthorized
	// 	}
	// }

	// // Admin can get any manager info
	// if g.IsAdmin {
	// 	allowedUUIDs = uuids
	// }

	query := database.GormDB.Where("uuid IN (?)", uuids).Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return managers, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to find managers")
		return managers, &errors.ErrWhileHandling
	}

	return managers, nil
}

func filterManager(query *gorm.DB, filter *gentypes.ManagersFilter) *gorm.DB {
	if filter != nil {
		query = middleware.FilterUser(query, &filter.UserFilter)

		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
	}

	return query
}

func (u *usersRepoImpl) Manager(UUID gentypes.UUID) (models.Manager, error) {
	// Admins can get any manager data
	// Managers can only get their own uuid
	// if g.IsAdmin || (g.IsManager && UUID == g.Claims.UUID) {

	var manager models.Manager
	query := database.GormDB.Where("uuid = ?", UUID).First(&manager)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Manager{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to get manager")
		return models.Manager{}, &errors.ErrWhileHandling
	}

	return manager, nil
}

func (u *usersRepoImpl) GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]models.Manager, gentypes.PageInfo, error) {
	// if !g.IsAdmin {
	// 	return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	// }

	var managers []models.Manager

	// Count the total filtered dataset
	var count int32
	query := filterManager(database.GormDB, filter)
	countErr := query.Model(&models.Manager{}).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		u.Logger.Log(sentry.LevelError, countErr, "Unable to count managers")
		return []models.Manager{}, gentypes.PageInfo{}, countErr
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title"}, "created_at DESC")
	if orderErr != nil {
		return []models.Manager{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)
	query = query.Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []models.Manager{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to find managers")
		return []models.Manager{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return managers, gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(managers)),
	}, nil
}

func (u *usersRepoImpl) CreateManager(managerDetails gentypes.CreateManagerInput, companyUUID gentypes.UUID) (models.Manager, error) {

	// check manager does not exist
	if u.managerEmailExists(managerDetails.Email) {
		return models.Manager{}, &errors.ErrUserExists
	}

	// Check if company exists
	if !u.CompanyExists(companyUUID) {
		return models.Manager{}, &errors.ErrCompanyNotFound
	}

	// TODO: Validate input better and return useful details
	manager := models.Manager{
		FirstName:   managerDetails.FirstName,
		LastName:    managerDetails.LastName,
		JobTitle:    managerDetails.JobTitle,
		Telephone:   managerDetails.Telephone,
		Password:    managerDetails.Password,
		Email:       managerDetails.Email,
		CompanyUUID: companyUUID,
	}
	createErr := database.GormDB.Create(&manager).Error
	if createErr != nil {
		u.Logger.Log(sentry.LevelError, createErr, "Unable to to create manager")
		return models.Manager{}, &errors.ErrWhileHandling
	}

	return manager, nil
}

func (u *usersRepoImpl) UpdateManager(input gentypes.UpdateManagerInput) (models.Manager, error) {

	var manager models.Manager
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&manager)
	if query.Error != nil {
		if query.RecordNotFound() {
			return models.Manager{}, &errors.ErrManagerNotFound
		}
		u.Logger.Logf(sentry.LevelError, query.Error, "Unable to find manager to update with UUID: %s", input.UUID)
		return models.Manager{}, &errors.ErrWhileHandling
	}

	if input.Email != nil {
		manager.Email = *input.Email
	}
	if input.FirstName != nil {
		manager.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		manager.LastName = *input.LastName
	}
	if input.Telephone != nil {
		manager.Telephone = *input.Telephone
	}
	if input.JobTitle != nil {
		manager.JobTitle = *input.JobTitle
	}

	save := database.GormDB.Save(&manager)
	if save.Error != nil {
		u.Logger.Logf(sentry.LevelError, save.Error, "Error updating manager with UUID: %s", input.UUID)
		return models.Manager{}, &errors.ErrWhileHandling
	}

	return manager, nil
}

func (u *usersRepoImpl) DeleteManager(uuid gentypes.UUID) (bool, error) {
	query := database.GormDB.Where("uuid = ?", uuid).Delete(models.Manager{})
	if query.Error != nil {
		u.Logger.Logf(sentry.LevelError, query.Error, "Error deleting manager with UUID: %s", uuid)
		return false, &errors.ErrWhileHandling
	}

	if query.RowsAffected == 0 {
		return false, &errors.ErrUserNotFound
	}
	return true, nil
}

func (u *usersRepoImpl) UpdateManagerProfileKey(managerUUID gentypes.UUID, newKey *string) error {
	query := database.GormDB.Model(&models.Manager{}).Where("uuid = ?", managerUUID).Update("profile_key", newKey)
	if query.Error != nil {
		if query.RecordNotFound() {
			return &errors.ErrManagerNotFound
		}

		u.Logger.Log(sentry.LevelError, query.Error, "Unable to update manager profile")
		return &errors.ErrWhileHandling
	}
	return nil
}

// ProfileUploadRequest generates a link that lets users upload a profile image to S3 directly
// Used by all user types
// func (g *Grant) ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
// 	if !g.IsManager && !g.IsAdmin {
// 		return "", "", &errors.ErrUnauthorized
// 	}

// 	url, successToken, err := uploads.GenerateUploadURL(
// 		imageMeta.FileType,      // The actual file type
// 		imageMeta.ContentLength, // The actual file content length
// 		[]string{"jpg", "png"},  // Allowed file types
// 		int32(20000000),         // Max file size = 20MB
// 		"profile",               // Save files in the "profile" s3 directory
// 		"profileImage",          // Unique identifier for this type of upload request
// 	)

// 	return url, successToken, err
// }

// ManagerProfileUploadSuccess checks the successToken and sets the profile image of the current manager
// func (g *Grant) ManagerProfileUploadSuccess(token string) error {
// 	if !g.IsManager {
// 		return &errors.ErrUnauthorized
// 	}

// 	s3Key, err := uploads.VerifyUploadSuccess(token, "profileImage")
// 	if err != nil {
// 		return err
// 	}

// 	query := database.GormDB.Model(&models.Manager{}).Where("uuid = ?", g.Claims.UUID).Update("profile_key", s3Key)
// 	if query.Error != nil {
// 		if query.RecordNotFound() {
// 			return &errors.ErrManagerNotFound
// 		}

// 		g.Logger.Log(sentry.LevelError, query.Error, "Unable to update manager profile")
// 		return &errors.ErrWhileHandling
// 	}

// 	return nil
// }
