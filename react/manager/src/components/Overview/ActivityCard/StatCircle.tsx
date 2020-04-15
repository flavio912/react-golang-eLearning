import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    padding: '0px 20px 0 25px'
  },
  circle: {
    width: '10px',
    height: '10px',
    borderRadius: '10px',
    marginRight: '15px',
  },
  heading: {
    fontSize: theme.fontSizes.small,
    fontWeight: 300,
    color: theme.colors.primaryBlack,
  },
  statValue: {
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: 800,
    color: theme.colors.primaryBlack,
  },
  borderLeft: {
    borderLeft: `1px solid ${theme.colors.borderGrey}`
  },
  borderRight: {
    borderRight: `1px solid ${theme.colors.borderGrey}`
  },
  green: {
    backgroundColor: theme.colors.primaryGreen,
  },
  red: {
    backgroundColor: theme.colors.primaryRed,
  },
  grey: {
    backgroundColor: theme.colors.secondaryGrey,
  }
}));

export type BorderOptions = "none" | "left" | "right";
export type ColorOptions = "green" | "red" | "grey";

type Props = {
  heading: String;
  value: Number;
  border?: BorderOptions;
  color?: ColorOptions;
  className?: string;
};

function StatCircle({ heading, value, border = 'none', color = 'green', className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const borderLink = {
    none: "",
    left: classes.borderLeft,
    right: classes.borderRight,
  };

  return (
    <div className={classNames(classes.root, borderLink[border], className)}>
        <div className={classNames(classes.circle, classes[color])}/>
        <div>
            <div className={classNames(classes.heading)}>{heading}</div>
            <div className={classNames(classes.statValue)}>{value}</div>
        </div>
    </div>
  );
}

export default StatCircle;
