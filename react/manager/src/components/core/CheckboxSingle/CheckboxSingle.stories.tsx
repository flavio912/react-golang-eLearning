import * as React from "react";
import CheckboxSingle from "./CheckboxSingle";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Core/CheckboxSingle",
  decorators: [withKnobs],
};

const Wrapper = () => {
  const [box, setBox] = React.useState<{ label: string; checked: boolean }>({
    label: "Checkbox 1",
    checked: true,
  });
  return <CheckboxSingle setBox={setBox} box={box} />;
};

export const normal = () => {
  return <Wrapper />;
};
