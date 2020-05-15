import * as React from "react";
import { withKnobs, number } from "@storybook/addon-knobs";
import ProgressBar from ".";

export default {
  title: "/ProgressBar",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <ProgressBar
      percent={number("Percent", 30)}
      width={number("Width (px)", 200)}
    />
  );
};