import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Button from 'sharedComponents/core/Input/Button';
import { createFragmentContainer, graphql } from 'react-relay';
import ActiveType from 'components/ActivityTable/ActiveType';
import theme, { Theme } from 'helpers/theme';
import Table from 'sharedComponents/core/Table';
import Text from 'sharedComponents/core/Table/Text/Text';
import Paginator from 'sharedComponents/Pagination/Paginator';
import TimeSpent from 'components/ActivityTable/TimeSpent';
import ActivityName from 'components/ActivityTable/ActivityName';
import classnames from 'classnames';
import { ActivityTable_activity } from './__generated__/ActivityTable_activity.graphql';
import moment from 'moment';

const useStyles = createUseStyles((theme: Theme) => ({
  rootActivityTable: {},
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
    marginLeft: theme.spacing(3),
    background: theme.colors.primaryWhite
  },
  activeType: {
    display: 'flex',
    alignItems: 'center',
    '& div': {
      marginLeft: theme.spacing(1)
    }
  },
  pagination: {
    marginTop: theme.spacing(4)
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
  avatar: string,
  classes?: any
): any => ({
  key,
  cells: [
    // {
    //   component: () => (
    //     <CheckboxSingle
    //       box={{
    //         label: "",
    //         checked: false,
    //       }}
    //       setBox={() => {}}
    //     />
    //   ),
    // },
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
      component: () => <ActivityName avatar={avatar} title={title} />
    },
    {
      component: () => (
        <ActiveType icon={activeType.icon} text={activeType.text} />
      )
    },
    { component: () => null },
    {
      component: () => null
    }
  ],
  onClick: () => {}
});

type Props = {
  className?: string;
  activity: ActivityTable_activity;
};

const ActivityTable = ({ activity, className }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  console.log('asda', activity);
  return (
    <div className={classnames(className, classes.rootActivityTable)}>
      <div className={classes.sectionTitleWrapper}>
        <h2>Your activity</h2>
        <div className={classes.courseDropdown}>
          <Button
            archetype={'default'}
            icon={{ right: 'FilterAdjust' }}
            children={'Filter Activities'}
          />
        </div>
      </div>
      <Table
        header={['ACTIVITY TIME', 'NAME', 'ACTIVE TYPE', '', '']}
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
            'https://picsum.photos/id/1/200/300',
            classes
          );
        })}
      />
      <div className={classes.pagination}>
        <Paginator
          currentPage={1}
          updatePage={() => {}}
          numPages={10}
          itemsPerPage={10}
        />
      </div>
    </div>
  );
};

export default createFragmentContainer(ActivityTable, {
  activity: graphql`
    fragment ActivityTable_activity on ActivityPage {
      edges {
        type
        createdAt
        course {
          ident: id
          name
        }
      }
      pageInfo {
        total
      }
    }
  `
});
