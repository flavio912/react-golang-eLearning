import * as React from "react";

export type IconNames =
  | "AddDelegateRepeater"
  | "ArrowLeft"
  | "ArrowRight"
  | "Card_SecondaryActon_Dots"
  | "CloseCourseManagementTray_X"
  | "CourseAccountActivated"
  | "CourseCertificates"
  | "CourseExpiringSoon"
  | "CourseFailed"
  | "CourseNewCourse"
  | "CourseStatus_Completed"
  | "CourseStatus_Incomplete"
  | "CourseStatus_NotStarted"
  | "CourseTimeTracked"
  | "Course_Calendar"
  | "DownloadCSV"
  | "EditDelegateProfilePic_Default"
  | "EditPicPencil"
  | "FilterAdjust"
  | "FormCheckbox_Checked"
  | "FormCheckbox_states"
  | "FormCheckbox_Unchecked"
  | "LeftNav_Icon_Courses"
  | "LeftNav_Icon_Dashboard"
  | "LeftNav_Icon_Delegates"
  | "Loading_Screen_Donut"
  | "Location_Pin"
  | "PDF_Icon"
  | "RemoveSelectedCourse_X"
  | "SampleImage_ClassroomCoursesDetail_Feat"
  | "SearchGlass";

const iconNameMap = {
  AddDelegateRepeater: require("../../../assets/AddDelegateRepeater.svg"),
  ArrowLeft: require("../../../assets/ArrowLeft.svg"),
  ArrowRight: require("../../../assets/ArrowRight.svg"),
  Card_SecondaryActon_Dots: require("../../../assets/Card_SecondaryActon_Dots.svg"),
  CloseCourseManagementTray_X: require("../../../assets/CloseCourseManagementTray_X.svg"),
  CourseAccountActivated: require("../../../assets/CourseAccountActivated.svg"),
  CourseCertificates: require("../../../assets/CourseCertificates.svg"),
  CourseExpiringSoon: require("../../../assets/CourseExpiringSoon.svg"),
  CourseFailed: require("../../../assets/CourseFailed.svg"),
  CourseNewCourse: require("../../../assets/CourseNewCourse.svg"),
  CourseStatus_Completed: require("../../../assets/CourseStatus_Completed.svg"),
  CourseStatus_Incomplete: require("../../../assets/CourseStatus_Incomplete.svg"),
  CourseStatus_NotStarted: require("../../../assets/CourseStatus_NotStarted.svg"),
  CourseTimeTracked: require("../../../assets/CourseTimeTracked.svg"),
  Course_Calendar: require("../../../assets/Course_Calendar.svg"),
  DownloadCSV: require("../../../assets/DownloadCSV.svg"),
  EditDelegateProfilePic_Default: require("../../../assets/EditDelegateProfilePic_Default.png"),
  EditPicPencil: require("../../../assets/EditPicPencil.svg"),
  FilterAdjust: require("../../../assets/FilterAdjust.svg"),
  FormCheckbox_Checked: require("../../../assets/FormCheckbox_Checked.svg"),
  FormCheckbox_states: require("../../../assets/FormCheckbox_states.png"),
  FormCheckbox_Unchecked: require("../../../assets/FormCheckbox_Unchecked.svg"),
  LeftNav_Icon_Courses: require("../../../assets/LeftNav_Icon_Courses.svg"),
  LeftNav_Icon_Dashboard: require("../../../assets/LeftNav_Icon_Dashboard.svg"),
  LeftNav_Icon_Delegates: require("../../../assets/LeftNav_Icon_Delegates.svg"),
  Loading_Screen_Donut: require("../../../assets/Loading_Screen_Donut.svg"),
  Location_Pin: require("../../../assets/Location_Pin.svg"),
  PDF_Icon: require("../../../assets/PDF_Icon.svg"),
  RemoveSelectedCourse_X: require("../../../assets/RemoveSelectedCourse_X.svg"),
  SampleImage_ClassroomCoursesDetail_Feat: require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"),
  SearchGlass: require("../../../assets/SearchGlass.svg"),
};

type Props = {
  name: IconNames;
  style?: React.CSSProperties;
  size?: number;
};

function Icon({
  name,
  size = 20,
  style,
  ...props
}: Props & React.ImgHTMLAttributes<HTMLImageElement>) {
  return (
    <img
      style={{ height: size, width: size, ...style }}
      src={iconNameMap[name]}
      alt={`${name} icon`}
      {...props}
    />
  );
}

export default Icon;