import * as React from "react";
import ProfileIcon from "./ProfileIcon";
import { withKnobs, text, number } from "@storybook/addon-knobs";

export default {
  title: "Core/ProfileIcon",
  decorators: [withKnobs],
};

export const plain = () => {
  const name: string = text('Name', "Fred Ecceleston")
  const size: number = number("Size", 30)
  const url: string = text("URL", require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"))
  return <ProfileIcon name={name} url={url} size={size} />;
};
