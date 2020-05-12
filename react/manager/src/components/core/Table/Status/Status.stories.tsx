import * as React from "react";
import Status from "./Status";
import { withKnobs, text, boolean } from "@storybook/addon-knobs";

export default {
  title: "/Table/UserLabel",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <Status
      isComplete={boolean("Is Complete", false)}
      expires={text("Expires", "Expires 20/02/202")}
    />
  );
};
