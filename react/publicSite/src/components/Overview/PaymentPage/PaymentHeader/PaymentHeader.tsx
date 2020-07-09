import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        backgroundColor: theme.colors.navyBlue,
        justifyContent: 'space-evenly',
        alignItems: 'center',
    },
    spacer: {
        display: 'flex',
        flex: 1,
        alignItems: 'flex-end',
    },
    header: {
        flex: 1,
        fontSize: theme.fontSizes.extraLargeHeading,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        textAlign: 'center',
        margin: '0 20px'
    },
    image: {
        width: '225px',
        marginTop: '40px',
        borderRadius: '20px 20px 0 0'
    }
}));

type Props = {
    text: string;
    imageURL: string;
    className?: string;
};

function PaymentHeader({ text, imageURL, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  return (
      <div className={classNames(classes.root, className)}>
          <div className={classes.spacer} />
          <div className={classes.header}>{text}</div>
          <div className={classes.spacer}>
            <img className={classes.image} src={imageURL} />
          </div>
      </div>
  );
}

export default PaymentHeader;