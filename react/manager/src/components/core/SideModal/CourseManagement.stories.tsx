import * as React from "react";
import SideModal from "./SideModal";
import { withKnobs, boolean } from "@storybook/addon-knobs";
import Button from "components/core/Button";
import Tabs, { TabContent, Footer, Body } from "./Tabs";
import SearchableDropdown, { CourseCategory } from "../SearchableDropdown";
import { ResultItem } from "components/UserSearch";
import MultiUserSearch from "components/UserSearch/MultiUserSearch";

export default {
  title: "Core/Side Modal",
  decorators: [withKnobs],
};

const categories: CourseCategory[] = [
  {
    title: "Dangerous Goods - Air",
    courses: [
      {
        id: 1,
        name: "Cargo Operative Screener (COS) – VC, HS, XRY, ETD",
        price: 200,
      },
      {
        id: 2,
        name: "Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD",
        price: 65,
      },
    ],
  },
  {
    title: "Known Consignor",
    courses: [
      {
        id: 3,
        name: "Known Consignor Responsible Person",
        price: 70,
      },
      {
        id: 4,
        name: "Known Consignor (Modules 1-7)",
        price: 55,
      },
      {
        id: 5,
        name: "Manual Handling Awareness",
        price: 25,
      },
      {
        id: 6,
        name: "Fire Safety Awareness",
        price: 25,
      },
      {
        id: 7,
        name: "Noise and Vibration Awareness",
        price: 25,
      },
      {
        id: 8,
        name: "Seat Belt Misuse Awareness",
        price: 300,
      },
    ],
  },
];

const courseSearchFunction = async (query: string) => {
  return await new Promise<CourseCategory[]>((resolve) =>
    setTimeout(
      () =>
        resolve(
          categories
            .map(({ title, courses }) => ({
              title,
              courses: courses.filter(({ name }) =>
                name.toLocaleLowerCase().includes(query.toLocaleLowerCase())
              ),
            }))
            .filter(({ courses }) => courses.length > 0)
        ),
      700
    )
  );
};

const items: ResultItem[] = [
  {
    key: "Jim Smith",
    value: "uuid-1",
  },
  {
    key: "Bruce Willis",
    value: "uuid-2",
  },
  {
    key: "Tony Stark",
    value: "uuid-3",
  },
];

const userSearchFunction = async (query: string) => {
  return await new Promise<ResultItem[]>((resolve) =>
    setTimeout(() => resolve(items), 700)
  );
};

const tabs: TabContent[] = [
  {
    key: "Courses",
    component: ({ state, setState, setTab, closeModal }) => (
      <>
        <Body>
          <h1 style={{ fontSize: 24, color: "#0C152E" }}>
            Start a Quick Booking
          </h1>
          <p style={{ fontSize: 15 }}>Select your Course</p>
          <SearchableDropdown
            selected={state.course ? [state.course] : []}
            searchQuery={courseSearchFunction}
            setSelected={(selected) =>
              setState((s: object) => ({ ...s, course: selected[0] }))
            }
          />
          <MultiUserSearch
            searchFunction={userSearchFunction}
            users={state.users}
            setUsers={(users) => setState((s: object) => ({ ...s, users }))}
            style={{ marginTop: 25 }}
          />
        </Body>
        <Footer>
          <div />
          <div style={{ display: "flex" }}>
            <Button archetype="default" onClick={() => closeModal()}>
              Cancel
            </Button>
            <Button
              archetype="submit"
              onClick={() => setTab("Terms of Business")}
              style={{ marginLeft: 20 }}
            >
              Next
            </Button>
          </div>
        </Footer>
      </>
    ),
  },
  {
    key: "Terms of Business",
    component: ({ state, setTab }) => (
      <div style={{ margin: "30px 40px" }}>
        <p>State is:</p>
        <pre>{JSON.stringify(state, null, 2)}</pre>
        <Button onClick={() => setTab("Payment")} style={{ marginTop: 20 }}>
          Next
        </Button>
      </div>
    ),
  },
  {
    key: "Payment",
    component: ({ setTab }) => (
      <div style={{ margin: "30px 40px" }}>
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

export const courseManagement = () => {
  return (
    <SideModal
      isOpen={boolean("Open", true)}
      title="Course Management"
      closeModal={() => alert("Close Function called")}
    >
      <Tabs
        content={tabs}
        closeModal={() => alert("Close Function called")}
        initialState={{ users: [undefined] }}
      />
    </SideModal>
  );
};
