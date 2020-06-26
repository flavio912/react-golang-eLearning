package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type QuestionResolver struct {
	question gentypes.Question
}

func (q *QuestionResolver) Text() string {
	return q.question.Text
}
func (q *QuestionResolver) RandomiseAnswers() *bool {
	return q.question.RandomiseAnswers
}
func (q *QuestionResolver) QuestionType() gentypes.QuestionType {
	return q.question.QuestionType
}
func (q *QuestionResolver) Answers() *[]*AnswerResolver {
	return nil
}
