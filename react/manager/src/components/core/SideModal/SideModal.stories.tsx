import * as React from "react";
import SideModal from "./SideModal";
import { withKnobs, boolean, text } from "@storybook/addon-knobs";
import FancyButton from "components/LoginDialogue/FancyButton";
import FancyInput from "components/LoginDialogue/FancyInput";
import Tabs, { TabContent } from "./Tabs";

export default {
  title: "Core/Side Modal",
  decorators: [withKnobs],
};

export const empty = () => {
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

const tabs = [
  {
    key: "First",
    component: ({ setState, setTab }: TabContent) => {
      return (
        <>
          <FancyInput label="State" onChange={(v) => setState(v)} />
          <FancyButton onClick={() => setTab("Second")} />
        </>
      );
    },
  },
  {
    key: "Second",
    component: ({ state }) => <p>State is: {state}</p>,
  },
];

export const tabbed = () => {
  return (
    <SideModal
      isOpen={boolean("Open", true)}
      title={text("Title", "Course Management")}
      closeModal={() => {
        /* close the modal here */
      }}
    >
      <Tabs>{}</Tabs>
    </SideModal>
  );
};
