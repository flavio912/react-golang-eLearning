import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import VideoPlayer from './VideoPlayer';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    '@media (max-width: 1000px)': {
      flexDirection: 'column',
      alignItems: 'center'
    }
  },
  floating: {
    position: 'absolute',
    borderRadius: '5px',
    boxShadow:
      '5px 15px 43px 8px rgba(143,143,173,0.28), 10px 27px 21px 3px rgba(0,0,0,0.06)',
    overflow: 'hidden',
    '@media (max-width: 1000px)': {
      position: 'relative'
    }
  },
  bottomBox: {
    display: 'flex',
    height: '334px',
    width: '100%',
    margin: '52px 0 0 316px',
    borderRadius: '6px',
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: '0 3px 15px 1px rgba(0,0,0,0.12)',
    '@media (max-width: 1000px)': {
      width: '100%',
      margin: '0'
    }
  },
  quoteContainer: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    margin: '0 68px 0 272px',
    '@media (max-width: 1000px)': {
      margin: '10%'
    }
  },
  quote: {
    fontSize: theme.fontSizes.smallHeading,
    color: theme.colors.textGrey,
    fontWeight: '300',
    marginBottom: '40px'
  },
  authorName: {
    fontSize: theme.fontSizes.tinyHeading,
    fontWeight: 'bold'
  },
  authorTitle: {
    fontSize: theme.fontSizes.extraLarge,
    color: theme.colors.textGrey
  },
  logo: {
    marginLeft: '34px'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between'
  }
}));

export type Author = {
  name: string;
  title: string;
  quote: string;
};

type Props = {
  source: string;
  height?: number;
  width?: number;
  thumbnail?: React.ReactNode;
  author?: Author;
  className?: string;
};

function FloatingVideo({
  source,
  height,
  width,
  author,
  thumbnail,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.root}>
      <div className={classes.floating}>
        <VideoPlayer
          source={source}
          height={height}
          width={width}
          thumbnail={thumbnail}
        />
      </div>
      {author && (
        <div className={classes.bottomBox}>
          <div className={classes.quoteContainer}>
            <div className={classes.quote}>"{author.quote}"</div>
            <div className={classes.row}>
              <div>
                <div className={classes.authorName}>{author.name}</div>
                <div className={classes.authorTitle}>{author.title}</div>
              </div>
              <Icon
                name="TTC_Logo_Icon"
                className={classes.logo}
                style={{ height: '39px', width: '127px' }}
              />
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default FloatingVideo;
