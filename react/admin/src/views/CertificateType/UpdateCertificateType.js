import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { CircularProgress } from '@material-ui/core';
import CertificateTypePage from './CertificateTypePage';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  }
}));

const GET_CERTIFICATE_TYPE = gql`
  query GetCertificateType($uuid: UUID!) {
    certificateType(uuid: $uuid) {
      uuid
      name
      certificateBodyImageURL
      regulationText
      requiresCAANo
      showTrainingSection
    }
  }
`;

const UPDATE_CERTIFICATE_TYPE = gql`
  mutation UpdateQuestion(
    $uuid: UUID!
    $name: String
    $regulationText: String
    $requiresCAANo: Boolean
    $showTrainingSection: Boolean
    $certificateBodyToken: String
  ) {
    updateCertificateType(
      input: {
        uuid: $uuid
        name: $name
        regulationText: $regulationText
        requiresCAANo: $requiresCAANo
        showTrainingSection: $showTrainingSection
        certificateBodyToken: $certificateBodyToken
      }
    ) {
      uuid
    }
  }
`;

var initState = {
  name: '',
  regulationText: '',
  requiresCAANo: false,
  showTrainingSection: false,
  certificateBodyToken: undefined
};

function UpdateCertificateType({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const { loading, error, data: queryData, refetch } = useQuery(
    GET_CERTIFICATE_TYPE,
    {
      variables: {
        uuid: ident
      },
      fetchPolicy: 'cache-and-network',
      skip: !ident
    }
  );
  const [updateCertificateType, { error: mutationErr }] = useMutation(
    UPDATE_CERTIFICATE_TYPE
  );

  const [state, setState] = useState(initState);

  const updateState = updates => {
    var updatedState = { ...state, ...updates };
    setState(updatedState);
  };

  useEffect(() => {
    if (loading || error) return;
    if (!queryData) return;

    const {
      name,
      regulationText,
      requiresCAANo,
      showTrainingSection,
      certificateBodyImageURL
    } = queryData.certificateType;
    setState({
      ...initState,
      name,
      regulationText,
      requiresCAANo,
      showTrainingSection,
      certificateBodyImageURL
    });
  }, [queryData, loading, error]);

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const onUpdate = async () => {
    try {
      const {
        name,
        regulationText,
        requiresCAANo,
        showTrainingSection,
        certificateBodyToken
      } = state;
      console.log('state', state);
      await updateCertificateType({
        variables: {
          uuid: ident,
          name,
          regulationText,
          requiresCAANo,
          showTrainingSection,
          certificateBodyToken
        }
      });
      refetch();
    } catch (err) {}
  };

  return (
    <CertificateTypePage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error || mutationErr}
      onSave={onUpdate}
      history={history}
      tabs={tabs}
      title="Edit Certificate Type"
    />
  );
}

export default UpdateCertificateType;
