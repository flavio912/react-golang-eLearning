import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";
import { createUseStyles, useTheme } from "react-jss";
import { useRouter } from "found";

type Props = {
  children?: React.ReactChildren;
};

const useStyles = createUseStyles(() => ({
  appHolder: {
    display: "flex",
    padding: "42px 60px",
    marginLeft: 93,
    marginTop: 87,
    justifyContent: "center",
  },
}));

export const AppHolder = ({ children }: Props) => {
  const classes = useStyles();
  const { match, router } = useRouter();
  const tabs: Tab[] = [
    {
      id: 0,
      icon: "LeftNav_Icon_Dashboard",
      size: 20,
      children: <div></div>,
    },
    {
      id: 1,
      icon: "LeftNav_Icon_Delegates",
      size: 28,
      children: <div></div>,
    },
    {
      id: 2,
      icon: "LeftNav_Icon_Courses",
      size: 22,
      children: <div></div>,
    },
  ];

  const selected = () => {
    switch (match.location.pathname) {
      case "/app":
        return tabs[0];
      case "/app/delegates":
        return tabs[1];
      case "/app/courses":
        return tabs[2];
      default:
        return tabs[0];
    }
  };

  return (
    <div>
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
