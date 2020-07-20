package course

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"
)

func (c *courseAppImpl) questionToGentype(question models.Question) gentypes.Question {
	if c.grant.IsAdmin {
		return gentypes.Question{
			UUID:             question.UUID,
			Text:             question.Text,
			RandomiseAnswers: &question.RandomiseAnswers,
			QuestionType:     question.QuestionType,
		}
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		return gentypes.Question{
			UUID:         question.UUID,
			Text:         question.Text,
			QuestionType: question.QuestionType,
		}
	}

	return gentypes.Question{}
}

func (c *courseAppImpl) questionsToGentypes(questions []models.Question) []gentypes.Question {
	genQuestions := make([]gentypes.Question, len(questions))
	for i, q := range questions {
		genQuestions[i] = c.questionToGentype(q)
	}
	return genQuestions
}

func (c *courseAppImpl) CreateQuestion(input gentypes.CreateQuestionInput) (gentypes.Question, error) {
	if !c.grant.IsAdmin {
		return gentypes.Question{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Question{}, err
	}

	var answerArgs []course.CreateAnswerArgs
	for _, answer := range input.Answers {
		var imageKey *string
		if answer.ImageToken != nil {
			key, err := uploads.VerifyUploadSuccess(*answer.ImageToken, "answerImages")
			if err != nil {
				return gentypes.Question{}, err
			}
			imageKey = &key
		}
		answerArgs = append(answerArgs, course.CreateAnswerArgs{
			Text:      answer.Text,
			ImageKey:  imageKey,
			IsCorrect: answer.IsCorrect,
		})
	}

	question, err := c.coursesRepository.CreateQuestion(course.CreateQuestionArgs{
		Text:             input.Text,
		RandomiseAnswers: input.RandomiseAnswers,
		QuestionType:     input.QuestionType,
		Answers:          answerArgs,
		Tags:             input.Tags,
	})
	return c.questionToGentype(question), err
}

func (c *courseAppImpl) UpdateQuestion(input gentypes.UpdateQuestionInput) (gentypes.Question, error) {
	if !c.grant.IsAdmin {
		return gentypes.Question{}, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return gentypes.Question{}, err
	}

	// TODO: If answer type is text delete any images associated with it

	// Validate image tokens if given
	var ans *[]course.UpdateAnswerArgs
	if input.Answers != nil {
		answers := []course.UpdateAnswerArgs{}
		for _, ans := range *input.Answers {
			var key *string
			if ans.ImageToken != nil && (ans.AnswerType == gentypes.ImageAnswer || ans.AnswerType == gentypes.TextImageAnswer) {
				imgKey, err := uploads.VerifyUploadSuccess(*ans.ImageToken, "answerImages")
				if err != nil {
					return gentypes.Question{}, err
				}

				key = &imgKey
			}

			answers = append(answers, course.UpdateAnswerArgs{
				UUID:       ans.UUID,
				IsCorrect:  ans.IsCorrect,
				Text:       ans.Text,
				ImageKey:   key,
				AnswerType: ans.AnswerType,
			})

		}

		ans = &answers
	}

	updateArgs := course.UpdateQuestionArgs{
		UUID:             input.UUID,
		Text:             input.Text,
		RandomiseAnswers: input.RandomiseAnswers,
		QuestionType:     input.QuestionType,
		Answers:          ans,
		Tags:             input.Tags,
	}

	question, err := c.coursesRepository.UpdateQuestion(updateArgs)
	return c.questionToGentype(question), err
}

func (c *courseAppImpl) answerToGentype(answer models.BasicAnswer) gentypes.Answer {
	var imageUrl *string

	if answer.ImageKey != nil {
		url := uploads.GetImgixURL(*answer.ImageKey)
		imageUrl = &url
	}

	if c.grant.IsAdmin {
		return gentypes.Answer{
			IsCorrect: &answer.IsCorrect,
			Text:      answer.Text,
			UUID:      answer.UUID,
			ImageURL:  imageUrl,
		}
	}
	return gentypes.Answer{
		Text:     answer.Text,
		UUID:     answer.UUID,
		ImageURL: imageUrl,
	}
}

func (c *courseAppImpl) answersToGentypes(answers []models.BasicAnswer) []gentypes.Answer {
	ans := make([]gentypes.Answer, len(answers))
	for i, answer := range answers {
		ans[i] = c.answerToGentype(answer)
	}
	return ans
}

func (c *courseAppImpl) ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]gentypes.Answer, error) {
	outputAns := make(map[gentypes.UUID][]gentypes.Answer)

	if !c.grantIsAllowedToViewQuestions(questionUUIDs) {
		return outputAns, &errors.ErrUnauthorized
	}

	ansMap, err := c.coursesRepository.ManyAnswers(questionUUIDs)

	for key, val := range ansMap {
		outputAns[key] = c.answersToGentypes(val)
	}

	return outputAns, err
}

