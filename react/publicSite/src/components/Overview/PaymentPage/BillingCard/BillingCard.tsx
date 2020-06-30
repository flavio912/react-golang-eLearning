import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CoreInput from 'sharedComponents/core/Input/CoreInput';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import CheckboxSingle from 'sharedComponents/core/Input/CheckboxSingle';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    padding: '25px',
    borderRadius: '14px',
    boxShadow: '0 1px 7px 3px rgba(0,0,0,0.11)'
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
  dropdown: {
    flex: 1.05
  },
  dropdownBox: {
    border: ['1px', 'solid', theme.colors.borderGrey],
    borderRadius: '6px'
  },
  dropdownText: {
    fontSize: theme.fontSizes.default,
    fontWeight: '400',
    color: theme.colors.textGrey
  },
  checkbox: {
    flex: 2,
    alignItems: 'flex-start',
    marginLeft: '13.5px'
  },
  checkboxText: {
    fontSize: theme.fontSizes.small,
    color: theme.colors.textGrey,
    margin: '0 15px'
  },
  last: {
    marginBottom: 0
  }
}));

export type BillingDetails = {
  firstName: string;
  lastName: string;
  addressOne: string;
  addressTwo: string;
  city: string;
  postcode: string;
  country: string;
  contact: boolean;
};

type Props = {
  billingDetails: BillingDetails;
  className?: string;
};

function BillingCard({ billingDetails, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Form Data
  const [firstName, setFirstName] = React.useState('');
  const [lastName, setLastName] = React.useState('');
  const [addressOne, setAddressOne] = React.useState('');
  const [addressTwo, setAddressTwo] = React.useState('');
  const [city, setCity] = React.useState('');
  const [postcode, setPostcode] = React.useState('');
  const [country, setCountry] = React.useState({ id: 0, title: 'Country' });
  const [contact, setContact] = React.useState(false);

  React.useEffect(() => {
    billingDetails = {
      firstName,
      lastName,
      addressOne,
      addressTwo,
      city,
      postcode,
      country: country.title !== 'Country' ? country.title : '',
      contact
    };
  });

  return (
    <form className={classNames(classes.root, className)}>
      <div className={classes.header}>Billing Address</div>

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
        placeholder="Address Line 1"
        type="text"
        onChange={(text: string) => setAddressOne(text)}
        value={addressOne}
        className={classes.input}
      />

      <CoreInput
        placeholder="Address Line 1"
        type="text"
        onChange={(text: string) => setAddressTwo(text)}
        value={addressTwo}
        className={classes.input}
      />

      <div className={classes.row}>
        <CoreInput
          placeholder="City"
          type="text"
          onChange={(text: string) => setCity(text)}
          value={city}
          className={classes.input}
        />

        <Dropdown
          placeholder="Country"
          options={[{ id: 0, title: 'U.K' }]}
          selected={country}
          setSelected={(selected: DropdownOption) => setCountry(selected)}
          fontStyle={classes.dropdownText}
          boxStyle={classes.dropdownBox}
          className={classNames(classes.dropdown, classes.marginLeft)}
        />
      </div>

      <div className={classes.row}>
        <CoreInput
          placeholder="Postcode"
          type="text"
          onChange={(text: string) => setPostcode(text)}
          value={postcode}
          className={classNames(classes.input, classes.last)}
        />
        <div style={{ flex: 1 }} />
        <CheckboxSingle
          className={classes.checkbox}
          fontStyle={classes.checkboxText}
          size={18}
          onChange={() => setContact(!contact)}
          label="By checking this box you confirm you are happy for our team to contact you during the registration period"
        />
      </div>
    </form>
  );
}

export default BillingCard;
