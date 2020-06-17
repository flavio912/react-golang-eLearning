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
    header: {
        flex: 1,
        fontSize: '46px',
        fontWeight: '800',
        marginTop: '10px',
        color: theme.colors.primaryWhite
    },
    image: {
        height: '140px',
        width: '240px',
        margin: '60px 0 0 50px',
        borderRadius: '20px 20px 0 0'
    },
    spacer: {
        flex: 1,
        height: '200px' // image height + padding
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