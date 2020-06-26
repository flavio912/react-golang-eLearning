import React, { useState, createRef, useEffect } from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CourseCard from 'sharedComponents/Overview/CourseCard';
import { Course as CourseCardProps } from 'sharedComponents/Overview/CourseCard/CourseCard';
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
    display: 'flex',
    justifyContent: 'flex-start',
    overflowX: 'scroll'
  },
  courseCard: {
    width: 325,
    marginRight: 46
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
    borderColor: theme.colors.borderGrey
  }
}));
export type Course = CourseCardProps & { progress: number };
type Props = {
  className?: string;
  courses: Course[];
  slidesToShow: number;
};

function CourseSlider({ className, courses, slidesToShow = 4 }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [active, setActive] = useState(0);
  const refs = courses.reduce((acc, value) => {
    acc[value.id] = createRef();
    return acc;
  }, {});

  const handleClick = (id: number) =>
    refs[id] &&
    refs[id].current &&
    refs[id].current.scrollIntoView({
      behavior: 'smooth',
      block: 'start'
    });
  useEffect(() => {
    handleClick(courses[active].id);
  });

  return (
    <div className={classNames(classes.sliderRoot, className)}>
      <div className={classes.sliderList}>
        <div className={classes.sliderButtons}>
          <Button
            onClick={() =>
              setActive(active - slidesToShow > 0 ? active - slidesToShow : 0)
            }
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
                active < courses.length - slidesToShow
                  ? active + slidesToShow
                  : courses.length - 1
              )
            }
            small={true}
            disabled={active === courses.length - 1}
            className={classNames(classes.nextButton, classes.navButton, {
              [classes.navButtonDisabled]: active === courses.length - 1
            })}
          >
            <Icon name="ArrowRight" size={15} pointer={true} />
          </Button>
        </div>
        <div className={classes.sliderTransition}>
          {courses.map((course: Course, index: number) => (
            <div key={course.id} ref={refs[course.id]}>
              <CourseCard
                key={index}
                course={course}
                onClick={() => {
                  //handle click here
                }}
                progress={course.progress}
                className={classNames(classes.courseCard, {
                  [classes.nextCourse]:
                    courses.length === slidesToShow &&
                    courses.length - 1 === index
                })}
              />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default CourseSlider;
