package middleware

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) managerToGentype(manager models.Manager) gentypes.Manager {
	profileURL := uploads.GetImgixURL(manager.ProfileKey)
	// Admins and managers themselves can get all info
	if g.IsAdmin || (g.IsManager && g.Claims.Company == manager.CompanyUUID) {
		createdAt := manager.CreatedAt.Format(time.RFC3339)
		return gentypes.Manager{
			CreatedAt:       &createdAt,
			UUID:            manager.UUID,
			FirstName:       manager.FirstName,
			LastName:        manager.LastName,
			JobTitle:        manager.JobTitle,
			Telephone:       manager.Telephone,
			Email:           manager.Email,
			CompanyUUID:     manager.CompanyUUID,
			ProfileImageURL: &profileURL,
		}
	}

	// Delegates can only get a subset of their manager's info
	if g.IsCompanyDelegate(manager.CompanyUUID) {
		return gentypes.Manager{
			FirstName:       manager.FirstName,
			LastName:        manager.LastName,
			JobTitle:        manager.JobTitle,
			Email:           manager.Email,
			CompanyUUID:     manager.CompanyUUID,
			ProfileImageURL: &profileURL,
		}
	}

	return gentypes.Manager{}
}

func (g *Grant) managersToGentype(managers []models.Manager) []gentypes.Manager {
	var genManagers []gentypes.Manager
	for _, manager := range managers {
		genManagers = append(genManagers, g.managerToGentype(manager))
	}
	return genManagers
}

func (g *Grant) managerEmailExists(email string) bool {
	query := database.GormDB.Where("email = ?", email).First(&models.Manager{})
	if query.Error != nil {
		if query.RecordNotFound() {
			return false
		}

		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find manager for Email: %s", email)
		return false
	}
	return true
}

func (g *Grant) GetManagersByUUID(uuids []string) ([]gentypes.Manager, error) {
	var managers []gentypes.Manager
	if !g.IsAdmin && !g.IsManager {
		return managers, &errors.ErrUnauthorized
	}

	var allowedUUIDs []string
	// Managers can only get their own info
	if g.IsManager {
		for _, uuid := range uuids {
			if g.Claims.UUID.String() == uuid {
				allowedUUIDs = append(allowedUUIDs, uuid)
			}
		}
		if len(uuids) > 0 && len(allowedUUIDs) == 0 {
			return managers, &errors.ErrUnauthorized
		}
	}

	// Admin can get any manager info
	if g.IsAdmin {
		allowedUUIDs = uuids
	}

	query := database.GormDB.Where("uuid IN (?)", allowedUUIDs).Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return managers, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to find managers")
		return managers, &errors.ErrWhileHandling
	}

	return managers, nil
}

func filterManager(query *gorm.DB, filter *gentypes.ManagersFilter) *gorm.DB {
	if filter != nil {
		query = filterUser(query, &filter.UserFilter)

		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
	}

	return query
}

func (g *Grant) Manager(UUID gentypes.UUID) (gentypes.Manager, error) {
	// Admins can get any manager data
	// Managers can only get their own uuid
	if g.IsAdmin || (g.IsManager && UUID == g.Claims.UUID) {
		var manager models.Manager
		query := database.GormDB.Where("uuid = ?", UUID).First(&manager)
		if query.Error != nil {
			if query.RecordNotFound() {
				return gentypes.Manager{}, &errors.ErrNotFound
			}

			g.Logger.Log(sentry.LevelError, query.Error, "Unable to get manager")
			return gentypes.Manager{}, &errors.ErrWhileHandling
		}

		return g.managerToGentype(manager), nil
	}
	return gentypes.Manager{}, &errors.ErrUnauthorized
}

