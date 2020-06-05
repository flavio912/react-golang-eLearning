import * as React from "react";
import Text from "./Text";
import { withKnobs, text, boolean } from "@storybook/addon-knobs";

export default {
  title: "Core/Table/Text",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <Text
      text={text("Text", "Cool text")}
      color={text("Text color", "#1081AA")}
      formatDate={boolean("Format as Date", false)}
    />
  );
};
