import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon/Icon';
import Button from 'sharedComponents/core/Input/Button';
import CheckoutPopup, { BasketItem } from './CheckoutPopup';
import TabOption, { Tab } from './TabOption';
import MobileMenu from './MobileMenu';

const useStyles = createUseStyles((theme: Theme) => ({
  headerRoot: {
    display: 'flex',
    justifyContent: 'center',
    boxShadow: '0px 7px 20px #00000012',
    background: 'white',
    position: 'fixed',
    width: '100%',
    zIndex: 100,
    '@media (max-width: 650px)': {
      display: 'none'
    }
  },
  centerer: {
    width: theme.centerColumnWidth,
    margin: '0 20px'
  },
  menu: {
    display: 'flex',
    flexDirection: 'row',
    padding: '20px 0px'
  },
  tabs: {
    display: 'grid',
    gridTemplateRows: '1fr',
    gridTemplateColumns: '1fr 1fr 1fr 1fr',
    '@media (max-width: 900px)': {
      gridTemplateColumns: '1fr 1fr',
      gridTemplateRows: '1fr 1fr',
      gridRowGap: '4px'
    }
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
    right: '155px'
  },
  login: {
    fontFamily: 'Muli',
    fontSize: theme.fontSizes.large,
    fontWeight: '300',
    marginRight: '30px',
    cursor: 'pointer'
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
    cursor: 'pointer',
    marginRight: '70px',
    width: 128
  },
  backgroundHider: {
    position: 'absolute',
    width: '100%',
    height: '100%',
    top: 0,
    left: 0
  },
  mobileMenu: {
    '@media (min-width: 650px)': {
      display: 'none'
    }
  }
}));

type Props = {
  tabs: Array<Tab>;
  selected?: Tab;
  setSelected: (tab?: Tab) => void;
  onClick: (link: string) => void;
  onRegisterClick?: () => void;
  onLogoClick?: () => void;
  basketItems?: BasketItem[];
  onCheckout: () => void;
  className?: string;
};

function HeaderMenu({
  tabs,
  selected,
  setSelected,
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

  const onHide = () => {
    setShowPopup(false);
    setSelected(undefined);
  };

  return (
    <>
      {!showPopup && (
        <div className={classes.backgroundHider} onClick={() => onHide()} />
      )}

      <MobileMenu
        tabs={tabs}
        selected={selected}
        setSelected={setSelected}
        onClick={onClick}
        onRegisterClick={onRegisterClick}
        onLogoClick={onLogoClick}
        basketItems={basketItems}
        onCheckout={onCheckout}
        showPopup={showPopup}
        setShowPopup={setShowPopup}
        className={classes.mobileMenu}
      />
      <div className={classNames(classes.headerRoot, className)}>
        <div className={classes.centerer}>
          <div className={classNames(classes.row, classes.menu)}>
            <div className={classes.row}>
              <img
                src={require('../../../assets/logo/ttc-logo.svg')}
                className={classes.logo}
                onClick={onLogoClick}
              />
              <div className={classes.tabs}>
                {tabs &&
                  tabs.map((tab) => (
                    <TabOption
                      tab={tab}
                      selected={selected}
                      setSelected={setSelected}
                      onClick={onClick}
                    />
                  ))}
              </div>
            </div>
            <div className={classes.row}>
              {basketItems && basketItems.length > 0 && (
                <div onClick={() => setShowPopup(!showPopup)}>
                  <div className={classes.notification}>
                    {basketItems.length}
                  </div>
                  <Icon
                    name="Basket"
                    className={classes.basket}
                    style={{ cursor: 'pointer' }}
                    size={50}
                  />
                  <CheckoutPopup
                    showPopup={showPopup}
                    className={classes.checkoutPopup}
                    basketItems={basketItems}
                    onCheckout={onCheckout}
                  />
                </div>
              )}
              <div className={classes.login}>Login</div>
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
    </>
  );
}

export default HeaderMenu;
