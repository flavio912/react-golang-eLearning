import * as React from "react";
import MultiUserSearch from "./MultiUserSearch";
import { withKnobs } from "@storybook/addon-knobs";
import { ResultItem } from "../UserSearch";

export default {
  title: "Search/MultiUserSearch",
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

const Wrapper = () => {
  const [users, setUsers] = React.useState<(ResultItem | undefined)[]>([undefined]);
  return (
    <MultiUserSearch
      searchFunction={searchFunc}
      users={users}
      setUsers={setUsers}
    />
  );
};

export const normal = () => {
  return <Wrapper />;
};
