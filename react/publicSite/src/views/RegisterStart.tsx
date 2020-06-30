import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import AccountTypeCard from 'components/Overview/Registration/AccountTypeCard';
import { Router } from 'found';

const useStyles = createUseStyles((theme: Theme) => ({
  registerRoot: {
    display: 'grid',
    minHeight: '100%',
    background: 'white',
    gridTemplateColumns: '600px 1fr',
    '@media (max-width: 800px)': {
      gridTemplateColumns: '1fr'
    }
  },
  fancyBackground: {
    background: theme.loginBackgroundGradient
  },
  picker: {
    padding: 48
  }
}));

type Props = {
  router: Router;
};

function RegisterStart({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <AccountTypeCard
          onIndividual={() => {
            router.push('/register/individual');
          }}
          onCompany={() => {
            router.push('/register/company');
          }}
          onLogoClick={() => {
            router.push('/');
          }}
        />
      </div>
      <div className={classes.fancyBackground}></div>
    </div>
  );
}

export default RegisterStart;
