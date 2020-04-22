import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";
import { createUseStyles, useTheme } from "react-jss";

type Props = {
  children?: React.ReactChildren;
};

const useStyles = createUseStyles(() => ({
  root: {
    marginLeft: 93,
    padding: 28,
  },
  background: {
    background: "#f8f9fb",
  },
}));

export const AppHolder = ({ children }: Props) => {
  const classes = useStyles();

  const tabs: Tab[] = [
    {
      id: 0,
      icon: "TTC_Logo_Icon",
      children: <div></div>,
    },
  ];
  return (
    <div className={classes.background}>
      <HeaderMenu logo={""} user={{ name: "Test", url: "My URL" }} />
      <SideMenu tabs={tabs} selected={tabs[0]} onClick={() => {}} />
      <div className={classes.root}>{children}</div>
    </div>
  );
};
