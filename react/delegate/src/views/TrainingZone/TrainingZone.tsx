import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { createFragmentContainer, graphql } from 'react-relay';
import Heading from 'components/core/Heading';
import TopInfo from './TopInfo';
import CardItem from 'components/core/Cards/CardItem';
import { useRouter } from 'found';
import Page from 'components/Page';
import {
  TrainingZone_user,
  CourseStatus
} from './__generated__/TrainingZone_user.graphql';
import CourseCard from 'sharedComponents/Overview/CourseCard';

const useStyles = createUseStyles((theme: Theme) => ({
  trainingZoneRoot: {
    position: 'relative',
    display: 'grid',
    gridTemplateColumns: 'repeat(6, 1fr)',
    gridTemplateRows: 'repeat(4, auto)',
    gridGap: theme.spacing(3),
    gridTemplateAreas: `
      "header header header topinf topinf topinf"
      "allcou allcou expire expire trainp trainp"
      "jumpba jumpba jumpba .      .      .     "
      "jumpco jumpco jumpco jumpco jumpco jumpco"
    `,
    '@media (max-width: 1350px)': {
      gridTemplateAreas: `
      "header"
      "topinf"
      "allcou"
      "expire"
      "trainp"
      "jumpba"
      "jumpco"
    `,
      gridTemplateColumns: 'repeat(1, 1fr)'
    }
  },
  trainingHeader: {
    gridArea: 'header'
  },
  topInfo: {
    gridArea: 'topinf'
  },
  card1: {
    gridArea: 'allcou'
  },
  card2: {
    gridArea: 'expire'
  },
  card3: {
    gridArea: 'trainp'
  },
  jumpHeader: {
    gridArea: 'jumpba'
  },
  courseHolder: {
    gridArea: 'jumpco',
    display: 'grid',
    gridGap: 26,
    gridTemplateColumns: 'repeat(auto-fit, 298px)'
  }
}));

type Props = {
  user: TrainingZone_user;
};

function TrainingZone({ user }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();

  const courses: { status: CourseStatus; enrolledAt: string }[] = [];
  if (user?.myCourses) {
    user?.myCourses.map((value) => {
      courses.push({ status: value.status, enrolledAt: value.enrolledAt });
    });
  }

  const incomplete = user?.myCourses?.filter(
    (course) => course.status === 'incomplete'
  );

  return (
    <Page>
      <div className={classes.trainingZoneRoot}>
        <div className={classes.trainingHeader}>
          <Heading text="Training Zone" size={'large'} />
          <Heading
            text={`You're doing great ${user?.firstName}, keep up the good work so you don't lose your momentum`}
            size={'medium'}
          />
        </div>
        <TopInfo className={classes.topInfo} courses={courses} />
        <CardItem
          className={classes.card1}
          title={'All Courses'}
          description={
            'Find all of the online courses you have been enrolled on here'
          }
          buttonProps={{
            title: 'View Courses',
            onClick: () => {
              router.push('/app/courses');
            }
          }}
        />
        <CardItem
          className={classes.card2}
          title={'Expiring Soon'}
          description={
            'See your previously completed courses that are due to expire'
          }
          buttonProps={{ title: 'Expiring Certificates', onClick: () => {} }}
        />
        <CardItem
          className={classes.card3}
          title={'Training Progress'}
          description={'Review your progress of all Courses in real time'}
          buttonProps={{ title: 'See Help Guides', onClick: () => {} }}
        />
        {incomplete && incomplete.length > 0 && (
          <div className={classes.jumpHeader}>
            <Heading text="Jump back in" size={'large'} />
            <Heading
              text={`${user?.firstName}, you have some unfinished courses,
            jump back in to continue learning now.`}
              size={'medium'}
            />
          </div>
        )}
        <div className={classes.courseHolder}>
          {incomplete &&
            incomplete.map((course, index) => (
              <CourseCard
                key={index}
                course={{
                  title: course.course.name,
                  id: course.course.ident,
                  type: course.course.category?.name ?? 'Online Course',
                  colour: course.course.category?.color ?? '',
                  description: course.course.excerpt ?? '',
                  url: course.course.bannerImageURL ?? ''
                }}
                onClick={() => {
                  router.push(`/app/courses/${course.course.ident}`);
                }}
                progress={course.progress?.percent ?? 0}
              />
            ))}
        </div>
      </div>
    </Page>
  );
}

export default createFragmentContainer(TrainingZone, {
  user: graphql`
    fragment TrainingZone_user on User {
      firstName
      myCourses {
        course {
          ident: id
          name
          excerpt
          category {
            color
            name
          }
          type
          bannerImageURL
        }
        status
        enrolledAt
        progress {
          percent
        }
      }
    }
  `
});
