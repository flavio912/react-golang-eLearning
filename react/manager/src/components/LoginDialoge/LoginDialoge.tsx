import * as React from 'react';
import Card from '../core/Card';
import FancyInput from './FancyInput';
import { createUseStyles, useTheme } from 'react-jss';

const useStyles = createUseStyles((theme: any) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    width: 340
  },
  heading: {
    fontWeight: 800,
    color: theme.color.primaryBlack
  },
  subheading: {
    color: '#22233194',
    fontWeight: 400,
    fontSize: 15,
    marginTop: 0,
    marginBottom: theme.spacing(2)
  }
}));

type Props = {};

function LoginDialoge(props: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <Card padding='medium' className={classes.root}>
      <h1 className={classes.heading}>Login to TTC Hub</h1>
      <p className={classes.subheading}>
        Glad to have you back, please enter your login details to proceed
      </p>
      <FancyInput label='Email' labelColor={'#5CC301'} type={'email'} />
      <FancyInput label='Password' labelColor={'#5CC301'} type={'password'} />
    </Card>
  );
}

export default LoginDialoge;
