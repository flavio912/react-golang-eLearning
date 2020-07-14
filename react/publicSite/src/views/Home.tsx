import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import Homepage from 'components/Overview/Homepage';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import { Router } from 'found';
import CarouselCourse from 'components/Misc/CarouselCourse';
import { Course } from 'sharedComponents/Overview/CourseCard';
import Button from 'sharedComponents/core/Input/Button';
import CarouselWithDemo from 'components/Misc/CarouselCourse/CarouselWithDemo';
import PageMargin from 'components/core/PageMargin';

const useStyles = createUseStyles((theme: Theme) => ({
  homeRoot: {
    width: '100%'
  },
  whiteSpacer: {
    background: theme.colors.primaryWhite,
    padding: '60px 0px 100px 0px'
  },
  heading: {
    fontSize: 32,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    padding: '60px 0px',
    textAlign: 'center'
  },
  explore: {
    paddingTop: 70
  },
  exploreCont: {
    maxWidth: '100%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center'
  },
  exploreText: {
    textAlign: 'center',
    fontSize: theme.fontSizes.extraLarge,
    maxWidth: 500
  },
  buttonHolder: {
    display: 'flex'
  },
  button: {
    height: 52,
    fontSize: 18,
    fontWeight: 800,
    boxShadow: '0px 2px 9px #00000014',
    padding: '0px 36px'
  }
}));

type Props = {
  router: Router;
};

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
  videoTime: 4
};

function Home({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.homeRoot}>
      <Homepage
        imageURL=""
        onView={() => {
          router.push('/courses');
        }}
        onDemo={() => {}}
      />
      <PageMargin centererStyle={classes.whiteSpacer}>
        <FloatingVideo
          width={560}
          source={require('assets/Stock_Video.mp4')}
          author={{
            name: 'Kristian Durhuus',
            title: 'Chief Executive Officer',
            quote:
              'TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore .'
          }}
        />
    </PageMargin>
    <PageMargin>
      <CarouselWithDemo
          className={classes.explore}
          heading="Explore our popular courses"
          description="It’s time to remove the headache for you and your team, with TTC
          you could be logged in and learning in 24 hours."
          courses={[1, 2, 3, 4, 5, 6, 7].map(() => defaultCourse)}
      />
      </PageMargin>
      <PageMargin centererStyle={classes.whiteSpacer}>
        <div className={classes.heading}>What you can do with TTC Hub</div>
        <ImageWithText
          title="Online courses"
          subtitle="Get certified online"
          description="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean quis tellus at quam aliquam aliquam. Phasellus consequat tincidunt ex nec blandit. Duis et sem lacus. "
          image={require('assets/getCertified.svg')}
          link={{ title: 'See Online Courses', link: '/' }}
        />
        <Spacer spacing={5} vertical />
        <ImageWithText
          title="Online courses"
          subtitle="Get certified online"
          description="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean quis tellus at quam aliquam aliquam. Phasellus consequat tincidunt ex nec blandit. Duis et sem lacus. "
          image={require('assets/manageTeam.svg')}
          link={{ title: 'See Online Courses', link: '/' }}
          textRight
        />
        <Spacer spacing={5} vertical />
        <ImageWithText
          title="Online courses"
          subtitle="Get certified online"
          description="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean quis tellus at quam aliquam aliquam. Phasellus consequat tincidunt ex nec blandit. Duis et sem lacus. "
          image={require('assets/trackCompliance.svg')}
          link={{ title: 'See Online Courses', link: '/' }}
        />
      </PageMargin>
    </div>
  );
}

export default Home;
