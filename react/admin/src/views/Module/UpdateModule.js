import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import ModulePage from './ModulePage';

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
      video {
          type
          url
      }
      syllabus
      complete
    }
  }
`;

const UPDATE_MODULE = gql`
  mutation UpdateModule(
    $uuid: UUID!
    $name: String!
    $bannerImageURL: String
    $description: String!
    $transcript: String!
    $voiceoverURL: String
    $video: Video
    $syllabus: [SyllabusItem!]
    $complete: Boolean
  ) {
    createModule(
      input: {
        name: $name
        bannerImageURL: $bannerImageURL
        description: $description
        transcript: $transcript
        voiceoverURL: $voiceoverURL
        video: $video
        syllabus: $syllabus
        complete: $complete
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
    bannerImageURL: '',
    description: '',
    transcript: '',
    voiceoverURL: '',
    video: { type: 'WISTIA', url: ''},
    syllabus: [],
    complete: false,
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

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;
    setState({
      ...initState,
      name: queryData.module.name,
      bannerImageURL: queryData.module.bannerImageURL,
      description: queryData.module.description,
      transcript: queryData.module.transcript,
      voiceoverURL: queryData.module.voiceoverURL,
      video: queryData.module.video,
      syllabus: queryData.module.syllabus,
      complete: queryData.module.complete,
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
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
          syllabus: state.syllabus
        }
      });
      refetch();
    } catch (err) {}
  };

  const onSaveDraft = () => {};

  return (
    <ModulePage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error || mutationErr}
      onSaveDraft={onSaveDraft}
      onPublish={onUpdate}
      history={history}
      tabs={tabs}
      title="Edit Module"
    />
  );
}

export default UpdateModule;
