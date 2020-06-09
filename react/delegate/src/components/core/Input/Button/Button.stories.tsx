import * as React from "react";
import { withKnobs, text } from "@storybook/addon-knobs";
import Button from "./Button";

export default {
  title: "Core/Input/Button",
  decorators: [withKnobs],
};

export const plain = () => {
  return <Button title={text("Title", "Next question")} onClick={() => {}} />;
};
