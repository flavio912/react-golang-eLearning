import * as React from "react";
import QuickInfo from "./QuickInfo";
import {
  withKnobs,
  text,
  select,
  number,
  boolean,
  object,
} from "@storybook/addon-knobs";
import { IconNames } from "sharedComponents/core/Icon";

export default {
  title: "Overview/Quick Info",
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
  "LeftNav_Icon_Courses",
  "LeftNav_Icon_Dashboard",
  "LeftNav_Icon_Delegates",
  "Loading_Screen_Donut",
  "Location_Pin",
  "PDF_Icon",
  "RemoveSelectedCourse_X",
  "SampleImage_ClassroomCoursesDetail_Feat",
  "SearchGlass",
];

export const normal = () => {
  const value = number("Value", 100);
  const valueTime = object("Value (time)", { h: 2, m: 20 });
  return (
    <QuickInfo
      icon={select("Icon", iconNames, "CourseExpiringSoon")}
      text={text("Text", "Expiring Soon")}
      value={boolean("UseTime", false) ? value : valueTime}
      footer={text("FooterText", "in March 2020")}
      valueArrow={text("ValueArrow", "up")}
    />
  );
};
