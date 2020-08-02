import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CoursePageHeader from 'components/core/PageHeader/CoursePageHeader';
import CoursePreview from 'components/Misc/CoursePreview';
import TickBullet from 'components/Misc/TickBullet';
import CarouselWithDemo from 'components/Misc/CarouselCourse/CarouselWithDemo';
import { Course } from 'sharedComponents/Overview/CourseCard';
import PageMargin from 'components/core/PageMargin';
import { createFragmentContainer, graphql } from 'react-relay';
import { CourseDetailsPage_course } from './__generated__/CourseDetailsPage_course.graphql';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

const useStyles = createUseStyles((theme: Theme) => ({
  courseRoot: {
    width: '100%',
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center',
  },
  centered: {
    flexDirection: 'row',
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    '@media (max-width: 1000px)': {
      flexDirection: 'column',
    },
  },
  column: {
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
  },
  heading: {
    fontSize: theme.fontSizes.tinyHeading,
    color: theme.colors.primaryBlack,
    fontWeight: 'bold',
    padding: '40px 0px 20px 0px',
  },
  text: {
    display: 'flex',
    alignItems: 'center',
    fontSize: theme.fontSizes.extraLarge,
    color: theme.colors.secondaryBlack,
  },
  bullet: {
    height: '10px',
    width: '11.5px',
    borderRadius: '10px',
    margin: '12.5px',
    backgroundColor: theme.colors.textBlue,
  },
  tickMargin: {
    margin: '15px',
    alignItems: 'center',
  },
  marginBottom: {
    marginBottom: '100px',
  },
  mainColumn: {
    flex: 1.8,
  },
  spacer: {
    flex: 1,
    '@media (min-width: 850px) and (max-width: 1300px)': {
      flex: 2,
    },
    '@media (max-width: 850px)': {
      display: 'none',
    },
  },
  carousel: {
    width: '90vw',
    marginTop: '50px',
    paddingBottom: '50px',
  },
  whiteBackground: {
    background: theme.colors.primaryWhite,
  },
}));

// Course Preview data
const defaultDetails: string[] = [
  ' 6 hours of on-demand video',
  '15 modules',
  '104 lessons',
  '4 examinations',
  'Full lifetime access',
  'Access on mobile/tablet/computer',
  'Industry-approved certificate',
];

const defaultCourse: Course = {
  id: 2,
  type: 'DANGEROUS GOODS AIR',
  colour: '#8C1CB4',
  url: require('../assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  title: 'Dangerous goods by air category 7',
  price: 60,
  description:
    'This course is for those involved in the handling, storage and loading of cargo or mail and baggage.',
  assigned: 40,
  expiring: 9,
  date: 'MAR 3rd 2020',
  location: 'TTC at Hilton T4',
  modules: 16,
  lessons: 144,
  videoTime: 4,
};

type Props = {
  course: CourseDetailsPage_course;
};

function CourseDetailsPage({ course }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.courseRoot}>
      <CoursePageHeader
        title={course?.name ?? ''}
        description={course?.excerpt ?? ''}
        history={['Courses', 'Aviation Security', 'Regulated Agents']}
        estimatedTime={`${course.hoursToComplete} hours`}
        lastUpdated="August 2020"
        price={`£${course.price}`}
        video={require('assets/Stock_Video.mp4')}
        onBuy={() => console.log('Buy')}
        onBasket={() => console.log('Basket')}
        sideComponent={
          <CoursePreview
            price={`£${course.price}`}
            details={defaultDetails}
            image={course.bannerImageURL ?? undefined}
            onBuy={() => console.log('Buy')}
            onBasket={() => console.log('Basket')}
          />
        }
      />
      <PageMargin
        centererStyle={classes.whiteBackground}
        centeredStyle={classes.centered}
      >
        <div className={classes.mainColumn}>
          <div className={classes.heading}>What you’ll learn</div>
          <div className={classes.row}>
            <div className={classes.column}>
              {course.whatYouLearn?.map((text) => (
                <TickBullet className={classes.tickMargin} text={text} />
              ))}
            </div>
            <div className={classes.column}></div>
          </div>
          <div className={classes.heading}>
            About this Course – Estimated time to complete{' '}
            {course.hoursToComplete} hours
          </div>
          <div className={classes.text}>{course.introduction}</div>
          <div className={classes.heading}>How to complete this Course</div>
          <div className={classes.text}>{course.howToComplete}</div>
          <div className={classes.heading}>Requirements</div>
          {course.requirements?.map((text) => (
            <div className={classes.text}>
              <div className={classes.bullet} />
              {text}
            </div>
          ))}
          <Spacer vertical spacing={3} />
        </div>
        <div className={classes.spacer} />
      </PageMargin>
      <div className={classes.centerer}>
        <CarouselWithDemo
          className={classes.carousel}
          heading="Explore our popular courses"
          description="It’s time to remove the headache for you and your team, with TTC
            you could be logged in and learning in 24 hours."
          courses={[1, 2, 3, 4, 5, 6, 7].map(() => defaultCourse)}
        />
      </div>
    </div>
  );
}

export default createFragmentContainer(CourseDetailsPage, {
  course: graphql`
    fragment CourseDetailsPage_course on Course {
      ident: id
      name
      price
      excerpt
      introduction
      bannerImageURL
      price
      howToComplete
      hoursToComplete
      whatYouLearn
      requirements
      category {
        name
        color
      }
    }
  `,
});
