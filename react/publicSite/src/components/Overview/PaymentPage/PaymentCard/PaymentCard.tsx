import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { useRouter } from 'found';
import { CardElement } from '@stripe/react-stripe-js';
import { Theme } from "helpers/theme";
import Button from "sharedComponents/core/Input/Button";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        borderRadius: '14px',
        boxShadow: '0 1px 7px 3px rgba(0,0,0,0.11)',
        
    },
    padded: {
        padding: '40px'
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
        flex: 1
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between'
    },
    header: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '900',
        marginBottom: '50px'
    },
    total: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '500',
        color: theme.colors.textGrey,
        marginBottom: '10px'
    },
    border: {
        padding: '10px 0 20px 0',
        borderBottom: ['1px', 'solid', theme.colors.borderGrey],
    },
    blue: {
        color: theme.colors.navyBlue,
        fontWeight: '700',
        margin: '30px 0 0 0'
    },
    cardType: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        margin: '15px 0 25px 0'
    },
    terms: {
        fontSize: theme.fontSizes.default,
        fontWeight: '800',
        color: theme.colors.textGrey
    },
    service: {
        cursor: 'pointer',
        textDecorationLine: 'underline',
        color: theme.colors.textBlue
    },
    background: {
        padding: '20px 40px 40px 40px',
        backgroundColor: theme.colors.backgroundGrey
    },
    input: {
        backgroundColor: theme.colors.primaryWhite,
        padding: '10px',
        border: ['2px', 'solid', theme.colors.borderGrey],
        borderRadius: '6px',
        margin: '30px 0 20px 0',
        alignItems: 'center'
    },
    error: {
        borderColor: theme.colors.primaryRed
    },
    errorText: {
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.primaryRed
    },
    button: {
        height: '55px',
        width: '100%',
        fontSize: theme.fontSizes.tinyHeading,
        fontWeight: '800',
        marginTop: '20px'
    }
}));

const CARD_ELEMENT_OPTIONS = {
    style: {
      base: {
        color: '#32325d',
        fontFamily: '"Helvetica Neue", Helvetica, sans-serif',
        fontSmoothing: 'antialiased',
        fontSize: '21px',
        '::placeholder': {
          color: '#aab7c4'
        }
      },
      invalid: {
        color: '#CB463A',
        iconColor: '#CB463A'
      }
    }
};

type Props = {
    total: number;
    className?: string;
};

function PaymentCard({ total, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const [error, setError] = React.useState(null);
    const handleChange = (event: any) => {
        console.log(event)
        if (event.error) {
            setError(event.error.message);
        } else {
            setError(null);
        }
    }

    const { router } = useRouter();
    const openTerms = () => {
        // TODO: add link to terms of service
        router.push('/');
    }

  return (
    <div className={classNames(classes.root, className)}>
        <div className={classes.padded}>
            <div className={classes.header}>Make Payment</div>

            <div className={classes.row}>
                <div className={classes.total}>Sub Total:</div>
                <div className={classes.total}>£{(total - (total * 0.2)).toFixed(2)}</div>
            </div>

            <div className={classNames(classes.row, classes.border)}>
                <div className={classes.total}>VAT (20%)</div>
                <div className={classes.total}>£{(total * 0.2).toFixed(2)}</div>
            </div>

            <div className={classes.row}>
                <div className={classNames(classes.total, classes.blue)}>Total Due Today:</div>
                <div className={classNames(classes.total, classes.blue)}>£{total.toFixed(2)}</div>
            </div>
        </div>

        <div className={classNames(classes.background, classes.padded)}>
            <div className={classes.cardType}>Credit or Debit Card</div>
            <div className={classes.terms}>{`By completing your purchase you agree to these `}
                <span
                    className={classes.service}
                    onClick={openTerms}
                >
                    Terms of Service
                </span>
            </div>

            <CardElement
                className={classNames(classes.input, error && classes.error)}
                options={CARD_ELEMENT_OPTIONS}
                onChange={handleChange}
            />
            {error && <div className={classes.errorText}>{error}</div>}

            <Button
                archetype="gradient"
                className={classes.button}
            >
                Complete Payment
            </Button>
        </div>
      </div>
  );
}

export default PaymentCard;