func (g *Grant) GetManagers(page *gentypes.Page, filter *gentypes.ManagersFilter, orderBy *gentypes.OrderBy) ([]gentypes.Manager, gentypes.PageInfo, error) {
	if !g.IsAdmin {
		return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	var managers []models.Manager

	// Count the total filtered dataset
	var count int32
	query := filterManager(database.GormDB, filter)
	countErr := query.Model(&models.Manager{}).Limit(MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		g.Logger.Log(sentry.LevelError, countErr, "Unable to count managers")
		return []gentypes.Manager{}, gentypes.PageInfo{}, countErr
	}

	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title"}, "created_at DESC")
	if orderErr != nil {
		return []gentypes.Manager{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	query = query.Find(&managers)
	if query.Error != nil {
		if query.RecordNotFound() {
			return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to find managers")
		return []gentypes.Manager{}, gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return g.managersToGentype(managers), gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  int32(len(managers)),
	}, nil
}

func (g *Grant) GetManagerSelf() (gentypes.Manager, error) {
	if !g.IsManager {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	manager, err := g.Manager(g.Claims.UUID)
	if err != nil {
		return gentypes.Manager{}, err
	}

	return manager, nil
}

func (g *Grant) CreateManager(managerDetails gentypes.CreateManagerInput) (gentypes.Manager, error) {
	var inputUUID gentypes.UUID
	// Auth
	if !g.IsAdmin && !g.IsManager {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	// check manager does not exist
	if g.managerEmailExists(managerDetails.Email) {
		return gentypes.Manager{}, &errors.ErrUserExists
	}

	// If you're an admin you need to provide the company UUID
	if g.IsAdmin {
		if managerDetails.CompanyUUID != nil {
			inputUUID = *managerDetails.CompanyUUID
		} else {
			return gentypes.Manager{}, &errors.ErrCompanyNotFound
		}
	}

	// If you're a manager the company UUID will be selected from the one in your JWT claims
	if g.IsManager {
		inputUUID = g.Claims.Company
	}

	// Check if company exists
	if !g.companyExists(inputUUID) {
		return gentypes.Manager{}, &errors.ErrCompanyNotFound
	}

	// TODO: Validate input better and return useful details
	manager := models.Manager{
		FirstName:   managerDetails.FirstName,
		LastName:    managerDetails.LastName,
		JobTitle:    managerDetails.JobTitle,
		Telephone:   managerDetails.Telephone,
		Password:    managerDetails.Password,
		Email:       managerDetails.Email,
		CompanyUUID: inputUUID,
	}
	createErr := database.GormDB.Create(&manager).Error
	if createErr != nil {
		g.Logger.Log(sentry.LevelError, createErr, "Unable to to create manager")
		return gentypes.Manager{}, &errors.ErrWhileHandling
	}

	return g.managerToGentype(manager), nil
}

func (g *Grant) UpdateManager(input gentypes.UpdateManagerInput) (gentypes.Manager, error) {
	if !g.IsAdmin && !(g.IsManager && g.Claims.UUID == input.UUID) {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	var manager models.Manager
	query := database.GormDB.Where("uuid = ?", input.UUID).First(&manager)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.Manager{}, &errors.ErrManagerNotFound
		}
		g.Logger.Logf(sentry.LevelError, query.Error, "Unable to find manager to update with UUID: %s", input.UUID)
		return gentypes.Manager{}, &errors.ErrWhileHandling
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
		g.Logger.Logf(sentry.LevelError, save.Error, "Error updating manager with UUID: %s", input.UUID)
		return gentypes.Manager{}, &errors.ErrWhileHandling
	}

	return g.managerToGentype(manager), nil
}

func (g *Grant) DeleteManager(uuid gentypes.UUID) (bool, error) {
	// managers can delete themselves
	// admins can delete any manager
	if (g.IsManager && g.Claims.UUID == uuid) || g.IsAdmin {
		// TODO: delete profile image from S3
		query := database.GormDB.Where("uuid = ?", uuid).Delete(models.Manager{})
		if query.Error != nil {
			g.Logger.Logf(sentry.LevelError, query.Error, "Error deleting manager with UUID: %s", uuid)
			return false, &errors.ErrWhileHandling
		}

		if query.RowsAffected == 0 {
			return false, &errors.ErrUserNotFound
		}
		return true, nil
	}

	return false, &errors.ErrUnauthorized
}

// ProfileUploadRequest generates a link that lets users upload a profile image to S3 directly
// Used by all user types
func (g *Grant) ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !g.IsManager && !g.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,      // The actual file type
		imageMeta.ContentLength, // The actual file content length
		[]string{"jpg", "png"},  // Allowed file types
		int32(20000000),         // Max file size = 20MB
		"profile",               // Save files in the "profile" s3 directory
		"profileImage",          // Unique identifier for this type of upload request
	)

	return url, successToken, err
}

// ManagerProfileUploadSuccess checks the successToken and sets the profile image of the current manager
func (g *Grant) ManagerProfileUploadSuccess(token string) error {
	if !g.IsManager {
		return &errors.ErrUnauthorized
	}

	s3Key, err := uploads.VerifyUploadSuccess(token, "profileImage")
	if err != nil {
		return err
	}

	query := database.GormDB.Model(&models.Manager{}).Where("uuid = ?", g.Claims.UUID).Update("profile_key", s3Key)
	if query.Error != nil {
		if query.RecordNotFound() {
			return &errors.ErrManagerNotFound
		}

		g.Logger.Log(sentry.LevelError, query.Error, "Unable to update manager profile")
		return &errors.ErrWhileHandling
	}

	return nil
}
