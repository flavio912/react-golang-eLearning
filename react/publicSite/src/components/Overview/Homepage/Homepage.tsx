import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Curve from 'components/core/Curve';
import TrustedCard from 'components/core/Cards/TrustedCard';
import Button from 'sharedComponents/core/Input/Button';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    background: theme.colors.lightBlue,
    paddingTop: 45
  },
  column: {
    display: 'flex',
    flexDirection: 'column'
  },
  centerColumn: {
    display: 'flex',
    alignItems: 'center',
    width: theme.centerColumnWidth
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
    '@media (max-width: 1200px)': {
      justifyContent: 'center',
      margin: '0 5% 0 5%'
    }
  },
  buttonRow: {
    display: 'flex',
    flexDirection: 'row',
    '@media (max-width: 600px)': {
      flexDirection: 'column'
    }
  },
  title: {
    fontSize: 60,
    fontWeight: '900',
    marginTop: '5%',
    maxWidth: '800px',
    color: theme.colors.primaryBlack
  },
  blueText: {
    color: theme.colors.navyBlue
  },
  description: {
    fontSize: theme.fontSizes.heading,
    fontWeight: '300',
    margin: '5% 0 5% 0',
    maxWidth: '750px'
  },
  button: {
    height: '52px',
    width: '206px',
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: '800',
    marginBottom: '5%',
    '@media (max-width: 600px)': {
      width: '100%'
    }
  },
  boxShadow: {
    marginLeft: '25px',
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)',
    '@media (max-width: 600px)': {
      marginLeft: '0'
    }
  },
  nameBox: {
    '@media (max-width: 1200px)': {
      opacity: 0,
      flex: 0,
      width: 0
    }
  },
  name: {
    fontSize: theme.fontSizes.large,
    fontWeight: 'bold'
  },
  jobTitle: {
    fontSize: theme.fontSizes.xSmall,
    color: theme.colors.textGrey,
    marginTop: '2px'
  }
}));

type Props = {
  imageURL: string;
  onView: () => void;
  onDemo: () => void;
  className?: string;
};

function Homepage({ imageURL, onView, onDemo, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div
      className={classNames(classes.root, className)}
      style={{ backgroundImage: `url(${imageURL})` }}
    >
      <div className={classes.row}>
        <div className={classes.centerColumn}>
          <div className={classes.column}>
            <div className={classes.title}>
              Redefining <span className={classes.blueText}>the future</span> of
              compliance training
            </div>
            <div className={classes.description}>
              Offering a full range of dangerous goods training and consultancy
              led by an expert team recognised in the UK and internationally
            </div>
            <div className={classes.buttonRow}>
              <Button
                onClick={onView}
                className={classes.button}
                archetype="submit"
              >
                View Courses
              </Button>
              <Button
                onClick={onDemo}
                className={classNames(classes.button, classes.boxShadow)}
              >
                Request Demo
              </Button>
            </div>
          </div>
          {/* <div className={classNames(classes.column, classes.nameBox)}>
            <div className={classes.name}>Jacob</div>
            <div className={classes.jobTitle}>Air Traffic Controller</div>
          </div> */}
        </div>
      </div>
      <Curve height={140} />
      <TrustedCard
        text="Trusted by more than 1,000 businesses in 120 countries."
        noShadow
      />
    </div>
  );
}

export default Homepage;
