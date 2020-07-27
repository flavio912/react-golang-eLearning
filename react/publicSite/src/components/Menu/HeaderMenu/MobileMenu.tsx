import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon/Icon';
import Button from 'sharedComponents/core/Input/Button';
import CheckoutPopup, { BasketItem } from './CheckoutPopup';
import { Tab } from './TabOption';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBars } from '@fortawesome/free-solid-svg-icons'

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        position: 'fixed',
        width: '100%',
        boxShadow: '0px 7px 20px #00000012',
        background: 'white',
        zIndex: 100,
        display: 'flex',
        justifyContent: 'space-between',
        padding: '20px'
    },
    logo: {
        cursor: 'pointer',
        marginRight: '10px',
        width: 128
    },
    openMenu: {
        position: 'absolute',
        top: '75px',
        left: '25px',
        right: '25px',
        zIndex: 10,
        border: ['0.5px', 'solid', theme.colors.borderGrey],
        borderRadius: '8px',
        backgroundColor: theme.colors.primaryWhite,
        padding: '5px 0',
        boxShadow: '0px 3px 10px #0000001f'
    },
    buttonWrapper: {
        display: 'flex',
        margin: '20px'
    },
    register: {
        height: '40px',
        width: '100%',
        boxShadow: '0px 3px 10px #0000001f',
        '@media (max-width: 225px)': {
            height: 'auto'
          }
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between'
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
    option: {
        margin: '20px'
    },
    title: {
        fontSize: theme.fontSizes.default,
        fontWeight: '400',
        color: theme.colors.textGrey,
        marginBottom: '10px'
    },
    text: {
        cursor: 'pointer',
        fontSize: theme.fontSizes.large,
        fontWeight: '500',
        marginBottom: '8px'
    },
    footer: {
        margin: '20px',
        fontSize: theme.fontSizes.small,
        fontWeight: '400',
        color: theme.colors.textGrey,
    },
    link: {
        textDecorationLine: 'underline',
        color: theme.colors.textBlue,
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
    showPopup: boolean;
    setShowPopup: (bool: boolean) => void;
    className?: string;
  };

function MobileMenu({
  tabs,
  selected,
  setSelected,
  onClick,
  onRegisterClick,
  onLogoClick,
  basketItems,
  onCheckout,
  showPopup,
  setShowPopup,
  className
}: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  return (
      <div className={classNames(classes.root, className)}>
          <img
            src={require('../../../assets/logo/ttc-logo.svg')}
            className={classes.logo}
            onClick={onLogoClick}
            />
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
                        className={classes.checkoutPopup}
                        basketItems={basketItems}
                        onCheckout={onCheckout}
                    />
                </div>
                )}
                <Icon
                    name={selected ? 'CloseHamburger' : 'MenuHamburger'}
                    onClick={() => (
                        selected ? setSelected(undefined) : setSelected(tabs[0])
                    )}
                    style={{ cursor: 'pointer' }}
                    size={25}
                />
            </div>
            {selected && (
                <div className={classes.openMenu}>
                    {tabs.map((tab: Tab) => (
                        <div className={classes.option}>
                            <div className={classes.title}>
                                {tab.title.toUpperCase()}
                            </div>
                            {tab.options ? (
                                tab.options.map((option: Tab) => (
                                    <div
                                        className={classes.text}
                                        onClick={() => option.link && onClick(option.link)}
                                    >
                                        {option.title}
                                    </div>
                                ))
                            ) : (
                                <div
                                    className={classes.text}
                                    onClick={() => tab.link && onClick(tab.link)}
                                >
                                    {tab.title}
                                </div>
                            )}
                        </div>
                    ))}
                    <div className={classes.buttonWrapper}>
                        <Button
                            archetype="gradient"
                            className={classes.register}
                            onClick={onRegisterClick}
                        >
                            Register with TTC
                        </Button>
                    </div>
                    <div className={classes.footer}>Already a customer? <span className={classes.link}>Login to My Account</span></div>
                </div>
            )}
      </div>
  );
}

export default MobileMenu;