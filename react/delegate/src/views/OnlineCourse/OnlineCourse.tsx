import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import Button from 'components/core/Input/Button';
import { createFragmentContainer, graphql } from 'react-relay';
import { Theme } from 'helpers/theme';
import PageTitle from 'components/PageTitle';
import FlatCard from 'components/core/Cards/FlatCard';
import CourseSyllabusCard from 'components/Overview/CourseSyllabusCard';
import Icon from 'sharedComponents/core/Icon';
import { useRouter } from 'found';
import { OnlineCourse_myActiveCourse } from './__generated__/OnlineCourse_myActiveCourse.graphql';
import moment from 'moment';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

const useStyles = createUseStyles((theme: Theme) => ({
  rootOnlineCourse: {
    gridTemplateColumns: '2fr 1fr',
    gridGap: theme.spacing(3),
    display: 'grid',
    maxWidth: 1040
  },
  courseHead: {
    display: 'flex',
    alignItems: 'center',
    marginBottom: 26,
    marginTop: 32
  },
  courseHeadItem: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: 80,
    '&:last-child': {
      marginRight: 0
    }
  },
  labelBold: {
    textTransform: 'uppercase',
    color: theme.colors.primaryBlack,
    fontWeight: 600,
    fontSize: theme.fontSizes.xSmall,
    letterSpacing: -0.3,
    lineHeight: `18px`,
    marginRight: 5
  },
  labelValue: {
    textTransform: 'uppercase',
    color: theme.colors.secondaryBlack,
    fontWeight: 400,
    fontSize: theme.fontSizes.xSmall,
    letterSpacing: -0.3,
    lineHeight: `18px`
  },
  flatCard: {
    alignItems: 'center',
    marginBottom: 28
  },
  flatCardText: {
    color: theme.colors.primaryBlack,
    fontWeight: 'bold',
    fontSize: theme.fontSizes.large,
    letterSpacing: -0.25,
    lineHeight: `22px`,
    margin: [0, 35, 0, 26]
  },
  courseContent: {},
  courseContentTitle: {
    color: theme.colors.primaryBlack,
    fontWeight: 'bold',
    fontSize: theme.fontSizes.tinyHeading,
    letterSpacing: -0.5,
    lineHeight: `51px`,
    margin: 0
  },
  courseContentText: {
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.extraLarge,
    letterSpacing: -0.45,
    lineHeight: `30px`,
    margin: 0,
    fontWeight: 300
  },
  howToCompleteDescription: {
    '& p': {
      marginBottom: 20,
      '&:last-child': {
        marginBottom: 0
      }
    }
  },
  keyThings: {
    padding: 0,
    margin: [4, 0],
    listStyle: 'none'
  },
  keyThingDot: {
    position: 'relative',
    paddingLeft: 21,
    marginBottom: 8,
    '&:last-child': {
      marginBottom: 0
    },
    '&:before': {
      content: "''",
      width: 10,
      height: 10,
      backgroundColor: theme.colors.textBlue,
      marginRight: 10,
      display: 'block',
      position: 'absolute',
      borderRadius: 50,
      left: 0,
      top: `calc(50% - 5px)`
    }
  },
  courseSyllabus: {
    marginTop: 52
  }
}));

const howToCompleteDescription = [
  'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
  'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
  'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
];
const keyThings = [
  'Bullet point 1 note about the Dangerous Goods course ',
  'Another interesting bullet point to consider',
  'A final example of a bullet point to consider when taking this course'
];
const courseFake = {
  id: 1,
  title: 'Dangerous Goods by Air Category 4',
  category: 'DANGEROUS GOODS',
  percentCompleted: 0,
  enrolled: '10/03/2020',
  aboutDescription:
    'This course is for staff of freight forwarders involved in processing cargo or mail (other than dangerous goods).',
  howToCompleteDescription,
  keyThings,
  estimateTimeComplete: 6
};
const defaultSyllabus = {
  completePercentage: 24,
  modules: [
    {
      sections: [
        {
          name: 'Lesson 1-1',
          uuid: '00000-0000-00000-0000',
          complete: false
        },
        {
          name: 'Lesson 1-2',
          uuid: '00000-0000-00000-0000',
          complete: false
        }
      ]
    },
    {
      sections: [
        {
          name: 'Lesson 2-1',
          uuid: '00000-0000-00000-0000',
          complete: false
        },
        {
          name: 'Lesson 2-2',
          uuid: '00000-0000-00000-0000',
          complete: false
        }
      ]
    },
    {
      sections: [
        {
          name: 'Lesson 3-1',
          uuid: '00000-0000-00000-0000',
          complete: false
        },
        {
          name: 'Lesson 3-2',
          uuid: '00000-0000-00000-0000',
          complete: false
        }
      ]
    }
  ]
};

