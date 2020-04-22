import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";
import ActvityCard from "components/Overview/ActivityCard";
import { createUseStyles, useTheme } from "react-jss";
import classes from "*.module.css";
import classNames from "classnames";
import Button from "components/core/Button";
import UserSearch from "components/UserSearch";

type Props = {};

const useStyles = createUseStyles(() => ({
  root: {
    display: "grid",
    gridGap: 17,
    gridTemplateRows: "0.1fr minmax(0, 0.1fr) 0.4fr 0.4fr auto",
    gridTemplateColumns: "repeat(auto-fill, 0.0833fr)",
  },
  activity: {
    gridRowStart: 4,
    gridColumn: "5 / 13",
  },
  title: {
    gridColumn: "1 / 5",
  },
  titleText: {
    margin: 0,
  },
  subtitleText: {
    margin: 0,
  },
  quickBookButton: {
    gridColumn: "9/11",
  },
  centerer: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  addDelegatesButton: {
    gridColumn: "11/13",
  },
  userSearch: {
    gridColumn: "1/5",
  },
}));

export const OrgOverview = () => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <div className={classes.title}>
        <h1 className={classes.titleText}>Fedex</h1>
        <p className={classes.subtitleText}>Organisation Overview</p>
      </div>
      <div className={classNames(classes.centerer, classes.quickBookButton)}>
        <Button bold archetype="submit">
          Quick Booking
        </Button>
      </div>
      <div className={classNames(classes.centerer, classes.addDelegatesButton)}>
        <Button bold archetype="submit">
          Add Delegates
        </Button>
      </div>
      <div className={classes.userSearch}>
        <UserSearch
          companyName="TESTcompany"
          searchFunction={async (query: string) => {
            return [];
          }}
        />
      </div>
      <ActvityCard
        className={classes.activity}
        padding={"medium"}
        leftHeading={"Delegates activity"}
        rightHeading={"Recent Updates"}
        options={["This month", "All Time"]}
        updates={[]}
        data={[
          { name: "name", value: 10 },
          { name: "time", value: 15 },
        ]}
      />
    </div>
  );
};
