import * as React from "react";
import { createUseStyles } from "react-jss";

import { Theme } from "helpers/theme";
import Dropdown, { DropdownOption } from "sharedComponents/core/Input/Dropdown";
import PageNumbers from "sharedComponents/Pagination/PageNumbers";
import { update } from "relay-runtime/lib/handlers/connection/ConnectionHandler";

const useStyles = createUseStyles((theme: Theme) => ({
  paginatorRoot: {
    display: "flex",
  },
  pageDropdown: {
    width: 117,
    backgroundColor: theme.colors.primaryWhite,
  },
}));

const defaultOptions: DropdownOption[] = [
  {
    id: 1,
    title: "Show 10",
    component: <div>Show 10</div>,
  },
  {
    id: 2,
    title: "Show 20",
    component: <div>Show 20</div>,
  },
  {
    id: 3,
    title: "Show 50",
    component: <div>Show 50</div>,
  },
];

type Props = {
  currentPage: number;
  showRange?: number;
  numPages: number;
  updatePage: (page: number) => void;
  itemsPerPage: 10 | 20 | 50; // TODO: Implement this
};

function Pageinator({
  currentPage,
  numPages,
  showRange = 4,
  updatePage,
}: Props) {
  const classes = useStyles();

  const [itemsPerPage, setItemsPerPage] = React.useState<DropdownOption>();
  return (
    <div className={classes.paginatorRoot}>
      <div className={classes.pageDropdown}>
        <Dropdown
          placeholder="Show 10"
          options={defaultOptions}
          selected={itemsPerPage}
          setSelected={setItemsPerPage}
        />
      </div>
      <PageNumbers
        numberOfPages={numPages}
        range={showRange}
        currentPage={currentPage}
        setCurrentPage={updatePage}
      />
    </div>
  );
}

export default Pageinator;
