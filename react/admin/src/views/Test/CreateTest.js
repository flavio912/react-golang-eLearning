import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import TestPage from './TestPage';

const CREATE_TEST = gql`
  mutation CreateTest(
    $name: String!
    $tags: [UUID!]
    $attemptsAllowed: Int!
    $passPercentage: Float!
    $questionsToAnswer: Int!
    $randomiseAnswers: Boolean!
    $questions: [UUID!]!
  ) {
    createTest(
      input: {
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

var initState = {
  name: '',
  tags: [],
  attemptsAllowed: false,
  passPercentage: false,
  questionsToAnswer: false,
  randomiseAnswers: undefined,
  randomiseAnswers: false,
  questions: []
};

function CreateTest({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'test-builder', label: 'Test Builder' }
  ];

  const [state, setState] = useState(initState);
  const [createTest, { error }] = useMutation(CREATE_TEST);

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const onSave = async () => {
    try {
      const res = await createTest({
        variables: {
          name: state.name,
          tags: state.tags,
          attemptsAllowed: state.attemptsAllowed,
          passPercentage: state.passPercentage,
          questionsToAnswer: state.questionsToAnswer,
          randomiseAnswers: state.randomiseAnswers,
          questions: state.questions
        }
      });
      if (res.data?.createTest?.test?.uuid) {
        history.push(`/test/${res.data?.createTest?.test?.uuid}/overview`);
      } else {
        console.warn('Unable to get save params');
      }
      console.log('REsp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  return (
    <TestPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Test"
    />
  );
}

export default CreateTest;
