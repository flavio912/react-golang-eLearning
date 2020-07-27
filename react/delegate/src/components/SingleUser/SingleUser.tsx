import * as React from 'react';
import SideModal from 'sharedComponents/SideModal/SideModal';
import Tabs from 'sharedComponents/SideModal/Tabs';
import { tabList } from './Tabs';
import { Course } from 'sharedComponents/core/Input/SearchableDropdown';

type Props = {
  isOpen: boolean;
  onClose: Function;
  initialCourses?: Course[];
  userUUID?: string;
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

const SingleUser = ({
  isOpen,
  onClose,
  initialCourses = [],
  userUUID
}: Props) => {
  return (
    <SideModal
      isOpen={isOpen}
      title="Book a new course"
      closeModal={() => onClose()}
    >
      <Tabs
        content={tabList}
        closeModal={() => onClose()}
        initialState={{
          courses: initialCourses,
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export default SingleUser;
