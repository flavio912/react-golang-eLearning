import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  subtitle: {
    fontSize: theme.fontSizes.tinyHeading,
    color: theme.colors.textGrey,
    marginBottom: '35px',
    textAlign: 'center'
  }
}));

type Props = {
  children?: React.ReactNode;
};

function Subtitle({ children }: Props) {
  const classes = useStyles();
  return <div className={classes.subtitle}>{children}</div>;
}

export default Subtitle;
