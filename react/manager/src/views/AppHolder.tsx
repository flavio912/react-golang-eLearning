import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";
import { createUseStyles, useTheme } from "react-jss";

type Props = {
  children?: React.ReactChildren;
};

const useStyles = createUseStyles(() => ({
  appHolder: {
    marginLeft: 93,
    padding: 28,
    justifyContent: "center",
    display: "flex",
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
      <HeaderMenu
        logo={
          "https://www.stickpng.com/assets/images/58428e7da6515b1e0ad75ab5.png"
        }
        user={{
          name: "Test",
          url:
            "https://www.stickpng.com/assets/images/58428e7da6515b1e0ad75ab5.png",
        }}
      />
      <SideMenu tabs={tabs} selected={tabs[0]} onClick={() => {}} />
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
