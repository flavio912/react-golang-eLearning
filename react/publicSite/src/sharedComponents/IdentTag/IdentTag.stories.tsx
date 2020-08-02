import * as React from "react";
import { withKnobs, text } from "@storybook/addon-knobs";
import IdentTag from "./IdentTag";

export default {
  title: "Misc/ID Tag",
  decorators: [withKnobs],
};

export const plain = () => {
  return <IdentTag ident={text("TTC ID", "Fedex_BruceWillis")} />;
};
