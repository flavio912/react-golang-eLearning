import * as React from "react";
import CourseTable from "./CourseTable";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "/Delegate/CourseTable",
  decorators: [withKnobs],
};

export const normal = () => {
  return <CourseTable />;
};
