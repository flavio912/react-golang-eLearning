import * as React from "react";
import SideModal from "./SideModal";
import { withKnobs, boolean, text } from "@storybook/addon-knobs";

export default {
  title: "Core/Side Modal",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <SideModal
      isOpen={boolean("Open", true)}
      title={text("Title", "Course Management")}
      closeModal={() => {
        /* close the modal here */
      }}
    />
  );
};
