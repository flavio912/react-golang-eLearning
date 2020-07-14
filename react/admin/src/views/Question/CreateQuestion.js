import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import QuestionPage from './QuestionPage';

const CREATE_QUESTION = gql`
  mutation CreateQuestion(
    $text: String!
    $randomise: Boolean!
    $answers: [CreateBasicAnswerInput!]!
    $tags: [UUID!]!
  ) {
    createQuestion(
      input: {
        text: $text
        randomiseAnswers: $randomise
        questionType: SINGLE_CHOICE
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

function CreateQuestion({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  var initState = {
    name: '',
    randomise: false,
    answers: [],
    tags: []
  };

  const [state, setState] = useState(initState);
  const [createQuestion, { error }] = useMutation(CREATE_QUESTION);

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const onSave = async () => {
    try {
      const res = await createQuestion({
        variables: {
          text: state.name,
          randomise: state.randomise,
          answers: state.answers,
          tags: state.tags
        }
      });
      if (res.data?.createQuestion?.question?.uuid) {
        history.push(`/question/${res.data?.createQuestion?.question?.uuid}`);
      } else {
        console.warn('Unable to get save params');
      }
      console.log('REsp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  return (
    <QuestionPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Question"
    />
  );
}

export default CreateQuestion;
