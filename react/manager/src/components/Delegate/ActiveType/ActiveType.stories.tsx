import * as React from "react";
import ActiveType from "./ActiveType";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "/Delegate/ActiveType",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <ActiveType icon={"CourseFailed"} text={text("Text", "Failed Course")} />
  );
};