func (c *courseAppImpl) grantIsAllowedToViewQuestions(questionUUIDs []gentypes.UUID) bool {
	if c.grant.IsAdmin {
		return true
	}

	if c.grant.IsDelegate || c.grant.IsIndividual {
		var takerId gentypes.UUID
		if c.grant.IsDelegate {
			delegate, _ := c.usersRepository.Delegate(c.grant.Claims.UUID)
			takerId = delegate.CourseTakerUUID
		}

		if c.grant.IsIndividual {
			individual, _ := c.usersRepository.Individual(c.grant.Claims.UUID)
			takerId = individual.CourseTakerUUID
		}

		activeCourses, _ := c.usersRepository.TakerActiveCourses(takerId)

		// TODO: Replace this with a big join
		// TODO: FIX This stipulates that all questions must be part of the same test
		allowed := false
		for _, activeCourse := range activeCourses {
			course, _ := c.coursesRepository.OnlineCourse(activeCourse.CourseID)
			tests, _ := c.coursesRepository.CourseTests(course.UUID)

			for _, test := range tests {
				questions, _ := c.coursesRepository.TestQuestions(test.UUID)
				for _, uuid := range questionUUIDs {
					found := false
					for _, question := range questions {
						if question.UUID == uuid {
							found = true
						}
					}
					if found {
						allowed = true
						break
					}
				}

				if allowed {
					break
				}
			}

			if allowed {
				break
			}
		}

		if allowed {
			return true
		}
	}
	return false
}

func (c *courseAppImpl) Question(uuid gentypes.UUID) (gentypes.Question, error) {
	if !c.grantIsAllowedToViewQuestions([]gentypes.UUID{uuid}) {
		return gentypes.Question{}, &errors.ErrUnauthorized
	}

	// Check user is assigned course with this question in
	question, err := c.coursesRepository.Question(uuid)

	return c.questionToGentype(question), err
}

func (c *courseAppImpl) Questions(
	page *gentypes.Page,
	filter *gentypes.QuestionFilter,
	orderBy *gentypes.OrderBy,
) ([]gentypes.Question, gentypes.PageInfo, error) {
	if !c.grant.IsAdmin {
		return []gentypes.Question{}, gentypes.PageInfo{}, &errors.ErrUnauthorized
	}

	questions, pageInfo, err := c.coursesRepository.Questions(page, filter, orderBy)
	if err != nil {
		return c.questionsToGentypes(questions), pageInfo, &errors.ErrWhileHandling
	}

	return c.questionsToGentypes(questions), pageInfo, nil
}

func (c *courseAppImpl) DeleteQuestion(input gentypes.DeleteQuestionInput) (bool, error) {
	if !c.grant.IsAdmin {
		return false, &errors.ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return false, err
	}

	return c.coursesRepository.DeleteQuestion(input.UUID)
}

func (c *courseAppImpl) TestQuestions(testUUID gentypes.UUID) ([]gentypes.Question, error) {
	if !c.grantCanViewSyllabusItems([]gentypes.UUID{testUUID}, gentypes.TestType) {
		return []gentypes.Question{}, &errors.ErrUnauthorized
	}

	questions, err := c.coursesRepository.TestQuestions(testUUID)
	if err != nil {
		return []gentypes.Question{}, &errors.ErrWhileHandling
	}

	return c.questionsToGentypes(questions), nil
}

// AnswerImageUploadRequest generates a link that lets users upload a profile image to S3 directly
// Used by all user types
func (c *courseAppImpl) AnswerImageUploadRequest(imageMeta gentypes.UploadFileMeta) (string, string, error) {
	if !c.grant.IsAdmin {
		return "", "", &errors.ErrUnauthorized
	}

	url, successToken, err := uploads.GenerateUploadURL(
		imageMeta.FileType,      // The actual file type
		imageMeta.ContentLength, // The actual file content length
		[]string{"jpg", "png"},  // Allowed file types
		int32(20000000),         // Max file size = 20MB
		"answers",               // Save files in the "answers" s3 directory
		"answerImages",          // Unique identifier for this type of upload request
	)

	return url, successToken, err
}
