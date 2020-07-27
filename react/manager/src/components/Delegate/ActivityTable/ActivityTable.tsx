import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Button from 'sharedComponents/core/Input/Button';
import ActiveType from 'components/Delegate/ActiveType';
import theme, { Theme } from 'helpers/theme';
import Table from 'sharedComponents/core/Table';
import Text from 'sharedComponents/core/Table/Text/Text';
import Paginator from 'sharedComponents/Pagination/Paginator';
import CheckboxSingle from 'components/core/Input/CheckboxSingle';
import TimeSpent from 'components/Delegate/TimeSpent';
import ActivityName from 'components/Delegate/ActivityName';
import { DelegateProfilePage_activity } from 'views/__generated__/DelegateProfilePage_activity.graphql';
import moment from 'moment';

const useStyles = createUseStyles((theme: Theme) => ({
  sectionTitleWrapper: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-start',
    '& h2': {
      fontSize: 15,
      color: theme.colors.primaryBlack,
      fontWeight: 300
    },
    marginBottom: 16
  },
  courseDropdown: {
    marginLeft: 49,
    background: theme.colors.primaryWhite
  },
  activeType: {
    display: 'flex',
    alignItems: 'center',
    '& div': {
      marginLeft: 10
    }
  }
}));
const activityRow = (
  key: string | number,
  activeTime: {
    time: string;
    date: string;
  },
  title: string,
  activeType: {
    icon: any;
    text: string;
  },
  timeSpent:
    | string
    | {
        h: number;
        m: number;
      },
  userName: string,
  classes?: any
): any => ({
  key,
  cells: [
    {
      component: () => (
        <CheckboxSingle/>
      ),
    },
    {
      component: () => {
        return (
          <Text
            text={`${activeTime.time} on ${activeTime.date}`}
            color={theme.colors.secondaryBlack}
          />
        );
      }
    },
    {
      component: () => <ActivityName userName={userName} title={title} />
    },
    {
      component: () => (
        <ActiveType icon={activeType.icon} text={activeType.text} />
      )
    },
    {
      component: () => <TimeSpent timeSpent={timeSpent} />
    },
    { component: () => null },
    {
      component: () => null
    }
  ],
  onClick: () => {}
});

type PageInfo = {
  currentPage: number;
  totalPages: number;
};

type Props = {
  activity: DelegateProfilePage_activity; 
  pageInfo: PageInfo;
  userName: string;
  onUpdatePage: (page: number) => void;
};

const ActivityTable = ({activity, userName, pageInfo, onUpdatePage}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div>
      <div className={classes.sectionTitleWrapper}>
        <h2>Bruce's activity</h2>
        <div className={classes.courseDropdown}>
          <Button
            archetype={'default'}
            icon={{ right: 'FilterAdjust' }}
            children={'Filter Activities'}
          />
        </div>
      </div>
      <Table
        header={['', 'ACTIVITY TIME', 'NAME', 'ACTIVE TYPE', 'TIME SPENT', '', '']}
        rows={(activity.edges ?? []).map((activity, index) => {
          if (!activity) return;

          var nameMap = {
            completedCourse: 'Completed',
            newCourse: 'Started',
            activated: 'Account was created',
            failedCourse: 'Failed'
          };

          var iconMap = {
            completedCourse: 'CourseStatus_Completed',
            newCourse: 'CourseNewCourseGreen',
            activated: 'CourseStatus_NotStarted',
            failedCourse: 'CourseFailed'
          };

          var iconTextMap = {
            completedCourse: 'Completed Course',
            newCourse: 'Started Course',
            activated: 'Account Activated',
            failedCourse: 'Failed course'
          };

          return activityRow(
            index,
            {
              time: moment(activity?.createdAt).format('hh:mm'),
              date: moment(activity?.createdAt).format('DD/MM/YY')
            },
            activity?.course
              ? `${nameMap[activity.type]} the ${activity?.course?.name} Course`
              : nameMap[activity?.type],
            {
              icon: iconMap[activity.type],
              text: iconTextMap[activity.type]
            },
            {
              h: 3,
              m: 15
            },
            userName,
            classes
            
          )
        })}
      />
      <Paginator
        currentPage={pageInfo.currentPage}
        updatePage={onUpdatePage}
        numPages={pageInfo.totalPages}
        itemsPerPage={10}
        showRange={pageInfo.totalPages > 4 ? 4 : pageInfo.totalPages}
      />
    </div>
  );
};

export default ActivityTable;
