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
};

function Progress({ activity }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  console.log('Activity', activity);
  const userName = 'James';
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
              text={`${userName}, here you can see your training progress,`}
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
          courses={[]}
          rowClicked={() => {
            router.push('/app/courses/1');
          }}
        />
        <ActivityTable className={classes.activeTable} />
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
          name
        }
      }
      pageInfo {
        total
      }
    }
  `
});
