import * as React from "react";
import Checkbox from "./Checkbox";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Core/Checkbox",
  decorators: [withKnobs],
};

const Wrapper = () => {
  const [boxes, setBoxes] = React.useState<
    { label: string; checked: boolean }[]
  >([
    { label: "Checkbox 1", checked: true },
    { label: "Checkbox 2", checked: false },
  ]);
  return <Checkbox setBoxes={setBoxes}>{boxes}</Checkbox>;
};

export const normal = () => {
  return <Wrapper />;
};
