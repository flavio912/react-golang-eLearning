import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon/Icon';
import Button from 'sharedComponents/core/Input/Button';
import CheckoutPopup, { BasketItem } from './CheckoutPopup';

const useStyles = createUseStyles((theme: Theme) => ({
  headerRoot: {
    display: 'flex',
    justifyContent: 'center',
    boxShadow: '0px 7px 20px #00000012',
    background: 'white',
    position: 'fixed',
    width: '100%',
    zIndex: 100
  },
  centerer: {
    width: theme.centerColumnWidth
  },
  menu: {
    display: 'flex',
    flexDirection: 'row',
    padding: '20px 0px'
  },
  tab: {
    fontFamily: 'Muli',
    fontSize: theme.fontSizes.large,
    fontWeight: '300',
    marginRight: '30px',
    cursor: 'pointer',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'flex-start'
  },
  basket: {
    paddingRight: '25px',
    marginRight: '25px',
    borderRight: [1, 'solid', theme.colors.borderGrey]
  },
  notification: {
    position: 'absolute',
    marginLeft: '35px',
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
  checkoutPopup: {
    right: 291
  },
  title: {
    fontSize: theme.fontSizes.large,
    fontWeight: 300
  },
  register: {
    height: '40px',
    width: '127.75px',
    boxShadow: '0px 3px 10px #0000001f'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between'
  },
  body: {
    backgroundColor: theme.colors.backgroundGrey,
    flexGrow: 1
  },
  logo: {
    marginRight: '70px',
    width: 128
  }
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
  onRegisterClick?: () => void;
  onLogoClick?: () => void;
  basketItems?: BasketItem[];
  onCheckout: () => void;
  children?: React.ReactNode;
  className?: string;
};

function HeaderMenu({
  tabs,
  selected,
  onClick,
  onRegisterClick,
  onLogoClick,
  basketItems,
  onCheckout,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [showPopup, setShowPopup] = React.useState(false);

  return (
    <div className={classNames(classes.headerRoot, className)}>
      <div className={classes.centerer}>
        <div className={classNames(classes.row, classes.menu)}>
          <div className={classes.row}>
            <img
              src={require('../../../assets/logo/ttc-logo.svg')}
              className={classes.logo}
              onClick={onLogoClick}
            />
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
                  {tab.options && (
                    <Icon
                      name="Down_Arrow"
                      size={10}
                      style={{ cursor: 'pointer', marginLeft: '5px' }}
                    />
                  )}
                </div>
              ))}
          </div>
          <div className={classes.row}>
            {basketItems && basketItems.length > 0 && (
              <div onClick={() => setShowPopup(!showPopup)}>
                <div className={classes.notification}>{basketItems.length}</div>
                <Icon
                  name="Basket"
                  className={classes.basket}
                  style={{ cursor: 'pointer' }}
                  size={50}
                />
                <CheckoutPopup
                  showPopup={showPopup}
                  onHide={() => setShowPopup(false)}
                  className={classes.checkoutPopup}
                  basketItems={basketItems}
                  onCheckout={onCheckout}
                />
              </div>
            )}
            <div className={classes.tab}>Login</div>
            <Button
              archetype="gradient"
              className={classes.register}
              onClick={onRegisterClick}
            >
              Register
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default HeaderMenu;
