import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { useRouter } from 'found';
import { loadStripe } from '@stripe/stripe-js';
import { Elements } from '@stripe/react-stripe-js';
import { Theme } from 'helpers/theme';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import PaymentHeader from 'components/Overview/PaymentPage/PaymentHeader';
import PaymentCard from 'components/Overview/PaymentPage/PaymentCard';
import BillingCard, {
  BillingDetails
} from 'components/Overview/PaymentPage/BillingCard';
import OrderCard, {
  OrderItem
} from 'components/Overview/PaymentPage/OrderCard';
import AccountCard, {
  AccountDetails
} from 'components/Overview/PaymentPage/AccountCard';

const useStyles = createUseStyles((theme: Theme) => ({
  courseRoot: {
    width: '100%',
    backgroundColor: theme.colors.primaryWhite
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  },
  centered: {
    display: 'flex',
    width: theme.centerColumnWidth
  },
  leftColumn: {
    display: 'flex',
    flexDirection: 'column',
    flex: 2
  },
  rightColumn: {
    display: 'flex',
    flexDirection: 'column',
    flex: 1,
    marginLeft: '30px'
  },
  privacy: {
    fontSize: theme.fontSizes.tiny,
    color: theme.colors.textGrey,
    margin: '40px 50px',
    textAlign: 'center'
  },
  underline: {
    cursor: 'pointer',
    textDecorationLine: 'underline'
  }
}));

const emptyAccountDetails: AccountDetails = {
  firstName: '',
  lastName: '',
  emailAddress: '',
  companyName: '',
  phoneNumber: ''
};

const emptyBillingDetails: BillingDetails = {
  addressOne: '',
  addressTwo: '',
  city: '',
  postcode: '',
  country: '',
  contact: false
};

const defaultOrderItems: OrderItem[] = [
  {
    id: '082739428373',
    name: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
    quantity: 1,
    price: 200,
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  },
  {
    id: '082739428374',
    name: 'Cargo Aircraft Protection',
    quantity: 3,
    price: 55,
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  },
  {
    id: '082739428375',
    name: 'Cargo Manager Recurrent (CM) – VC, HS, XRY, EDS',
    quantity: 2,
    price: 25,
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  }
];

// Setup Stripe.js and the Elements provider
const stripePromise = loadStripe('pk_test_TYooMQauvdEDq54NiTphI7jx');

type Props = {};

function PaymentPage({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { router } = useRouter();
  const onPrivacy = () => {
    // TODO: add link to terms of service
    router.push('/');
  };

  return (
    <div className={classes.courseRoot}>
      <PaymentHeader
        text="Checkout Securely"
        imageURL={require('assets/StripePayment.svg')}
      />
      <Spacer spacing={4} vertical />
      <div className={classes.centerer}>
        <div className={classes.centered}>
          <div className={classes.leftColumn}>
            <AccountCard accountDetails={emptyAccountDetails} />
            <Spacer vertical spacing={3} />
            <BillingCard billingDetails={emptyBillingDetails} />
            <Spacer vertical spacing={3} />
            <OrderCard orderItems={defaultOrderItems} />
            <Spacer vertical spacing={3} />
          </div>
          <div className={classes.rightColumn}>
            <Elements stripe={stripePromise}>
              <PaymentCard total={336} card={null} />
            </Elements>
            <div className={classes.privacy}>
              Pay securely using our credt card. All payments are 128-Bit
              secured by Global Payments. Your personal data will be used to
              process your order, support your experience through this website,
              and for other purposes described in our
              <span className={classes.underline} onClick={onPrivacy}>
                privacy policy
              </span>
              .
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PaymentPage;
