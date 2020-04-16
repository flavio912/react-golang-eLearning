import * as React from "react";
import SearchableDropdown from "./SearchableDropdown";
import { withKnobs, text, object } from "@storybook/addon-knobs";

export default {
  title: "Core/SearchableDropdown",
  decorators: [withKnobs],
};

const options = [
  {
    title: "Dangerous Goods - Air",
    content: [
      {
        name: "Cargo Operative Screener (COS) – VC, HS, XRY, ETD",
        price: 200,
      },
      {
        name: "Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD",
        price: 65,
      },
    ],
  },
  {
    title: "Known Consignor",
    content: [
      {
        name: "Known Consignor Responsible Person",
        price: 70,
      },
      {
        name: "Known Consignor (Modules 1-7)",
        price: 55,
      },
      {
        name: "Manual Handling Awareness",
        price: 25,
      },
      {
        name: "Fire Safety Awareness",
        price: 25,
      },
      {
        name: "Noise and Vibration Awareness",
        price: 25,
      },
      {
        name: "Seat Belt Misuse Awareness",
        price: 300,
      },
    ],
  },
];

const DropdownExample = ({ placeholder, options }: any) => {
  const [selected, setSelected] = React.useState<string | null>(null);
  return (
    <SearchableDropdown
      placeholder={placeholder}
      selected={selected}
      setSelected={setSelected}
      options={options}
    />
  );
};

export const normal = () => {
  return (
    <DropdownExample
      placeholder={text("Placeholder", "Placeholder text goes here")}
      options={object("Options", options)}
    />
  );
};
