import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import PageHeader from 'components/core/PageHeader';
import Quote from 'components/Overview/Article/Quote';
import Card from 'sharedComponents/core/Cards/Card';
import PageMargin from 'components/core/PageMargin';

const useStyles = createUseStyles((theme: Theme) => ({
  aboutRoot: {
    width: '100%'
  },
  textCenter: {
    maxWidth: '821px',
  },
  row: {
    display: 'flex',
    flexDirection: 'row'
  },
  blueLogo: {
    width: 70
  },
  centeredLogoItems: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    textAlign: 'center',
    marginBottom: '50px'
  },
  logoText: {
    fontSize: theme.fontSizes.heading,
    maxWidth: 500,
  },
  subheading: {
    color: '#34373A',
    fontSize: theme.fontSizes.heading,
    fontWeight: 600,
    marginBottom: '50px'
  },
  heading: {
    fontSize: theme.fontSizes.heading,
    fontWeight: 600,
    marginBottom: '25px'
  },
  text: {
    color: theme.colors.textGrey,
    fontSize: theme.fontSizes.tinyHeading,
  },
  link: {
    cursor: 'pointer',
    color: theme.colors.textBlue
  },
  image: {
    height: 'auto',
    maxWidth: '100%',
    maxHeight: '879px',
    margin: '90px 0 100px 0'
  },
  video: {
    position: 'relative',
    width: '100vw',
    overflow: 'hidden',
    margin: '100px 0px',
    '@media (max-width: 750px)': {
      width: '95vw',
      boxShadow: 'none',
      margin: '5% 0',
    }
  },
  quote: {
    margin: '0 0 70px 89px',
    flex: 1
  },
  card: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '100px 0',
    textAlign: 'center'
  },
  cardText: {
    maxWidth: '575px'
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
      <PageMargin>
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
          <div className={classes.textCenter}>
            <div className={classes.subheading}>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipiscing. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore </div>
            <div className={classes.heading}>A new league of compliance</div>
            <div className={classes.text}>Enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in <span className={classes.link}>reprehenderit</span> in voluptate. <br/><br/> Velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est.</div>
          </div>
          <img className={classes.image} src={require('assets/AboutUsStockImage.png')}/>
          <div className={classes.textCenter}>
            <div className={classes.heading}>Your success means everything</div>
            <div className={classes.text}>Enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea <span className={classes.link}>commodo consequat</span>. Duis aute irure dolor in reprehenderit in voluptate.<br/><br/>Velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. <br/><br/>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat <span className={classes.link}>non proident</span>, sunt in culpa qui officia deserunt mollit anim id est laborum </div>
            <FloatingVideo
              className={classes.video}
              height={462}
              width={821}
              source={require('assets/Stock_Video.mp4')}
            />
            <div className={classes.heading}>We’re building the future of online training</div>
            <div className={classes.row}>
                <div
                  className={classes.text}
                  style={{ flex: 2 }}
                >
                  Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ipsum est lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
                </div>
                <Quote
                  className={classes.quote}
                  borderSide="top"
                  quote="“Whilst growth is key, development and compliance within your team are paramount.”"
                />
            </div>
            <div className={classes.text}>Minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat.</div>
            <Card className={classes.card} padding="large">
              <div className={classes.heading}>We’re hiring.</div>
              <div className={classnames(classes.text, classes.cardText)}>We’re always looking for people to raise the bar and help us deliver an even better service to customers. <span className={classes.link}>Join us!</span></div>
            </Card>
          </div>
      </PageMargin>
    </div>
  );
}

export default AboutUs;
