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
  backButton: {
    width: '100%',
    height: '52px',
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
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

const initState = {
  address1: '',
  address2: '',
  county: '',
  postcode: '',
  country: '',
};

type StateUpdate = {
  address1?: string;
  address2?: string;
  county?: string;
  postcode?: string;
  country?: string;
};

type Props = {
  onLogoClick?: () => void;
  onChange: (value: {
    address1: string;
    address2: string;
    county: string;
    postcode: string;
    country: string;
  }) => void;
  onNext: () => void;
  onBack: () => void;
};

function RegisterCompanyPart2({
  onLogoClick,
  onChange,
  onNext,
  onBack,
}: Props) {
  const classes = useStyles();

  const [inputData, setInput] = React.useState(initState);
  const [termsAgreed, setTermsAgreed] = React.useState(false);

  const onUpdate = (updates: StateUpdate) => {
    const newState = { ...inputData, ...updates };
    setInput(newState);
    onChange(newState);
  };

  let isComplete = true;

  for (const key in inputData) {
    if (!termsAgreed) {
      isComplete = false;
      break;
    }
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
          placeholder="Address Line 1"
          type="text"
          onChange={(val) => onUpdate({ address1: val })}
          value={'address1'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Address Line 2"
          type="text"
          onChange={(val) => onUpdate({ address2: val })}
          value={'address2'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="County"
          type="text"
          onChange={(val) => onUpdate({ county: val })}
          value={'county'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Post Code"
          type="text"
          onChange={(val) => onUpdate({ postcode: val })}
          value={'postcode'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CoreInput
          placeholder="Country"
          type="text"
          onChange={(val) => onUpdate({ country: val })}
          value={'country'}
          className={classnames(classes.input, classes.fullWidth)}
        />
        <CheckboxSingle
          size={18}
          onChange={(termsAgreed) => {
            setTermsAgreed(termsAgreed);
          }}
          fontStyle={classes.checkboxText}
          className={classes.fullWidth}
          label="By checking this box you confirm you are happy for our team to contact you during the registration period"
        />
      </div>
      <Spacer vertical spacing={3} />
      <Button
        archetype="submit"
        className={classes.button}
        onClick={onNext}
        disabled={!isComplete}
      >
        Register with TTC
      </Button>
      <Spacer vertical spacing={2} />
      <Button
        archetype="default"
        className={classes.backButton}
        onClick={onBack}
      >
        Back
      </Button>
      <Spacer vertical spacing={3} />
    </div>
  );
}

export default RegisterCompanyPart2;
