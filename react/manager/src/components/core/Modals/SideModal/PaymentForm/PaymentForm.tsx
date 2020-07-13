import * as React from 'react';
import { createUseStyles } from 'react-jss';
import theme, { Theme } from 'helpers/theme';
import { loadStripe, ConfirmCardPaymentData } from '@stripe/stripe-js';
import {
  CardElement,
  Elements,
  useStripe,
  useElements
} from '@stripe/react-stripe-js';
import Button from 'components/core/Input/Button';
import { OnPurchase } from '../CourseManagement/MultiUser/Tabs';

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
const stripePromise = loadStripe('pk_test_T5ZBhTO9Lq709gdga8c9aoPN00PnTm0tfU');

type Props = {
  onPurchase: OnPurchase;
};

function PaymentForm({ onPurchase }: Props) {
  const classes = useStyles();
  const stripe = useStripe();
  const elements = useElements();
  const handleSubmit = async (event: any) => {
    event.preventDefault();

    // Get clientsecret
    onPurchase(async (resp, err) => {
      console.log('RESP', resp);
      console.log('purerr', err);

      if (resp.purchaseCourses?.transactionComplete) {
        console.log('Transaction unexpectly complete');
        return;
      }

      if (!resp.purchaseCourses?.stripeClientSecret) return;

      const cardEl = elements?.getElement(CardElement);
      if (!cardEl) return;

      const paymentData: ConfirmCardPaymentData = {
        payment_method: {
          card: cardEl
        }
      };

      const result = await stripe?.confirmCardPayment(
        resp.purchaseCourses?.stripeClientSecret,
        paymentData
      );

      console.log('res', result);
    });
  };
  return (
    <div className={classes.paymentFormRoot}>
      <label className={classes.cardTitle}>Credit or debit card*</label>
      <form className={classes.form} onSubmit={handleSubmit}>
        <CardElement
          options={{
            style: {
              base: {
                fontSize: `${theme.fontSizes.default}px`,
                fontFamily: `'Muli', sans-serif`,
                lineHeight: `31px`,
                color: theme.colors.secondaryBlack
              }
            }
          }}
        />
        <Button
          title={'Place Order'}
          onClick={() => {}}
          type="submit"
          className={classes.buttonOrder}
          padding="medium"
          disabled={!stripe}
        />
      </form>
    </div>
  );
}
function PaymentStripeProvider({ onPurchase }: Props) {
  return (
    <Elements stripe={stripePromise}>
      <PaymentForm onPurchase={onPurchase} />
    </Elements>
  );
}
export default PaymentStripeProvider;
