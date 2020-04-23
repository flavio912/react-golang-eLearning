import * as React from "react";
import { withKnobs, text } from "@storybook/addon-knobs";
import PageTitle from "./PageTitle";

export default {
  title: "Pages/Title",
  decorators: [withKnobs],
};

export const plain = () => {
  return (
    <PageTitle
      title={text("Title", "Fedex")}
      subTitle={text("Subtitle", "Organisation Overview")}
    />
  );
};
