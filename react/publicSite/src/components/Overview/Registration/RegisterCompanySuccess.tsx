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
  checkboxText: {
    color: theme.colors.textGrey,
  },
}));

type Props = {
  onLogoClick?: () => void;
  onNext: () => void;
};

function RegisterCompanySuccess({ onLogoClick, onNext }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.registerCompanyRoot}>
      <Icon name="TTC_Logo_Icon" size={44} onClick={onLogoClick} />
      <Spacer vertical spacing={4} />
      <Title>Thanks for registering</Title>
      <Subtitle>
        Thank you for registering with TTC. We'll contact you in the next 5 - 7
        days to finalise your account
      </Subtitle>
      <Spacer vertical spacing={3} />
      <Spacer vertical spacing={3} />
      <Button archetype="submit" className={classes.button} onClick={onNext}>
        Back to TTC Hub
      </Button>
      <Spacer vertical spacing={3} />
    </div>
  );
}

export default RegisterCompanySuccess;
