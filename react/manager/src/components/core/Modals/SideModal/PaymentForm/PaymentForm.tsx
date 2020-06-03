import * as React from 'react';
import { createUseStyles } from 'react-jss';
import theme, { Theme } from 'helpers/theme';
import {
  CardElement,
  injectStripe,
  StripeProvider
} from 'react-stripe-elements';
import { Elements } from 'react-stripe-elements';
import Button from 'components/core/Input/Button';
const useStyles = createUseStyles((theme: Theme) => ({
  paymentFormRoot: {
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGreyBold}`,
    borderRadius: 4,
    padding: [14.5, 17.5, 43.5, 20.5],
    marginBottom: 25
  },
  cardTitle: {
    lineHeight: `41px`,
    letterSpacing: -0.35,
    fontSize: theme.fontSizes.default,
    fontWeight: 900,
    color: theme.colors.primaryBlack,
    marginBottom: 5
  },
  buttonOrder: {
    background: theme.paymentButtonBackgroundGradient,
    boxShadow: `0 2px 9px 0 rgba(14,99,232,0.35)`,
    width: 192,
    border: 'none',
    height: 39,
    marginLeft: 23
  },
  form: {
    display: 'flex',
    '& .StripeElement': {
      padding: [3, 23, 5, 5],
      flex: 1,
      background: theme.colors.primaryWhite,
      borderRadius: 5,
      border: `1px solid ${theme.colors.textGrey3}`,
      boxShadow: `0 2px 4px 0 rgba(0,0,0,0.08)`
    }
  }
}));
function PaymentForm({}: any) {
  const classes = useStyles();
  return (
    <div className={classes.paymentFormRoot}>
      <label className={classes.cardTitle}>Credit or debit card*</label>
      <div className={classes.form}>
        <CardElement
          style={{
            base: {
              fontSize: `${theme.fontSizes.default}px`,
              fontFamily: `'Muli', sans-serif`,
              lineHeight: `31px`,
              color: theme.colors.secondaryBlack
            }
          }}
        />
        <Button
          title={'Place Order'}
          onClick={() => {}}
          className={classes.buttonOrder}
          padding="medium"
        />
      </div>
    </div>
  );
}
const InjectPaymentForm = injectStripe(PaymentForm);
function PaymentStripProvider() {
  return (
    <StripeProvider apiKey="pk_test_12345">
      <Elements>
        <InjectPaymentForm />
      </Elements>
    </StripeProvider>
  );
}
export default PaymentStripProvider;
