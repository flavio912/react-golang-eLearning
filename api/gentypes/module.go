package gentypes

type VideoType string

const (
	WistiaVideo VideoType = "wistia"
)

type Video struct {
	Type VideoType
	URL  string
}

type Module struct {
	UUID           UUID
	Name           string
	BannerImageURL *string
	Description    string
	Transcript     string
	VoiceoverURL   *string
	Video          *Video
}

type ModuleElement string

const (
	ModuleTest   ModuleElement = "test"
	ModuleLesson ModuleElement = "lesson"
)

type ModuleItem struct {
	Type ModuleElement
	UUID UUID
}

type CreateModuleInput struct {
	Name                    string
	Description             string
	Transcript              string
	BannerImageSuccessToken *string
	VoiceoverSuccessToken   *string
	Video                   *Video
	Tags                    *[]UUID
	Syllabus                *[]ModuleItem
}

type UpdateModuleInput struct {
	UUID                    UUID
	Name                    *string
	Description             *string
	Transcript              *string
	BannerImageSuccessToken *string
	VoiceoverSuccessToken   *string
	Video                   *Video
	Tags                    *[]UUID
	Syllabus                *[]ModuleItem
}
