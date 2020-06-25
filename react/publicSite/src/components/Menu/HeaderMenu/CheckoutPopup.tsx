import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Button from 'sharedComponents/core/Input/Button';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    borderRadius: '10px',
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: '2px 7px 20px 2px rgba(0,0,0,0.15)',
    position: 'absolute',
    marginTop: '18px',
    zIndex: 10
  },
  column: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'
  },
  triangle: {
    width: 0,
    height: 0,
    borderRight: '12px solid transparent',
    borderLeft: '12px solid transparent',
    borderBottom: ['20px', 'solid', theme.colors.primaryWhite],
    filter: 'drop-shadow(1px -1px 1px rgba(0,0,0,0.09))',
    position: 'absolute',
    top: -18,
    right: 75
  },
  item: {
    padding: '25px 25px 6.25px 25px'
  },
  image: {
    height: '70px',
    width: '80px',
    borderRadius: '3px',
    marginRight: '18.5px'
  },
  text: {
    fontSize: 15,
    fontWeight: 'bold'
  },
  name: {
    fontWeight: '700',
    marginBottom: '7px',
    maxWidth: '250px'
  },
  price: {
    color: theme.colors.navyBlue,
    fontWeight: '600'
  },
  total: {
    fontSize: theme.fontSizes.large,
    margin: '10px',
    color: theme.colors.textGrey
  },
  button: {
    height: '38.5px',
    width: '122px',
    borderRadius: '5px'
  },
  background: {
    backgroundColor: theme.colors.backgroundGrey,
    padding: '15px',
    marginTop: '18.75px',
    borderRadius: '0 0 10px 10px'
  },
  black: {
    color: theme.colors.primaryBlack
  },
  start: {
    justifyContent: 'flex-start'
  },
  backgroundHider: {
    position: 'absolute',
    width: '100%',
    height: '100%',
    top: 0,
    left: 0
  }
}));

export type BasketItem = {
  id: string | number;
  name: string;
  price: number;
  imageURL: string;
};

type Props = {
  showPopup: boolean;
  onHide: () => void;
  basketItems: BasketItem[];
  onCheckout: () => void;
  className?: string;
};

const CheckoutPopup = ({
  basketItems,
  onCheckout,
  className,
  onHide,
  showPopup
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  let total: number = 0;
  if (basketItems) {
    basketItems.map((item: BasketItem) => {
      total += item.price;
    });
  }

  return (
    <>
      <div className={classes.backgroundHider} onClick={() => onHide()} />
      <div
        className={classNames(classes.root, className)}
        onClick={(evt) => evt.stopPropagation()}
      >
        <div className={classes.triangle} />
        {basketItems &&
          basketItems.map((item: BasketItem) => (
            <div
              key={item.id}
              className={classNames(classes.row, classes.item, classes.start)}
            >
              <img src={item.imageURL} className={classes.image} />
              <div className={classes.column}>
                <div className={classNames(classes.text, classes.name)}>
                  {item.name}
                </div>
                <div className={classNames(classes.text, classes.price)}>
                  £{item.price.toFixed(2)}
                </div>
              </div>
            </div>
          ))}
        <div className={classNames(classes.row, classes.background)}>
          <div className={classNames(classes.text, classes.total)}>
            Total: <span className={classes.black}>£{total.toFixed(2)}</span>
          </div>
          <Button
            archetype="submit"
            className={classNames(classes.text, classes.button)}
            onClick={onCheckout}
          >
            Checkout
          </Button>
        </div>
      </div>
    </>
  );
};

export default CheckoutPopup;
