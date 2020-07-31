import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import QuickInfo from 'components/Overview/QuickInfo';
import classnames from 'classnames';
import { CourseStatus } from './__generated__/TrainingZone_user.graphql';
import moment from 'moment';

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
  courses: {
    status: CourseStatus;
    enrolledAt: string;
  }[]
};

function TopInfo({ className, courses }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const month =moment().format('MMMM');
  const year = moment().format('YYYY');

  const currentMonthCourses = courses.filter(x => moment(x.enrolledAt).isSameOrAfter(moment().format("YYYY-MM-01")));
  const newCourses = currentMonthCourses.length;
  const certificates = currentMonthCourses.filter(x => x.status !== 'failed').length;
  
  return (
    <div className={classnames(classes.outer, className)}>
      <div className={classes.cont}>
        <QuickInfo
          icon={'CourseNewCourseGreen'}
          text={'Courses'}
          value={newCourses}
          footer={`in ${month} ${year}`}
          valueArrow={'up'}
        />
        <div className={classes.divider} />
      </div>
      <div className={classes.cont}>
        <QuickInfo
          icon={'CourseCertificates'}
          text={'Certificates'}
          value={certificates}
          footer={`in ${month} ${year}`}
          valueArrow={'up'}
        />
        {/* <div className={classes.divider} /> */}
      </div>
      {/* <QuickInfo
        icon={'CourseTimeTrackedGreen'}
        text={'Time Tracked'}
        value={{ h: 14, m: 2 }}
        footer={'in March 2020'}
        valueArrow={'up'}
      /> */}
    </div>
  );
}
export default TopInfo;
