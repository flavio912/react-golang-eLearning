import { commitMutation, graphql } from 'react-relay';
import environment from 'api/environment';
import { GraphError } from 'types/general';
import { mutations_CreateDelegateMutationResponse } from './__generated__/mutations_CreateDelegateMutation.graphql';
import { DelegateInfo } from './DelegateSlideIn';
import { mutations_UpdateDelegateMutationResponse } from './__generated__/mutations_UpdateDelegateMutation.graphql';

const createMutation = graphql`
  mutation mutations_CreateDelegateMutation(
    $firstName: String!
    $lastName: String!
    $jobTitle: String!
    $email: String!
    $phone: String!
    $generatePassword: Boolean
  ) {
    createDelegate(
      input: {
        firstName: $firstName
        lastName: $lastName
        email: $email
        jobTitle: $jobTitle
        telephone: $phone
        generatePassword: $generatePassword
      }
    ) {
      delegate {
        firstName
        lastName
        jobTitle
        TTC_ID
        email
        telephone
      }
      generatedPassword
    }
  }
`;

export const CreateDelegate = (
  delegate: {
    firstName: string;
    lastName: string;
    jobTitle: string;
    email: string;
    phone: string;
  },
  generatePassword = false,
  errorCallback: (err: string) => void,
  successCallback: (response: mutations_CreateDelegateMutationResponse) => void
) => {
  const variables = {
    firstName: delegate.firstName,
    lastName: delegate.lastName,
    jobTitle: delegate.jobTitle,
    email: delegate.email,
    phone: delegate.phone,
    generatePassword: generatePassword
  };
  console.log('Variables', variables);
  commitMutation(environment, {
    mutation: createMutation,
    variables,
    onCompleted: (
      response: mutations_CreateDelegateMutationResponse,
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        errorCallback(`${errors[0]?.extensions?.message}`);
        return;
      }
      console.log('Response received from server.', response, errors);
      successCallback(response);
    },
    onError: (err) => console.error(err)
  });
};

const updateMutation = graphql`
  mutation mutations_UpdateDelegateMutation(
    $uuid: UUID!
    $firstName: String
    $lastName: String
    $jobTitle: String
    $email: String
    $phone: String
  ){
    updateDelegate(input: {
      uuid: $uuid
      firstName: $firstName
      lastName: $lastName
      email: $email
      jobTitle: $jobTitle
      telephone: $phone
    }) {
      firstName
      lastName
      jobTitle
      TTC_ID
      email
      telephone
    }
  }
`;

export const UpdateDelegate = (
  delegate: {
    uuid: string;
    firstName: string;
    lastName: string;
    jobTitle: string;
    email: string;
    phone: string;
  },
  errorCallback: (err: string) => void,
  successCallback: (response: mutations_UpdateDelegateMutationResponse) => void
) => {
  const variables = {
    uuid: delegate.uuid,
    firstName: delegate.firstName,
    lastName: delegate.lastName,
    jobTitle: delegate.jobTitle,
    email: delegate.email,
    phone: delegate.phone,
  };

  commitMutation(environment, {
    mutation: updateMutation,
    variables,
    onCompleted: (
      response: mutations_UpdateDelegateMutationResponse,
      errors: GraphError[]
    ) => {
      if (errors) {
        errorCallback(`${errors[0]?.extensions?.message}`);
        return;
      }

      successCallback(response);
    }
  });
};