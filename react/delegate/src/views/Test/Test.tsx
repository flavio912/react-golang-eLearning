import React, { useState } from 'react';
import { createUseStyles, useTheme } from 'react-jss';
// import { Grid } from '@material-ui/core';
import { Theme } from 'helpers/theme';
import Heading from 'components/core/Heading';
import Question from 'components/Misc/Question';
import Button from 'components/core/Input/Button';
import Page from 'components/Page';
import { createFragmentContainer, graphql, commitMutation } from 'react-relay';
import { Test_test } from './__generated__/Test_test.graphql';
import { Test_myActiveCourse } from './__generated__/Test_myActiveCourse.graphql';
import { stringify } from 'querystring';
import environment from 'api/environment';
import { GraphError } from 'types/general';
import CourseSlider from 'components/Overview/CourseSlider';
import { Test_SubmitAnswersMutationVariables } from './__generated__/Test_SubmitAnswersMutation.graphql';
import { Match } from 'found';

const useStyles = createUseStyles((theme: Theme) => ({
  questionsRoot: {
    display: 'flex',
    flexGrow: 1,
    maxWidth: 1275
  },
  viewPanel: {
    flex: 2
  },
  infoPanel: {
    flex: 1,
    marginLeft: 30,
    borderTop: '1px solid',
    paddingTop: 30
  },
  mainHeading: {
    gridArea: 'headin'
  },
  subHeading: {
    gridArea: 'subhea',
    maxWidth: 466
  },
  nextQuestionWrap: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: 150,
    background: 'white',
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 6
  },
  infoHeading: {
    color: `${theme.colors.primaryGreen}`
  },
  questionStep: {
    margin: 0,
    letterSpacing: 0
  }
}));

const mutation = graphql`
  mutation Test_SubmitAnswersMutation(
    $courseID: Int!
    $testUUID: UUID!
    $answers: [QuestionAnswer!]!
  ) {
    submitTest(
      input: { courseID: $courseID, testUUID: $testUUID, answers: $answers }
    ) {
      courseStatus
      passed
    }
  }
`;

const submitAnswers = (
  courseID: Test_SubmitAnswersMutationVariables['courseID'],
  testUUID: Test_SubmitAnswersMutationVariables['testUUID'],
  answers: Test_SubmitAnswersMutationVariables['answers']
) => {
  const variables: Test_SubmitAnswersMutationVariables = {
    courseID,
    testUUID,
    answers
  };

  return new Promise((resolve, reject) => {
    commitMutation(environment, {
      mutation,
      variables,
      onCompleted: (
        response: { delegateLogin: { token: string } },
        errors: GraphError[]
      ) => {
        if (errors) {
          // Display error
          reject(`${errors[0]?.extensions?.message}`);
          return;
        }
        console.log('Response received from server.', response, errors);
        resolve(response);
      },
      onError: (err) => {
        reject(err);
      }
    });
  });
};

type Props = {
  className?: string;
  test?: Test_test;
  myActiveCourse?: Test_myActiveCourse;
  match: Match;
};

function Test({ className, test, myActiveCourse, match }: Props) {
  const { courseID, testUUID } = match.params;

  const questions = (test?.questions ?? []).map((question, index) => {
    return {
      id: index,
      uuid: question.uuid,
      type: 'text',
      head: '',
      title: question.text,
      options: (question.answers ?? []).map((answer, index) => {
        return {
          id: answer.uuid,
          title: answer.text ?? '',
          image: answer.imageURL ?? undefined,
          index: String(index)
        };
      })
    };
  });

  const theme = useTheme();
  const classes = useStyles({ theme });
  const [curQuestion, setCurQuestion] = useState(questions[0]);
  const [answers, setAnswers] = useState<{ [key: string]: string }>({});
  const handleNextQuestion = async () => {
    if (curQuestion.id >= questions.length - 1) {
      // Submit answers
      const normalisedAnswers = Object.keys(answers).map((key) => ({
        questionUUID: key,
        answerUUID: answers[key]
      }));

      try {
        const resp = await submitAnswers(
          parseInt(courseID),
          testUUID,
          normalisedAnswers
        );
        console.log('RESP', resp);
        setAnswers({});
      } catch (err) {
        console.error('Unable to submit answers', err);
      }
      return;
    }

    setCurQuestion(questions[curQuestion.id + 1]);
  };

  return (
    <Page>
      <div className={classes.questionsRoot}>
        <div className={classes.viewPanel}>
          <Heading
            text={test?.name ?? ''}
            size={'large'}
            className={classes.mainHeading}
          />
          <Question
            question={curQuestion}
            type="text"
            onSelected={(option) => {
              setAnswers({
                ...answers,
                [curQuestion.uuid]: option.id
              });
            }}
          />
          <div className={classes.nextQuestionWrap}>
            <Button
              title={
                curQuestion.id >= questions.length - 1
                  ? 'Submit Answers'
                  : 'Next Question'
              }
              padding="large"
              onClick={handleNextQuestion}
            />
          </div>
        </div>
        <div className={classes.infoPanel}>
          <span className={classes.infoHeading}>Question</span>
          <h1 className={classes.questionStep}>
            {curQuestion.id + 1} / {test?.questions?.length}
          </h1>
        </div>
      </div>
    </Page>
  );
}

export default createFragmentContainer(Test, {
  myActiveCourse: graphql`
    fragment Test_myActiveCourse on MyCourse {
      course {
        ...CourseSyllabusCardFrag_course
      }
    }
  `,
  test: graphql`
    fragment Test_test on Test {
      name
      uuid
      questions {
        uuid
        text
        questionType
        answers {
          uuid
          text
          imageURL
        }
      }
    }
  `
});
