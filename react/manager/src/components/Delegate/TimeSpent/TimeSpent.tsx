import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Text from 'sharedComponents/core/Table/Text/Text';
import themeConfig, { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import { IconNames } from 'sharedComponents/core/Icon/Icon';
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    '& div': {
      marginLeft: 7
    }
  }
}));

type Props = {
  timeSpent:
    | {
        h: string | number;
        m: string | number;
      }
    | string;
};

function TimeSpent({ timeSpent }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  let text;
  let timeTrackedIcon = 'CourseTimeTrackedGreen';
  if (typeof timeSpent === 'string') {
    text = timeSpent;
    timeTrackedIcon = 'CourseTimeTrackedGrey';
  } else {
    text = `${timeSpent.h}hr ${timeSpent.m}mins`;
  }
  return (
    <span className={classes.root}>
      <Icon name={timeTrackedIcon as IconNames} size={16} />
      <Text text={text} color={themeConfig.colors.secondaryBlack} />
    </span>
  );
}

export default TimeSpent;
