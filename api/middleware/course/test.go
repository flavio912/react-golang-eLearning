package course

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/dbutils"

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

func (c *coursesRepoImpl) createQuestionLinks(tx *gorm.DB, testUUID gentypes.UUID, questionUUIDs []gentypes.UUID) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create question links
	for i, uuid := range questionUUIDs {
		link := models.TestQuestionsLink{
			TestUUID:     testUUID,
			QuestionUUID: uuid,
			Rank:         strconv.Itoa(i),
		}
		if err := tx.Create(&link).Error; err != nil {
			c.Logger.Log(sentry.LevelWarning, err, "Unable to create test question link")
			return &errors.ErrWhileHandling
		}
	}

	return nil
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

	err := c.createQuestionLinks(tx, test.UUID, input.Questions)
	if err != nil {
		tx.Rollback()
		return models.Test{}, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit test")
		return models.Test{}, &errors.ErrWhileHandling
	}

	return test, nil
}

type UpdateTestInput struct {
	UUID              gentypes.UUID
	Name              *string
	AttemptsAllowed   *uint
	PassPercentage    *float64
	QuestionsToAnswer *uint
	RandomiseAnswers  *bool
	Questions         *[]gentypes.UUID
	Tags              *[]gentypes.UUID
}

func (c *coursesRepoImpl) UpdateTest(input UpdateTestInput) (models.Test, error) {
	test, err := c.Test(input.UUID)
	if err != nil {
		c.Logger.Log(sentry.LevelWarning, err, "Unable to get test")
		return models.Test{}, &errors.ErrNotFound
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.AttemptsAllowed != nil {
		updates["attempts_allowed"] = *input.AttemptsAllowed
	}
	if input.PassPercentage != nil {
		updates["pass_percentage"] = *input.PassPercentage
	}
	if input.QuestionsToAnswer != nil {
		updates["questions_to_answer"] = *input.QuestionsToAnswer
	}
	if input.RandomiseAnswers != nil {
		updates["randomise_answers"] = *input.RandomiseAnswers
	}

	tx := database.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if input.Tags != nil {
		tags, err := c.CheckTagsExist(*input.Tags)
		if err != nil {
			return test, err
		}
		if err := tx.Model(&test).Association("Tags").Replace(tags).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to replace tags")
			return test, &errors.ErrWhileHandling
		}
	}

	if input.Questions != nil {
		// Remove old links + add new ones
		if err := tx.Where("test_uuid = ?", test.UUID).Delete(&models.TestQuestionsLink{}).Error; err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to delete test links")
			return test, &errors.ErrWhileHandling
		}

		err := c.createQuestionLinks(tx, test.UUID, *input.Questions)
		if err != nil {
			tx.Rollback()
			c.Logger.Log(sentry.LevelError, err, "Unable to create question links")
			return test, &errors.ErrWhileHandling
		}
	}

	if err := tx.Model(&test).Updates(updates).Error; err != nil {
		tx.Rollback()
		c.Logger.Log(sentry.LevelError, err, "Unable to update test")
		return test, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.Logger.Log(sentry.LevelError, err, "Unable to commit test update")
		return test, &errors.ErrWhileHandling
	}

	// Get the updated test
	updatedTest, err := c.Test(input.UUID)
	if err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to get test after update")
		return models.Test{}, &errors.ErrWhileHandling
	}

	return updatedTest, nil
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

func filterTest(query *gorm.DB, filter *gentypes.TestFilter) *gorm.DB {
	if filter != nil {
		if filter.UUID != nil {
			query = query.Where("uuid = ?", *filter.UUID)
		}

		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name ILIKE ?", "%%"+*filter.Name+"%%")
		}
	}

	return query
}

func (c *coursesRepoImpl) Tests(
	page *gentypes.Page,
	filter *gentypes.TestFilter,
	orderBy *gentypes.OrderBy,
) ([]models.Test, gentypes.PageInfo, error) {
	var tests []models.Test
	utils := dbutils.NewDBUtils(c.Logger)
	pageInfo, err := utils.GetPageOf(
		&models.Test{},
		&tests,
		page,
		orderBy,
		[]string{"created_at", "name"},
		"created_at DESC",
		func(db *gorm.DB) *gorm.DB {
			return filterTest(db, filter)
		},
	)
	pageInfo.Given = int32(len(tests))

	return tests, pageInfo, err
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

func (c *coursesRepoImpl) DeleteTest(uuid gentypes.UUID) (bool, error) {
	tx := database.GormDB.Begin()

	var course_structure models.CourseStructure
	if !tx.Model(&models.CourseStructure{}).Where("test_uuid = ?", uuid).Find(&course_structure).RecordNotFound() {
		glog.Info("Cannot delete test that is part of a course")
		return false, &errors.ErrWhileHandling
	}

	if err := tx.Delete(models.Test{}, "uuid = ?", uuid).Error; err != nil {
		c.Logger.Logf(sentry.LevelError, err, "Unable to delete test: %s", uuid)
		return false, &errors.ErrWhileHandling
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Log(sentry.LevelError, err, "Unable to commit transaction")
		return false, &errors.ErrWhileHandling
	}

	return true, nil
}
