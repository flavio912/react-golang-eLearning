package models

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type Test struct {
	UUID              gentypes.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name              string
	Tags              []Tag `gorm:"many2many:test_tags_link;"`
	AttemptsAllowed   *uint
	PassPercentage    float32
	QuestionsToAnswer uint
	RandomiseAnswers  bool
}

type TestQuestionsLink struct {
	TestUUID     gentypes.UUID `gorm:"primary_key;type:uuid;"`
	QuestionUUID gentypes.UUID
	Rank         string `gorm:"primary_key"`
}

type Question struct {
	UUID             gentypes.UUID `gorm:"primary_key;default:uuid_generate_v4()"`
	Text             string
	RandomiseAnswers bool
	QuestionType     gentypes.QuestionType // e.g singleChoice
	Answers          []BasicAnswer         `gorm:"foreignkey:QuestionUUID;association_autoupdate:false"`
	Tags             []Tag                 `gorm:"many2many:question_tags_link;"`
}

type BasicAnswer struct {
	UUID         gentypes.UUID `gorm:"primary_key;default:uuid_generate_v4()"`
	QuestionUUID gentypes.UUID
	IsCorrect    bool
	Text         *string
	ImageKey     *string // S3 image key
	Rank         string
}
