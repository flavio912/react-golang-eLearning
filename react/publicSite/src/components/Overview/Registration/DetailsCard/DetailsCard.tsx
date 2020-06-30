import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import Button from 'sharedComponents/core/Input/Button';
import CoreInput from 'sharedComponents/core/Input/CoreInput';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import CheckboxSingle from 'sharedComponents/core/Input/CheckboxSingle';
import Title from '../Title';
import Subtitle from '../Subtitle';
import Footer from '../Footer';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between'
  },
  logo: {
    margin: '83px 0 73px 0'
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    maxWidth: '400px'
  },
  input: {
    fontSize: theme.fontSizes.default,
    padding: '11px',
    border: ['2px', 'solid', theme.colors.borderGrey],
    borderRadius: '6px',
    marginBottom: '15px'
  },
  marginLeft: {
    marginLeft: '27px'
  },
  dropdown: {
    flex: 1
  },
  dropdownBox: {
    border: ['2px', 'solid', theme.colors.borderGrey],
    borderRadius: '6px',
    marginBottom: '15px'
  },
  dropdownText: {
    fontSize: theme.fontSizes.default,
    fontWeight: '400',
    color: theme.colors.textGrey
  },
  checkbox: {
    alignItems: 'flex-start',
    marginBottom: '15px'
  },
  checkboxText: {
    fontSize: theme.fontSizes.small,
    color: theme.colors.textGrey,
    margin: '0 15px'
  },
  button: {
    height: '52px',
    width: '396px',
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
    color: theme.colors.primaryWhite,
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)',
    marginBottom: '37.5px'
  },
  center: {
    textAlign: 'center',
    maxWidth: '460px'
  }
}));

export type CompanyDetails = {
  firstName: string;
  lastName: string;
  companyName: string;
  companyEmail: string;
  number: string;
  companyType: string;
  role: string;
  contact: boolean;
};

export type IndividualDetails = {
  firstName: string;
  lastName: string;
  email: string;
  number: string;
  role: string;
  contact: boolean;
};

type Props = {
  title: string;
  subtitle: string;
  onRegister: () => void;
  onBook: () => void;
  onIndividual: () => void;
  onCompany: () => void;
  individual?: boolean;
  company?: boolean;
  individualDetails?: IndividualDetails;
  companyDetails?: CompanyDetails;
  className?: string;
};

function DetailsCard({
  title,
  subtitle,
  onRegister,
  onBook,
  onIndividual,
  onCompany,
  individual,
  company,
  individualDetails,
  companyDetails,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Form Data
  const [firstName, setFirstName] = React.useState('');
  const [lastName, setLastName] = React.useState('');
  const [email, setEmail] = React.useState('');
  const [companyName, setCompanyName] = React.useState('');
  const [companyEmail, setCompanyEmail] = React.useState('');
  const [number, setNumber] = React.useState('');
  const [companyType, setCompanyType] = React.useState({
    id: 0,
    title: 'Type of Company'
  });
  const [role, setRole] = React.useState({
    id: 0,
    title: 'What role best describes you'
  });
  const [contact, setContact] = React.useState(false);

  React.useEffect(() => {
    if (individualDetails) {
      individualDetails = {
        firstName,
        lastName,
        email,
        number,
        role: role.title !== 'Country' ? role.title : '',
        contact
      };
    }
    if (companyDetails) {
      companyDetails = {
        firstName,
        lastName,
        companyName,
        companyEmail,
        number,
        companyType:
          companyType.title !== 'Type of Company' ? companyType.title : '',
        role: role.title !== 'What role best describes you' ? role.title : '',
        contact
      };
    }
  });

  return (
    <div className={classNames(classes.root, className)}>
      <Icon className={classes.logo} name="TTC_Logo_Icon" size={44} />
      <Title>{title}</Title>
      <Subtitle>{subtitle}</Subtitle>
      {!individual && !company && (
        <div className={classes.form}>
          <Button
            archetype="submit"
            className={classes.button}
            onClick={onIndividual}
          >
            Register as an Individual
          </Button>

          <Button
            archetype="submit"
            className={classes.button}
            onClick={onCompany}
          >
            Register as a Company
          </Button>
        </div>
      )}
      {(individual || company) && (
        <form className={classes.form}>
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

          {individual && (
            <CoreInput
              placeholder="Email"
              type="text"
              onChange={(text: string) => setEmail(text)}
              value={email}
              className={classes.input}
            />
          )}

          {company && (
            <CoreInput
              placeholder="Company Name"
              type="text"
              onChange={(text: string) => setCompanyName(text)}
              value={companyName}
              className={classes.input}
            />
          )}

          {company && (
            <CoreInput
              placeholder="Company Email"
              type="text"
              onChange={(text: string) => setCompanyEmail(text)}
              value={companyEmail}
              className={classes.input}
            />
          )}

          <CoreInput
            placeholder="Telephone Number"
            type="text"
            onChange={(text: string) => setNumber(text)}
            value={number}
            className={classes.input}
          />

          {company && (
            <Dropdown
              placeholder="Type of company"
              options={[{ id: 0, title: 'Default Option' }]}
              selected={companyType}
              setSelected={(selected: DropdownOption) =>
                setCompanyType(selected)
              }
              fontStyle={classes.dropdownText}
              boxStyle={classes.dropdownBox}
              className={classes.dropdown}
            />
          )}

          <Dropdown
            placeholder="What role best describes you"
            options={[{ id: 0, title: 'Default Option' }]}
            selected={role}
            setSelected={(selected: DropdownOption) => setRole(selected)}
            fontStyle={classes.dropdownText}
            boxStyle={classes.dropdownBox}
            className={classes.dropdown}
          />

          <CheckboxSingle
            className={classes.checkbox}
            fontStyle={classes.checkboxText}
            size={18}
            onChange={() => setContact(!contact)}
            label="By checking this box you confirm you are happy for our team to contact you during the registration period"
          />
          <Button
            archetype="submit"
            className={classes.button}
            onClick={onRegister}
          >
            Register with TTC
          </Button>
        </form>
      )}
      <Footer onBook={onBook} />
    </div>
  );
}

export default DetailsCard;
