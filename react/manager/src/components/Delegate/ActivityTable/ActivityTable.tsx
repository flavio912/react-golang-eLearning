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
import { createFragmentContainer, graphql } from 'react-relay';

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

type Props = {
  activity: any;
};

const ActivityTable = ({ activity }: Props) => {
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
        header={['ACTIVITY TIME', 'NAME', 'ACTIVE TYPE', 'TIME SPENT', '', '']}
        rows={[
          activityRow(
            1,
            {
              time: '10:29',
              date: '01/02/2020'
            },
            'Bruce failed the Dangerous Goods by Road Awareness Course',
            {
              icon: 'CourseFailed',
              text: 'Failed Course'
            },
            {
              h: 3,
              m: 15
            },
            'Bruce Willis',
            classes
          ),
          activityRow(
            2,
            {
              time: '10:29',
              date: '01/02/2020'
            },
            'Bruce failed the Dangerous Goods by Road Awareness Course',
            {
              icon: 'CourseNewCourse',
              text: 'New Course'
            },
            {
              h: 3,
              m: 15
            },
            'Bruce Willis',
            classes
          ),
          activityRow(
            3,
            {
              time: '10:29',
              date: '01/02/2020'
            },
            'Bruce failed the Dangerous Goods by Road Awareness Course',
            {
              icon: 'CourseCertificates',
              text: 'New New certificate!'
            },
            {
              h: 3,
              m: 15
            },
            'Bruce Willis',
            classes
          ),
          activityRow(
            4,
            {
              time: '10:29',
              date: '01/02/2020'
            },
            'Bruce failed the Dangerous Goods by Road Awareness Course',
            {
              icon: 'CourseAccountActivated',
              text: 'Account active'
            },
            'n/a',
            'Bruce Willis',
            classes
          )
        ]}
      />
      <Paginator
        currentPage={1}
        updatePage={() => {}}
        numPages={10}
        itemsPerPage={10}
      />
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
