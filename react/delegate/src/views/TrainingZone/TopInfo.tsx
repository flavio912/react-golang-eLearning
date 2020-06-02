import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import QuickInfo from 'components/Overview/QuickInfo';
import classnames from 'classnames';

const useStyles = createUseStyles((theme: Theme) => ({
  divider: {
    height: 110,
    width: 1,
    position: 'relative',
    background: theme.colors.borderGrey,
    margin: '0px 56px'
  },
  cont: {
    display: 'flex',
    justifyContent: 'space-around',
    alignItems: 'center'
  },
  outer: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  }
}));

type Props = {
  className?: string;
};

function TopInfo({ className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classnames(classes.outer, className)}>
      <div className={classes.cont}>
        <QuickInfo
          icon={'CourseNewCourseGreen'}
          text={'Courses'}
          value={10}
          footer={'in March 2020'}
          valueArrow={'up'}
        />
        <div className={classes.divider} />
      </div>
      <div className={classes.cont}>
        <QuickInfo
          icon={'CourseCertificates'}
          text={'Certificates'}
          value={7}
          footer={'in March 2020'}
          valueArrow={'up'}
        />
        <div className={classes.divider} />
      </div>
      <QuickInfo
        icon={'CourseTimeTrackedGreen'}
        text={'Time Tracked'}
        value={{ h: 14, m: 2 }}
        footer={'in March 2020'}
        valueArrow={'up'}
      />
    </div>
  );
}
export default TopInfo;
