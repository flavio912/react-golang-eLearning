import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import TestPage from './TestPage';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_QUESTION = gql`
  query GetQuestion($uuid: UUID!) {
    question(uuid: $uuid) {
      uuid
      text
      randomiseAnswers
      answers {
        uuid
        isCorrect
        text
        imageURL
      }
    }
  }
`;

const UPDATE_QUESTION = gql`
  mutation UpdateQuestion(
    $uuid: UUID!
    $text: String
    $randomise: Boolean
    $answers: [UpdateBasicAnswerInput!]
    $tags: [UUID!]
  ) {
    updateQuestion(
      input: {
        uuid: $uuid
        text: $text
        randomiseAnswers: $randomise
        answers: $answers
        tags: $tags
      }
    ) {
      question {
        uuid
      }
    }
  }
`;

const initState = {
  name: '',
  randomise: false,
  answers: [],
  tags: []
};

function UpdateTest({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const { loading, error, data: queryData, refetch } = useQuery(GET_QUESTION, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateQuestion, { error: mutationErr }] = useMutation(UPDATE_QUESTION);

  const [state, setState] = useState(initState);

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;
    setState({
      ...initState,
      name: queryData.question.text,
      randomise: queryData.question.randomiseAnswers,
      answers: queryData.question.answers.map(ans => {
        var display = 'TEXT';
        if (ans.text && ans.imageURL) {
          display = 'TEXT_IMAGE';
        } else if (ans.imageURL) {
          display = 'IMAGE';
        }
        return {
          uuid: ans.uuid,
          isCorrect: ans.isCorrect,
          text: ans.text,
          imageURL: ans.imageURL,
          answerType: display
        };
      })
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    try {
      await updateQuestion({
        variables: {
          uuid: ident,
          text: state.name,
          randomise: state.randomise,
          answers: state.answers.map(ans => ({
            uuid: ans.uuid || undefined,
            isCorrect: ans.isCorrect,
            text: ans.text,
            imageToken: ans.imageToken || undefined,
            answerType: ans.answerType
          }))
        }
      });
      refetch();
    } catch (err) {}
  };

  return (
    <TestPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error || mutationErr}
      onSave={onUpdate}
      history={history}
      tabs={tabs}
      title="Edit Test"
    />
  );
}

export default UpdateTest;
