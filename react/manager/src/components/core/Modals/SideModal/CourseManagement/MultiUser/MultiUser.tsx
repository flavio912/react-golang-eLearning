import * as React from 'react';
import { commitMutation, graphql } from 'react-relay';
import { GraphError } from 'types/general';
import SideModal from 'components/core/Modals/SideModal';
import Tabs from 'components/core/Modals/SideModal/Tabs';
import { tabList } from './Tabs';
import environment from 'api/environment';
import { MultiUser_PurchaseMutationResponse } from './__generated__/MultiUser_PurchaseMutation.graphql';

type Props = {
  isOpen: boolean;
  onClose: Function;
};

const trainingInitialValue = [
  {
    label:
      'I Fred Eccs can confirm that the above delegates have undergone the necessary 5-year background checks for this Training Programme',
    checked: false
  }
];

const ToBInitialValue = [
  {
    label:
      'I Fred Eccs can confirm that I have read and understood the TTC Terms of Business',
    checked: false
  }
];

const mutation = graphql`
  mutation MultiUser_PurchaseMutation(
    $courses: [Int!]!
    $users: [UUID!]!
    $extraEmail: String
  ) {
    purchaseCourses(
      input: {
        courses: $courses
        users: $users
        extraInvoiceEmail: $extraEmail
        acceptedTerms: true
        backgroundCheckConfirm: true
      }
    ) {
      transactionComplete
      stripeClientSecret
    }
  }
`;

export type PurchaseCoursesType = (
  courses: number[],
  users: string[],
  extraEmail: string,
  callback?: (
    response: MultiUser_PurchaseMutationResponse,
    error: string | undefined
  ) => void
) => void;

const PurchaseCourses = (
  courses: number[],
  users: string[],
  extraEmail: string,
  callback?: (
    response: MultiUser_PurchaseMutationResponse,
    error: string | undefined
  ) => void
) => {
  const variables = {
    courses,
    users,
    extraEmail
  };

  commitMutation(environment, {
    mutation,
    variables,
    onCompleted: async (
      response: MultiUser_PurchaseMutationResponse,
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        if (callback) {
          callback(response, `${errors[0]?.extensions?.message}`);
        }
        return;
      }

      if (callback) {
        callback(response, undefined);
      }
      console.log('Response received from server.', response, errors);
    },
    onError: (err) => console.error(err)
  });
};

const multipleUsers = ({ isOpen, onClose }: Props) => {
  return (
    <SideModal
      isOpen={isOpen}
      title="Course Management"
      closeModal={() => onClose()}
    >
      <Tabs
        content={tabList(PurchaseCourses, false)}
        closeModal={() => onClose()}
        initialState={{
          courses: [],
          users: [undefined],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export default multipleUsers;
