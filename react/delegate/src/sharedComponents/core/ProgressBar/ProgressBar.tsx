import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  progressBarRoot: {
    display: 'flex',
    alignItems: 'center',
    position: 'relative',
    overflow: 'hidden'
  },
  progressSlider: {
    position: 'absolute',
    left: 0,
    right: 0,
    height: '100%',
    overflow: 'hidden',
    animationFillMode: 'forwards',
    animation: '$slideup 1.2s ease'
  },
  '@keyframes slideup': {
    from: {
      width: 0
    },
    to: {
      width: '100%'
    }
  }
}));

type Props = {
  percent: number;
  width?: number;
  height?: number;
};

function ProgressBar({ percent, width = 200, height = 7 }: Props) {
  const theme: any = useTheme();
  const classes = useStyles({ percent, theme });

  return (
    <div
      className={classes.progressBarRoot}
      style={{
        width: width,
        height: height,
        backgroundColor: '#d3d6db',
        borderRadius: `${height / 2}px`
      }}
    >
      <div
        className={classes.progressSlider}
        style={{
          maxWidth: `${percent}%`,
          borderRadius: `${height / 2}px`,
          backgroundImage: `linear-gradient(to right, ${theme.colors.primaryBlue} 0px, ${theme.colors.primaryGreen} ${width}px)`
        }}
      />
    </div>
  );
}

export default ProgressBar;
