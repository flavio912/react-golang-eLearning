import * as React from 'react';
import { createUseStyles } from 'react-jss';

const useStyles = createUseStyles({
  title: {
    fontSize: '37px',
    fontWeight: '800',
    marginBottom: '30px',
    textAlign: 'center'
  }
});

type Props = {
  children?: React.ReactNode;
};

function Title({ children }: Props) {
  const classes = useStyles();
  return <div className={classes.title}>{children}</div>;
}

export default Title;
