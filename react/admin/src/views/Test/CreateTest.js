import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import TestPage from './TestPage';

const CREATE_TEST = gql`
  mutation CreateTest {
    createTest(input: {}) {
      test {
        uuid
      }
    }
  }
`;

function CreateTest({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'test-builder', label: 'Test Builder' }
  ];

  var initState = {
    name: '',
    randomise: false,
    answers: [],
    tags: []
  };

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
        variables: {}
      });
      if (res.data?.createTest?.test?.uuid) {
        history.push(`/test/${res.data?.createTest?.test?.uuid}`);
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
