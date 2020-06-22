import * as React from 'react';
import Card from '../../sharedComponents/core/Cards/Card';
import FancyInput from './FancyInput';
import FancyButton from './FancyButton';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { ReactComponent as Logo } from '../../assets/logo/ttc-logo.svg';
import { Link } from 'react-router-dom';
const useStyles = createUseStyles((theme: Theme) => ({
  loginDialogueRoot: {},
  root: {
    display: 'flex',
    width: 378,
    flexDirection: 'column',
    background: 'white'
  },
  logoContainer: {
    padding: [30, 0, 20]
  },
  logo: {
    height: 70
  },
  heading: {
    fontWeight: 800,
    color: theme.colors.primaryBlack
  },
  subheading: {
    color: theme.colors.textGrey,
    fontWeight: 300,
    fontSize: 15,
    marginTop: 0,
    marginBottom: theme.spacing(2)
  },
  link: {
    margin: [15, 0, 30, 0],
    textAlign: 'center',
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.small
  },
  errMessage: {
    color: '#43454a',
    fontWeight: 200,
    fontSize: 15,
    textAlign: 'center',
    margin: 3
  },
  links: {
    marginTop: 8,
    display: 'inline-flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    width: '100%'
  },
  linkText: {
    fontSize: 15,
    letterSpacing: '-0.38',
    lineHeight: '25px',
    color: theme.colors.textBlue,
    textDecoration: 'none'
  },
  helpIcon: {
    color: theme.colors.textBlue,
    fontWeight: 'bold',
    fontSize: theme.fontSizes.xSmallHeading,
    lineHeight: `25px`,
    width: 32,
    height: 32,
    border: `1px solid rgba(16,129,170,0.17)`,
    marginRight: 14.5,
    display: 'inline-flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 50
  }
}));

type Props = {
  onSubmit: (
    email: string,
    password: string,
    errorCallback: (err: string) => void
  ) => void;
};

function LoginDialogue({ onSubmit }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [error, setError] = React.useState('');
  const onLogin = () => {
    onSubmit(email, password, (err) => {
      setError(err);
    });
  };
  return (
    <div className={classes.loginDialogueRoot}>
      <Card padding="medium" className={classes.root}>
        <div className={classes.logoContainer}>
          <Logo className={classes.logo} />
        </div>
        <h1 className={classes.heading}>Login to TTC Hub</h1>
        <p className={classes.subheading}>
          Glad to have you back, please enter your login details to proceed
        </p>
        <p className={classes.errMessage}>{error}</p>
        <FancyInput
          label="Email"
          labelColor={'#5CC301'}
          type={'email'}
          onChange={setEmail}
        />
        <FancyInput
          label="Password"
          labelColor={'#5CC301'}
          type={'password'}
          onChange={setPassword}
        />
        <FancyButton text="Login to TTC" onClick={onLogin} />
        <a className={classes.link} href="https://example.com">
          I don't have a TTC Hub account
        </a>
      </Card>
      <div className={classes.links}>
        <a href="/help" className={classes.linkText}>
          <span className={classes.helpIcon}>?</span>Need Help?
        </a>
        <a href="/forgot-password" className={classes.linkText}>
          Forgot password
        </a>
      </div>
    </div>
  );
}

export default LoginDialogue;
