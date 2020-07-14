package users

import (
	"time"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/email"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (u *usersAppImpl) delegateToGentype(delegate models.Delegate) gentypes.Delegate {
	var profileURL *string
	if delegate.ProfileKey != nil {
		url := uploads.GetImgixURL(*delegate.ProfileKey)
		profileURL = &url
	}

	createdAt := delegate.CreatedAt.Format(time.RFC3339)
	lastLogin := delegate.LastLogin.Format(time.RFC3339)
	return gentypes.Delegate{
		CreatedAt:       &createdAt,
		UUID:            delegate.UUID,
		FirstName:       delegate.FirstName,
		LastName:        delegate.LastName,
		JobTitle:        delegate.JobTitle,
		Telephone:       delegate.Telephone,
		Email:           delegate.Email,
		CompanyUUID:     delegate.CompanyUUID,
		TTC_ID:          delegate.TtcId,
		ProfileImageURL: profileURL,
		LastLogin:       lastLogin,
		CourseTakerUUID: delegate.CourseTakerUUID,
	}
}

func (u *usersAppImpl) delegatesToGentype(delegates []models.Delegate) []gentypes.Delegate {
	var genDelegates []gentypes.Delegate
	for _, delegate := range delegates {
		genDelegates = append(genDelegates, u.delegateToGentype(delegate))
	}

	return genDelegates
}

func (u *usersAppImpl) GetDelegates(page *gentypes.Page, filter *gentypes.DelegatesFilter, orderBy *gentypes.OrderBy) ([]gentypes.Delegate, gentypes.PageInfo, error) {
	if u.grant.IsAdmin {
		delegates, pageInfo, err := u.usersRepository.GetDelegates(page, filter, orderBy, nil)
		return u.delegatesToGentype(delegates), pageInfo, err
	}

	if u.grant.IsManager {
		delegates, pageInfo, err := u.usersRepository.GetDelegates(page, filter, orderBy, &u.grant.Claims.Company)
		return u.delegatesToGentype(delegates), pageInfo, err
	}

	return []gentypes.Delegate{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
}

func (u *usersAppImpl) CreateDelegate(delegateDetails gentypes.CreateDelegateInput) (gentypes.Delegate, *string, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager {
		return gentypes.Delegate{}, nil, &errors.ErrUnauthorized
	}

	if err := delegateDetails.Validate(); err != nil {
		return gentypes.Delegate{}, nil, err
	}

	var (
		needsGeneratePass = delegateDetails.GeneratePassword != nil && *delegateDetails.GeneratePassword
		companyUUID       gentypes.UUID
		s3UploadKey       *string
		realPass          *string
		password          string
	)

	if u.grant.IsManager {
		companyUUID = u.grant.Claims.Company
	}
	if u.grant.IsAdmin {
		if delegateDetails.CompanyUUID == nil {
			return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
		}
		companyUUID = *delegateDetails.CompanyUUID
	}

	// Check if company exists
	if !u.usersRepository.CompanyExists(companyUUID) {
		return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
	}

	// Check if autogenerating password is required
	if needsGeneratePass {
		pass, err := auth.GenerateSecurePassword(10)
		if err != nil {
			u.grant.Logger.Log(sentry.LevelError, err, "Unable to generate secure password")
			return gentypes.Delegate{}, nil, &errors.ErrWhileHandling
		}
		password = pass
		realPass = &pass
	}

	// Check if upload token is valid
	if delegateDetails.ProfileImageUploadToken != nil {
		tmpUploadKey, err := uploads.VerifyUploadSuccess(*delegateDetails.ProfileImageUploadToken, "profileImage")
		if err != nil {
			return gentypes.Delegate{}, nil, &errors.ErrUploadTokenInvalid
		}

		s3UploadKey = &tmpUploadKey
	}

	comp, err := u.usersRepository.Company(companyUUID)
	if err != nil {
		return gentypes.Delegate{}, nil, &errors.ErrCompanyNotFound
	}

	sendEmails := func(delegate models.Delegate) bool {
		// Send transactional email
		// If not generated password, send an email to the user
		if !needsGeneratePass {
			token, err := auth.GenerateFinaliseDelegateToken(auth.FinaliseDelegateClaims{
				UUID: delegate.UUID,
			})
			if err != nil {
				u.grant.Logger.Log(sentry.LevelError, err, "Unable to generate finalise delegate token")
				return false
			}

			if delegate.Email == nil {
				u.grant.Logger.LogMessage(sentry.LevelError, "Delegate email is nil")
				return false
			}

			err = email.SendFinaliseAccountEmail(token, delegate.FirstName, *delegate.Email)
			if err != nil {
				u.grant.Logger.Log(sentry.LevelWarning, err, "Unable to send finalise account email")
				return false
			}
		}

		return true
	}
	delegate, err := u.usersRepository.CreateDelegate(
		delegateDetails,
		s3UploadKey,
		&password,
		comp,
		&sendEmails,
	)

	return u.delegateToGentype(delegate), realPass, err
}

func (u *usersAppImpl) UpdateDelegate(input gentypes.UpdateDelegateInput) (gentypes.Delegate, error) {
	if !u.grant.IsAdmin && !u.grant.IsManager {
		return gentypes.Delegate{}, &errors.ErrUnauthorized
	}

	var (
		password    *string
		s3UploadKey *string
	)

	if u.grant.IsManager {
		input.CompanyUUID = &u.grant.Claims.Company
	}

	if u.grant.IsAdmin {
		password = input.NewPassword
	}

	if input.ProfileImageUploadToken != nil {
		tmpUploadKey, err := uploads.VerifyUploadSuccess(*input.ProfileImageUploadToken, "profileImage")
		if err != nil {
			return gentypes.Delegate{}, &errors.ErrUploadTokenInvalid
		}

		s3UploadKey = &tmpUploadKey
	}

	delegate, err := u.usersRepository.UpdateDelegate(input, s3UploadKey, password)

	return u.delegateToGentype(delegate), err
}
