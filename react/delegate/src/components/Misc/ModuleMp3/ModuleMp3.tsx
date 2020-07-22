import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import ReactPlayer, { ReactPlayerProps } from 'react-player';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  moduleMp3Root: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between'
  },
  moduleNamWrapper: {
    marginRight: theme.spacing(3)
  },
  processWrapper: {
    flex: 1
  },
  moduleName: {
    fontSize: theme.fontSizes.default,
    color: theme.colors.secondaryBlack,
    fontWeight: 500,
    letterSpacing: -0.35,
    display: 'block'
  },
  subTitle: {
    fontSize: theme.fontSizes.default,
    color: theme.colors.secondaryBlack,
    letterSpacing: -0.35,
    display: 'block'
  },
  mp3Wrapper: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between'
  },
  mp3Icon: {
    cursor: 'pointer',
    display: 'flex',
    alignItems: 'center'
  },
  mp3Trail: {
    height: 6,
    opacity: 0.3,
    borderRadius: 12.5,
    backgroundColor: theme.colors.textBlue,
    width: '100%'
  },
  mp3TrailWrapper: {
    width: '100%',
    marginLeft: 20,
    marginRight: 29,
    position: 'relative'
  },
  mp3Speaker: {
    display: 'flex',
    alignItems: 'center'
  },
  mp3Loaded: {
    backgroundColor: theme.colors.blueRibbon,
    height: 6,
    borderRadius: 12.5,
    display: 'flex',
    alignItems: 'center',
    position: 'absolute',
    top: 0,
    transition: 'width 0.5s ease-in-out',
    minWidth: 14,
    '&:after': {
      content: "''",
      display: 'inline-block',
      width: 14,
      height: 14,
      backgroundColor: theme.colors.primaryWhite,
      border: `1px solid rgba(151,151,151,0.4)`,
      boxShadow: `0 0 4px 0 rgba(0,0,0,0.17)`,
      position: 'absolute',
      borderRadius: 50,
      right: 0
    }
  },
  mp3SpeakerMuted: {
    position: 'relative',
    cursor: 'pointer',
    '&:after': {
      content: "''",
      display: 'block',
      backgroundColor: theme.colors.primaryBlack,
      width: 25,
      height: 2,
      transform: 'rotate(45deg)',
      position: 'absolute'
    }
  },
  noDisplay: {
    display: 'none'
  },
  playedText: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    fontSize: 12,
    fontWeight: 300,
    letterSpacing: -0.3,
    color: theme.colors.secondaryBlack
  }
}));

type Props = {
  className?: string;
  mp3Url: string;
  name: string;
  subTitle: string;
};
const convertSecondToHHSS = (second: number): string =>
  new Date(second * 1000).toISOString().substr(14, 5);
function ModuleMp3({ className, name, subTitle, mp3Url }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [process, setProcess] = React.useState<ReactPlayerProps>({
    loaded: 0,
    loadedSeconds: 0,
    played: 0,
    playedSeconds: 0
  });
  const [playing, setPlaying] = React.useState<boolean>(false);
  const [muted, setMuted] = React.useState<boolean>(false);
  const percentLoaded = (process.playedSeconds / process.loadedSeconds) * 100;
  return (
    <div className={classNames(classes.moduleMp3Root, className)}>
      <div className={classes.moduleNamWrapper}>
        <span className={classes.moduleName}>{name}</span>
        <span className={classes.subTitle}>- {subTitle}</span>
      </div>
      <div className={classes.processWrapper}>
        <div className={classes.playedText}>
          <span>
            {convertSecondToHHSS(process.playedSeconds)} -{' '}
            {convertSecondToHHSS(process.loadedSeconds)}
          </span>
        </div>
        <div className={classes.mp3Wrapper}>
          <div className={classes.mp3Icon} onClick={() => setPlaying(!playing)}>
            {playing ? (
              <Icon name="Mp3_Pause" pointer={true} />
            ) : (
              <Icon name="Mp3_Play" pointer={true} />
            )}
          </div>
          <div className={classes.mp3TrailWrapper}>
            <div className={classes.mp3Trail} />
            <span
              className={classes.mp3Loaded}
              style={{ width: `${percentLoaded}%` }}
            ></span>
          </div>
          <div
            className={classNames(classes.mp3Speaker, {
              [classes.mp3SpeakerMuted]: muted
            })}
            onClick={() => setMuted(!muted)}
          >
            <Icon name="Mp3_Speaker" pointer={true} />
          </div>
        </div>
      </div>
      <div className={classes.noDisplay}>
        <ReactPlayer
          url={mp3Url}
          controls={true}
          playing={playing}
          muted={muted}
          onProgress={setProcess}
          height={50}
          width={`100%`}
          onEnded={() => setPlaying(false)}
        />
      </div>
    </div>
  );
}

export default ModuleMp3;
