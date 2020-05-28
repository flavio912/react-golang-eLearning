import * as React from "react";
import EasyInput from "./EasyInput";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Core/EasyInput",
  decorators: [withKnobs],
};

export const input = () => {
  return (
    <EasyInput
      label={text('Label', 'First Name')}
      placeholder={text('Placeholder', 'e.g. John')}
    />
  )
};