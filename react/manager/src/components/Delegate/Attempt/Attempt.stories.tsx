import * as React from "react";
import Attempt from "./Attempt";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "/Delegate/Attempt",
  decorators: [withKnobs],
};

export const normal = () => {
  return <Attempt attempt={text("Text", "1")} />;
};
