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
}));

type Props = {
  TTC_ID: string;
  onSubmit: (
    password: string,
    passwordRepeat: string,
    errorCallback: (err: string) => void
  ) => void;
};

function FinaliseDialogue({ TTC_ID, onSubmit }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [password, setPassword] = React.useState('');
  const [passwordRepeat, setPasswordRepeat] = React.useState('');
  const [error, setError] = React.useState('');
  const onLogin = () => {
    onSubmit(password, passwordRepeat, (err) => {
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
            label="TTC ID"
            labelColor={'#5CC301'}
            type={'text'}
            placeholder={TTC_ID}
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
