import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
const useStyles = createUseStyles((theme: Theme) => ({
  paymentSuccessRoot: {},
  successHeading: {
    background: `${theme.paymentSuccessBackgroundGradient}`,
    height: 30,
    paddingLeft: 17,
    borderRadius: [4, 4, 0, 0],
    fontSize: theme.fontSizes.small,
    fontWeight: 'bold',
    letterSpacing: -0.33,
    color: theme.colors.primaryWhite,
    display: 'flex',
    alignItems: 'center'
  },
  paymentSuccess: {
    border: `1px solid ${theme.colors.borderGreyBold}`,
    borderRadius: 4,
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: `2px 2px 10px 0 rgba(0, 0, 0, 0.07)`
  },
  paymentSuccessContent: {
    padding: [21, 23, 19, 19.5],
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
  },
  trans: {
    backgroundColor: `rgba(21,195,36,0.05)`,
    borderRadius: 3,
    display: 'inline-flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: [5, 8]
  },
  transTag: {
    height: 16,
    width: 53,
    border: `1px solid #81BE86`,
    borderRadius: 3,
    backgroundColor: `rgba(208, 243, 211, 0)`,
    boxShadow: `0 1px 0 0 rgba(0, 0, 0, 0.1)`,
    fontSize: 9,
    lineHeight: `11px`,
    fontWeight: 500,
    letterSpacing: -0.48,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  transID: {
    fontSize: theme.fontSizes.tiny,
    letterSpacing: -0.59,
    lineHeight: '14px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginLeft: 11,
    '& strong': {
      fontWeight: 500
    }
  }
}));
type Props = {
  total: number;
  transactionId: number | string;
};
export default function PaymentSuccess({ total, transactionId }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.paymentSuccessRoot}>
      <div className={classes.paymentSuccess}>
        <div className={classes.successHeading}>Transaction Successful</div>
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
            <p className={classes.paymentSuccessHead}>
              Amount recieved: £{total}
            </p>
            <div className={classes.trans}>
              <div className={classes.transTag}>TRANS-ID</div>
              <div className={classes.transID}>
                <strong>REF-</strong>
                {transactionId}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
