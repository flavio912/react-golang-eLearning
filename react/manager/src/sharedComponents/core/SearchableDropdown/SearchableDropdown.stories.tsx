import * as React from "react";
import SearchableDropdown, {
  CourseCategory,
  Course,
} from "./SearchableDropdown";
import { withKnobs, boolean, number } from "@storybook/addon-knobs";

export default {
  title: "Core/SearchableDropdown",
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
  {
    title: "Dangerous Goods - Land",
    courses: [
      {
        id: 9,
        name: "Coffee Machine (COFF) – VC, HS, XRY, ETD",
        price: 360,
      },
      {
        id: 10,
        name: "The Heavy Door in the Office (HDO) Recurrent – VC, HS, XRY, ETD",
        price: 65,
      },
    ],
  },
  {
    title: "Dangerous Goods - Sea",
    courses: [
      {
        id: 11,
        name: "Squids - Giant and Normal Size",
        price: 200,
      },
      {
        id: 12,
        name: "Sharks and Other Toothy Creatures",
        price: 70,
      },
      {
        id: 13,
        name: "Seaweed and Seagulls",
        price: 20,
      },
    ],
  },
];

const searchFunc = async (query: string) => {
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

const DropdownExample = ({ multiselect, debounceTime }: any) => {
  const [selected, setSelected] = React.useState<Course[]>([]);
  return (
    <SearchableDropdown
      selected={selected}
      setSelected={setSelected}
      multiselect={multiselect}
      searchQuery={searchFunc}
      debounceTime={debounceTime}
    />
  );
};

export const normal = () => {
  return (
    <DropdownExample
      multiselect={boolean("Multiselect", false)}
      debounceTime={number("Debounce Time", 500)}
    />
  );
};
