import * as React from "react";
import InfoField, { PaddingOptions } from "./InfoField";
import { withKnobs, select, text } from "@storybook/addon-knobs";

export default {
  title: "Core/InfoField",
  decorators: [withKnobs],
};

const paddingOptions: PaddingOptions[] = ["none", "small", "medium", "large"];

export const plain = () => {
  const padding: PaddingOptions = select("Padding", paddingOptions, "medium");
  const fieldNameText: string = text("Field Name", "");
  const valueText: string = text("Value", "");
  return <InfoField fieldName={fieldNameText || "Field Name"} value={valueText || "Value"} padding={padding} />;
};
