import * as React from "react";
import { withKnobs, text, select } from "@storybook/addon-knobs";
import Heading from "./Heading";

export default {
  title: "Misc/ID Tag",
  decorators: [withKnobs],
};

export const plain = () => {
  return (
    <Heading
      text={text("Text", "Training Zone")}
      size={select("Size", ["large", "medium"], "large")}
    />
  );
};
