import * as React from "react";
import LoginDialogue from "./LoginDialogue";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Login/Login Dialogue",
  decorators: [withKnobs],
};

export const normal = () => {
  return <LoginDialogue onSubmit={() => {}} />;
};
