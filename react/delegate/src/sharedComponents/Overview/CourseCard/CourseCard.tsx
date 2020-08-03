import * as React from 'react';
import Card, { PaddingOptions } from '../../core/Cards/Card';
import { createUseStyles, useTheme } from 'react-jss';
import { Background } from 'react-imgix';
import classNames from 'classnames';
import Button from '../../core/Input/Button';
import Icon from '../../core/Icon';
import { Theme } from 'helpers/theme';
import FooterIcon from './FooterIcon';
import CourseCompletion from 'sharedComponents/core/CourseCompletion';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    transition: '0.2s all',
    '&:hover': {
      boxShadow: '0 2px 12px 0 rgba(0,0,0,0.18)'
    },
    justifyContent: 'space-between'
  },
  noBorder: {
    borderRadius: `0 ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px`
  },
  mainContainer: {
    borderRadius: `0 ${theme.primaryBorderRadius}px 0 0`,
    height: 186
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between'
  },
  column: {
    display: 'flex',
    flexDirection: 'column'
  },
  heading: {
    alignSelf: 'flex-start',
    fontSize: theme.fontSizes.small,
    fontWeight: '700',
    color: theme.colors.primaryWhite,
    borderRadius: `0 0 ${theme.secondaryBorderRadius}px 0`,
    padding: `${theme.spacing(1)}px ${theme.spacing(2)}px`
  },
  icon: {
    alignSelf: 'flex-start',
    margin: `${theme.spacing(1)}px`
  },
  price: {
    color: theme.colors.primaryWhite,
    fontSize: theme.fontSizes.large,
    margin: '20px 20px 5px 20px',
    fontWeight: '800'
  },
  title: {
    color: theme.colors.primaryWhite,
    margin: '0 20px 30px 20px',
    fontSize: '23px',
    fontWeight: '900'
  },
  lecture: {
    display: 'flex',
    positio: 'relative',
    width: '100%',
    fontSize: theme.fontSizes.xSmall,
    lineHeight: '1.5em',
    margin: `${theme.spacing(2)}px ${theme.spacing(2)}px 0`,
    color: theme.colors.textGrey,
    padding: 0
  },
  lectureItem: {
    position: 'relative',
    marginLeft: '16px',
    '&:after': {
      content: '" "',
      position: 'absolute',
      top: '50%',
      left: '-8px',
      transform: 'translate(-50%, -50%)',
      width: '4px',
      height: '4px',
      backgroundColor: theme.colors.textGrey,
      borderRadius: '4px'
    },
    '&:first-child': {
      marginLeft: 0,
      '&:after': {
        display: 'none'
      }
    }
  },
  description: {
    flex: 3,
    lineHeight: '1.5em',
    overflow: 'hidden',
    fontWeight: '300',
    fontSize: theme.fontSizes.small,
    margin: `10px ${theme.spacing(2)}px ${theme.spacing(2)}px`,
    color: theme.colors.secondaryBlack,
    height: 56
  },
  progress: {
    color: theme.colors.primaryBlack,
    margin: '15px 0',
    fontSize: theme.fontSizes.tiny,
    fontWeight: '400'
  },
  button: {
    flex: 1,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  viewButton: {
    width: '100%',
    paddingTop: theme.spacing(1),
    paddingBottom: theme.spacing(1)
  },
  viewCourseButton: {
    width: '100%'
  },
  footer: {
    padding: `0 ${theme.spacing(2)}px`,
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGrey}`,
    borderWidth: '1px 0 0 0',
    borderRadius: `0 0 ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px`
  },
  filler: {
    flex: 0.5
  },
  small: {
    width: '298px',
    cursor: 'pointer'
  },
  large: {
    width: '610px'
  }
}));

export type SizeOptions = 'small' | 'large';

export interface Course {
  id: number;
  type: string;
  colour: string;
  url: string;
  title: string;
  price?: number;
  description: string;
  assigned?: number;
  expiring?: number;
  date?: string;
  location?: string;
  modules?: number;
  lessons?: number;
  videoTime?: number;
}

type Props = {
  course: Course;
  size?: SizeOptions;
  progress?: number;
  onClick: Function;
  padding?: PaddingOptions;
  className?: string;
  isShowViewButton?: boolean;
};

function CourseCard({
  course,
  onClick,
  size = 'small',
  progress,
  className,
  isShowViewButton = false
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const backgroundColor = { backgroundColor: course.colour };

  const description =
    course.description.length > 100
      ? `${course.description.substr(0, 100)}...`
      : course.description;
  return (
    <Card
      className={classNames(
        classes.root,
        classes.noBorder,
        classes[size],
        className
      )}
      onClick={() => {
        onClick();
      }}
    >
      <Background
        className={classNames(classes.mainContainer)}
        src={course.url}
        imgixParams={{ w: 300, h: 200, bri: -23 }}
      >
        <div className={classNames(classes.row)}>
          <div className={classNames(classes.heading)} style={backgroundColor}>
            {course.type}
          </div>
          <Icon
            className={classNames(classes.icon)}
            name="Card_SecondaryActon_Dots"
            size={18}
          />
        </div>
        {course.price !== undefined && (
          <div className={classNames(classes.price)}>
            £{course.price?.toFixed(2)}
          </div>
        )}
        <div className={classNames(classes.title)}>{course.title}</div>
      </Background>

      <div className={classNames(classes.column)}>
        {(course.modules || course.lessons || course.videoTime) && (
          <div className={classNames(classes.row)}>
            <div className={classNames(classes.lecture)}>
              {course.modules && (
                <span
                  className={classNames(classes.lectureItem)}
                >{`${course.modules} modules`}</span>
              )}
              {course.lessons && (
                <span
                  className={classNames(classes.lectureItem)}
                >{`${course.lessons} lessons`}</span>
              )}
              {course.videoTime && (
                <span
                  className={classNames(classes.lectureItem)}
                >{`${course.videoTime} hours of video`}</span>
              )}
            </div>
          </div>
        )}
        <div className={classNames(classes.row)}>
          <div
            className={classNames(classes.description)}
            style={{
              marginTop:
                course.modules || course.lessons || course.videoTime
                  ? '10px'
                  : '22px'
            }}
          >
            {description}
          </div>
          {size === 'large' && (
            <div className={classNames(classes.button)}>
              <Button
                archetype="submit"
                onClick={() => {
                  onClick();
                }}
              >
                Book Now
              </Button>
            </div>
          )}
        </div>
      </div>

      {progress !== undefined ? (
        <div className={classNames(classes.row, classes.footer)}>
          <div className={classes.progress}>PROGRESS</div>
          <CourseCompletion
            complete={progress}
            total={100}
            width={125}
            fraction={false}
          />
        </div>
      ) : isShowViewButton ? (
        <div className={classNames(classes.row, classes.footer)}>
          <div className={classNames(classes.viewButton)}>
            <Button
              archetype="submit"
              onClick={() => {
                onClick();
              }}
              bold
              className={classes.viewCourseButton}
            >
              View Course
            </Button>
          </div>
        </div>
      ) : size === 'small' ? (
        <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="Icon_Delegates" size={20} value={course.assigned} />
          <FooterIcon
            name="CourseExpiringSoon"
            size={20}
            value={course.expiring}
          />
          <div className={classNames(classes.filler)} />
        </div>
      ) : (
        <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="Course_Calendar" size={20} text={course.date} />
          <FooterIcon name="Location_Pin" size={20} text={course.location} />
          <div className={classNames(classes.filler)} />
        </div>
      )}
    </Card>
  );
}

export default CourseCard;
