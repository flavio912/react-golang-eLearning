import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import TestPage from './TestPage';
import ErrorModal from 'src/components/ErrorModal';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_TEST = gql`
  query GetTest($uuid: UUID!) {
    test(uuid: $uuid) {
      uuid
      name
      tags {
        uuid
        name
        color
      }
      attemptsAllowed
      passPercentage
      questionsToAnswer
      randomiseAnswers
      questions {
        uuid
        text
      }
    }
  }
`;

const UPDATE_TEST = gql`
  mutation UpdateTest(
    $uuid: UUID!
    $name: String!
    $tags: [UUID!]
    $attemptsAllowed: Int!
    $passPercentage: Float!
    $questionsToAnswer: Int!
    $randomiseAnswers: Boolean!
    $questions: [UUID!]!
  ) {
    updateTest(
      input: {
        uuid: $uuid
        name: $name
        tags: $tags
        attemptsAllowed: $attemptsAllowed
        passPercentage: $passPercentage
        questionsToAnswer: $questionsToAnswer
        randomiseAnswers: $randomiseAnswers
        questions: $questions
      }
    ) {
      test {
        uuid
      }
    }
  }
`;

const initState = {
  name: '',
  tags: [],
  attemptsAllowed: false,
  passPercentage: false,
  questionsToAnswer: false,
  randomiseAnswers: false,
  questions: []
};

function UpdateTest({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const { loading, error, data: queryData, refetch } = useQuery(GET_TEST, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateTest, { error: mutationErr }] = useMutation(UPDATE_TEST);

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
      name: queryData.test.name,
      tags: queryData.test.tags,
      attemptsAllowed: queryData.test.attemptsAllowed,
      passPercentage: queryData.test.passPercentage,
      questionsToAnswer: queryData.test.questionsToAnswer,
      randomiseAnswers: queryData.test.randomiseAnswers,
      questions: queryData.test.questions || []
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    try {
      await updateTest({
        variables: {
          uuid: ident,
          name: state.name,
          tags: state.tags,
          attemptsAllowed: parseInt(state.attemptsAllowed),
          passPercentage: state.passPercentage,
          questionsToAnswer: state.questionsToAnswer,
          randomiseAnswers: state.randomiseAnswers,
          questions: state.questions
        }
      });
      refetch();
    } catch (err) {}
  };

  return (
    <>
      <ErrorModal error={error || mutationErr} />
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
    </>
  );
}

export default UpdateTest;
