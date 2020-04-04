import * as React from "react";
import UserSearch, { ResultItem } from "./UserSearch";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Search/UserSearch",
  decorators: [withKnobs],
};

const items: ResultItem[] = [
  {
    key: "Jim Smith",
    value: "uuid",
  },
  {
    key: "Bruce Willis",
    value: "uuid",
  },
  {
    key: "Tony Stark",
    value: "uuid",
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
    />
  );
};
