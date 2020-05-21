import * as React from "react";
import CircleBorder, { User } from "./CircleBorder";
import { withKnobs, text, number } from "@storybook/addon-knobs";

export default {
  title: "core/CircleBorder",
  decorators: [withKnobs],
};

export const normal = () => {
  const colour: string = text("Colour", "#FFFFFF");
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text("Profile Image", require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"));
  const user: User = { name, url };
  const size: number = number("Size", 44);
  const fontSize: number = number("fontSize", 18);
  const userMode: string = text("UserMode", "Manager");
  const borderSize: number = number("BorderSize", 6);

  return <CircleBorder user={user} size={size} fontSize={fontSize} colour={colour} userMode={userMode} borderSize={borderSize} />;
};