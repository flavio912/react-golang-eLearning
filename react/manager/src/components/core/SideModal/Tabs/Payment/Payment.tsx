import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import CoreInput from 'components/core/CoreInput';
import PaymentForm from 'components/core/SideModal/PaymentForm';
import PaymentSuccess from 'components/core/SideModal/PaymentSuccess';

const useStyles = createUseStyles((theme: Theme) => ({
  paymentRoot: {},
  paymentItems: {
    marginBottom: 25,
    '& table': {
      width: '100%'
    },
    '& thead': {
      '& td': {
        padding: [10, 20],
        borderBottom: `1px solid ${theme.colors.borderGrey}`,
        '&:first-child': {
          paddingLeft: 0
        },
        '&:last-child': {
          paddingRight: 0
        }
      }
    },
    '& tbody': {
      '& td': {
        padding: [16, 20],
        borderBottom: `1px dashed ${theme.colors.borderGreyBold}`,
        '&:first-child': {
          paddingLeft: 0
        },
        '&:last-child': {
          paddingRight: 0
        }
      }
    }
  },
  tableHeadText: {
    fontSize: theme.fontSizes.tiny,
    color: theme.colors.secondaryBlack,
    letterSpacing: -0.28,
    textTransform: 'uppercase'
  },
  productText: {
    fontSize: theme.fontSizes.default,
    letterSpacing: -0.35,
    color: theme.colors.secondaryBlack,
    display: 'block'
  },
  productSku: {
    display: 'block',
    fontSize: theme.fontSizes.default,
    letterSpacing: -0.35,
    color: theme.colors.textIron
  },
  summaryBox: {
    display: 'flex',
    justifyContent: 'space-between',
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGreyBold}`,
    borderRadius: 4,
    padding: [20, 25]
  },
  summaryBoxLeft: {
    flex: 1
  },
  summaryBoxRight: {
    flex: 1,
    paddingLeft: 52
  },
  cardInfo: {
    marginBottom: 42
  },
  summaryHeading: {
    fontSize: theme.fontSizes.default,
    fontWeight: 900,
    letterSpacing: -0.35,
    color: theme.colors.primaryBlack,
    lineHeight: '41px',
    margin: 0
  },
  inputWrapper: {
    marginTop: 11,
    backgroundColor: theme.colors.primaryWhite,
    borderRadius: 4,
    border: `1px solid ${theme.colors.borderGrey}`,
    // height: 37,
    padding: [8, 14],
    '& input': {
      width: '100%',
      height: '100%',
      padding: 0,
      border: 'none',
      fontSize: theme.fontSizes.default,
      color: theme.colors.secondaryBlack,
      lineHeight: `20px`,
      letterSpacing: -0.35
    }
  },
  summaryTable: {
    width: '100%'
  },
  summaryText: {
    color: theme.colors.secondaryBlack,
    fontSize: theme.fontSizes.smallLarge,
    lineHeight: `29px`
  },
  totalDueText: {
    color: theme.colors.textNavyBlue2,
    fontSize: theme.fontSizes.smallLarge,
    fontWeight: 800,
    lineHeight: `29px`
  },
  smallHeading: {
    fontSize: theme.fontSizes.smallHeading,
    color: theme.colors.primaryBlack,
    fontWeight: 900,
    letterSpacing: -0.6,
    marginBottom: 25,
    lineHeight: `41px`
  },
  cardContent: {
    display: 'flex',
    justifyContent: 'space-between'
  },
  cardRight: {
    flex: 2,
    textAlign: 'right'
  },
  cardRightImage: {
    display: 'inline-block'
  },
  cardLeft: {
    fontSize: theme.fontSizes.tiny,
    color: theme.colors.secondaryBlack,
    lineHeight: `16px`,
    letterSpacing: -0.28,
    flex: 3
  },
  stripeImage: {
    width: 164,
    marginBottom: 14,
    height: 45,
    '& img': {
      width: '100% !important',
      height: 'auto !important'
    }
  },
  paymentImage: {
    width: 196,
    height: 36,
    '& img': {
      width: '100% !important',
      height: 'auto !important'
    }
  },
  breakLine: {
    height: 1,
    width: '100%',
    borderBottom: `1px dashed ${theme.colors.borderGreyBold}`,
    marginTop: 24,
    marginBottom: 42
  },
  paymentPlace: {
    border: `1px solid ${theme.colors.borderGreyBold}`,
    borderRadius: 4,
    backgroundColor: theme.colors.backgroundGrey,
    padding: [14.5, 17.5, 43.5, 20.5],
    marginBottom: 45
  },
  placeHeading: {
    fontSize: theme.fontSizes.default,
    color: theme.colors.primaryBlack,
    fontWeight: 900,
    letterSpacing: -0.35,
    lineHeight: '41px',
    margin: 0
  }
}));
export type Course = {
  id: number;
  name: string;
  sku: string;
  price: number;
  qty: number;
  subtotal: number;
};
export default function Payment({ courses }: { courses: Course[] }) {
  const classes = useStyles();
  const subTotal = courses
    .map(({ subtotal }: { subtotal: number }) => subtotal)
    .reduce(
      (prevValue: number, currentValue: number) => prevValue + currentValue,
      0
    );
  const vat = subTotal * (20 / 100);
  const totalDue = subTotal + vat;
  return (
    <>
      <div className={classes.paymentRoot}>
        <div className={classes.paymentItems}>
          <table>
            <thead>
              <tr>
                <td colSpan={2}>
                  <span className={classes.tableHeadText}>PRODUCT</span>
                </td>
                <td colSpan={1}>
                  <span className={classes.tableHeadText}>QUANTITY</span>
                </td>
                <td colSpan={1}>
                  <span className={classes.tableHeadText}>PRICE</span>
                </td>
                <td colSpan={1}>
                  <span className={classes.tableHeadText}>TOTAL</span>
                </td>
              </tr>
            </thead>
            <tbody>
              {courses.map(
                (
                  { name, qty, subtotal, sku, price }: any,
                  index: string | number
                ) => (
                  <tr key={index}>
                    <td colSpan={2}>
                      <span className={classes.productText}>
                        <strong>{qty}x</strong>
                        {name}
                      </span>
                      <span className={classes.productSku}>SKU: {sku}</span>
                    </td>
                    <td colSpan={1}>
                      <span className={classes.productText}>{qty}</span>
                    </td>
                    <td colSpan={1}>
                      <span className={classes.productText}>£{price}</span>
                    </td>
                    <td colSpan={1}>
                      <span className={classes.productText}>£{subtotal}</span>
                    </td>
                  </tr>
                )
              )}
            </tbody>
          </table>
        </div>
        <div className={classes.summaryBox}>
          <div className={classes.summaryBoxLeft}>
            <p className={classes.summaryHeading}>
              Please send a copy of my Invoice to
            </p>
            <div className={classes.inputWrapper}>
              <CoreInput type="text" onChange={() => {}} />
            </div>
          </div>
          <div className={classes.summaryBoxRight}>
            <p className={classes.summaryHeading}>Order Summary</p>
            <table className={classes.summaryTable}>
              <tbody>
                <tr>
                  <td>
                    <span className={classes.summaryText}>Sub Total:</span>
                  </td>
                  <td>
                    <span className={classes.summaryText}>£{subTotal}</span>
                  </td>
                </tr>
                <tr>
                  <td>
                    <span className={classes.summaryText}>VAT (20%):</span>
                  </td>
                  <td>
                    <span className={classes.summaryText}>£{vat}</span>
                  </td>
                </tr>
                <tr>
                  <td>
                    <span className={classes.totalDueText}>Total Due:</span>
                  </td>
                  <td>
                    <span className={classes.totalDueText}>£{totalDue}</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div className={classes.breakLine} />{' '}
        <div className={classes.cardInfo}>
          <h1 className={classes.smallHeading}>Pay by Card</h1>
          <div className={classes.cardContent}>
            <div className={classes.cardLeft}>
              Pay securely using our credt card. All payments are 128-Bit
              secured by Global Payments. Your personal data will be used to
              process your order, support your experience through this website,
              and for other purposes described in our privacy policy.
            </div>
            <div className={classes.cardRight}>
              <div className={classes.cardRightImage}>
                <div className={classes.stripeImage}>
                  <Icon name="Stripe" />
                </div>
                <div className={classes.paymentImage}>
                  <Icon name="Payments_Method" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <PaymentForm />
        <PaymentSuccess total={totalDue} transactionId={`34239rCD`} />
      </div>
    </>
  );
}
