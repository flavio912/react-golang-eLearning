import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CoreInput from 'sharedComponents/core/Input/CoreInput';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    padding: '25px',
    borderRadius: '10px',
    boxShadow: '2px 2px 10px rgba(0, 0, 0, 0.11)'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between'
  },
  header: {
    fontSize: theme.fontSizes.heading,
    fontWeight: '900',
    marginBottom: '30px'
  },
  input: {
    fontSize: theme.fontSizes.default,
    padding: '11px',
    border: ['1px', 'solid', theme.colors.borderGrey],
    borderRadius: '6px',
    marginBottom: '15px'
  },
  marginLeft: {
    marginLeft: '27px'
  },
  marginRight: {
    marginRight: '27px'
  }
}));

export type AccountDetails = {
  firstName: string;
  lastName: string;
  emailAddress: string;
  companyName: string;
  phoneNumber: string;
};

type Props = {
  accountDetails: AccountDetails;
  className?: string;
};

function AccountCard({ accountDetails, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Form Data
  const [firstName, setFirstName] = React.useState('');
  const [lastName, setLastName] = React.useState('');
  const [emailAddress, setEmailAddress] = React.useState('');
  const [companyName, setCompanyName] = React.useState('');
  const [phoneNumber, setPhoneNumber] = React.useState('');

  React.useEffect(() => {
    accountDetails = {
      firstName,
      lastName,
      emailAddress,
      companyName,
      phoneNumber
    };
  });

  return (
    <form className={classNames(classes.root, className)}>
      <div className={classes.header}>Create An Account</div>

      <div className={classes.row}>
        <CoreInput
          placeholder="First Name"
          type="text"
          onChange={(text: string) => setFirstName(text)}
          value={firstName}
          className={classes.input}
        />

        <CoreInput
          placeholder="Last Name"
          type="text"
          onChange={(text: string) => setLastName(text)}
          value={lastName}
          className={classNames(classes.input, classes.marginLeft)}
        />
      </div>

      <CoreInput
        placeholder="Email Address"
        type="text"
        onChange={(text: string) => setEmailAddress(text)}
        value={emailAddress}
        className={classes.input}
      />

      <CoreInput
        placeholder="Company Name (if applicable)"
        type="text"
        onChange={(text: string) => setCompanyName(text)}
        value={companyName}
        className={classes.input}
      />

      <div className={classes.row}>
        <CoreInput
          placeholder="Contact Telephone Number"
          type="text"
          onChange={(text: string) => setPhoneNumber(text)}
          value={phoneNumber}
          className={classNames(classes.input, classes.marginRight)}
        />
        <div style={{ flex: 1.05 }} />
      </div>
    </form>
  );
}

export default AccountCard;
