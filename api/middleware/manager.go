package middleware

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func (g *Grant) managerToGentype(manager models.Manager) gentypes.Manager {
	profileURL := uploads.GetImgixURL(manager.ProfileKey)
	// Admins and managers themselves can get all info
	if g.IsAdmin || (g.IsManager && g.Claims.Company.UUID == manager.CompanyID) {
		createdAt := manager.CreatedAt.Format(time.RFC3339)
		return gentypes.Manager{
			User: gentypes.User{
				CreatedAt: &createdAt,
				UUID:      gentypes.UUID{UUID: manager.UUID},
				Email:     manager.Email,
				FirstName: manager.FirstName,
				LastName:  manager.LastName,
				JobTitle:  manager.JobTitle,
				Telephone: manager.Telephone,
			},
			CompanyID:       gentypes.UUID{UUID: manager.CompanyID},
			ProfileImageURL: &profileURL,
		}
	}

	// Delegates can only get a subset of their manager's info
	if g.IsCompanyDelegate(gentypes.UUID{UUID: manager.CompanyID}) {
		return gentypes.Manager{
			User: gentypes.User{
				Email:     manager.Email,
				FirstName: manager.FirstName,
				LastName:  manager.LastName,
				JobTitle:  manager.JobTitle,
			},
			CompanyID:       gentypes.UUID{UUID: manager.CompanyID},
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
		// If some other error occurs log it
		glog.Errorf("Unable to find manager for Email: %s - error: %s", email, query.Error.Error())
		return false
	}
	return true
}

func (g *Grant) GetManagersByUUID(uuids []string) ([]gentypes.Manager, error) {
	var managers []gentypes.Manager

	var allowedUUIDs []string

	// Managers can only get their own info
	if g.IsManager {
		for _, uuid := range uuids {
			if g.Claims.UUID.String() == uuid {
				allowedUUIDs = append(allowedUUIDs, uuid)
			}
		}
	}

	// Admin can get any manager info
	if g.IsAdmin {
		allowedUUIDs = uuids
	}

	if !g.IsAdmin && !g.IsManager {
		return managers, &errors.ErrUnauthorized
	}

	db := database.GormDB.Where("uuid IN (?)", allowedUUIDs).Find(&managers)
	if db.Error != nil {
		if db.RecordNotFound() {
			return managers, &errors.ErrNotFound
		}
		glog.Errorf("DB Error: %s", db.Error.Error())
		return managers, &errors.ErrWhileHandling
	}

	return managers, nil
}

func filterManager(query *gorm.DB, filter *gentypes.ManagersFilter) *gorm.DB {
	if filter != nil {
		if filter.Email != nil && *filter.Email != "" {
			query = query.Where("email ILIKE ?", "%%"+*filter.Email+"%%")
		}
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("first_name || ' ' || last_name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.UUID != nil && *filter.UUID != "" {
			query = query.Where("uuid = ?", *filter.UUID)
		}
		if filter.JobTitle != nil && *filter.JobTitle != "" {
			query = query.Where("job_title ILIKE ?", "%%"+*filter.JobTitle+"%%")
		}
	}

	return query
}

func (g *Grant) GetManagerByUUID(UUID gentypes.UUID) (gentypes.Manager, error) {
	// Admins can get any manager data
	// Managers can only get their own uuid
	if g.IsAdmin || (g.IsManager && UUID == g.Claims.UUID) {
		var manager models.Manager
		err := database.GormDB.Where("uuid = ?", UUID).First(&manager).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return gentypes.Manager{}, &errors.ErrNotFound
			}

			return gentypes.Manager{}, err
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
		glog.Errorf("Count query failed: %s", countErr.Error())
		return []gentypes.Manager{}, gentypes.PageInfo{}, countErr
	}
	fmt.Printf("COUNT: %d", count)

	query = query.Order("created_at DESC")
	query, orderErr := getOrdering(query, orderBy, []string{"created_at", "email", "first_name", "job_title"})
	if orderErr != nil {
		return []gentypes.Manager{}, gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := getPage(query, page)
	err := query.Find(&managers).Error
	if err != nil {
		return []gentypes.Manager{}, gentypes.PageInfo{}, err
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

	manager, err := g.GetManagerByUUID(g.Claims.UUID)
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
	if !g.CompanyExists(inputUUID) {
		return gentypes.Manager{}, &errors.ErrCompanyNotFound
	}

	// TODO: Validate input better and return useful details
	manager := models.Manager{
		User: models.User{
			FirstName: managerDetails.FirstName,
			LastName:  managerDetails.LastName,
			Email:     managerDetails.Email,
			JobTitle:  managerDetails.JobTitle,
			Telephone: managerDetails.Telephone,
			Password:  managerDetails.Password,
		},
		CompanyID: inputUUID.UUID,
	}
	createErr := database.GormDB.Create(&manager).Error
	if createErr != nil {
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
		glog.Errorf("Unable to find manager to update with UUID: %s - error: %s", input.UUID, query.Error.Error())
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

	save := database.GormDB.Save(manager)
	if save.Error != nil {
		glog.Errorf("Error updating manager with UUID: %s - error: %s", input.UUID, save.Error.Error())
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
			glog.Errorf("Unable to delete manager: %s", query.Error.Error())
			return false, &errors.ErrDeleteFailed
		}

		if query.RowsAffected == 0 {
			return false, &errors.ErrUserNotFound
		}
		return true, nil
	}

	return false, &errors.ErrUnauthorized
}

// ManagerProfileUploadRequest generates a link that lets users upload a profile image to S3 directly
func (g *Grant) ManagerProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !g.IsManager && !g.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,      // The actual file type
		imageMeta.ContentLength, // The actual file content length
		[]string{"jpg", "png"},  // Allowed file types
		int32(20000000),         // Max file size = 20MB
		"profile",               // Save files in the "profile" s3 directory
		"managerProfile",        // Unique identifier for this type of upload request
	)

	return url, successToken, err
}

// ManagerProfileUploadSuccess checks the successToken and sets the profile image of the current manager
func (g *Grant) ManagerProfileUploadSuccess(token string) error {
	if !g.IsManager {
		return &errors.ErrUnauthorized
	}

	s3Key, err := uploads.VerifyUploadSuccess(token, "managerProfile")
	if err != nil {
		return err
	}

	query := database.GormDB.Model(&models.Manager{}).Where("uuid = ?", g.Claims.UUID).Update("profile_key", s3Key)
	if query.Error != nil {
		return getDBErrorType(query)
	}

	return nil
}
