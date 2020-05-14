import * as React from "react";
import TimeSpent from "./TimeSpent";
import { withKnobs, number } from "@storybook/addon-knobs";

export default {
  title: "/Delegate/ActiveType",
  decorators: [withKnobs],
};

export const normal = () => {
  return <TimeSpent timeSpent={{ h: number("h", 1), m: number("m", 2) }} />;
};
