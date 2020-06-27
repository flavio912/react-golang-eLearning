package gentypes

type VideoType string

const (
	WistiaVideo VideoType = "wistia"
)

type VideoInput struct {
	VideoType VideoType
	VideoURL  string
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
	Name                  string
	Tags                  *[]UUID
	Description           string
	Transcript            string
	VoiceoverSuccessToken string
	Video                 *VideoInput
	Syllabus              *[]ModuleItem
}
