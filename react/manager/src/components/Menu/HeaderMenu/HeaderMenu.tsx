import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from '../../core/Icon/Icon';
import ProfileIcon from 'components/core/ProfileIcon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    borderBottom: [1, 'solid', theme.colors.borderGrey],
    top: 0,
    width: '100%',
    position: 'fixed'
  },
  menu: {
    backgroundColor: theme.colors.primaryWhite,
    padding: '17px 25px'
  },
  logo: {
    cursor: 'pointer',
    height: '50px',
    width: '140px',
    borderRadius: theme.primaryBorderRadius,
    border: [1, 'solid', theme.colors.borderGrey],
    boxShadow: theme.shadows.primary,
    objectFit: 'cover'
  },
  profileCircle: {
    height: 47,
    width: 47,
    borderRadius: 47,
    backgroundColor: theme.colors.primaryWhite
  },
  profileBorder: {
    height: 49,
    width: 49,
    borderRadius: 49,
    backgroundImage: `linear-gradient(45deg,
      ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'
  },
  center: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  body: {
    backgroundColor: theme.colors.backgroundGrey,
    flexGrow: 1
  }
}));

export interface User {
  name: string;
  url: string;
}

type Props = {
  logo: string;
  user: User;
  children?: React.ReactNode;
  onLogoClick?: Function;
  onProfileClick?: Function;
  className?: string;
};

function HeaderMenu({
  logo,
  user,
  children,
  onLogoClick,
  onProfileClick,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classNames(classes.row, classes.menu)}>
        <Icon name="TTC_Logo_Icon" size={55} />
        <img
          className={classNames(classes.logo)}
          onClick={() => onLogoClick && onLogoClick()}
          src={logo}
          alt="Logo"
        />
        <div className={classNames(classes.row)}>
          <div className={classNames(classes.center, classes.profileBorder)}>
            <div className={classNames(classes.center, classes.profileCircle)}>
              <ProfileIcon
                name={user?.name}
                url={user?.url}
                size={45}
                fontSize={18}
              />
            </div>
          </div>
          <Icon
            name="Card_SecondaryActon_Dots"
            size={20}
            style={{ cursor: 'pointer', margin: '0 20px' }}
            onClick={() => onProfileClick && onProfileClick()}
          />
        </div>
      </div>
      <div className={classes.body}>{children}</div>
    </div>
  );
}

export default HeaderMenu;
