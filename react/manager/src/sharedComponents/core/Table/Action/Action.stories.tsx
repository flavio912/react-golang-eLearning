import * as React from "react";
import Action from "./Action";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Core/Table/Action",
  decorators: [withKnobs],
};

export const normal = () => {
  return <Action />;
};
