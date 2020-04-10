import * as React from "react";
import Table from "./Table";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "/Table",
  decorators: [withKnobs],
};

export const normal = () => {
  return <Table />;
};