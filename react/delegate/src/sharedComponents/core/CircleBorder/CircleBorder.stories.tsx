import * as React from "react";
import CircleBorder, { User } from "./CircleBorder";
import { withKnobs, text, number } from "@storybook/addon-knobs";

export default {
  title: "core/CircleBorder",
  decorators: [withKnobs],
};

export const normal = () => {
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text("Profile Image", require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"));
  const user: User = { name, url };
  const size: number = number("Size", 45);
  const fontSize: number = number("fontSize", 18);
  return <CircleBorder user={user} size={size} fontSize={fontSize} />;
};