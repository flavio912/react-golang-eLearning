import * as React from 'react';
import SideModal from 'components/core/Modals/SideModal';
import Tabs from 'components/core/Modals/SideModal/Tabs';
import { tabList } from './Tabs';

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

const multipleUsers = ({ isOpen, onClose }: Props) => {
  return (
    <SideModal
      isOpen={isOpen}
      title="Course Management"
      closeModal={() => onClose()}
    >
      <Tabs
        content={tabList}
        closeModal={() => onClose()}
        initialState={{
          course: undefined,
          users: [undefined],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export default multipleUsers;
