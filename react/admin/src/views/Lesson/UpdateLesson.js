import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import LessonPage from './LessonPage';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_LESSON = gql`
  query GetLesson($uuid: UUID!) {
    lesson(uuid: $uuid) {
      uuid
      name
      type
      complete
      description
      tags {
        uuid
        name
        color
      }
      bannerImageURL
      voiceoverURL
      transcript
      video {
        url
      }
    }
  }
`;

const UPDATE_LESSON = gql`
  mutation UpdateLesson(
    $uuid: UUID!,
    $name: String,
    $description: String
    $tags: [UUID!]
    $bannerImageToken: String
    $voiceoverToken: String
    $transcript: String
    $video: VideoInput
  ) {
    updateLesson(
      input: {
        uuid: $uuid,
        name: $name,
        description: $description,
        tags: $tags,
        bannerImageToken: $bannerImageToken,
        voiceoverToken: $voiceoverToken,
        transcript: $transcript,
        video: $video,
      }
    ) {
      uuid
    }
  }
`;

const initState = {
  name: '',
  description: '',
  tags: [],
  bannerImageURL: '',
  bannerImageToken: undefined,
  voiceoverToken: undefined,
  transcript: '',
  video: { type: 'WISTIA', url: '' },
};

function UpdateLesson({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'audioVideo', label: 'Audio/Video' }
  ];

  const { loading, error, data: queryData, refetch } = useQuery(GET_LESSON, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateLesson, { error: mutationErr }] = useMutation(UPDATE_LESSON);

  const [state, setState] = useState(initState);

  const updateState = updates => {
    var updatedState = { ...state, ...updates };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;
    setState({
      ...initState,
      name: queryData.lesson.name,
      tags: queryData.lesson.tags,
      bannerImageURL: queryData.lesson.bannerImageURL,
      description: queryData.lesson.description,
      transcript: queryData.lesson.transcript,
      voiceoverURL: queryData.lesson.voiceoverURL,
      video: queryData.lesson.video,
      complete: queryData.lesson.complete
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    console.log(state)
    try {
      await updateLesson({
        variables: {
          uuid: ident,
          name: state.name,
          type: state.type,
          complete: state.complete,
          description: state.description,
          tags: state.tags,
          bannerImageURL: state.bannerImageURL,
          voiceoverURL: state.voiceoverURL,
          transcript: state.transcript,
          video: state.video,
        }
      });
      refetch();
    } catch (err) {}
  };

  return (
    <LessonPage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error || mutationErr}
      onSave={onUpdate}
      history={history}
      tabs={tabs}
      title="Edit Lesson"
    />
  );
}

export default UpdateLesson;
