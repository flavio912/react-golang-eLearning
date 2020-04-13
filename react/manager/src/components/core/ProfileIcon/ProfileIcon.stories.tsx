import * as React from "react";
import ProfileIcon from "./ProfileIcon";
import { withKnobs, text, number } from "@storybook/addon-knobs";

export default {
  title: "Core/ProfileIcon",
  decorators: [withKnobs],
};

export const plain = () => {
  const name: string = text('Name', "")
  const size: number = number("Size", 30)
  return <ProfileIcon name={name || 'Fred Ecceleston'} size={size} />;
};
