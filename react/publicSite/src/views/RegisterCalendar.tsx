import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import RegisterComp from 'components/Overview/Registration/RegisterCompany';
import { Router } from 'found';
import InfoCard from 'components/Overview/Registration/InfoCard';
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
    background: 'white'
  }
}));

type Props = {
  router: Router;
};

function RegisterCalendar({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <InfoCard
            title="New to TTC Hub?"
            subtitle="Book a 30 minute demo and get your questions answered with one of our customer champions"
            imageURL={require('assets/newToTTC.svg')}
            imageTitle="How does it work?"
            imageSubtitle="Simply select a date and time thats suitable for you and we’ll send you through the video call details to confirm."
        />
      </div>
      <div className={classes.fancyBackground}></div>
    </div>
  );
}

export default RegisterCalendar;
