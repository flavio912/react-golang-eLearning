import * as React from "react";
import LoginDialogue from "./LoginDialogue";
import { withKnobs } from "@storybook/addon-knobs";
import FinaliseDialogue from "./FinaliseDialogue";
import PasswordDialogue from "./PasswordDialogue";

export default {
  title: "Login/Login Dialogue",
  decorators: [withKnobs],
};

export const normal = () => {
  return <LoginDialogue onSubmit={() => {}} />;
};

export const finalise = () => {
  return <FinaliseDialogue TTC_ID="TTC_ID" onSubmit={() => {}} />;
};

export const password = () => {
  return <PasswordDialogue onBack={() => null} onSubmit={() => {}} />;
};
