import * as React from 'react';
import Card from '../core/Card';
import FancyInput from './FancyInput';
import { createUseStyles, useTheme } from 'react-jss';

const useStyles = createUseStyles({
  root: {
    display: 'flex',
    flexDirection: 'column'
  }
});

type Props = {};

function LoginDialoge(props: Props) {
  const classes = useStyles();
  return (
    <Card padding='medium' className={classes.root}>
      <h1>Login to TTC Hub</h1>
      <p>Glad to have you back, please enter your login details to proceed</p>
      <FancyInput label='Email' />
    </Card>
  );
}

export default LoginDialoge;
