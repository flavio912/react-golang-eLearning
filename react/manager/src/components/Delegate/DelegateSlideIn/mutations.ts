import { commitMutation, graphql } from 'react-relay';
import environment from 'api/environment';
import { GraphError } from 'types/general';
import { mutations_CreateDelegateMutationResponse } from './__generated__/mutations_CreateDelegateMutation.graphql';
import { DelegateInfo } from './DelegateSlideIn';

const mutation = graphql`
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
    mutation,
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
