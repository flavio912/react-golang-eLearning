import * as React from "react";
import ActiveCoursesEmpty from "./ActiveCoursesEmpty";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Delegate/ActiveCoursesEmpty",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <ActiveCoursesEmpty
      title={text("Title", '"Book John on their first Course"')}
    />
  );
};
