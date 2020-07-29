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
    alignItems: 'center',
  },
  button: {
    width: '100%',
    height: '52px',
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
    color: theme.colors.primaryWhite,
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)',
  },
  form: {
    display: 'grid',
    width: '100%',
    gridGap: theme.spacing(2),
    gridTemplateColumns: '1fr 1fr',
  },
  fullWidth: {
    gridColumn: '1 / 3',
  },
  input: {
    fontSize: theme.fontSizes.default,
    padding: '14px',
    border: ['1px', 'solid', theme.colors.borderGrey],
    borderRadius: theme.buttonBorderRadius,
    '&::placeholder': {
      color: theme.colors.secondaryBlack,
    },
  },
  dropdownText: {
    fontSize: theme.fontSizes.default,
    fontWeight: '400',
    color: theme.colors.secondaryBlack,
  },
  dropdown: {
    flex: 1,
  },
  checkboxText: {
    color: theme.colors.textGrey,
  },
}));

const initState: State = {
  fname: '',
  lname: '',
  email: '',
  password: '',
  telephone: '',
  companyName: '',
};

type StateUpdate = {
  fname?: string;
  lname?: string;
  email?: string;
  password?: string;
  telephone?: string;
  companyName?: string;
};

type State = {
  fname: string;
  lname: string;
  email: string;
  password: string;
  telephone: string;
  companyName: string;
};

type Props = {
  onLogoClick?: () => void;
  onChange: (value: State) => void;
  onNext: () => void;
};

function RegisterCompany({ onLogoClick, onChange, onNext }: Props) {
  const classes = useStyles();

  const [inputData, setInput] = React.useState(initState);

  const onUpdate = (updates: StateUpdate) => {
    const newState = { ...inputData, ...updates };
    setInput(newState);
    onChange(newState);
  };

  let isComplete = true;

  for (const key in inputData) {
    if (!inputData[key]) {
      isComplete = false;
      break;
    }
  }

  return (
    <div className={classes.registerCompanyRoot}>
      <Icon name="TTC_Logo_Icon" size={44} onClick={onLogoClick} />
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
          onChange={(val) => onUpdate({ fname: val })}
          value={'fname'}
          className={classes.input}
        />
        <CoreInput
          placeholder="Last Name"
          type="text"
          onChange={(val) => onUpdate({ lname: val })}
          value={'lname'}
          className={classes.input}
        />
        <CoreInput
          placeholder="Company Name"
          type="text"
          onChange={(val) => onUpdate({ companyName: val })}
          value={'company'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Company Email"
          type="email"
          onChange={(val) => onUpdate({ email: val })}
          value={'email'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Password"
          type="password"
          onChange={(val) => onUpdate({ password: val })}
          value={'password'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Telephone Number"
          type="tel"
          onChange={(val) => onUpdate({ telephone: val })}
          value={'telephone'}
          className={classnames(classes.input, classes.fullWidth)}
        />
      </div>
      <Spacer vertical spacing={3} />
      <Button
        archetype="submit"
        className={classes.button}
        disabled={!isComplete}
        onClick={onNext}
      >
        Next
      </Button>
      <Spacer vertical spacing={3} />
    </div>
  );
}

export default RegisterCompany;
