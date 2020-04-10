import * as React from "react";
import Table from "./Table";
import { withKnobs, array, select, text } from "@storybook/addon-knobs";
import Icon from "../Icon";

export default {
  title: "/Table",
  decorators: [withKnobs],
};

export const normal = () => {
  const header = array("Header", ["Name", "Progress", "Options"]);
  return (
    <Table
      header={header}
      sort={{
        by: text("Sort by", "Name"),
        dir: select("Sort dir", ["ASC", "DESC"], "DESC"),
      }}
      rows={[
        {
          key: "Dave",
          cells: [
            {
              component: "Dave",
              sort: "Dave",
            },
            {
              component: "5/10",
              sort: 5 / 10,
            },
            {
              component: () => <Icon name="Card_SecondaryActon_Dots" />,
            },
          ],
        },
        {
          key: "Steve",
          cells: [
            {
              component: "Steve",
              sort: "Steve",
            },
            {
              component: "6/10",
              sort: 6 / 10,
            },
            {
              component: () => <Icon name="Card_SecondaryActon_Dots" />,
            },
          ],
        },
        {
          key: "Mary",
          cells: [
            {
              component: "Mary",
              sort: "Mary",
            },
            {
              component: "14/15",
              sort: 14 / 15,
            },
            {
              component: () => <Icon name="Card_SecondaryActon_Dots" />,
            },
          ],
        },
      ]}
    />
  );
};
