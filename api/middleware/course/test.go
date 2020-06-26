package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

type CreateTestInput struct {
	Name              string
	AttemptsAllowed   uint
	PassPercentage    float64
	QuestionsToAnswer uint
	RandomiseAnswers  bool
	Questions         []gentypes.UUID
	Tags              *[]gentypes.UUID
}

func (c *coursesRepoImpl) CreateTest(input CreateTestInput) (models.Test, error) {

	test := models.Test{
		Name:              input.Name,
		AttemptsAllowed:   input.AttemptsAllowed,
		PassPercentage:    input.PassPercentage,
		QuestionsToAnswer: input.QuestionsToAnswer,
		RandomiseAnswers:  input.RandomiseAnswers,
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&test).Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create test")
		tx.Rollback()
		return models.Test{}, &errors.ErrWhileHandling
	}

	// Assert that test uuid is valid
	if test.UUID == (gentypes.UUID{}) {
		tx.Rollback()
		c.Logger.LogMessage(sentry.LevelError, "Create test uuid blank")
		return models.Test{}, &errors.ErrWhileHandling
	}

	// Create question links
	for i, uuid := range input.Questions {
		link := models.TestQuestionsLink{
			TestUUID:     test.UUID,
			QuestionUUID: uuid,
			Rank:         strconv.Itoa(i),
		}
		if err := tx.Create(&link).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelWarning, err, "Unable to create test question link")
			return models.Test{}, &errors.ErrWhileHandling
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit test")
		return models.Test{}, &errors.ErrWhileHandling
	}

	return test, nil
}

func (c *coursesRepoImpl) Test(testUUID gentypes.UUID) (models.Test, error) {
	testMap, err := c.ManyTests([]gentypes.UUID{testUUID})
	if _, ok := testMap[testUUID]; ok {
		return testMap[testUUID], nil
	}

	if err == &errors.ErrNotAllFound {
		return models.Test{}, &errors.ErrNotFound
	}

	return models.Test{}, err
}

// ManyTests maps testUUIDs to their respective test
func (c *coursesRepoImpl) ManyTests(testUUIDs []gentypes.UUID) (map[gentypes.UUID]models.Test, error) {
	var tests []models.Test
	query := database.GormDB.Where("uuid IN (?)", testUUIDs).Find(&tests)
	if query.Error != nil {
		if query.RecordNotFound() {
			return map[gentypes.UUID]models.Test{}, &errors.ErrNotAllFound
		}

		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get many tests")
		return map[gentypes.UUID]models.Test{}, &errors.ErrWhileHandling
	}

	var err error
	if len(tests) < len(testUUIDs) {
		err = &errors.ErrNotAllFound
	}

	var uuidToTest = make(map[gentypes.UUID]models.Test)
	for _, test := range tests {
		uuidToTest[test.UUID] = test
	}

	return uuidToTest, err
}

// TestQuestions gets slice of questions for a test (in rank order)
func (c *coursesRepoImpl) TestQuestions(testUUID gentypes.UUID) ([]models.Question, error) {
	var questions []models.Question
	query := database.GormDB.Table("questions").
		Joins("JOIN test_questions_links ON test_questions_links.question_uuid = questions.uuid AND test_questions_links.test_uuid = ?", testUUID).
		Order("rank ASC").
		Find(&questions)

	if query.Error != nil && !query.RecordNotFound() {
		return []models.Question{}, &errors.ErrWhileHandling
	}

	return questions, nil
}

// ManyAnswers gets a mapping between questionUUIDs and their respective answers
func (c *coursesRepoImpl) ManyAnswers(questionUUIDs []gentypes.UUID) (map[gentypes.UUID][]models.BasicAnswer, error) {
	var answers []models.BasicAnswer
	query := database.GormDB.Where("question_uuid IN (?)", questionUUIDs).Order("question_uuid, rank ASC").Find(&answers)
	if query.Error != nil && !query.RecordNotFound() {
		c.Logger.Log(sentry.LevelError, query.Error, "Unable to get many answers")
		return map[gentypes.UUID][]models.BasicAnswer{}, &errors.ErrWhileHandling
	}

	var output = make(map[gentypes.UUID][]models.BasicAnswer)
	for _, answer := range answers {
		output[answer.QuestionUUID] = append(output[answer.QuestionUUID], answer)
	}

	return output, nil
}

// CourseTests gets all the tests in a course (including ones nested in modules), in no particular order
func (c *coursesRepoImpl) CourseTests(onlineCourseUUID gentypes.UUID) ([]models.Test, error) {
	// Get outer course structure
	structures, err := c.OnlineCourseStructure(onlineCourseUUID)
	if err != nil {
		return []models.Test{}, err
	}

	var testIDs []gentypes.UUID
	var moduleIDs []gentypes.UUID
	for _, item := range structures {
		if item.TestUUID != nil {
			testIDs = append(testIDs, *item.TestUUID)
		}
		if item.ModuleUUID != nil {
			moduleIDs = append(moduleIDs, *item.ModuleUUID)
		}
	}

	// Get get structures of all modules given
	moduleMap, err := c.ManyModuleItems(moduleIDs)
	if err != nil && err != &errors.ErrNotFound {
		c.Logger.Log(sentry.LevelWarning, err, "Unable to get module items")
		return []models.Test{}, &errors.ErrWhileHandling
	}

	for _, moduleItems := range moduleMap {
		for _, item := range moduleItems {
			if item.Type == gentypes.ModuleTest {
				testIDs = append(testIDs, item.UUID)
			}
		}
	}

	// Fetch the tests
	testMap, err := c.ManyTests(testIDs)
	if err != nil {
		c.Logger.Log(sentry.LevelWarning, err, "Unable to get tests")
		return []models.Test{}, &errors.ErrWhileHandling
	}

	var outputTests []models.Test
	for _, test := range testMap {
		outputTests = append(outputTests, test)
	}

	return outputTests, nil
}

func (c *coursesRepoImpl) CreateTestMarks(mark models.TestMark) error {
	err := database.GormDB.Create(&mark).Error
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to create test marks")
		return &errors.ErrWhileHandling
	}

	return nil
}
