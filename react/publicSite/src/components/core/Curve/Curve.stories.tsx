import * as React from "react";
import Curve from "./Curve";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "core/Curve",
  decorators: [withKnobs],
};

export const normal = () => {
  return <Curve />;
};