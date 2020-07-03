import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import Homepage from 'components/Overview/Homepage';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import PageHeader from 'components/core/PageHeader';

const useStyles = createUseStyles((theme: Theme) => ({
  aboutRoot: {
    width: '100%'
  },
  blueLogo: {
    width: 70
  },
  centeredLogoItems: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    maxWidth: 500,
    textAlign: 'center'
  },
  logoText: {
    fontSize: theme.fontSizes.heading
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  }
}));

type Props = {};

function AboutUs({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.aboutRoot}>
      <PageHeader
        title="About Us"
        description="Our mission is to create the highest quality safety &amp; compliance training in the world"
      />
      <Spacer spacing={4} vertical />
      <div className={classes.centerer}>
        <div className={classes.centeredLogoItems}>
          <img
            src={require('assets/logo/blueCircleLogo.svg')}
            className={classes.blueLogo}
          />
          <h2 className={classes.logoText}>
            TTC was created with a passion for excellence and a sense of
            superior quality
          </h2>
        </div>
      </div>
    </div>
  );
}

export default AboutUs;
