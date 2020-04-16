import * as React from "react";
import SearchableDropdown from "./SearchableDropdown";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Core/SearchableDropdown",
  decorators: [withKnobs],
};

export const normal = () => {
  return <SearchableDropdown placeholder={text("Placeholder", "Placeholder text goes here")} />;
};