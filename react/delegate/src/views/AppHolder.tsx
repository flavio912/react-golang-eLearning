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
  },
  appHolderRoot: {
    display: "grid",
    height: "100vh",
    gridTemplateColumns: "auto 1fr",
    gridTemplateRows: "82px auto",
  },
}));

export const AppHolder = ({ children }: Props) => {
  const classes = useStyles();
  const { match, router } = useRouter();
  const tabs: Tab[] = [
    { id: 0, icon: "LeftNav_Icon_Dashboard", title: "Dashboard" },
    { id: 1, icon: "LeftNav_Icon_Courses", title: "Online Courses", size: 23 },
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
    <div className={classes.appHolderRoot}>
      <SideMenu
        tabs={tabs}
        selected={selected()}
        logo={require("../assets/logo/ttc-logo.svg")}
      />
      <HeaderMenu user={{ name: "James Smith", url: "" }} />
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
