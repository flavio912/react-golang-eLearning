import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from '../../../sharedComponents/core/Icon/Icon';
import CircleBorder, { User } from 'sharedComponents/core/CircleBorder';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import CoreInput from 'sharedComponents/core/Input/CoreInput';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    borderBottom: [1, 'solid', theme.colors.borderGrey],
    gridArea: '1 / 2',
    zIndex: 10
  },
  menu: {
    display: 'flex',
    flexDirection: 'row',
    justfiyContent: 'flex-start',
    backgroundImage: `linear-gradient(90deg,
      ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`,
    padding: '17px 25px'
  },
  search: {
    display: 'flex',
    alignItems: 'center',
    flex: 1,
    padding: theme.spacing(1)
  },
  searchInput: {
    color: theme.colors.primaryWhite,
    fontSize: theme.fontSizes.large,
    fontWeight: 300,
    paddingLeft: theme.spacing(1),
    backgroundColor: 'transparent',
    '&::placeholder': {
      color: theme.colors.primaryWhite
    }
  },
  profile: {
    display: 'flex',
    alignSelf: 'flex-end'
  },
  name: {
    fontSize: theme.fontSizes.tinyHeading,
    fontWeight: 800,
    color: theme.colors.primaryWhite
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
  },
  hidden: {
    visibility: 'hidden'
  }
}));

type Props = {
  user: User;
  onProfileClick?: Function;
  className?: string;
  onToggleSearchModal?: Function;
  showSearch?: boolean;
};

function HeaderMenu({
  user,
  onProfileClick,
  onToggleSearchModal,
  className,
  showSearch = false
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classNames(classes.row, classes.menu)}>
        <div
          className={classNames(
            classes.search,
            showSearch ? '' : classes.hidden
          )}
          onClick={() => onToggleSearchModal && onToggleSearchModal()}
        >
          <Icon name="SearchGlass" size={15} />
          <CoreInput
            type="search"
            placeholder="Search"
            disabled={true}
            className={classes.searchInput}
            onChange={(text: string) => {}}
            /*
            onFocus={onFocus}
            onBlur={onBlur}
            value={input}
            setValue={setInput}*/
          />
        </div>
        <div
          className={classNames(classes.row, classes.profile)}
          onClick={() => onProfileClick && onProfileClick()}
        >
          <div className={classes.name}>{user.name}</div>
          <Spacer horizontal spacing={2} />
          <CircleBorder user={user} borderType="plain" />
        </div>
      </div>
    </div>
  );
}

export default HeaderMenu;
