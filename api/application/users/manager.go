package users

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (u *usersAppImpl) managerToGentype(manager models.Manager) gentypes.Manager {
	profileURL := uploads.GetImgixURL(manager.ProfileKey)
	// Admins and managers themselves can get all info
	if u.grant.IsAdmin || (u.grant.IsManager && u.grant.Claims.Company == manager.CompanyUUID) {
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
	if u.grant.IsCompanyDelegate(manager.CompanyUUID) {
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

func (u *usersAppImpl) managersToGentype(managers []models.Manager) []gentypes.Manager {
	var genManagers []gentypes.Manager
	for _, manager := range managers {
		genManagers = append(genManagers, u.managerToGentype(manager))
	}
	return genManagers
}

func (u *usersAppImpl) GetManagersByUUID(uuids []gentypes.UUID) ([]gentypes.Manager, error) {
	// Manager can get own uuid, admin can get any
	if !(len(uuids) == 1 && u.grant.Claims.UUID == uuids[0] && u.grant.IsManager) && !u.grant.IsAdmin {
		return []gentypes.Manager{}, &errors.ErrUnauthorized
	}

	managers, err := u.usersRepository.GetManagersByUUID(uuids)
	return u.managersToGentype(managers), err
}

func (u *usersAppImpl) GetManagerIDsByCompany(
	companyUUID gentypes.UUID,
	page *gentypes.Page,
	filter *gentypes.ManagersFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.UUID, gentypes.PageInfo, error) {
	if !u.grant.IsAdmin {
		return []gentypes.UUID{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	return u.usersRepository.GetManagerIDsByCompany(companyUUID, page, filter, orderBy)
}

func (u *usersAppImpl) Delegate(uuid gentypes.UUID) (gentypes.Delegate, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager && !u.grant.IsDelegate {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	delegate, err := u.usersRepository.Delegate(uuid)

	if !u.grant.IsAdmin &&
		!(u.grant.IsManager && u.grant.Claims.Company == delegate.CompanyUUID) &&
		!(u.grant.IsDelegate && u.grant.Claims.UUID == delegate.UUID) {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	return u.delegateToGentype(delegate), err
}

func (u *usersAppImpl) CreateManager(managerDetails gentypes.CreateManagerInput) (gentypes.Manager, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	var compUUID gentypes.UUID
	// If you're an admin you need to provide the company UUID
	if u.grant.IsAdmin {
		if managerDetails.CompanyUUID != nil {
			compUUID = *managerDetails.CompanyUUID
		} else {
			return gentypes.Manager{}, &errors.ErrCompanyNotFound
		}
	}

	// If you're a manager the company UUID will be selected from the one in your JWT claims
	if u.grant.IsManager {
		compUUID = u.grant.Claims.Company
	}

	manager, err := u.usersRepository.CreateManager(managerDetails, compUUID)
	return u.managerToGentype(manager), err
}

func (u *usersAppImpl) UpdateManager(input gentypes.UpdateManagerInput) (gentypes.Manager, error) {
	if !(u.grant.IsManager && u.grant.Claims.UUID == input.UUID) && !u.grant.IsAdmin {
		return gentypes.Manager{}, &errors.ErrUnauthorized
	}

	manager, err := u.usersRepository.UpdateManager(input)
	return u.managerToGentype(manager), err
}

func (u *usersAppImpl) DeleteManager(uuid gentypes.UUID) (bool, error) {
	// managers can delete themselves
	// admins can delete any manager
	if !(u.grant.IsManager && u.grant.Claims.UUID == uuid) && !u.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	// TODO: delete profile image from S3

	return u.usersRepository.DeleteManager(uuid)
}

// ProfileUploadRequest generates a link that lets users upload a profile image to S3 directly
// Used by all user types
func (u *usersAppImpl) ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !u.grant.IsManager && !u.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,             // The actual file type
		imageMeta.ContentLength,        // The actual file content length
		[]string{"jpg", "png", "jpeg"}, // Allowed file types
		int32(20000000),                // Max file size = 20MB
		"profile",                      // Save files in the "profile" s3 directory
		"profileImage",                 // Unique identifier for this type of upload request
	)

	return url, successToken, err
}

// ManagerProfileUploadSuccess checks the successToken and sets the profile image of the current manager
func (u *usersAppImpl) ManagerProfileUploadSuccess(token string) error {
	if !u.grant.IsManager {
		return &errors.ErrUnauthorized
	}

	s3Key, err := uploads.VerifyUploadSuccess(token, "profileImage")
	if err != nil {
		return err
	}

	err = u.usersRepository.UpdateManagerProfileKey(u.grant.Claims.UUID, &s3Key)
	if err != nil {
		//TODO Delete s3 image
		return &errors.ErrWhileHandling
	}

	return nil
}
