import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon, { IconNames } from "../Icon/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
  },
  menu: {
    width: '124px',
    height: '906px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    alignItems: 'center',
    backgroundColor: theme.colors.primaryWhite,
    borderRight: `1px solid ${theme.colors.borderGrey}`
  },
  tab: {
    width: '100%',
    cursor: 'pointer',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '30px 0',
    opacity: 0.3,
    transition: 'background-color 0.3s linear, opacity 0.3s linear'
  },
  selected: {
    backgroundColor: theme.colors.hoverGreen,
    justifyContent: 'space-between',
    opacity: 1,
    transition: 'background-color 0.3s linear, opacity 0.3s linear'
  },
  fold: {
    height: '40px',
    width: '5px',
    borderRadius: '0 9px 9px 0',
    backgroundColor: theme.colors.secondaryGreen,
    opacity: 1,
    transition: 'opacity 1s linear'
  },
  noFold: {
    height: '40px',
    width: '5px',
    opacity: 0,
    transition: 'visibility 0s 1s, opacity 1s linear'
  },
  body: {
    padding: '30px 30px',
    backgroundColor: theme.colors.primaryWhite,
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
    onClick: Function;
    className?: string;
};

function SideMenu({ tabs, selected, onClick, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    return (
        <div className={classNames(classes.root, className)}>
            <div className={classNames(classes.menu)}>
                {tabs.map(tab => (
                    <div
                        className={classNames(
                            classes.tab,
                            selected.id === tab.id && classes.selected
                        )}
                        onClick={() => onClick(tab)}
                    >
                        <div className={classNames(
                            selected.id === tab.id
                            ? classes.fold
                            : classes.noFold
                        )}/>
                        <Icon
                            name={tab.icon}
                            size={20}
                            style={{ cursor: 'pointer' }}
                        />
                        <div />
                    </div> 
                ))}
            </div>
            <div className={classes.body}>{selected.children}</div>
        </div>
    );
}

export default SideMenu;