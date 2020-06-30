import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';
import Button from 'sharedComponents/core/Input/Button';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import Icon from 'sharedComponents/core/Icon';
import Title from './Title';
import Subtitle from './Subtitle';
import Footer from './Footer';

const useStyles = createUseStyles((theme: Theme) => ({
  accountTypeRoot: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'space-between',
    minHeight: '100%'
  },
  button: {
    width: '100%',
    height: '52px',
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
    color: theme.colors.primaryWhite,
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)'
  }
}));

type Props = {
  onIndividual: () => void;
  onCompany: () => void;
  onLogoClick?: () => void;
};

function AccountTypeCard({ onIndividual, onCompany, onLogoClick }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.accountTypeRoot}>
      <Icon name="TTC_Logo_Icon" size={44} onClick={onLogoClick} />

      <div>
        <Spacer vertical spacing={4} />
        <Title>Register with TTC today</Title>
        <Subtitle>
          If you’re looking to get access to the finest level of compliance
          training register with TTC Hub below
        </Subtitle>
      </div>
      <div style={{ width: '100%' }}>
        <Button
          archetype="submit"
          className={classes.button}
          onClick={onIndividual}
        >
          Register as an Individual
        </Button>
        <Spacer vertical spacing={3} />
        <Button
          archetype="submit"
          className={classes.button}
          onClick={onCompany}
        >
          Register as a Company
        </Button>
        <Spacer vertical spacing={3} />
      </div>
      <Footer onBook={() => {}} />
    </div>
  );
}

export default AccountTypeCard;
