import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import CoreInput, { InputTypes } from 'sharedComponents/core/CoreInput';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import Heading from 'components/Heading';
import QuickInfo from 'components/Overview/QuickInfo';
import TopInfo from './TopInfo';
import CardItem from 'components/CardItem';

const useStyles = createUseStyles((theme: Theme) => ({
  trainingZoneRoot: {
    position: 'relative',
    display: 'grid',
    gridTemplateColumns: 'repeat(6, 1fr)',
    gridTemplateRows: 'repeat(4, auto)',
    gridGap: theme.spacing(3),
    gridTemplateAreas: `
      "header header header topinf topinf topinf"
      "allcou allcou expire expire trainp trainp"
      "jumpba jumpba jumpba .      .      .     "
    `
  },
  trainingHeader: {
    gridArea: 'header'
  },
  topInfo: {
    gridArea: 'topinf'
  },
  card1: {
    gridArea: 'allcou'
  },
  card2: {
    gridArea: 'expire'
  },
  card3: {
    gridArea: 'trainp'
  },
  jumpHeader: {
    gridArea: 'jumpba'
  }
}));

type Props = {};

function TrainingZone({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const userName = 'James';
  return (
    <div className={classes.trainingZoneRoot}>
      <div className={classes.trainingHeader}>
        <Heading text="Training Zone" size={'large'} />
        <Heading
          text={`You're doing great ${userName}, keep up the good work so you don't loose your momentum`}
          size={'medium'}
        />
      </div>
      <TopInfo className={classes.topInfo} />
      <CardItem
        className={classes.card1}
        title={'All Courses'}
        description={
          'Find all of the online courses you have been enrolled on here'
        }
        buttonProps={{ title: 'View Courses', onClick: () => {} }}
      />
      <CardItem
        className={classes.card2}
        title={'Expiring Soon'}
        description={
          'See your previously completed courses that are due to expire'
        }
        buttonProps={{ title: 'Expiring Certificates', onClick: () => {} }}
      />
      <CardItem
        className={classes.card3}
        title={'Training Progress'}
        description={'Review your progress of all Courses in real time'}
        buttonProps={{ title: 'See Help Guides', onClick: () => {} }}
      />
      <div className={classes.jumpHeader}>
        <Heading text="Jump back in" size={'large'} />
        <Heading
          text={`${userName}, you have some unfinished courses,
          jump back in to continue learning now.`}
          size={'medium'}
        />
      </div>
    </div>
  );
}

export default TrainingZone;
