import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import LessonPage from './LessonPage';

const CREATE_LESSON = gql`
  mutation CreateLesson(
    $name: String!
    $tags: [UUID!]
    $description: String!
    $bannerImageToken: String
    $voiceoverToken: String
    $transcript: String
    $video: VideoInput
  ) {
    createLesson(
      input: {
        name: $name
        tags: $tags
        description: $description
        bannerImageToken: $bannerImageToken
        voiceoverToken: $voiceoverToken
        transcript: $transcript
        video: $video
      }
    ) {
      uuid
    }
  }
`;

function CreateLesson({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'audioVideo', label: 'Audio/Video' }
  ];

  var initState = {
    name: '',
    description: '',
    tags: [],
    bannerImageURL: '',
    bannerImageToken: undefined,
    voiceoverToken: undefined,
    transcript: '',
    video: { type: 'WISTIA', url: '' }
  };

  const [state, setState] = useState(initState);
  const [createLesson, { error }] = useMutation(CREATE_LESSON);

  const updateState = updates => {
    var updatedState = { ...state, ...updates };
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
          type: state.type,
          complete: state.complete,
          description: state.description,
          tags: state.tags,
          bannerImageURL: state.bannerImageURL,
          voiceoverURL: state.voiceoverURL,
          transcript: state.transcript,
          video: state.video
        }
      });
      if (res.data?.createLesson?.uuid) {
        console.log(res.data);
        history.push(`/lesson/${res.data?.createLesson?.uuid}/overview`);
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
