import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import TutorPage from './TutorPage';

const CREATE_TUTOR = gql`
  mutation CreateTutor(
    $name: String!
    $cin: String!
    $signatureToken: String
  ) {
    createTutor(
      input: {
        name: $name
        cin: $cin
        signatureToken: $signatureToken
      }
    ) {
        uuid
    }
  }
`;

function CreateTutor({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const initState = {
    name: '',
    cin: '',
    signatureToken: undefined,
  };

  const [state, setState] = useState(initState);
  const [createTutor, { error }] = useMutation(CREATE_TUTOR);

  const updateState = (item) => {
    const updatedState = { ...state, ...item };
    setState(updatedState);
  };

  if (!tabs.find((tab) => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const onSave = async () => {
    try {
      const res = await createTutor({
        variables: {
          name: state.name,
          cin: state.cin,
          signatureToken: state.signatureToken,
        }
      });
      if (res.data?.createTutor?.uuid) {
        history.push(`/tutor/${res.data?.createTutor?.uuid}/overview`);
      } else {
        console.warn('Unable to get save params');
      }
      console.log('REsp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  return (
    <TutorPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Tutor"
    />
  );
}

export default CreateTutor;
