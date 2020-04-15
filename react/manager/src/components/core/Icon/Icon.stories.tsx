import * as React from "react";
import Icon, { IconNames } from "./Icon";
import { withKnobs, select, number } from "@storybook/addon-knobs";

export default {
  title: "Core/Icon",
  decorators: [withKnobs],
};

const iconNames: IconNames[] = [
  "AddDelegateRepeater",
  "ArrowLeft",
  "ArrowRight",
  "Card_SecondaryActon_Dots",
  "CloseCourseManagementTray_X",
  "CourseAccountActivated",
  "CourseCertificates",
  "CourseExpiringSoon",
  "CourseFailed",
  "CourseNewCourse",
  "CourseStatus_Completed",
  "CourseStatus_Incomplete",
  "CourseStatus_NotStarted",
  "CourseTimeTracked",
  "Course_Calendar",
  "DownloadCSV",
  "EditDelegateProfilePic_Default",
  "EditPicPencil",
  "FilterAdjust",
  "FormCheckbox_Checked",
  "FormCheckbox_states",
  "FormCheckbox_Unchecked",
  "Icon_Delegates",
  "LeftNav_Icon_Courses",
  "LeftNav_Icon_Dashboard",
  "LeftNav_Icon_Delegates",
  "Loading_Screen_Donut",
  "Location_Pin",
  "PDF_Icon",
  "RemoveSelectedCourse_X",
  "SampleImage_ClassroomCoursesDetail_Feat",
  "SearchGlass",
  "TTC_Logo_Icon",
];

export const plain = () => {
  const name: IconNames = select(
    "Name",
    iconNames,
    "EditDelegateProfilePic_Default"
  );
  const size: number = number("Size", NaN)
  return <Icon name={name} size={size} />;
};
