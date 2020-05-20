import * as React from "react";
import classNames from 'classnames';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from "helpers/theme";
import ProfileIcon from 'sharedComponents/core/ProfileIcon';

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
    backgroundColor: `${theme.colors.primaryWhite}`,
    boxShadow: theme.shadows.primary
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
    colour?: string;
    className?: string;
};

function CircleBorder({ user, text, size = 45, fontSize = 18, colour, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div
        className={classNames(classes.center, classes.profileBorder, className)}
        style={{ height: size + 8, width: size + 8, borderRadius: size + 8, backgroundImage: colour && `linear-gradient(45deg, ${colour}, ${colour})` }}
    >
        <div
            className={classNames(classes.center, classes.profileCircle)}
            style={{ height: size, width: size, borderRadius: size }}
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