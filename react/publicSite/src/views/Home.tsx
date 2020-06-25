import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import Homepage from 'components/Overview/Homepage';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';

const useStyles = createUseStyles((theme: Theme) => ({
  homeRoot: {
    width: '100%'
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  },
  centered: {
    width: theme.centerColumnWidth
  },
  whiteSpacer: {
    background: theme.colors.primaryWhite,
    padding: '60px 0px 100px 60px'
  },
  heading: {
    fontSize: 32,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    padding: '60px 0px',
    textAlign: 'center'
  },
  explore: {
    paddingTop: 70
  }
}));

type Props = {};

function Home({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.homeRoot}>
      <Homepage imageURL="" onView={() => {}} onDemo={() => {}} />
      <div className={classnames(classes.centerer, classes.whiteSpacer)}>
        <div className={classes.centered}>
          <FloatingVideo
            width={528}
            source={require('assets/Stock_Video.mp4')}
            author={{
              name: 'Kristian Durhuus',
              title: 'Chief Executive Officer',
              quote:
                'TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore .'
            }}
          />
        </div>
      </div>
      <div className={classes.explore}>
        <div className={classes.heading}>Explore our popular courses</div>
      </div>
    </div>
  );
}

export default Home;
