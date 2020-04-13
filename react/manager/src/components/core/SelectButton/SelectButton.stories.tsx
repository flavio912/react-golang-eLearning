import * as React from "react";
import SelectButton from "./SelectButton";
import {
  withKnobs,
  select,
  text,
  array,
  number,
} from "@storybook/addon-knobs";

export default {
  title: "Core/SelectButton",
  decorators: [withKnobs],
};

const defaultOptions = ['Online Courses', 'Classroom Courses'];

export const plain = () => React.createElement(() => {
  const [selected, setSelected] = React.useState(defaultOptions[0]);
  const options: Array<string> = array("Options", defaultOptions);

  return (
    <SelectButton
      selected={selected}
      options={options}
      onClick={(option: string) => setSelected(option)}
    />
  );
});
