import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
import IdentTag from 'components/IdentTag';
import LabelledCard from 'sharedComponents/core/Cards/LabelledCard';
const useStyles = createUseStyles((theme: Theme) => ({
  paymentSuccessRoot: {},
  paymentSuccessContent: {
    display: 'flex',
    justifyContent: 'space-between'
  },
  paymentSuccessLeft: {
    flex: 2
  },
  paymentSuccessRight: {
    flex: 1,
    textAlign: 'right'
  },
  paymentSuccessLeftHead: {
    fontSize: theme.fontSizes.large,
    color: theme.colors.primaryBlack,
    fontWeight: 'bold',
    lineHeight: '21px',
    letterSpacing: -0.4,
    marginTop: 0,
    marginBottom: 10
  },
  paymentSuccessLeftDescription: {
    fontSize: theme.fontSizes.large,
    color: theme.colors.textGrey2,
    fontWeight: 'bold',
    lineHeight: '21px',
    letterSpacing: -0.4
  },
  paymentSuccessHead: {
    fontSize: theme.fontSizes.smallLarge,
    letterSpacing: -0.38,
    lineHeight: `29px`,
    color: theme.colors.secondaryBlack,
    marginBottom: 17,
    marginTop: 0
  }
}));
type Props = {
  total: number;
  transactionId: number | string;
};
export default function PaymentSuccess({ total, transactionId }: Props) {
  const classes = useStyles();
  return (
    <LabelledCard label={'Transaction Successful'}>
      <div className={classes.paymentSuccessContent}>
        <div className={classes.paymentSuccessLeft}>
          <h6 className={classes.paymentSuccessLeftHead}>
            Thanks, we’ve received your payment successfully.
          </h6>
          <p className={classes.paymentSuccessLeftDescription}>
            Your receipt and invoice has been emailed to you, feel free now to
            close this window and access your new course.
          </p>
        </div>
        <div className={classes.paymentSuccessRight}>
          {/* <p className={classes.paymentSuccessHead}>
            Amount recieved: £{total}
          </p> */}
          <IdentTag label={'TRANS-ID'} ident={'REF-1231432'} />
        </div>
      </div>
    </LabelledCard>
  );
}
