import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  thumbnail: {
    display: 'flex',
    position: 'absolute',
    zIndex: 10
  },
  playCircle: {
    width: '76px',
    height: '76px',
    borderRadius: '76px',
    backgroundColor: theme.colors.navyBlue,
    opacity: 0.8
  },
  playTriangle: {
    width: 0,
    height: 0,
    marginLeft: 6,
    borderTop: '13.5px solid transparent',
    borderBottom: '13.5px solid transparent',
    borderLeft: ['27px', 'solid', theme.colors.primaryWhite]
  }
}));

type Props = {
  source: string;
  height?: number;
  width?: number;
  thumbnail?: React.ReactNode;
  className?: string;
};

function VideoPlayer({ source, height, width, thumbnail, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const videoRef = React.useRef<HTMLVideoElement>(null);
  const [started, setStarted] = React.useState(false);

  const playVideo = () => {
    if (videoRef != null && videoRef.current) {
      setStarted(true);
      videoRef.current.play();
    }
  };

  // Will show the thumbnail rather than controls if played for less than a second
  const checkShowThumbnail = () => {
    if (
      videoRef != null &&
      videoRef.current &&
      videoRef.current.currentTime < 1
    ) {
      setStarted(false);
      videoRef.current.currentTime = 0;
    }
  };

  return (
    <div className={classNames(classes.root, className)}>
      {!started && (
        <div
          className={classNames(!thumbnail && classes.root, classes.thumbnail)}
          onClick={() => playVideo()}
          style={{ maxHeight: height, maxWidth: width }}
        >
          {thumbnail ? (
            thumbnail
          ) : (
            <div className={classNames(classes.root, classes.playCircle)}>
              <div className={classes.playTriangle} />
            </div>
          )}
        </div>
      )}
      <video
        src={source}
        style={{ height, width, backgroundSize: 'cover', overflow: 'hidden' }}
        controls={started}
        ref={videoRef}
        onPause={() => checkShowThumbnail()}
        onEnded={() => setStarted(false)}
        preload="metadata"
      >
        Your Browser Does Not Support This Video Type
      </video>
    </div>
  );
}

export default VideoPlayer;
