import * as React from "react";
import { withKnobs, number } from "@storybook/addon-knobs";
import CourseCompletion from ".";

export default {
  title: "Core/Table/CourseCompletion",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <CourseCompletion
      complete={number("Courses completed", 4)}
      total={number("Total courses", 6)}
    />
  );
};
