import { commitMutation, graphql } from 'react-relay';
import environment from 'api/environment';
import { GraphError } from 'types/general';

const mutation = graphql`
  mutation mutations_CreateDelegateMutation(
    $firstName: String!
    $lastName: String!
    $jobTitle: String!
    $email: String!
    $phone: String!
  ) {
    createDelegate(
      input: {
        firstName: $firstName
        lastName: $lastName
        email: $email
        jobTitle: $jobTitle
        telephone: $phone
      }
    ) {
      firstName
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
  errorCallback: (err: string) => void
) => {
  const variables = delegate;

  commitMutation(environment, {
    mutation,
    variables,
    onCompleted: (
      response: { createDelegate: { firstName: string } },
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        errorCallback(`${errors[0]?.extensions?.message}`);
        return;
      }
      console.log('Response received from server.', response, errors);
    },
    onError: (err) => console.error(err)
  });
};
