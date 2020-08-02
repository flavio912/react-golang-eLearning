import * as React from 'react';
import SideModal from 'sharedComponents/SideModal/SideModal';
import Tabs from 'sharedComponents/SideModal/Tabs';
import { tabList } from './Tabs';
import { Course } from 'sharedComponents/core/Input/SearchableDropdown/SearchableDropdown';


type Props = {
  isOpen: boolean;
  onClose: Function;
  courses?: Course[];
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

const multipleUsers = ({ isOpen, onClose, courses }: Props) => {
  return (
    <SideModal
      isOpen={isOpen}
      title="Course Management"
      closeModal={() => onClose()}
    >
      <Tabs
        content={tabList(false)}
        closeModal={() => onClose()}
        initialState={{
          courses: courses ? courses : [],
          users: [undefined],
          training: trainingInitialValue,
          ToB: ToBInitialValue
        }}
      />
    </SideModal>
  );
};

export default multipleUsers;
