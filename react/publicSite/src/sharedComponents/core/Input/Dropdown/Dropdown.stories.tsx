import * as React from "react";
import Dropdown, { DropdownOption } from "./Dropdown";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Core/Input/Dropdown",
  decorators: [withKnobs],
};

const defaultComponent = () => (
  <div>
    PlaceHolder Name
  </div>
)

const defaultOption: DropdownOption = {
  id: 1,
  title: "Test",
  component: defaultComponent()
}

const defaultOptions = [defaultOption, defaultOption, defaultOption]

const DropdownExample = () => {
  const [selected, setSelected] = React.useState<DropdownOption>();
  return (
    <Dropdown
      placeholder="Test"
      options={defaultOptions}
      selected={selected}
      setSelected={setSelected}
    />
  );
};

export const normal = () => {
  return (
    <DropdownExample />
  );
};
