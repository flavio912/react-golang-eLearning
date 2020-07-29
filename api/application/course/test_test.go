package course_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	courseMocks "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/course/mocks"
	userMocks "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware/user/mocks"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

var (
	adminGrant = middleware.Grant{
		Claims: auth.UserClaims{
			UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Role:    auth.AdminRole,
		},
		Logger:  logging.Logger{},
		IsAdmin: true,
	}
	managerGrant = middleware.Grant{
		Claims: auth.UserClaims{
			UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Role:    auth.ManagerRole,
		},
		Logger:    logging.Logger{},
		IsManager: true,
	}
	delegateGrant = middleware.Grant{
		Claims: auth.UserClaims{
			UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Role:    auth.DelegateRole,
		},
		Logger:     logging.Logger{},
		IsDelegate: true,
	}
	individualGrant = middleware.Grant{
		Claims: auth.UserClaims{
			UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000012"),
			Role: auth.IndividualRole,
		},
		Logger:       logging.Logger{},
		IsIndividual: true,
	}
	publicGrant = middleware.Grant{IsPublic: true}
)

func TestSubmitTest(t *testing.T) {

	q1 := gentypes.MustParseToUUID("2d4e2e2a-6d9b-4de8-b9d0-bfd3323222bd")
	q2 := gentypes.MustParseToUUID("7ae7184c-caa7-42b1-8d0b-abcbf4f669ce")

	answers := []gentypes.QuestionAnswer{
		gentypes.QuestionAnswer{
			QuestionUUID: q1,
			AnswerUUID:   gentypes.MustParseToUUID("24a084b2-aed5-40d2-abb3-633bdeb3cbde"),
		},
		gentypes.QuestionAnswer{
			QuestionUUID: q2,
			AnswerUUID:   gentypes.MustParseToUUID("f2038eb8-3da3-469d-888b-6d9f429f318f"),
		},
	}

	input1Correct := gentypes.SubmitTestInput{
		CourseID: 3,
		TestUUID: gentypes.MustParseToUUID("94c12222-ba61-4547-a230-7d5f076be492"),
		Answers:  answers,
	}

	answersNotEnough := []gentypes.QuestionAnswer{
		gentypes.QuestionAnswer{
			QuestionUUID: q1,
			AnswerUUID:   gentypes.MustParseToUUID("24a084b2-aed5-40d2-abb3-633bdeb3cbde"),
		},
	}

	inputNotEnoughQuestions := gentypes.SubmitTestInput{
		CourseID: 3,
		TestUUID: gentypes.MustParseToUUID("94c12222-ba61-4547-a230-7d5f076be492"),
		Answers:  answersNotEnough,
	}

	testItems := []struct {
		name            string
		testMarks       []models.TestMark
		hasActiveCourse bool
		input           gentypes.SubmitTestInput
		grant           middleware.Grant
		success         bool
		err             error
	}{
		{
			name:            "Works on delegate grant",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           delegateGrant,
			success:         true,
		},
		{
			name:            "Works on individual grant",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           individualGrant,
			success:         true,
		},
		{
			name:            "Managers can't submit tests",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           managerGrant,
			success:         false,
			err:             &errors.ErrUnauthorized,
		},
		{
			name:            "Not assigned to course",
			testMarks:       []models.TestMark{},
			hasActiveCourse: false,
			input:           input1Correct,
			grant:           delegateGrant,
			success:         false,
			err:             &errors.ErrUnauthorized,
		},
		{
			name: "Already submitted test",
			testMarks: []models.TestMark{
				models.TestMark{
					TestUUID: input1Correct.TestUUID,
				},
			},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           delegateGrant,
			success:         true,
			err:             nil,
		},
		{
			name:            "Not enough questions given to complete",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           inputNotEnoughQuestions,
			grant:           delegateGrant,
			success:         false,
			err:             &errors.ErrNotEnoughAnswersGiven,
		},
		{
			name:            "Answers given not part of same test",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input: gentypes.SubmitTestInput{
				CourseID: 3,
				TestUUID: gentypes.MustParseToUUID("94c12222-ba61-4547-a230-7d5f076be492"),
				Answers: []gentypes.QuestionAnswer{
					gentypes.QuestionAnswer{
						QuestionUUID: q1,
						AnswerUUID:   gentypes.MustParseToUUID("95f846c6-df0e-4654-bd11-595126fac82a"),
					},
					gentypes.QuestionAnswer{
						QuestionUUID: gentypes.MustParseToUUID("1b66d76e-2d02-408f-8475-651d7b1207d6"),
						AnswerUUID:   gentypes.MustParseToUUID("fd2b484b-dd37-475a-a7d3-fa7004723903"),
					},
					gentypes.QuestionAnswer{
						QuestionUUID: gentypes.MustParseToUUID("1b66d76e-2d02-408f-8475-651d7b1207d6"),
						AnswerUUID:   gentypes.MustParseToUUID("fd2b484b-dd37-475a-a7d3-fa7004723903"),
					},
				},
			},
			grant:   delegateGrant,
			success: false,
			err:     &errors.ErrNotEnoughAnswersGiven,
		},
		{
			name:            "Repeated correct answer still just gives one mark",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input: gentypes.SubmitTestInput{
				CourseID: 3,
				TestUUID: gentypes.MustParseToUUID("94c12222-ba61-4547-a230-7d5f076be492"),
				Answers: []gentypes.QuestionAnswer{
					gentypes.QuestionAnswer{
						QuestionUUID: q1,
						AnswerUUID:   gentypes.MustParseToUUID("95f846c6-df0e-4654-bd11-595126fac82a"),
					},
					gentypes.QuestionAnswer{
						QuestionUUID: q2,
						AnswerUUID:   gentypes.MustParseToUUID("e31ac3ab-b7fa-46dc-af3f-c61297b6d77d"),
					},
					gentypes.QuestionAnswer{
						QuestionUUID: q2,
						AnswerUUID:   gentypes.MustParseToUUID("e31ac3ab-b7fa-46dc-af3f-c61297b6d77d"),
					},
					gentypes.QuestionAnswer{
						QuestionUUID: q2,
						AnswerUUID:   gentypes.MustParseToUUID("e31ac3ab-b7fa-46dc-af3f-c61297b6d77d"),
					},
				},
			},
			grant:   delegateGrant,
			success: true,
		},
		{
			name:            "Admins can't submit tests",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           adminGrant,
			success:         false,
			err:             &errors.ErrUnauthorized,
		},
		{
			name:            "Public can't submit tests",
			testMarks:       []models.TestMark{},
			hasActiveCourse: true,
			input:           input1Correct,
			grant:           publicGrant,
			success:         false,
			err:             &errors.ErrUnauthorized,
		},
	}

	for _, testItem := range testItems {
		t.Run(testItem.name, func(t *testing.T) {
			// Mock the repositories
			var coursesRepo = new(courseMocks.CoursesRepository)
			var usersRepo = new(userMocks.UsersRepository)

			takerUUID := gentypes.MustParseToUUID("98ba9cf8-fddb-4ec1-866f-3f4740d30e98")

			if testItem.grant.IsDelegate {
				usersRepo.On("Delegate", testItem.grant.Claims.UUID).Return(models.Delegate{
					Base: models.Base{
						UUID: testItem.grant.Claims.UUID,
					},
					CourseTakerUUID: takerUUID,
				}, nil)
			}
			if testItem.grant.IsIndividual {
				usersRepo.On("Individual", testItem.grant.Claims.UUID).Return(models.Individual{
					Base: models.Base{
						UUID: testItem.grant.Claims.UUID,
					},
					CourseTakerUUID: takerUUID,
				}, nil)
			}

			usersRepo.On("TakerTestMarks", takerUUID, uint(testItem.input.CourseID)).Return(testItem.testMarks, nil)

			// Setup test with two questions and two answers in each
			coursesRepo.On("TestQuestions", testItem.input.TestUUID).Return([]models.Question{
				models.Question{
					UUID: q1,
				},
				models.Question{
					UUID: q2,
				},
			}, nil)

			questionsToAnswers := map[gentypes.UUID][]models.BasicAnswer{
				q1: []models.BasicAnswer{
					models.BasicAnswer{
						UUID:         gentypes.MustParseToUUID("95f846c6-df0e-4654-bd11-595126fac82a"),
						QuestionUUID: q1,
						IsCorrect:    false,
					},
					models.BasicAnswer{
						UUID:         gentypes.MustParseToUUID("24a084b2-aed5-40d2-abb3-633bdeb3cbde"),
						QuestionUUID: q1,
						IsCorrect:    true,
					},
				},
				q2: []models.BasicAnswer{
					models.BasicAnswer{
						UUID:         gentypes.MustParseToUUID("e31ac3ab-b7fa-46dc-af3f-c61297b6d77d"),
						QuestionUUID: q2,
						IsCorrect:    true,
					},
					models.BasicAnswer{
						UUID:         gentypes.MustParseToUUID("f2038eb8-3da3-469d-888b-6d9f429f318f"),
						QuestionUUID: q2,
						IsCorrect:    false,
					},
				},
			}

			if testItem.hasActiveCourse {
				usersRepo.On("TakerActiveCourse", mock.Anything, mock.Anything).Return(models.ActiveCourse{
					Status:         gentypes.CourseIncomplete,
					MinutesTracked: 23,
				}, nil)
			} else {
				usersRepo.On("TakerActiveCourse", mock.Anything, mock.Anything).Return(models.ActiveCourse{}, &errors.ErrNotFound)
			}

			coursesRepo.On("Test", testItem.input.TestUUID).Return(models.Test{
				UUID:              testItem.input.TestUUID,
				QuestionsToAnswer: 2,
				AttemptsAllowed:   1,
				PassPercentage:    9,
			}, nil)
			coursesRepo.On("ManyAnswers", mock.Anything).Return(questionsToAnswers, nil)

			onlineCourseUUID := gentypes.MustParseToUUID("1fe014a2-2633-4103-94fa-ceb514141e4b")
			coursesRepo.On("OnlineCourse", uint(testItem.input.CourseID)).Return(models.OnlineCourse{
				Base: models.Base{
					UUID: onlineCourseUUID,
				},
			}, nil)
			coursesRepo.On("CourseTests", onlineCourseUUID).Return([]models.Test{
				models.Test{},
			}, nil)
			coursesRepo.On("TakerTestMarks", takerUUID, uint(testItem.input.CourseID)).Return([]models.TestMark{
				models.TestMark{},
				models.TestMark{},
			}, nil)
			usersRepo.On("SaveTestMarks", mock.Anything).Return(nil)
			marks := models.TestMark{
				TestUUID:        testItem.input.TestUUID,
				CourseTakerUUID: takerUUID,
				CourseID:        uint(testItem.input.CourseID),
				Passed:          testItem.success,
				CurrentAttempt:  1,
			}

			var courseApp = course.NewCourseApp(&testItem.grant)
			courseApp.SetCoursesRepository(coursesRepo)
			courseApp.SetUsersRepository(usersRepo)

			success, stat, err := courseApp.SubmitTest(testItem.input)
			assert.Equal(t, testItem.success, success)
			assert.Equal(t, testItem.err, err)

			if !testItem.success && testItem.err == nil {
				assert.Equal(t, gentypes.CourseFailed, stat)
			} else {
				assert.Equal(t, gentypes.CourseIncomplete, stat)
			}

			if testItem.success {
				usersRepo.AssertCalled(t, "SaveTestMarks", marks)
			} else {
				usersRepo.AssertNotCalled(t, "SaveTestMarks")
			}
		})
	}
}
