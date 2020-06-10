import * as React from "react";
import UserLabel from "./UserLabel";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Core/Table/UserLabel",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <UserLabel
      name={text("User Name", "Bill Gates")}
      profileUrl={text("Profile URL", "")}
    />
  );
};
