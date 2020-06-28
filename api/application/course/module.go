package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) moduleToGentype(module models.Module) gentypes.Module {

	var bannerUrl *string
	if module.BannerKey != nil {
		url := uploads.GetImgixURL(*module.BannerKey)
		bannerUrl = &url
	}

	var voiceoverUrl *string
	if module.VoiceoverKey != nil {
		url := uploads.GetImgixURL(*module.VoiceoverKey)
		voiceoverUrl = &url
	}

	var video *gentypes.Video
	if module.VideoURL != nil && module.VideoType != nil {
		video = &gentypes.Video{
			Type: *module.VideoType,
			URL:  *module.VideoURL,
		}
	}

	return gentypes.Module{
		UUID:           module.UUID,
		Title:          module.Name,
		BannerImageURL: bannerUrl,
		Description:    module.Description,
		Transcript:     module.Transcript,
		VoiceoverURL:   voiceoverUrl,
		Video:          video,
	}
}

func getUploadKey(token *string, uploadIdent string) (*string, error) {
	var uploadKey *string
	if token != nil {
		key, err := uploads.VerifyUploadSuccess(*token, "moduleImage")
		if err != nil {
			return nil, &errors.ErrUploadTokenInvalid
		}
		uploadKey = &key
	}
	return uploadKey, nil
}

func (c *courseAppImpl) CreateModule(input gentypes.CreateModuleInput) (gentypes.Module, error) {
	if !c.grant.IsAdmin {
		return gentypes.Module{}, &errors.ErrUnauthorized
	}

	var (
		video *course.VideoInput
	)

	bannerKey, err := getUploadKey(input.BannerImageSuccessToken, "moduleImages")
	if err != nil {
		return gentypes.Module{}, err
	}
	voiceoverKey, err := getUploadKey(input.VoiceoverSuccessToken, "voiceoverUploads")
	if err != nil {
		return gentypes.Module{}, err
	}

	if input.Video != nil {
		video = &course.VideoInput{
			Type: (*input.Video).Type,
			URL:  (*input.Video).URL,
		}
	}

	createInput := course.CreateModuleInput{
		Name:         input.Name,
		Description:  input.Description,
		Transcript:   input.Transcript,
		Tags:         input.Tags,
		Syllabus:     input.Syllabus,
		Video:        video,
		BannerKey:    bannerKey,
		VoiceoverKey: voiceoverKey,
	}

	module, err := c.coursesRepository.CreateModule(createInput)
	return c.moduleToGentype(module), err
}

func (c *courseAppImpl) UpdateModule(input gentypes.UpdateModuleInput) (gentypes.Module, error) {
	if !c.grant.IsAdmin {
		return gentypes.Module{}, &errors.ErrUnauthorized
	}

	bannerKey, err := getUploadKey(input.BannerImageSuccessToken, "moduleImages")
	if err != nil {
		return gentypes.Module{}, err
	}
	voiceoverKey, err := getUploadKey(input.VoiceoverSuccessToken, "voiceoverUploads")
	if err != nil {
		return gentypes.Module{}, err
	}

	var video *course.VideoInput
	if input.Video != nil {
		video = &course.VideoInput{
			Type: (*input.Video).Type,
			URL:  (*input.Video).URL,
		}
	}

	inp := course.UpdateModuleInput{
		UUID:         input.UUID,
		Name:         input.Name,
		Description:  input.Description,
		Transcript:   input.Transcript,
		BannerKey:    bannerKey,
		VoiceoverKey: voiceoverKey,
		Video:        video,
		Tags:         input.Tags,
		Syllabus:     input.Syllabus,
	}

	module, err := c.coursesRepository.UpdateModule(inp)
	// TODO: Delete S3 images on success

	return c.moduleToGentype(module), err
}
