import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
import classnames from 'classnames';
import Button from 'sharedComponents/core/Input/Button';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import Icon from 'sharedComponents/core/Icon';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import CoreInput from 'sharedComponents/core/Input/CoreInput';
import Title from './Title';
import Subtitle from './Subtitle';
import Footer from './Footer';
import CheckboxSingle from 'sharedComponents/core/Input/CheckboxSingle';

const useStyles = createUseStyles((theme: Theme) => ({
  registerCompanyRoot: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center'
  },
  button: {
    width: '100%',
    height: '52px',
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
    color: theme.colors.primaryWhite,
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)'
  },
  form: {
    display: 'grid',
    width: '100%',
    gridGap: theme.spacing(2),
    gridTemplateColumns: '1fr 1fr'
  },
  fullWidth: {
    gridColumn: '1 / 3'
  },
  input: {
    fontSize: theme.fontSizes.default,
    padding: '14px',
    border: ['1px', 'solid', theme.colors.borderGrey],
    borderRadius: theme.buttonBorderRadius,
    '&::placeholder': {
      color: theme.colors.secondaryBlack
    }
  },
  dropdownText: {
    fontSize: theme.fontSizes.default,
    fontWeight: '400',
    color: theme.colors.secondaryBlack
  },
  dropdown: {
    flex: 1
  },
  checkboxText: {
    color: theme.colors.textGrey
  }
}));

type Props = {
  onSubmit: (
    fname: string,
    lname: string,
    email: string,
    telephone: string,
    role: string
  ) => void;
};

function RegisterCompany({ onSubmit }: Props) {
  const classes = useStyles();

  const [firstName, setFirstName] = React.useState('');
  const [lastName, setLastName] = React.useState('');
  const [email, setEmail] = React.useState('');
  const [telephone, setTelephone] = React.useState('');
  const [role, setRole] = React.useState({
    id: 0,
    title: 'What role best describes you'
  });
  const [agreedTerms, setAgreedTerms] = React.useState(false);

  const submitInfo = () => {
    onSubmit(firstName, lastName, email, telephone, role.title);
  };

  return (
    <div className={classes.registerCompanyRoot}>
      <Icon name="TTC_Logo_Icon" size={44} />
      <Spacer vertical spacing={4} />
      <Title>Register your team today</Title>
      <Subtitle>
        If you’re looking to get access to the finest level of compliance
        training register with TTC Hub below
      </Subtitle>
      <Spacer vertical spacing={3} />
      <div className={classes.form}>
        <CoreInput
          placeholder="First Name"
          type="text"
          onChange={setFirstName}
          value={'fname'}
          className={classes.input}
        />
        <CoreInput
          placeholder="Last Name"
          type="text"
          onChange={setLastName}
          value={'lname'}
          className={classes.input}
        />
        <CoreInput
          placeholder="Company Email"
          type="email"
          onChange={setEmail}
          value={'email'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Telephone Number"
          type="tel"
          onChange={setTelephone}
          value={'telephone'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <Dropdown
          placeholder="Type of Company"
          options={[{ id: 0, title: 'Default Option' }]}
          selected={role}
          fontStyle={classes.dropdownText}
          setSelected={(selected: DropdownOption) => setRole(selected)}
          className={classnames(classes.dropdown, classes.fullWidth)}
        />
        <Dropdown
          placeholder="What role best describes you"
          options={[{ id: 0, title: 'Default Option' }]}
          selected={role}
          fontStyle={classes.dropdownText}
          setSelected={(selected: DropdownOption) => setRole(selected)}
          className={classnames(classes.dropdown, classes.fullWidth)}
        />
        <CheckboxSingle
          size={18}
          onChange={() => {}}
          fontStyle={classes.checkboxText}
          label="By checking this box you confirm you are happy for our team to contact you during the registration period"
        />
      </div>
      <Spacer vertical spacing={3} />
      <Button
        archetype="submit"
        className={classes.button}
        onClick={submitInfo}
      >
        Register with TTC
      </Button>
      <Spacer vertical spacing={3} />
    </div>
  );
}

export default RegisterCompany;
