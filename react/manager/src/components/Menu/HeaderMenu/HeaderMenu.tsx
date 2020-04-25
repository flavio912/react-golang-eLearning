import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from '../../core/Icon/Icon';
import CircleBorder, { User } from 'components/core/CircleBorder';

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
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'
  },
  body: {
    backgroundColor: theme.colors.backgroundGrey,
    flexGrow: 1
  }
}));

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
          <CircleBorder user={user} />
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
