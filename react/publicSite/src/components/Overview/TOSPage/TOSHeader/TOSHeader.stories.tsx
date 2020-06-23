import * as React from "react";
import TOSHeader from "./TOSHeader";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Overview/TOSPage/TOSHeader",
  decorators: [withKnobs],
};

export const normal = () => {
  const title: string = text("Title", "Policies and Procedures");
  const subtitle: string = text("`Subtitle`", "We’re committed to keeping your data secure, your private information private, and being transparent about our practices as a business.");
  return <TOSHeader
    title={title}
    subtitle={subtitle}
  />;
};