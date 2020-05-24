import * as React from 'react';
import classNames from 'classnames';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import ProfileIcon from 'sharedComponents/core/ProfileIcon';

const useStyles = createUseStyles((theme: Theme) => ({
  center: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    fontWeight: '900',
  },
  profileCircle: {
    backgroundColor: theme.colors.primaryWhite,
  },
  profileBorder: {
    backgroundColor: `${theme.colors.primaryWhite}`,
    boxShadow: `rgba(0, 0, 0, 0.24) 0px 0px 3px`,
    backgroundImage: ({ type }: { type: BorderType }) =>
      type == 'fancy'
        ? `linear-gradient(45deg, ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`
        : '',
  },
}));

export interface User {
  name?: string;
  url?: string;
}

export type BorderType = 'fancy' | 'plain' | undefined;

type Props = {
  user?: User;
  text?: string | number;
  size?: number;
  fontSize?: number;
  borderType?: BorderType;
  className?: string;
};

function CircleBorder({
  user,
  text,
  size = 44,
  fontSize = 18,
  borderType = 'fancy',
  className,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ type: borderType, theme });
  const borderTypes = {
    fancy: {
      borderSize: size + 4,
      circleSize: size + 2,
    },
    plain: {
      borderSize: size + 6,
      circleSize: size,
    },
  };

  return (
    <div
      className={classNames(classes.center, classes.profileBorder, className)}
      style={{
        height: borderTypes[borderType].borderSize,
        width: borderTypes[borderType].borderSize,
        borderRadius: borderTypes[borderType].borderSize,
      }}
    >
      <div
        className={classNames(classes.center, classes.profileCircle)}
        style={{
          width: borderTypes[borderType].circleSize,
          height: borderTypes[borderType].circleSize,
          borderRadius: borderTypes[borderType].circleSize,
        }}
      >
        {user && (
          <ProfileIcon
            name={user?.name}
            url={user?.url}
            size={size}
            fontSize={fontSize}
          />
        )}
        {text && (
          <div className={classes.center} style={{ fontSize }}>
            {text}
          </div>
        )}
      </div>
    </div>
  );
}

export default CircleBorder;
