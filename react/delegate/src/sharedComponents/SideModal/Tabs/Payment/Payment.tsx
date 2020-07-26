import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import CoreInput from 'sharedComponents/core/Input/CoreInput';
import PaymentForm from 'sharedComponents/SideModal/PaymentForm';
import PaymentSuccess from 'sharedComponents/SideModal/PaymentSuccess';
import Button from 'sharedComponents/core/Input/Button';
import { OnPurchase } from 'sharedComponents/SideModal/PaymentForm/PaymentForm';
import environment from 'api/environment';
import { commitMutation, graphql } from 'react-relay';
import { GraphError } from 'types/general';
import { Payment_PurchaseMutationResponse } from './__generated__/Payment_PurchaseMutation.graphql';

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
        borderBottom: `1px dashed #E9EBEB`,
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
    display: 'block',
    fontWeight: 300
  },
  productSku: {
    display: 'block',
    fontSize: theme.fontSizes.default,
    letterSpacing: -0.35,
    color: '#CCCDCD'
  },
  summaryBox: {
    display: 'flex',
    justifyContent: 'space-between',
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid #E9EBEB`,
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
    fontSize: 15,
    lineHeight: `29px`
  },
  totalDueText: {
    color: '#0E63E8',
    fontSize: 15,
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
    borderBottom: `1px dashed #E9EBEB`,
    marginTop: 24,
    marginBottom: 42
  },
  paymentPlace: {
    border: `1px solid #E9EBEB`,
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
};

const mutation = graphql`
  mutation Payment_PurchaseMutation(
    $courses: [Int!]!
    $users: [UUID!]!
    $extraEmail: String
  ) {
    purchaseCourses(
      input: {
        courses: $courses
        users: $users
        extraInvoiceEmail: $extraEmail
        acceptedTerms: true
        backgroundCheckConfirm: true
      }
    ) {
      transactionComplete
      stripeClientSecret
    }
  }
`;

type PurchaseCallback = (
  response: Payment_PurchaseMutationResponse,
  error: string | undefined
) => void;

const PurchaseCourses = (
  courses: number[],
  users: string[],
  extraEmail: string,
  callback?: PurchaseCallback
) => {
  const variables = {
    courses,
    users,
    extraEmail
  };

  if (!users) {
    if (callback)
      callback(
        { purchaseCourses: null },
        'No users given, cannot purchase course without users'
      );
  }

  if (!courses) {
    if (callback)
      callback(
        { purchaseCourses: null },
        'Please select at least one course to purchase'
      );
  }

  commitMutation(environment, {
    mutation,
    variables,
    onCompleted: async (
      response: Payment_PurchaseMutationResponse,
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        if (callback) {
          callback(response, `${errors[0]?.extensions?.message}`);
        }
        return;
      }

      if (callback) {
        callback(response, undefined);
      }
      console.log('Response received from server.', response, errors);
    },
    onError: (err) => {
      console.log('ERR', err);
      if (callback) callback({ purchaseCourses: null }, err.message);
    }
  });
};

export default function Payment({
  courses,
  userUUIDs,
  isContract,
  onSuccess,
  onError
}: {
  courses: Course[];
  userUUIDs: string[];
  isContract: boolean;
  onSuccess: () => void;
  onError: (message: string) => void;
}) {
  const classes = useStyles();
  const subTotal = courses
    .map(({ price }: { price: number }) => price)
    .reduce(
      (prevValue: number, currentValue: number) => prevValue + currentValue,
      0
    );
  const vat = subTotal * (20 / 100);
  const totalDue = subTotal + vat;

  const onPurchase = (callback?: PurchaseCallback) =>
    PurchaseCourses(
      courses.map((course: Course) => course.id),
      userUUIDs,
      'email@email.com',
      callback
    );

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
              {courses.map(({ name, sku, price }, index: string | number) => (
                <tr key={index}>
                  <td colSpan={2}>
                    <span className={classes.productText}>
                      <strong>{userUUIDs.length}x </strong>
                      {name}
                    </span>
                    <span className={classes.productSku}>SKU: {sku}</span>
                  </td>
                  <td colSpan={1}>
                    <span className={classes.productText}>
                      {userUUIDs.length}
                    </span>
                  </td>
                  <td colSpan={1}>
                    <span className={classes.productText}>£{price}</span>
                  </td>
                  <td colSpan={1}>
                    <span className={classes.productText}>£{price}</span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div className={classes.summaryBox}>
          <div className={classes.summaryBoxLeft}>
            <p className={classes.summaryHeading}>
              Please send a copy of my Invoice to
            </p>
            <div className={classes.inputWrapper}>
              <CoreInput
                type="text"
                onChange={() => {}}
                placeholder={'email@email.com'}
              />
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
        {!isContract && (
          <>
            <div className={classes.cardInfo}>
              <h1 className={classes.smallHeading}>Pay by Card</h1>
              <div className={classes.cardContent}>
                <div className={classes.cardLeft}>
                  Pay securely using our credt card. All payments are 128-Bit
                  secured by Global Payments. Your personal data will be used to
                  process your order, support your experience through this
                  website, and for other purposes described in our privacy
                  policy.
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
            <PaymentForm
              onPurchase={onPurchase}
              onSuccess={onSuccess}
              onError={onError}
            />
          </>
        )}
        {isContract && (
          <Button onClick={() => onPurchase()}>Place Order</Button>
        )}
      </div>
    </>
  );
}
