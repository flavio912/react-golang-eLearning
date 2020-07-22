package course

import (
	"github.com/getsentry/sentry-go"
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
		Name:           module.Name,
		BannerImageURL: bannerUrl,
		Description:    module.Description,
		Transcript:     module.Transcript,
		VoiceoverURL:   voiceoverUrl,
		Video:          video,
	}
}

func (c *courseAppImpl) modulesToGentypes(modules []models.Module) []gentypes.Module {
	var genModules = make([]gentypes.Module, len(modules))
	for i, module := range modules {
		genModules[i] = c.moduleToGentype(module)
	}
	return genModules
}

func (c *courseAppImpl) Modules(
	page *gentypes.Page,
	filter *gentypes.ModuleFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Module, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Module{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	modules, pageInfo, err := c.coursesRepository.Modules(page, filter, orderBy)
	if err != nil {
		return c.modulesToGentypes(modules), pageInfo, &errors.ErrWhileHandling
	}

	return c.modulesToGentypes(modules), pageInfo, nil
}

func (c *courseAppImpl) ModulesByUUIDs(uuids []gentypes.UUID) ([]gentypes.Module, error) {
	if !c.grant.IsAdmin && !c.grant.IsDelegate && !c.grant.IsIndividual {
		return []gentypes.Module{}, &errors.ErrUnauthorized
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		// Check user is taking a course with those modules in it
		var courseTakerID gentypes.UUID
		if c.grant.IsDelegate {
			delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
			courseTakerID = delegate.CourseTakerUUID
		}

		if c.grant.IsIndividual {
			individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
			courseTakerID = individual.CourseTakerUUID
		}

		activeCourses, err := c.usersRepository.TakerActiveCourses(courseTakerID)
		if err != nil {
			return []gentypes.Module{}, &errors.ErrWhileHandling
		}

		var courseIds = make([]uint, len(activeCourses))
		for i, activeCourse := range activeCourses {
			courseIds[i] = activeCourse.CourseID
		}

		areModsInCourses, err := c.coursesRepository.AreInCourses(courseIds, uuids, gentypes.ModuleType)
		if err != nil {
			return []gentypes.Module{}, &errors.ErrWhileHandling
		}

		if !areModsInCourses {
			return []gentypes.Module{}, &errors.ErrWhileHandling
		}
	}

	modules, err := c.coursesRepository.ModulesByUUIDs(uuids)
	return c.modulesToGentypes(modules), err
}

func (c *courseAppImpl) Module(uuid gentypes.UUID) (gentypes.Module, error) {
	if !c.grant.IsAdmin && !c.grant.IsDelegate && !c.grant.IsIndividual {
		return gentypes.Module{}, &errors.ErrUnauthorized
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		// Check user is taking a course with this module in it
		var courseTakerID gentypes.UUID
		if c.grant.IsDelegate {
			delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
			courseTakerID = delegate.CourseTakerUUID
		}

		if c.grant.IsIndividual {
			individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
			courseTakerID = individual.CourseTakerUUID
		}

		activeCourses, err := c.usersRepository.TakerActiveCourses(courseTakerID)
		if err != nil {
			return gentypes.Module{}, &errors.ErrWhileHandling
		}

		var courseIds = make([]uint, len(activeCourses))
		for i, activeCourse := range activeCourses {
			courseIds[i] = activeCourse.CourseID
		}

		moduleInCourses, err := c.coursesRepository.IsModuleInCourses(courseIds, uuid)
		if err != nil {
			return gentypes.Module{}, &errors.ErrWhileHandling
		}

		if !moduleInCourses {
			return gentypes.Module{}, &errors.ErrUnauthorized
		}
	}

	// Get module
	module, err := c.coursesRepository.GetModuleByUUID(uuid)
	return c.moduleToGentype(module), err
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

func (c *courseAppImpl) ModuleSyllabus(moduleUUID gentypes.UUID) ([]gentypes.ModuleItem, error) {
	if !c.grantCanViewSyllabusItems([]gentypes.UUID{moduleUUID}, gentypes.ModuleType) {
		return []gentypes.ModuleItem{}, &errors.ErrUnauthorized
	}

	return c.coursesRepository.GetModuleStructure(moduleUUID)
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

func (c *courseAppImpl) DeleteModule(input gentypes.DeleteModuleInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return false, err
	}

	return c.coursesRepository.DeleteModule(input.UUID)
}

func (c *courseAppImpl) ModuleBannerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,      // The actual file type
		imageMeta.ContentLength, // The actual file content length
		[]string{"jpg", "png"},  // Allowed file types
		int32(20000000),         // Max file size = 20MB
		"moduleImages",          // Save files in this s3 directory
		"moduleImages",          // Unique identifier for this type of upload request
	)

	return url, successToken, err
}

func (c *courseAppImpl) VoiceoverUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,      // The actual file type
		imageMeta.ContentLength, // The actual file content length
		[]string{"mp3"},         // Allowed file types
		int32(20000000),         // Max file size = 20MB
		"voiceoverUploads",      // Save files in this s3 directory
		"voiceoverUploads",      // Unique identifier for this type of upload request
	)

	return url, successToken, err
}

func (c *courseAppImpl) ManyModuleTags(moduleUUIDs []gentypes.UUID) (map[gentypes.UUID][]gentypes.Tag, error) {
	modulesToTags, err := c.coursesRepository.ManyModuleTags(moduleUUIDs)
	if err != nil {
		c.grant.Logger.Log(sentry.LevelWarning, err, "ManyModuleTags: Unable to get tags")
		return map[gentypes.UUID][]gentypes.Tag{}, &errors.ErrWhileHandling
	}

	var genTags = map[gentypes.UUID][]gentypes.Tag{}
	for key, element := range modulesToTags {
		genTags[key] = tagsToGentypes(element)
	}
	return genTags, nil
}
