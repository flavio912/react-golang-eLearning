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

func (c *courseAppImpl) CreateModule(input gentypes.CreateModuleInput) (gentypes.Module, error) {
	if !c.grant.IsAdmin {
		return gentypes.Module{}, &errors.ErrUnauthorized
	}

	var (
		bannerKey    *string
		voiceoverKey *string
		video        *course.VideoInput
	)
	if input.BannerImageSuccessToken != nil {
		key, err := uploads.VerifyUploadSuccess(*input.BannerImageSuccessToken, "moduleImage")
		if err != nil {
			return gentypes.Module{}, &errors.ErrUploadTokenInvalid
		}
		bannerKey = &key
	}

	if input.VoiceoverSuccessToken != nil {
		key, err := uploads.VerifyUploadSuccess(*input.VoiceoverSuccessToken, "mp3Uploads")
		if err != nil {
			return gentypes.Module{}, &errors.ErrUploadTokenInvalid
		}
		voiceoverKey = &key
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
