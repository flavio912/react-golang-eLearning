import * as React from "react";
import classNames from 'classnames';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from "helpers/theme";
import ProfileIcon from 'components/core/ProfileIcon';

const useStyles = createUseStyles((theme: Theme) => ({
  center: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    fontWeight: '900'
  },
  profileCircle: {
    backgroundColor: theme.colors.primaryWhite
  },
  profileBorder: {
    backgroundImage: `linear-gradient(45deg,
        ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`
  },
}));

export interface User {
    name?: string;
    url?: string;
  }

type Props = {
    user?: User;
    text?: string | number;
    size?: number;
    fontSize?: number;
    className?: string;
};

function CircleBorder({ user, text, size = 45, fontSize = 18, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div
        className={classNames(classes.center, classes.profileBorder, className)}
        style={{ height: size + 4, width: size + 4, borderRadius: size + 4 }}
    >
        <div
            className={classNames(classes.center, classes.profileCircle)}
            style={{ height: size + 2, width: size + 2, borderRadius: size + 2 }}
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
                <div
                    className={classes.center}
                    style={{ fontSize }}
                >
                {text}
            </div>
            )}
        </div>
    </div>
  );
}

export default CircleBorder;