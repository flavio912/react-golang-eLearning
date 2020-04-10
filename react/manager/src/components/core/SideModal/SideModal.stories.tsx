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

const tabs: TabContent[] = [
  {
    key: "First",
    component: ({ setState, setTab }) => (
      <div style={{ padding: 15 }}>
        <FancyInput label="State" onChange={(v) => setState(v)} />
        <FancyButton text="Next tab" onClick={() => setTab("Second")} />
      </div>
    ),
  },
  {
    key: "Second",
    component: ({ state }) => (
      <div style={{ padding: 15 }}>
        <p>State is: {state}</p>
      </div>
    ),
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
      <Tabs content={tabs} />
    </SideModal>
  );
};
