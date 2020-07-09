import React, { ReactElement } from 'react';
import { ExpandItemType } from './ExpandableListView';
import { createUseStyles, useTheme } from 'react-jss';
import themeRoot, { Theme } from 'helpers/theme';
import classNames from 'classnames';
import Icon from 'sharedComponents/core/Icon';
const useStyles = createUseStyles((theme: Theme) => ({
  rootExpandItemView: {
    padding: [19.5, 28],
    marginBottom: 20,
    backgroundColor: theme.colors.primaryWhite,
    border: `1px solid ${theme.colors.approxZircon}`,
    boxShadow:
      '0 2px 10px 0 rgba(0,0,0,0.15), 4px 2px 10px -2px rgba(0,0,0,0.06)',
  },
  titleWrapper: {
    flexDirection: 'row',
    display: 'flex',
    alignItems: 'center',
    cursor: 'pointer',
    paddingLeft: 5.5,
  },
  title: {
    paddingLeft: 33,
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.heading,
    fontWeight: 800,
    letterSpacing: -0.63,
    lineHeight: `60.02px`,
    margin: 0,
  },
  description: {
    padding: [35, 0],
    margin: 0,
    color: theme.colors.textGrey,
    fontSize: theme.fontSizes.heading,
    lineHeight: `42px`,
    letterSpacing: -0.63,
  },
  iconExpanded: {
    transform: `rotate(90deg)`,
  },
}));

type Props = {
  item: ExpandItemType;
  onClickItem: () => void;
};
export const ExpandItemView = ({ item, onClickItem }: Props): ReactElement => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.rootExpandItemView}>
      <div onClick={onClickItem} className={classes.titleWrapper}>
        <Icon
          name={'Arrow_Right_Blue'}
          size={25}
          pointer
          color={themeRoot.colors.navyBlue2}
          className={classNames({
            [classes.iconExpanded]: item.isExpanded,
          })}
        />
        <p className={classes.title}>{item.title}</p>
      </div>

      {item.isExpanded && (
        <div
          className={classes.description}
          dangerouslySetInnerHTML={{ __html: item.description }}
        ></div>
      )}
    </div>
  );
};
