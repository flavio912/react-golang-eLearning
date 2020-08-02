import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import CertificateTypePage from './CertificateTypePage';

const CREATE_CERTIFICATE_TYPE = gql`
  mutation CreateCertificateType(
    $name: String!
    $regulationText: String!
    $requiresCAANo: Boolean
    $showTrainingSection: Boolean
    $certificateBodyToken: String
  ) {
    createCertificateType(
      input: {
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

function CreateCertificateType({ match, history }) {
  const { tab: currentTab } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  var initState = {
    name: '',
    regulationText: '',
    requiresCAANo: false,
    showTrainingSection: false,
    certificateBodyToken: undefined
  };

  const [state, setState] = useState(initState);
  const [createCertificateType, { error }] = useMutation(
    CREATE_CERTIFICATE_TYPE
  );

  const updateState = updates => {
    var updatedState = { ...state, ...updates };
    setState(updatedState);
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const onSave = async () => {
    try {
      const res = await createCertificateType({
        variables: {
          name: state.name,
          regulationText: state.regulationText,
          requiresCAANo: state.requiresCAANo,
          showTrainingSection: state.showTrainingSection,
          certificateBodyToken: state.certificateBodyToken
        }
      });

      const uuid = res.data?.createCertificateType?.uuid;
      if (uuid) {
        history.push(`/certificateTypes/${uuid}/overview`);
      } else {
        console.warn('Unable to get save params');
      }
      console.log('REsp', res);
    } catch ({ graphQLErrors, networkError }) {
      console.log('ERR', graphQLErrors);
    }
  };

  return (
    <CertificateTypePage
      state={state}
      setState={updateState}
      currentTab={currentTab}
      error={error}
      onSave={onSave}
      history={history}
      tabs={tabs}
      title="Create Certificate Type"
    />
  );
}

export default CreateCertificateType;
