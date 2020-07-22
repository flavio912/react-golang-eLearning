import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import ModulePage from './ModulePage';
import ErrorModal from 'src/components/ErrorModal';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_MODULE = gql`
  query GetModule($uuid: UUID!) {
    module(uuid: $uuid) {
      uuid
      name
      bannerImageURL
      description
      transcript
      voiceoverURL
      tags {
        uuid
        name
        color
      }
      video {
        type
        url
      }
      syllabus {
        name
        uuid
      }
      complete
    }
  }
`;

const UPDATE_MODULE = gql`
  mutation UpdateModule(
    $uuid: UUID!
    $tags: [UUID!]
    $name: String!
    $description: String!
    $transcript: String!
    $bannerImageSuccessToken: String
    $voiceoverSuccessToken: String
    $video: VideoInput
    $syllabus: [ModuleItem!]
  ) {
    updateModule(
      input: {
        uuid: $uuid
        tags: $tags
        name: $name
        bannerImageSuccessToken: $bannerImageSuccessToken
        description: $description
        transcript: $transcript
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

const initState = {
  name: '',
  tags: [],
  bannerImageSuccessToken: undefined,
  bannerImageURL: false,
  description: '',
  transcript: '',
  voiceoverSuccessToken: undefined,
  video: { type: 'WISTIA', url: '' },
  syllabus: [{ type: 'lesson', uuid: '00000000-0000-0000-0000-000000000002' }],
  complete: false
};

function UpdateModule({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'audiovideo', label: 'Audio/Video' },
    { value: 'modulebuilder', label: 'Module Builder' }
  ];

  const { loading, error, data: queryData, refetch } = useQuery(GET_MODULE, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateModule, { error: mutationErr }] = useMutation(UPDATE_MODULE);

  const [state, setState] = useState(initState);

  const updateState = stateUpdate => {
    var updatedState = { ...state, ...stateUpdate };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;
    setState({
      ...initState,
      name: queryData.module.name,
      tags: queryData.module.tags,
      bannerImageSuccessToken: queryData.module.bannerImageSuccessToken,
      description: queryData.module.description,
      transcript: queryData.module.transcript,
      voiceoverSuccessToken: queryData.module.voiceoverSuccessToken,
      video: queryData.module.video,
      syllabus: queryData.module.syllabus,
      complete: queryData.module.complete
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    console.log(state.tags);
    try {
      await updateModule({
        variables: {
          uuid: ident,
          name: state.name,
          tags: state.tags,
          description: state.description,
          transcript: state.transcript,
          bannerImageSuccessToken: state.bannerImageSuccessToken,
          voiceoverSuccessToken: state.voiceoverSuccessToken,
          video: state.video,
          syllabus: state.syllabus.map(({ uuid }) => ({
            type: 'lesson',
            uuid
          })),
          complete: state.complete
        }
      });
      refetch();
    } catch (err) {}
  };

  const onPublish = () => {};

  return (
    <>
      <ErrorModal error={mutationErr} />
      <ModulePage
        state={state}
        setState={updateState}
        currentTab={currentTab}
        error={error || mutationErr}
        onSave={onUpdate}
        history={history}
        tabs={tabs}
        title="Edit Module"
      />
    </>
  );
}

export default UpdateModule;
