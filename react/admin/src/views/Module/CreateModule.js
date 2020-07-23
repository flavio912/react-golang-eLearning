import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import ModulePage from './ModulePage';

const CREATE_MODULE = gql`
  mutation PublishModule(
    $name: String!
    $tags: [UUID!]
    $description: String!
    $transcript: String!
    $bannerImageSuccessToken: String
    $voiceoverSuccessToken: String
    $video: VideoInput
    $syllabus: [ModuleItem!]
  ) {
    createModule(
      input: {
        name: $name
        tags: $tags
        description: $description
        transcript: $transcript
        bannerImageSuccessToken: $bannerImageSuccessToken
        voiceoverSuccessToken: $voiceoverSuccessToken
        video: $video
        syllabus: $syllabus
      }
    ) {
      module {
        uuid
      }
    }
  }
`;

function CreateQuestion({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'audiovideo', label: 'Audio/Video' },
    { value: 'modulebuilder', label: 'Module Builder' }
  ];

  var initState = {
    name: '',
    tags: [],
    description: '',
    transcript: '',
    bannerImageSuccessToken: undefined,
    voiceoverSuccessToken: undefined,
    video: { type: 'WISTIA', url: '' },
    syllabus: [],
  };

  const [state, setState] = useState(initState);
  const [createModule, { error }] = useMutation(CREATE_MODULE);
  const updateState = (item) => {
    var updatedState = {...state, ...item};
    setState(updatedState);
  };

  const onSave = async () => {
    try {
      const res = await createModule({
        variables: {
          name: state.name,
          tags: state.tags,
          description: state.description,
          transcript: state.transcript,
          bannerImageSuccessToken: state.bannerImageSuccessToken,
          voiceoverSuccessToken: state.voiceoverSuccessToken,
          video: state.video,
          syllabus: state.syllabus.map(({ uuid, type }) => ({ type, uuid }))
        }
      });
      if (res.data?.createModule?.module?.uuid) {
        history.push(
          `/module/${res.data?.createModule?.module?.uuid}/overview`
        );
      } else {
        console.warn('Unable to get save params');
      }
      console.log('Resp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  return (
    <ModulePage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Module"
    />
  );
}

export default CreateQuestion;
