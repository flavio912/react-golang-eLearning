import * as React from 'react';
import SideModal from 'sharedComponents/SideModal/SideModal';
import Tabs from 'sharedComponents/SideModal/Tabs';
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

const SingleUser = ({ isOpen, onClose }: Props) => {
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
          courses: [],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export default SingleUser;
