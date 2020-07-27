import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: theme.colors.navyBlue,
        justifyContent: 'center',
        alignItems: 'center',
        padding: '87px 0'
    },
    title: {
        fontSize: theme.fontSizes.extraLargeHeading,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        marginBottom: '21px',
        textAlign: 'center'
    },
    subtitle: {
        fontSize: theme.fontSizes.tinyHeading,
        fontWeight: '200',
        color: theme.colors.primaryWhite,
        maxWidth: '711px',
        textAlign: 'center'
    }
}));

type Props = {
    title: string;
    subtitle: string;
    className?: string;
};

function TOSHeader({ title, subtitle, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  return (
      <div className={classNames(classes.root, className)}>
          <div className={classes.title}>{title}</div>
          <div className={classes.subtitle}>{subtitle}</div>
      </div>
  );
}

export default TOSHeader;