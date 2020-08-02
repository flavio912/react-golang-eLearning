import * as React from "react";
import Spinner from "./Spinner";
import { withKnobs, number } from "@storybook/addon-knobs";

export default {
  title: "Core/Spinner",
  decorators: [withKnobs],
};

export const normal = () => {
  return <Spinner size={number("Size", NaN)} />;
};
