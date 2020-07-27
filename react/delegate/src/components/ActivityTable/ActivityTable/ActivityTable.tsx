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
import moment from 'moment';
import { ActivityTable_activity } from './__generated__/ActivityTable_activity.graphql';

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

type PageInfo = {
  currentPage: number;
  totalPages: number;
};

type Props = {
  className?: string;
  activity: ActivityTable_activity;
  onUpdatePage: (page: number, limit: number) => void;
};

const ActivityTable = ({ activity, className, onUpdatePage }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  
  const pageProps = {
    total: activity.pageInfo?.total ?? 0,
    limit: activity.pageInfo?.limit ?? 10,
    offset: activity.pageInfo?.offset ?? 0
  };

  const pageInfo = {
    currentPage: Math.ceil(pageProps.offset/ pageProps.limit),
    totalPages: Math.ceil(pageProps.total/ pageProps.limit)
  };

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
              ? `You ${nameMap[activity.type]} the ${activity?.course?.name} Course`
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
          currentPage={pageInfo.currentPage}
          updatePage={(page) => onUpdatePage(page, pageProps.limit)}
          numPages={pageInfo.totalPages}
          itemsPerPage={10}
          showRange={pageInfo.totalPages > 4 ? 4 : pageInfo.totalPages}
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
        limit
        offset
      }
    }
  `
});
