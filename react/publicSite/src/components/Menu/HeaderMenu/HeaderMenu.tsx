import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon/Icon';
import Button from 'sharedComponents/core/Input/Button';
import CircleBorder from 'sharedComponents/core/CircleBorder';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    borderBottom: [1, 'solid', theme.colors.borderGrey],
    gridArea: '1 / 2',
    zIndex: 10,
  },
  menu: {
    display: 'flex',
    flexDirection: 'row',
    padding: '20px 95px',
  },
  tab: {
    fontFamily: 'Muli',
    fontSize: theme.fontSizes.large,
    fontWeight: "500",
    marginRight: "30px",
    cursor: "pointer",
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start",
  },
  basket: {
    paddingRight: '25px',
    marginRight: '25px',
    borderRight: [1, 'solid', theme.colors.borderGrey]
  },
  notification: {
    position: 'absolute',
    marginLeft: '25px',
    backgroundColor: theme.colors.navyBlue,
    color: theme.colors.primaryWhite,
    fontSize: theme.fontSizes.default,
    height: '20px',
    width: '20px',
    borderRadius: '20px',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  title: {
    fontSize: theme.fontSizes.large,
    fontWeight: "500",
  },
  register: {
    height: '40px',
    width: '127.75px'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: "space-between",
  },
  body: {
    backgroundColor: theme.colors.backgroundGrey,
    flexGrow: 1,
  },
}));

export interface Tab {
  id: number;
  title: string;
  options?: string[];
}

type Props = {
  tabs: Array<Tab>;
  selected: Tab;
  onClick?: (tab: Tab) => void;
  basketItems?: number;
  className?: string;
};

function HeaderMenu({ tabs, selected, onClick, basketItems, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classNames(classes.row, classes.menu)}>
        <div className={classes.row}>
          <Icon name="TTC_Logo_Icon" size={44} style={{ marginRight: '70px' }} />
          {tabs &&
              tabs.map((tab) => (
                <div
                  key={tab.id}
                  className={classNames(classes.tab)}
                  onClick={() => {
                    if (onClick) onClick(tab);
                  }}
                >
                  <div className={classes.title}>{tab.title}</div>
                  {tab.options && <Icon name="Down_Arrow" size={10} style={{ cursor: "pointer", marginLeft: '5px' }} />}
                </div>
              ))}
          </div>
          <div className={classes.row}>
            {basketItems && basketItems > 0 && 
              <div>
                <div className={classes.notification}>{basketItems}</div>
                <Icon name="Basket" className={classes.basket} size={38} />
              </div>
            }
            <div className={classes.tab}>Login</div>
            <Button archetype="gradient" className={classes.register}>Register</Button>
          </div>
        </div>
      </div>
  );
}

export default HeaderMenu;
