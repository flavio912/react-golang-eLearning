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

// ProfileUploadRequest generates a link that lets users upload a profile image to S3 directly
// Used by all user types
func (u *usersAppImpl) ProfileUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !u.grant.IsManager && !u.grant.IsAdmin {
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
