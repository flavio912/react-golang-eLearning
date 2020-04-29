import * as React from "react";
import UserSearch, { ResultItem } from "./UserSearch";
import { withKnobs, text, number } from "@storybook/addon-knobs";

export default {
  title: "Search/UserSearch",
  decorators: [withKnobs],
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

const searchFunc = async (query: string) => {
  return await new Promise<ResultItem[]>((resolve) =>
    setTimeout(() => resolve(items), 700)
  );
};

export const plain = () => {
  return (
    <UserSearch
      companyName={text("Company name", "Fedex")}
      searchFunction={searchFunc}
      debounceTime={number("Debounce time (ms)", 500)}
    />
  );
};
