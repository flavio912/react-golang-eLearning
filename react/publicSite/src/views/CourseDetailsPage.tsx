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
    display: 'flex',
    flexDirection: 'column',
    flex: 1,
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

type Props = {};

function CourseDetailsPage(props: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.courseRoot}>
      <CoursePageHeader
        title="Cargo Operative Screener (COS) Recurrent – VC, HS, MDE"
        description="This recurrent course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search and metal detection equipment."
        history={['Courses', 'Aviation Security', 'Regulated Agents']}
        estimatedTime="6 hours"
        lastUpdated="May 2020"
        price="£310.00"
        video={require('assets/Stock_Video.mp4')}
        onBuy={() => console.log('Buy')}
        onBasket={() => console.log('Basket')}
        sideComponent={
          <CoursePreview
            price="£310.00"
            details={defaultDetails}
            onBuy={() => console.log('Buy')}
            onBasket={() => console.log('Basket')}
            video={require('assets/Stock_Video.mp4')}
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
              <TickBullet
                className={classes.tickMargin}
                text="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"
              />
              <TickBullet
                className={classes.tickMargin}
                text="Quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea"
              />
              <TickBullet
                className={classes.tickMargin}
                text="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"
              />
            </div>
            <div className={classes.column}>
              <TickBullet
                className={classes.tickMargin}
                text="Punt in culpa qui officia deserunt mollit anim id est laboru"
              />
              <TickBullet
                className={classes.tickMargin}
                text="Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ip"
              />
              <TickBullet
                className={classes.tickMargin}
                text="Punt in culpa qui officia deserunt mollit anim id est laboru"
              />
            </div>
          </div>
          <div className={classes.heading}>
            About this Course – Estimated time to complete 6 hours
          </div>
          <div className={classes.text}>
            This recurrent course is for those who screen air cargo and mail, to
            provide them with the knowledge and skills needed to deliver
            effective screening in visual check, hand search and metal detection
            equipment.
          </div>
          <div className={classes.heading}>How to complete this Course</div>
          <div className={classes.text}>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
            aliquip ex ea commodo consequat.
            <br />
            <br />
            Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
            nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
            pariatur.
            <br />
            <br />
            Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
            officia deserunt mollit anim id est laborum.
          </div>
          <div className={classes.heading}>Requirements</div>
          <div className={classes.text}>
            <div className={classes.bullet} />
            Have a computer with Internet
          </div>
          <div className={classes.text}>
            <div className={classes.bullet} />
            Have a pair of working speakers or headphones
          </div>
          <div className={classNames(classes.text, classes.marginBottom)}>
            <div className={classes.bullet} />
            Brace yourself for a 6 hour stint of education!
          </div>
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

export default CourseDetailsPage;
