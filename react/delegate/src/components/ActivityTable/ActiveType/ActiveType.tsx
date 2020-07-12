import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Text from 'sharedComponents/core/Table/Text/Text';
import themeConfig, { Theme } from 'helpers/theme';
import Icon, { IconNames } from 'sharedComponents/core/Icon';
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    alignItems: 'center',
    '& div': {
      marginLeft: 10
    }
  }
}));

type Props = {
  icon: IconNames;
  text: string;
};
const colorActiveTypes = {
  CourseCertificates: themeConfig.colors.secondaryGreen,
  CourseFailed: themeConfig.colors.secondaryDanger,
  CourseNewCourse: themeConfig.colors.primaryBlack
};
function ActiveType({ icon, text }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <span className={classes.root}>
      <Icon name={icon} size={18} />
      <Text text={text} color={colorActiveTypes[icon]} />
    </span>
  );
}

export default ActiveType;
