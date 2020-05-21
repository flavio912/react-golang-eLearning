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
    backgroundImage: `linear-gradient(45deg,
        ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`,
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
  userMode?: string; 
  colour?: string;
  borderSize?: number;
  className?: string;
};

function CircleBorder({ user, text, size = 44, fontSize = 18, userMode = 'manager', borderSize = 6, colour, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const profileBorderSize = (userMode.toLowerCase() === "manager" ? size + borderSize + 2 : size + borderSize);
  const pofileCircleSize = (userMode.toLowerCase() === "manager" ? size + 2 : size);

  let profileBackImg = {};
  if (colour && userMode.toLowerCase() !== "manager")
    profileBackImg['backgroundImage'] = `linear-gradient(45deg, ${colour}, ${colour})`;

  return (
    <div
      className={classNames(classes.center, classes.profileBorder, className)}
      style={{ 
        height: profileBorderSize, width: profileBorderSize, 
        borderRadius: profileBorderSize, 
        ...profileBackImg
      }}
    >
        <div
          className={classNames(classes.center, classes.profileCircle)}
          style={{ width: pofileCircleSize, height: pofileCircleSize, borderRadius: pofileCircleSize }}
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