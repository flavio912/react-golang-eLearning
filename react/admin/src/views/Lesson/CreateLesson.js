import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import LessonPage from './LessonPage';

const CREATE_LESSON = gql`
  mutation CreateLesson($name: String!, $text: String!) {
    createLesson(input: { name: $name, text: $text }) {
      uuid
      name
      text
    }
  }
`;

function CreateLesson({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  var initState = {
    name: '',
    description: '',
    tags: []
  };

  const [state, setState] = useState(initState);
  const [createLesson, { error }] = useMutation(CREATE_LESSON);

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const onSave = async () => {
    try {
      const res = await createLesson({
        variables: {
          name: state.name,
          text: state.description
        }
      });
      if (res.data?.createLesson?.uuid) {
        console.log(res.data);
        history.push('/lessons');
      } else {
        console.warn('Unable to get save params');
      }
      console.log('REsp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  return (
    <LessonPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Lesson"
    />
  );
}

export default CreateLesson;
