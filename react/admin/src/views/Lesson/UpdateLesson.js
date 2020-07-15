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
      text
    }
  }
`;

const UPDATE_LESSON = gql`
  mutation UpdateLesson($uuid: UUID!, $name: String, $text: String) {
    updateLesson(input: { uuid: $uuid, name: $name, text: $text }) {
      uuid
      name
      text
    }
  }
`;

const initState = {
  name: '',
  text: '',
  tags: []
};

function UpdateLesson({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const { loading, error, data: queryData, refetch } = useQuery(GET_LESSON, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateLesson, { error: mutationErr }] = useMutation(UPDATE_LESSON);

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
      name: queryData.lesson.name,
      text: queryData.lesson.text
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    try {
      await updateLesson({
        variables: {
          uuid: ident,
          name: state.name,
          text: state.text
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
