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
  activeTable: {}
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
        total: 100,
        completed: myCourse.status === 'complete' ? 100 : 0
      },
      attempt: 1,
      status: {
        isComplete: myCourse.status === 'complete'
      },
      onClick: () => {}
    })) ?? [];

  const pageProps = {
    total: activity.pageInfo?.total ?? 0,
    limit: activity.pageInfo?.limit ?? 10,
    offset: activity.pageInfo?.offset ?? 0
  };

  const pageInfo = {
    currentPage: Math.ceil(pageProps.offset/ pageProps.limit),
    totalPages: Math.ceil(pageProps.total/ pageProps.limit)
  };

  const onUpdatePage = (page: number) => {
    router.push(`/app/progress?offset=${(page - 1) * pageProps.limit}&limit=${pageProps.limit}`);
  }

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
          EmptyComponent={<div>No Courses to show</div>}
          className={classes.courseTable}
          courses={myCourses}
          rowClicked={() => {
            router.push('/app/courses/1');
          }}
        />
        <ActivityTable className={classes.activeTable} activity={activity} pageInfo={pageInfo} onUpdatePage={onUpdatePage} />
      </div>
    </Page>
  );
}

export default createFragmentContainer(Progress, {
  activity: graphql`
    fragment TrainingProgress_activity on ActivityPage {
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
  `,
  user: graphql`
    fragment TrainingProgress_user on User {
      firstName
      myCourses {
        status
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
