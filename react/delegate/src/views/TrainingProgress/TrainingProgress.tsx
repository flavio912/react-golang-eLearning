import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { createFragmentContainer, graphql } from 'react-relay';
import { useRouter } from 'found';
import Heading from 'components/core/Heading';
import CourseTable from 'sharedComponents/CourseTable';
import ActivityTable from 'components/ActivityTable/ActivityTable';
import Page from 'components/Page';
import { TrainingProgress_activity } from './__generated__/TrainingProgress_activity.graphql';
import { TrainingProgress_user } from './__generated__/TrainingProgress_user.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  progressRoot: {
    width: `100%`
  },
  subHeading: {
    marginBottom: 0
  },
  headingDescription: {
    marginBottom: theme.spacing(4)
  },
  heading: {
    marginBottom: `${2 * theme.spacing(2)}px`
  },
  header: {},
  courseTable: {
    marginBottom: theme.spacing(4)
  },
  activeTable: {},
  noCourses: {
    textAlign: 'center',
    fontWeight: 300,
    color: theme.colors.textGrey
  }
}));

type Props = {
  activity: TrainingProgress_activity;
  user: TrainingProgress_user;
};

function Progress({ activity, user }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  const myCourses =
    user?.myCourses?.map((myCourse) => ({
      title: myCourse.course.name,
      categoryName: myCourse.course.category?.name ?? '',
      progress: {
        total: myCourse.progress?.total ?? 0,
        completed:
          myCourse.status === 'complete'
            ? 100
            : myCourse.progress?.completed ?? 0
      },
      attempt: 1,
      status: {
        isComplete: myCourse.status === 'complete'
      },
      onClick: () => {}
    })) ?? [];

  const onUpdatePage = (page: number, limit: number) => {
    router.push(`/app/progress?offset=${(page - 1) * limit}&limit=${limit}`);
  };

  return (
    <Page>
      <div className={classes.progressRoot}>
        <div className={classes.header}>
          <Heading
            text="Training Progress"
            size={'large'}
            className={classes.heading}
          />
          <div className={classes.headingDescription}>
            <Heading
              text={`${user.firstName}, here you can see your training progress,`}
              size={'medium'}
              className={classes.subHeading}
            />
            <Heading
              text={`and keep up to date with your daily activity.`}
              size={'medium'}
              className={classes.subHeading}
            />
          </div>
        </div>
        <CourseTable
          EmptyComponent={
            <div className={classes.noCourses}>No Courses to show</div>
          }
          showCertificates={user.type === 'individual'}
          className={classes.courseTable}
          courses={myCourses}
          rowClicked={() => {
            router.push('/app/courses/1');
          }}
        />
        <ActivityTable
          className={classes.activeTable}
          activity={activity}
          onUpdatePage={onUpdatePage}
        />
      </div>
    </Page>
  );
}

export default createFragmentContainer(Progress, {
  activity: graphql`
    fragment TrainingProgress_activity on ActivityPage {
      ...ActivityTable_activity
    }
  `,
  user: graphql`
    fragment TrainingProgress_user on User {
      firstName
      type
      myCourses {
        status
        progress {
          total
          completed
        }
        course {
          name
          category {
            name
          }
        }
      }
    }
  `
});
