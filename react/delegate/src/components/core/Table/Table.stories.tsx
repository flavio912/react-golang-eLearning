import * as React from "react";
import Table from "./Table";
import {
  withKnobs,
  array,
  select,
  boolean,
  number,
} from "@storybook/addon-knobs";
import Icon from "../../../sharedComponents/core/Icon";

export default {
  title: "Core/Table",
  decorators: [withKnobs],
};

export const normal = () => {
  const header = array("Header", ["Name", "Progress", "Status", "Options"]);
  const sort = boolean("Sort", false);
  const sortBy = number("Sort index", 0);
  const sortDir = select("Sort dir", ["ASC", "DESC"], "DESC");
  const filter = boolean("Filter", false);
  const filterBy = number("Filter index", 2);
  const filterFunc = boolean("Filter true/false", true);
  return (
    <Table
      header={header}
      sort={
        sort
          ? {
              column: sortBy,
              dir: sortDir,
            }
          : undefined
      }
      filter={
        filter
          ? {
              column: filterBy,
              filterFunc: filterFunc ? (v) => !!v : (v) => !v,
            }
          : undefined
      }
      rows={[
        {
          key: "Dave",
          cells: [
            {
              component: "Dave",
              sort: "Dave",
            },
            {
              component: "7/10",
              sort: 7 / 10,
            },
            {
              component: () => (
                <p>
                  <Icon size={12} name="CourseStatus_Incomplete" /> Incomplete
                </p>
              ),
              sort: false,
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
              component: () => (
                <p>
                  <Icon size={12} name="CourseStatus_Incomplete" /> Incomplete
                </p>
              ),
              sort: false,
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
              component: "15/15",
              sort: 15 / 15,
            },
            {
              component: () => (
                <p>
                  <Icon size={12} name="CourseStatus_Completed" /> Complete
                </p>
              ),
              sort: true,
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
