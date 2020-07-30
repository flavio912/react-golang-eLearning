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
    fontSize: 25,
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
    margin: '22px 0 19px 0',
    textAlign: 'center',
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.small
  },
  errMessage: {
    color: theme.colors.secondaryDanger,
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
  email: string;
  onSubmit: (
    email: string,
    password: string,
    passwordRepeat: string,
    errorCallback: (err: string) => void
  ) => void;
};

function FinaliseDialogue({ email, onSubmit }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [password, setPassword] = React.useState('');
  const [passwordRepeat, setPasswordRepeat] = React.useState('');
  const [error, setError] = React.useState('');
  const onLogin = () => {
    onSubmit(email, password, passwordRepeat, (err) => {
      setError(err);
    });
  };
  return (
    <div className={classes.loginDialogueRoot}>
      <Card padding="medium" className={classes.root}>
        <div className={classes.logoContainer}>
          <Logo className={classes.logo} />
        </div>
        <h1 className={classes.heading}>Finalise your account</h1>
        <p className={classes.subheading}>
          Please enter your new password
        </p>
        <form
          onSubmit={(evt) => {
            evt.preventDefault();
            onLogin();
          }}
        >
          <FancyInput
            label="Email"
            labelColor={'#5CC301'}
            type={'text'}
            placeholder={email}
            disabled
          />
          <FancyInput
            label="New Password"
            labelColor={'#5CC301'}
            type={'password'}
            onChange={setPassword}
          />
          <FancyInput
            label="New Password Repeated"
            labelColor={'#5CC301'}
            type={'password'}
            onChange={setPasswordRepeat}
          />
          <p className={classes.errMessage}>{error}</p>
          <FancyButton text="Save and Login" onClick={onLogin} />
        </form>
      </Card>
    </div>
  );
}

export default FinaliseDialogue;
