import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CourseCard from 'sharedComponents/Overview/CourseCard';
import { Course as CourseCardProps } from 'sharedComponents/Overview/CourseCard/CourseCard';
import { Transition, animated } from 'react-spring/renderprops';
import Button from 'sharedComponents/core/Input/Button';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  sliderRoot: {
    maxHeight: 428,
    overflow: 'hidden'
  },
  sliderList: {
    display: 'flex',
    flexDirection: 'column'
  },
  sliderTransition: {
    position: 'relative'
  },
  courseCard: {
    width: 325,
    marginRight: 46
  },
  currentCourses: {
    display: 'flex',
    justifyContent: 'flex-end'
  },
  sliderButtons: {
    display: 'flex',
    marginBottom: 41,
    paddingRight: 45,
    justifyContent: 'flex-end'
  },
  navButton: {
    width: 40
  },
  prevButton: {
    marginRight: 15
  },
  nextButton: {},
  nextCourse: {
    position: 'absolute',
    top: 0,
    left: `calc(100% - 24px)`
  },
  navButtonDisabled: {
    backgroundColor: theme.colors.approxZircon,
    borderColor: theme.colors.borderGrey,
  }
}));
export type Course = CourseCardProps & { progress: number };
type Props = {
  className?: string;
  courses: Course[];
  slidesToShow: number;
};
function arrayChunk(array: Course[] = [], chunk = 4) {
  let result = [];
  for (let i = 0; i < array.length; i += chunk) {
    result.push(array.slice(i, i + chunk));
  }
  return result;
}
function CourseSlider({ className, courses, slidesToShow = 4 }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [active, setActive] = React.useState(0);
  const sliders = arrayChunk(courses, slidesToShow);
  return (
    <div className={classNames(classes.sliderRoot, className)}>
      <div className={classes.sliderList}>
        <div className={classes.sliderButtons}>
          <Button
            onClick={() => setActive(active > 0 ? active - 1 : 0)}
            small={true}
            disabled={active === 0}
            className={classNames(classes.prevButton, classes.navButton, {
              [classes.navButtonDisabled]: active === 0
            })}
          >
            <Icon name="ArrowLeft" size={15} pointer={true} />
          </Button>
          <Button
            onClick={() =>
              setActive(
                active < sliders.length - 1 ? active + 1 : sliders.length - 1
              )
            }
            small={true}
            disabled={active === sliders.length - 1}
            className={classNames(classes.nextButton, classes.navButton, {
              [classes.navButtonDisabled]: active === sliders.length - 1
            })}
          >
            <Icon name="ArrowRight" size={15} pointer={true} />
          </Button>
        </div>
        <div className={classes.sliderTransition}>
          <Transition
            native
            reset
            unique
            items={active}
            from={{ opacity: 0, transform: 'translate3d(100%,0,0)' }}
            enter={{ opacity: 1, transform: 'translate3d(0%,0,0)' }}
            leave={{ opacity: 0, transform: 'translate3d(-50%,0,0)' }}
          >
            {(index: any) => (style: any) => {
              const newCourses = sliders[index] as any;
              return (
                <animated.div
                  style={{ ...style }}
                  className={classes.currentCourses}
                >
                  {newCourses.map((course: Course, index: number) => (
                    <CourseCard
                      key={index}
                      course={course}
                      onClick={() => {}}
                      progress={course.progress}
                      className={classNames(classes.courseCard, {
                        [classes.nextCourse]:
                          newCourses.length === slidesToShow &&
                          newCourses.length - 1 === index
                      })}
                    />
                  ))}
                </animated.div>
              );
            }}
          </Transition>
        </div>
      </div>
    </div>
  );
}

export default CourseSlider;
