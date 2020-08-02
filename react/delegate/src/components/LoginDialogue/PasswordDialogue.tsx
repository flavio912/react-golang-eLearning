import * as React from 'react';
import Card from '../../sharedComponents/core/Cards/Card';
import FancyInput from './FancyInput';
import FancyButton from './FancyButton';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { ReactComponent as Logo } from '../../assets/logo/ttc-logo.svg';

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
    cursor: 'pointer',
    margin: '22px 0 19px 0',
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
}));

type Props = {
  onBack: () => void;
  onSubmit: (
    email: string,
    errorCallback: (err: string) => void
  ) => void;
};

function PasswordDialogue({ onBack, onSubmit }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [email, setEmail] = React.useState('');
  const [error, setError] = React.useState('');
  const onRecover = () => {
    onSubmit(email, (err) => {
      setError(err);
    });
  };
  return (
    <div className={classes.loginDialogueRoot}>
      <Card padding="medium" className={classes.root}>
        <div className={classes.logoContainer}>
          <Logo className={classes.logo} />
        </div>
        <h1 className={classes.heading}>Password recovery</h1>
        <p className={classes.subheading}>
          Please enter your e-mail address below to recover your password
        </p>
        <p className={classes.errMessage}>{error}</p>
        <form
          onSubmit={(evt) => {
            evt.preventDefault();
            onRecover();
          }}
        >
          <FancyInput
            label="Email"
            labelColor={'#5CC301'}
            type={'text'}
            onChange={setEmail}
          />
          <FancyButton text="Recover password" onClick={onRecover} />
        </form>
        <a className={classes.link} onClick={onBack}>
          Back
        </a>
      </Card>
    </div>
  );
}

export default PasswordDialogue;
