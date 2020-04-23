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
      children: <div></div>,
    },
    {
      id: 1,
      icon: "LeftNav_Icon_Delegates",
      children: <div></div>,
    },
    {
      id: 2,
      icon: "LeftNav_Icon_Courses",
      children: <div></div>,
    },
  ];

  const selected = () => {
    switch (match.location.pathname) {
      case "/app":
        return tabs[0];
      case "/app/delegates":
        return tabs[1];
      default:
        return tabs[0];
    }
  };

  return (
    <div>
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
      <SideMenu
        tabs={tabs}
        selected={selected()}
        onClick={(tab) => {
          console.log("tab", tab);
          switch (tab.id) {
            case 0:
              router.push("/app");
              break;
            case 1:
              router.push("/app/delegates");
              break;
            default:
              break;
          }
        }}
      />
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
