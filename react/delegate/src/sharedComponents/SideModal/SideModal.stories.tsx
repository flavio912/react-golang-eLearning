import * as React from "react";
import SideModal from "./SideModal";
import { withKnobs, boolean, text } from "@storybook/addon-knobs";
import Button from "sharedComponents/core/Input/Button";
import FancyInput from "components/LoginDialogue/FancyInput";
import Tabs, { TabContent } from "./Tabs";

export default {
  title: "Core/Modals/Side Modal",
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
        <Button onClick={() => setTab("Second")} style={{ marginTop: 20 }}>
          Next
        </Button>
      </div>
    ),
  },
  {
    key: "Second",
    component: ({ state, setTab }) => (
      <div style={{ padding: 15 }}>
        <p>State is:</p>
        <pre>{JSON.stringify(state, null, 2)}</pre>
        <Button onClick={() => setTab("Third")} style={{ marginTop: 20 }}>
          Next
        </Button>
      </div>
    ),
  },
  {
    key: "Third",
    component: ({ setTab }) => (
      <div style={{ padding: 15 }}>
        <p>
          Storybook doesn't allow you to change the value of the knobs, but this
          component takes closeModal as a prop so you can close the modal
        </p>
        <Button onClick={() => setTab("First")} style={{ marginTop: 20 }}>
          Back to Start
        </Button>
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
      <Tabs content={tabs} closeModal={() => {}} initialState={""} />
    </SideModal>
  );
};