export type Course = {
  id: number;
  title: string;
  category: string;
  percentCompleted: number;
  enrolled: string | Date;
  aboutDescription?: string | React.ReactNode;
  howToCompleteDescription?: string[];
  keyThings?: string[];
  estimateTimeComplete?: number;
};

export type OnlineCourseProps = {
  course: Course;
  // onClick: Function;
  className?: string;
  myActiveCourse?: OnlineCourse_myActiveCourse;
};

function OnlineCourse({
  course = courseFake,
  myActiveCourse,
  className
}: OnlineCourseProps) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  console.log('course', myActiveCourse);

  const syllabusSections = (myActiveCourse?.course.syllabus ?? []).map(
    (courseElement) => {
      if (courseElement.type == 'module') {
        return {
          sections: (courseElement?.syllabus ?? []).map((moduleItem) => {
            return {
              name: moduleItem.name ?? '',
              uuid: moduleItem.uuid ?? '',
              complete: moduleItem.complete ?? false
            };
          })
        };
      } else {
        return {
          sections: [
            {
              name: courseElement.name ?? '',
              uuid: courseElement.uuid ?? '',
              complete: courseElement.complete ?? false
            }
          ]
        };
      }
    }
  );

  const syllabus = {
    completePercentage: 23,
    modules: syllabusSections
  };
  return (
    <div className={classes.rootOnlineCourse}>
      <div>
        <PageTitle
          title={myActiveCourse?.course.name ?? ''}
          backProps={{
            text: 'all Online Courses',
            onClick: () => {
              router.push('/app/courses');
            }
          }}
        />
        <div className={classes.courseHead}>
          <div className={classes.courseHeadItem}>
            <span className={classes.labelBold}>Category: </span>
            <span className={classes.labelValue}>{course.category}</span>
          </div>
          <div className={classes.courseHeadItem}>
            <span className={classes.labelBold}>
              {course.percentCompleted}% Completed
            </span>
          </div>
          <div className={classes.courseHeadItem}>
            <span className={classes.labelBold}>Enrolled:</span>
            <span className={classes.labelValue}>
              {moment(myActiveCourse?.enrolledAt).format('DD/MM/YYYY')}
            </span>
          </div>
        </div>
        <FlatCard shadow padding="large" className={classes.flatCard}>
          <Icon name="Volume" size={51} />
          <p className={classes.flatCardText}>
            Make sure your speakers are turned on before you start this course
          </p>
          <Button
            title="Begin Course"
            onClick={() => {}}
            padding="massive"
            noWrap
          />
        </FlatCard>
        <div className={classes.courseContent}>
          <h6 className={classes.courseContentTitle}>
            About this Course
            {myActiveCourse?.course.hoursToComplete
              ? `– Estimated time to complete ${myActiveCourse?.course.hoursToComplete} hours`
              : ''}
          </h6>
          <p className={classes.courseContentText}>
            {myActiveCourse?.course.excerpt}
          </p>
        </div>
        <Spacer vertical spacing={3} />
        {myActiveCourse?.course.howToComplete && (
          <div className={classes.courseContent}>
            <h6 className={classes.courseContentTitle}>
              How to complete this Course
            </h6>
            <div className={classes.howToCompleteDescription}>
              {myActiveCourse?.course.howToComplete}
            </div>
          </div>
        )}
        {myActiveCourse?.course.whatYouLearn &&
          myActiveCourse?.course.whatYouLearn.length > 0 && (
            <div className={classes.courseContent}>
              <h6 className={classes.courseContentTitle}>
                Key things to consider
              </h6>
              <ul className={classes.keyThings}>
                {myActiveCourse?.course.whatYouLearn.map((text, index) => (
                  <li
                    className={classNames(
                      classes.courseContentText,
                      classes.keyThingDot
                    )}
                    key={index}
                  >
                    {text}
                  </li>
                ))}
              </ul>
            </div>
          )}
      </div>
      <div className={classes.courseSyllabus}>
        <CourseSyllabusCard courseSyllabus={syllabus} />
      </div>
    </div>
  );
}

export default createFragmentContainer(OnlineCourse, {
  myActiveCourse: graphql`
    fragment OnlineCourse_myActiveCourse on MyCourse {
      status
      enrolledAt
      course {
        name
        excerpt
        introduction
        howToComplete
        whatYouLearn
        hoursToComplete
        syllabus {
          name
          type
          uuid
          complete
          ... on Module {
            syllabus {
              name
              type
              uuid
              complete
            }
          }
        }
      }
    }
  `
});
