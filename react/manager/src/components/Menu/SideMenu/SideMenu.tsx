import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon, { IconNames } from "../../core/Icon/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    zIndex: 10,
    top: 89,
    position: "fixed",
    height: "100%",
    flexDirection: "row",
    boxShadow: theme.shadows.primary,
  },
  menu: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "flex-start",
    alignItems: "center",
    backgroundColor: theme.colors.primaryWhite,
  },
  tab: {
    width: "93px",
    cursor: "pointer",
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    padding: "25px 0",
    opacity: 0.3,
    transition: "background-color 0.3s linear, opacity 0.3s linear",
  },
  selected: {
    backgroundColor: theme.colors.hoverGreen,
    justifyContent: "space-between",
    opacity: 1,
    transition: "background-color 0.3s linear, opacity 0.3s linear",
  },
  fold: {
    height: "40px",
    width: "5px",
    borderRadius: "0 9px 9px 0",
    backgroundColor: theme.colors.secondaryGreen,
    opacity: 1,
    transition: "opacity 1s linear",
  },
  noFold: {
    height: "40px",
    width: "5px",
    opacity: 0,
    transition: "visibility 0s 1s, opacity 1s linear",
  },
  body: {
    padding: "30px 30px",
    backgroundColor: theme.colors.backgroundGrey,
    boxShadow: theme.shadows.body,
    flexGrow: 1,
  },
}));

export interface Tab {
  id: number;
  icon: IconNames;
  children: React.ReactNode;
}

type Props = {
  tabs: Array<Tab>;
  selected: Tab;
  onClick: (tab: Tab) => void;
  className?: string;
};

function SideMenu({ tabs, selected, onClick, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classNames(classes.menu)}>
        {tabs &&
          tabs.map((tab) => (
            <div
              key={tab.id}
              className={classNames(
                classes.tab,
                selected && selected.id === tab.id && classes.selected
              )}
              onClick={() => onClick(tab)}
            >
              <div
                className={classNames(
                  selected && selected.id === tab.id
                    ? classes.fold
                    : classes.noFold
                )}
              />
              <Icon name={tab.icon} size={20} style={{ cursor: "pointer" }} />
              <div />
            </div>
          ))}
      </div>
    </div>
  );
}

export default SideMenu;
