import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import theme, { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';
import CourseCard, { Course } from 'sharedComponents/Overview/CourseCard';
const useStyles = createUseStyles((theme: Theme) => ({
  rooCarouselCourse: {},
  carousel: {},
  containerClass: {
    padding: theme.spacing(1),
    paddingBottom: 44,
    maxWidth: '100vw'
  },
  carouselCard: {
    borderRadius: [0, 9, 9, 9],
    boxShadow: `0 2px 12px 0 rgba(0, 0, 0, 0.18)`,
    marginRight: 34
  },
  arrow: {
    height: 57,
    width: 57,
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: `0 2px 11px 0 rgba(0, 0, 0, 0.16)`,
    position: 'absolute',
    outline: 'none',
    transition: 'all 0.5s',
    borderRadius: 50,
    zIndex: 10,
    opacity: 1,
    cursor: 'pointer',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  arrowRight: {
    right: 10.5
  },
  arrowLeft: {
    left: 8.5
  },
  dot: {
    position: 'relative',
    width: 14,
    height: 14,
    border: `1px solid ${theme.colors.silver}`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 0,
    borderRadius: 50,
    margin: [0, 7.5],
    cursor: 'pointer',
    background: 'transparent',
    '&:focus': {
      outline: 'none'
    }
  },
  dotActive: {
    '&:after': {
      content: "''",
      display: 'block',
      width: 8,
      height: 8,
      borderRadius: 50,
      backgroundColor: theme.colors.navyBlue2
    }
  }
}));
type Props = {
  courses: Course[];
};
const responsive = {
  desktop: {
    breakpoint: { max: 90000, min: 2000 },
    items: 5,
    slidesToSlide: 4, // optional, default to 1.
    partialVisibilityGutter: 10
  },
  desktop1: {
    breakpoint: { max: 3000, min: 1250 },
    items: 3,
    slidesToSlide: 4, // optional, default to 1.
    partialVisibilityGutter: 10
  },
  tablet: {
    breakpoint: { max: 1250, min: 900 },
    items: 3,
    slidesToSlide: 3,
    partialVisibilityGutter: 30
  },
  medium: {
    breakpoint: { max: 900, min: 618 },
    items: 2,
    slidesToSlide: 1,
    partialVisibilityGutter: 0
  },
  small: {
    breakpoint: { max: 618, min: 0 },
    items: 1,
    slidesToSlide: 1,
    partialVisibilityGutter: 0
  }
};
type ArrowType = {
  children: JSX.Element;
  onClick?: Function;
  className?: any;
};
const Arrow = ({ children, className, onClick }: ArrowType) => (
  <div className={className} onClick={() => onClick && onClick()}>
    {children}
  </div>
);
const CustomDot = ({ onClick, active, classes }: any) => {
  return (
    <li>
      <button
        className={classNames({ [classes.dotActive]: active }, classes.dot)}
        onClick={() => onClick()}
      />
    </li>
  );
};
function CarouselCourse({ courses }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.rooCarouselCourse}>
      <div className={classes.carousel}>
        <Carousel
          additionalTransfrom={0}
          arrows
          autoPlaySpeed={3000}
          centerMode={false}
          className=""
          containerClass={classes.containerClass}
          draggable
          focusOnSelect={false}
          infinite={false}
          itemClass=""
          keyBoardControl
          minimumTouchDrag={80}
          renderButtonGroupOutside={false}
          renderDotsOutside={false}
          responsive={responsive}
          showDots={true}
          sliderClass=""
          slidesToSlide={1}
          customDot={<CustomDot classes={classes} />}
          swipeable
          customLeftArrow={
            <Arrow className={classNames(classes.arrow, classes.arrowLeft)}>
              <Icon name="ArrowLeftNavyBlue" />
            </Arrow>
          }
          customRightArrow={
            <Arrow className={classNames(classes.arrow, classes.arrowRight)}>
              <Icon name="ArrowRightNavyBlue" />
            </Arrow>
          }
        >
          {courses.map((course, index) => (
            <CourseCard
              key={index}
              course={course}
              onClick={() => console.log('Pressed')}
              size={'small'}
              className={classes.carouselCard}
              isShowViewButton={true}
            />
          ))}
        </Carousel>
      </div>
    </div>
  );
}

export default CarouselCourse;
