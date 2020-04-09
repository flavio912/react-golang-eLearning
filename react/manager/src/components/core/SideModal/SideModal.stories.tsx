import * as React from "react";
import SideModal from "./SideModal";
import { withKnobs, boolean } from "@storybook/addon-knobs";

export default {
  title: "Core/Side Modal",
  decorators: [withKnobs],
};

export const normal = () => {
  const isOpen = boolean("Open", true);
  return (
    <SideModal
      isOpen={isOpen}
      closeModal={() => {
        /* close the modal here */
      }}
    />
  );
};
