import * as React from "react";
import Button, { Archetypes } from "./Button";
import { withKnobs, select, text, boolean } from "@storybook/addon-knobs";
import { IconNames } from "../Icon";

export default {
  title: "Core/Button",
  decorators: [withKnobs],
};

const archetypes: Archetypes[] = ["default", "grey", "submit"];

const iconNames: (IconNames|"None")[] = [
  "None",
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

export const flexible = () => {
  const archetypeTexts = {
    default: "Cancel",
    grey: "Add Delegate",
    submit: "Continue to Terms",
  };
  const archetype: Archetypes = select("Type", archetypes, "default");
  const btnText: string = text("Text", "");
  const bold: boolean = boolean("Bold", false);
  const small: boolean = boolean("Small", false);
  const iconLeft: (IconNames|"None") = select("Left Icon", iconNames, "None");
  const iconRight: (IconNames|"None") = select("Right Icon", iconNames, "None");

  return (
    <Button
      archetype={archetype}
      iconLeft={iconLeft !== "None" ? iconLeft : undefined}
      iconRight={iconRight !== "None" ? iconRight : undefined}
      bold={bold}
      small={small}
    >
      {btnText || archetypeTexts[archetype]}
    </Button>
  );
};
