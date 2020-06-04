import * as React from "react";
import PageHeader from "./PageHeader";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "core/PageHeader",
  decorators: [withKnobs],
};

export const normal = () => {
  const title: string = text("Title", "About Us");
  const description: string = text("Description", "Our mission is to create the highest quality safety & compliance training in the world");
  return <PageHeader title={title} description={description} />;
};