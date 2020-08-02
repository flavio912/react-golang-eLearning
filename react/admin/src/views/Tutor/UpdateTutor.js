import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import ErrorModal from 'src/components/ErrorModal';
import TutorPage from './TutorPage';

const useStyles = makeStyles((theme) => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_TUTOR = gql`
  query GetTutor($uuid: UUID!) {
    tutor(uuid: $uuid) {
      uuid
      name
      cin
      signatureURL
    }
  }
`;

const UPDATE_TUTOR = gql`
  mutation UpdateTutor(
    $uuid: UUID!
    $name: String
    $cin: String
    $signatureToken: String
  ) {
    updateTutor(
      input: {
        uuid: $uuid
        name: $name
        cin: $cin
        signatureToken: $signatureToken
      }
    ) {
      uuid
    }
  }
`;

const initState = {
  name: '',
  cin: '',
  signatureURL: '',
  signatureToken: undefined,
};

function UpdateTutor({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' }
  ];

  const {
    loading, error, data: queryData, refetch
  } = useQuery(GET_TUTOR, {
    variables: {
      uuid: ident
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  const [updateTutor, { error: mutationErr }] = useMutation(UPDATE_TUTOR);

  const [state, setState] = useState(initState);

  const updateState = (stateUpdate) => {
    const updatedState = { ...state, ...stateUpdate };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;
    setState({
      ...initState,
      name: queryData.tutor.name,
      cin: queryData.tutor.cin,
      signatureURL: queryData.tutor.signatureURL,
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    try {
      await updateTutor({
        variables: {
          uuid: ident,
          name: state.name,
          cin: state.cin,
          signatureToken: state.signatureToken,
        }
      });
      refetch();
    } catch (err) {
      console.log('Failed to update:', err);
    }
  };

  return (
    <>
      <ErrorModal error={mutationErr} />
      <TutorPage
        state={state}
        setState={updateState}
        currentTab={currentTab}
        error={error || mutationErr}
        onSave={onUpdate}
        history={history}
        tabs={tabs}
        title="Edit Tutor"
      />
    </>
  );
}

export default UpdateTutor;
