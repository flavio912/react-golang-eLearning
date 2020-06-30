import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex'
    },
    row: {
        flexDirection: 'row'
    },
    column: {
        flexDirection: 'column',
    },
    left: {
        marginRight: '25px',
        borderLeft: ['3px', 'solid', theme.colors.primaryBlack]
    },
    top: {
        width: '105px',
        borderBottom: ['3px', 'solid', theme.colors.primaryBlack]
    },
    topQuote: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '800',
        marginTop: '15px'
    },
    leftQuote: {
        fontSize: theme.fontSizes.tinyHeading,
        fontStyle: 'italic',
        color: theme.colors.textGrey,
        margin: '15px 0'
    }
}));

export type BorderSide = "top" | "left";

type Props = {
    borderSide: BorderSide;
    quote: string;
    className?: string;
};

function Quote({ borderSide, quote, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const direction = {
        top: 'column', left: 'row'
    }

  return (
      <div
        className={
            classNames(classes.root, classes[direction[borderSide]], className)
        }
      >
          <div className={classes[borderSide]} />
          <div className={classes[borderSide + 'Quote']}>{quote}</div>
      </div>
  );
}

export default Quote;